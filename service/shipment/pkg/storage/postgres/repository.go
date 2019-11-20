package postgres

import (
	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(host, port, user, dbname, password string) (*Storage, error) {
	db, err := connect(host, port, user, dbname, password)
	if err != nil {

		return nil, err
	}
	err = createSchema(db)
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func Close(s *Storage) {
	s.db.Close()
}

func connect(host, port, user, dbname, password string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host="+host+" port="+port+" user="+user+" dbname="+dbname+" password="+password+" sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createSchema(db *gorm.DB) error {
	if db.HasTable(&pb.Shipment{}) {
		return nil
	}

	db.AutoMigrate(&pb.Shipment{})

	return nil
}

func (storage *Storage) CreateNewShipment(shipment *pb.CreateShipmentRequest, timestamp string) (string, error) {
	newShipment := &pb.Shipment{
		CurrentHolderID: shipment.CurrentHolderID,
		WasteType:       shipment.WasteType,
		History: []*pb.HistoryItem{
			&pb.HistoryItem{
				Event:      0,
				OwnerID:    shipment.CurrentHolderID,
				ReceiverID: "0",
				TimeStamp:  timestamp,
				Location:   shipment.Location,
			},
		},
	}

	id, _ := uuid.NewV4()
	newShipment.ID = id.String()

	storage.db.Create(newShipment)
	return newShipment.ID, nil
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

	for _, shipment := range shipments {
		shipment = getAllShipmentData(storage.db, shipment)
	}
	return nil, shipments
}

func getAllShipmentData(db *gorm.DB, shipment *pb.Shipment) *pb.Shipment {
	var history []*pb.HistoryItem
	db.Where("id = ?", shipment.ID).Find(&history)
	shipment.History = history

	return shipment
}

func (storage *Storage) ProcessShipment(processRequest *pb.ProcessShipmentRequest, timeStamp string) error {
	var shipment pb.Shipment
	storage.db.Where("id = ?", processRequest.ID).First(&shipment)
	shipment = *getAllShipmentData(storage.db, &shipment)

	history := shipment.History

	processEvent := &pb.HistoryItem{
		Event:      2,
		OwnerID:    processRequest.OwnerID,
		ReceiverID: "0",
		TimeStamp:  timeStamp,
		Location:   processRequest.Location,
	}
	history = append(history, processEvent)
	shipment.History = history

	storage.db.Update(shipment)
	return nil
}

func (storage *Storage) TransferShipment(transferRequest *pb.TransferShipmentRequest, timestamp string) error {
	var shipment pb.Shipment
	storage.db.Where("id = ?", transferRequest.ShipmentID).First(&shipment)
	shipment = *getAllShipmentData(storage.db, &shipment)

	history := shipment.History

	processEvent := &pb.HistoryItem{
		Event:      1,
		OwnerID:    transferRequest.OwnerID,
		ReceiverID: transferRequest.ReceiverID,
		TimeStamp:  timestamp,
		Location:   transferRequest.Location,
	}
	history = append(history, processEvent)
	shipment.History = history

	storage.db.Update(shipment)
	return nil

}
