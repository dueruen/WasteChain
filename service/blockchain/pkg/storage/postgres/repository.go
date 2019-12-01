package postgres

import (
	"errors"
	"fmt"
	"time"

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
