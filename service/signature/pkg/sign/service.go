package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	"github.com/dueruen/WasteChain/service/signature/pkg/key"
	"github.com/gofrs/uuid"
)

type Repository interface {
	StoreDoubleSignProgres(id, signature, dataHash string) error
}

type Service interface {
}

type service struct {
	repo       Repository
	keyService key.Service
}

func NewService(repo Repository, keyService key.Service) Service {
	return &service{repo, keyService}
}

func (service *service) StartDoubleSign(data []byte, currentHolderID, password string) error {
	//Hash data
	dataHash, hashType := service.hashData(data)

	//Get currentHolder private key
	privateKey, err := service.keyService.GetPrivateKey(currentHolderID, password)
	if err != nil {
		return err
	}

	//The currentHolder signs the data hash
	signature, err := sign(dataHash, privateKey, hashType)

	//Store progres for the next to sign
	id, _ := uuid.NewV4()
	err = service.repo.StoreDoubleSignProgres(id.String(), string(signature), string(dataHash))
	if err != nil {
		return err
	}

}

func (service *service) SingleSign() {

}

func (service *service) hashData(data []byte) (dataHash []byte, hashType crypto.Hash) {
	hash := sha256.New()
	hash.Write(data)

	PSSdata := data
	hashType = crypto.SHA256
	pssh := hashType.New()
	pssh.Write(PSSdata)
	dataHash = pssh.Sum(nil)
	return dataHash, hashType
}

func sign(data []byte, privateKey *rsa.PrivateKey, hashType crypto.Hash) (signature []byte, err error) {
	opts := rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto}
	signature, err = rsa.SignPSS(
		rand.Reader,
		privateKey,
		hashType,
		data,
		&opts,
	)

	if err != nil {
		return nil, err
	}
	return signature, nil
}

func Verify(publicKey *rsa.PublicKey, hashed, signature []byte) (err error) {
	opts := rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto}
	err = rsa.VerifyPSS(
		publicKey,
		crypto.SHA256,
		hashed,
		signature,
		&opts,
	)
	if err != nil {
		return err
	}
	return nil
}
