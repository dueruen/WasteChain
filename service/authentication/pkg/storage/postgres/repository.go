package postgres

import (
	"errors"

	pb "github.com/dueruen/WasteChain/service/authentication/gen/proto"
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
	if db.HasTable(&pb.Credentials{}) {
		return nil
	}

	db.AutoMigrate(&pb.Credentials{})
	return nil
}

func (storage *Storage) SaveCredentials(userID, hashedPassword, username string) error {
	newCredential := &pb.Credentials{
		UserID:   userID,
		Password: hashedPassword,
		Username: username,
	}
	storage.db.Create(newCredential)
	return nil
}

func (storage *Storage) GetPassword(username string) (string, error) {
	var credentials pb.Credentials
	storage.db.Where("username = ?", username).First(&credentials)
	if credentials.Username == "" {
		return "", errors.New("No such thing")
	}
	return credentials.Password, nil
}
