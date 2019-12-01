package postgres

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Storage struct {
	db *gorm.DB
}

type KeyData struct {
	UserID              string
	EncryptedPrivateKey []byte
	PublicKey           []byte
}

type ProgressData struct {
	ProgressID      string
	CurrentHolderID string
	ShipmentID      string
	Signature       []byte
	DataHash        []byte
}

func NewStorage(host, port, user, dbname, password string) (*Storage, error) {
	db := connect(host, port, user, dbname, password)

	err := createSchema(db)
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func Close(s *Storage) {
	s.db.Close()
}

func connect(host, port, user, dbname, password string) *gorm.DB {
	i := 5
	for i > 0 {
		db, err := gorm.Open("postgres", "host="+host+" port="+port+" user="+user+" dbname="+dbname+" password="+password+" sslmode=disable")
		if err != nil {
			fmt.Println("Can't connect to db, sleeping for 2 sec")
			time.Sleep(2 * time.Second)
			i--
			continue
		} else {
			return db
		}
	}
	panic("Not connected to storage")
}

func createSchema(db *gorm.DB) error {
	if db.HasTable(&KeyData{}) && db.HasTable(&ProgressData{}) {
		return nil
	}

	db.AutoMigrate(&KeyData{}, &ProgressData{})

	return nil
}

func (storage *Storage) SaveKeys(userID string, encryptedPrivateKey, publicKey []byte) error {
	data := &KeyData{
		UserID:              userID,
		EncryptedPrivateKey: encryptedPrivateKey,
		PublicKey:           publicKey,
	}
	storage.db.Create(data)
	return nil
}

func (storage *Storage) GetPublicKey(userID string) (publicKey []byte, err error) {
	var data KeyData
	storage.db.Where("user_id = ?", userID).First(&data)
	return data.PublicKey, nil
}

func (storage *Storage) GetPrivateKey(userID string) (encryptedPrivateKey []byte, err error) {
	var data KeyData
	storage.db.Where("user_id = ?", userID).First(&data)
	return data.EncryptedPrivateKey, nil
}

func (storage *Storage) StoreDoubleSignProgress(id, currentHolderID, shipmentID string, signature, dataHash []byte) error {
	data := &ProgressData{
		ProgressID:      id,
		CurrentHolderID: currentHolderID,
		ShipmentID:      shipmentID,
		Signature:       signature,
		DataHash:        dataHash,
	}
	storage.db.Create(data)
	return nil
}

func (storage *Storage) GetStoredDoubleSignProgress(id string) (currentHolderID, shipmentID string, signature, dataHash []byte, err error) {
	var data ProgressData
	storage.db.Where("progress_id = ?", id).First(&data)
	return data.CurrentHolderID, data.ShipmentID, data.Signature, data.DataHash, nil
}
