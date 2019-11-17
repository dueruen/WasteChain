package key

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"

	ept "github.com/dueruen/WasteChain/service/signature/pkg/encrypt"
)

type Repository interface {
	SaveKeys(userID string, passphrase string, encryptedPrivateKey string, publicKey string) error
	GetPublicKey(userID string) (publicKey string, err error)
	GetPrivateKey(userID string) (privateKey string, err error)
}

type Service interface {
	CreateKeyPair(userID string, passphrase string) error
	GetPublicKey(userID string) (*rsa.PublicKey, error)
	GetPrivateKey(userID string, passphrase string) (*rsa.PrivateKey, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (service *service) CreateKeyPair(userID string, passphrase string) error {
	//Generate keys
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	//private key to json string
	privateKeyJson, err := json.Marshal(&privateKey)
	if err != nil {
		return err
	}

	//public key to json string
	publicKeyJson, _ := json.Marshal(&privateKey.PublicKey)
	if err != nil {
		return err
	}

	//Encrypt private key
	ciphertext, err := ept.Encrypt(privateKeyJson, []byte(passphrase))
	if err != nil {
		return err
	}

	err = service.repo.SaveKeys(userID, passphrase, string(ciphertext), string(publicKeyJson))
	if err != nil {
		return err
	}

	return nil
}

func (service *service) GetPublicKey(userID string) (*rsa.PublicKey, error) {
	publicKeyJSON, err := service.repo.GetPublicKey(userID)
	if err != nil {
		return nil, err
	}
	publicKey := rsa.PublicKey{}
	json.Unmarshal([]byte(publicKeyJSON), &publicKey)
	return &publicKey, nil
}

func (service *service) GetPrivateKey(userID string, passphrase string) (*rsa.PrivateKey, error) {
	//Get ciphertext
	ciphertext, err := service.repo.GetPrivateKey(userID)
	if err != nil {
		return nil, err
	}
	//decrypt ciphertext
	privateKeyJSON, err := ept.Decrypt([]byte(ciphertext), []byte(passphrase))
	if err != nil {
		return nil, err
	}

	privateKey := rsa.PrivateKey{}
	json.Unmarshal([]byte(privateKeyJSON), &privateKey)

	return &privateKey, nil
}
