package postgres

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ShipmentData struct {
	ShipmentID string
	Addr       string
	Seed       string
}

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
	if db.HasTable(&ShipmentData{}) {
		return nil
	}

	db.AutoMigrate(&ShipmentData{})
	return nil
}

func (storage *Storage) SaveShipmentInfo(shipmentID, addr, seed string) {
	newShipmentData := &ShipmentData{
		ShipmentID: shipmentID,
		Addr:       addr,
		Seed:       seed,
	}
	storage.db.Create(newShipmentData)
}

func (storage *Storage) GetShipmentAddress(shipmentID string) (addr string, err error) {
	var shipmentData ShipmentData
	storage.db.Where("shipment_id = ?", shipmentID).First(&shipmentData)
	if shipmentData.Addr == "" {
		return "", errors.New("No such thing")
	}
	return shipmentData.Addr, nil
}
