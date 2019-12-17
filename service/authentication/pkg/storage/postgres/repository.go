package postgres

import (
	"errors"
	"fmt"
	"time"

	pb "github.com/dueruen/WasteChain/service/authentication/gen/proto"
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

func (storage *Storage) GetPassword(username string) (id, passward string, err error) {
	var credentials pb.Credentials
	storage.db.Where("username = ?", username).First(&credentials)
	if credentials.Username == "" {
		return "", "", errors.New("No such thing")
	}
	return credentials.UserID, credentials.Password, nil
}
