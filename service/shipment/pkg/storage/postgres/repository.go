package postgres

import (
	"errors"
	"fmt"
	"time"

	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(db_string string) (*Storage, error) {
	db := connect(db_string)

	err := createSchema(db)
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func Close(s *Storage) {
	s.db.Close()
}

func connect(db_string string) *gorm.DB {
	i := 5
	for i > 0 {
		db, err := gorm.Open("postgres", db_string)
		if err != nil {
			fmt.Println("Can't connect to db, sleeping for 2 sec, err: ", err)
			time.Sleep(2 * time.Second)
			i--
			continue
		} else {
			fmt.Println("Connected to storage")
			return db
		}
	}
	panic("Not connected to storage")
}

func createSchema(db *gorm.DB) error {
	if db.HasTable(&pb.Shipment{}) {
		return nil
	}

	db.AutoMigrate(&pb.Shipment{}, &pb.HistoryItem{})

	db.Model(&pb.HistoryItem{}).AddForeignKey("shipment_id", "shipments(id)", "CASCADE", "CASCADE")

	return nil
}

func (storage *Storage) CreateNewShipment(shipment *pb.CreateShipmentRequest, timestamp string, companyID string) (string, *pb.HistoryItem, error) {
	newShipment := &pb.Shipment{
		CurrentHolderID:    shipment.CurrentHolderID,
		ProducingCompanyID: companyID,
		WasteType:          shipment.WasteType,
		History: []*pb.HistoryItem{
			&pb.HistoryItem{
				Event:      0,
				OwnerID:    shipment.CurrentHolderID,
				ReceiverID: "",
				TimeStamp:  timestamp,
				Location:   shipment.Location,
				Published:  false,
			},
		},
	}

	id, _ := uuid.NewV4()
	newShipment.ID = id.String()

	historyItemID, _ := uuid.NewV4()
	newShipment.History[0].ID = historyItemID.String()
	newShipment.History[0].ShipmentID = id.String()

	historyItem := newShipment.History[0]

	storage.db.Create(newShipment)
	return newShipment.ID, historyItem, nil
}

func (storage *Storage) GetShipmentDetails(getRequest *pb.GetShipmentDetailsRequest) (error, *pb.Shipment) {
	var shipment pb.Shipment
	storage.db.Where("id = ?", getRequest.ID).First(&shipment)
	shipment = *getAllShipmentData(storage.db, &shipment)
	return nil, &shipment
}

func (storage *Storage) ListAllShipments() (error, []*pb.Shipment) {
	var shipments []*pb.Shipment
	storage.db.Find(&shipments)
	var shipmentsToBeReturned []*pb.Shipment

	for _, shipment := range shipments {
		shipment = getAllShipmentData(storage.db, shipment)
	}

	for _, shipment := range shipments {
		if len(shipment.History) == 0 {
			continue
		}
		shipment = getAllShipmentData(storage.db, shipment)
		shipmentsToBeReturned = append(shipmentsToBeReturned, shipment)
	}

	return nil, shipmentsToBeReturned
}

func (storage *Storage) ListUsersShipments(request *pb.ListUsersShipmentsRequest) (error, []*pb.Shipment) {
	var shipments []*pb.Shipment
	storage.db.Where("current_holder_id = ?", request.ID).Find(&shipments)
	var shipmentsToBeReturned []*pb.Shipment

	for _, shipment := range shipments {
		shipment = getAllShipmentData(storage.db, shipment)
	}

	for _, shipment := range shipments {
		if len(shipment.History) == 0 {
			continue
		}
		shipment = getAllShipmentData(storage.db, shipment)
		shipmentsToBeReturned = append(shipmentsToBeReturned, shipment)
	}

	return nil, shipmentsToBeReturned
}

func getAllShipmentData(db *gorm.DB, shipment *pb.Shipment) *pb.Shipment {
	var history []*pb.HistoryItem
	db.Where("shipment_id = ? AND published = true", shipment.ID).Find(&history)
	shipment.History = history

	return shipment
}

func (storage *Storage) ProcessShipment(processRequest *pb.ProcessShipmentRequest, timeStamp string) (*pb.HistoryItem, error) {

	if storage.shipmentHasBeenProcessed(processRequest.ShipmentID) {
		return nil, errors.New("Shipment has already been processed, and can therefore not be processed again")
	}

	processEvent := &pb.HistoryItem{
		Event:      2,
		OwnerID:    processRequest.OwnerID,
		ReceiverID: "",
		TimeStamp:  timeStamp,
		Location:   processRequest.Location,
		Published:  false,
	}
	processEventID, _ := uuid.NewV4()

	processEvent.ID = processEventID.String()
	processEvent.ShipmentID = processRequest.ShipmentID

	storage.db.Create(processEvent)
	return processEvent, nil
}

func (storage *Storage) TransferShipment(transferRequest *pb.TransferShipmentRequest, timestamp string) (*pb.HistoryItem, error) {

	if storage.shipmentHasBeenProcessed(transferRequest.ShipmentID) {
		return nil, errors.New("Shipment has been processed, and can therefore not be transfered")
	}

	transferEvent := &pb.HistoryItem{
		Event:      1,
		OwnerID:    transferRequest.OwnerID,
		ReceiverID: transferRequest.ReceiverID,
		TimeStamp:  timestamp,
		Location:   transferRequest.Location,
		Published:  false,
	}

	transferEventID, _ := uuid.NewV4()

	transferEvent.ID = transferEventID.String()
	transferEvent.ShipmentID = transferRequest.ShipmentID

	storage.db.Create(transferEvent)
	return transferEvent, nil
}

func (storage *Storage) LatestHistoryEventIsPublished(shipmentID string) error {
	var hi pb.HistoryItem
	storage.db.Where("shipment_id = ? AND published = false", shipmentID).Last(&hi)
	hi.Published = true

	if hi.Event == 1 {
		var shipment pb.Shipment
		storage.db.Where("id = ?", shipmentID).First(&shipment)
		shipment.CurrentHolderID = hi.ReceiverID
		storage.db.Save(shipment)
	}

	storage.db.Save(hi)
	return nil
}

func (storage *Storage) shipmentHasBeenProcessed(shipmentID string) bool {
	var hi pb.HistoryItem
	storage.db.Where("shipment_id = ?", shipmentID).Order("time_stamp desc").First(&hi)

	if hi.Event == 2 {
		return true
	}
	return false
}
