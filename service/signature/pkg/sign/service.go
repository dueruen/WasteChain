package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	pb "github.com/dueruen/WasteChain/service/signature/gen/proto"

	"github.com/dueruen/WasteChain/service/signature/pkg/key"
	"github.com/gofrs/uuid"
)

type Repository interface {
	StoreDoubleSignProgress(id, signature, dataHash string) error
	GetStoredDoubleSignProgress(id string) (signature, dataHash string, err error)
}

type EventHandler interface {
	DoubleSignNeeded(event *pb.DoubleSignNeededEvent)
	DoubleSignDone(event *pb.DoneEvent)
	SingleSignDone(event *pb.DoneEvent)
}

type Service interface {
}

type service struct {
	repo         Repository
	keyService   key.Service
	eventHandler EventHandler
}

func NewService(repo Repository, keyService key.Service, eventHandler EventHandler) Service {
	return &service{repo, keyService, eventHandler}
}

func (service *service) StartDoubleSign(data []byte, currentHolderID, password string) error {
	//sign data
	dataHash, _, signature, err := service.sign(data, currentHolderID, password)
	if err != nil {
		return err
	}

	//Store progres for the next to sign
	id, _ := uuid.NewV4()
	err = service.repo.StoreDoubleSignProgress(id.String(), string(signature), string(dataHash))
	if err != nil {
		return err
	}

	//Public event DoubleSignNeeded
	service.eventHandler.DoubleSignNeeded(&pb.DoubleSignNeededEvent{
		CurrentHolderID: currentHolderID,
		QRCode:          nil,
	})

	return nil
}

func (service *service) ContinueDoubleSign(continueID, newHolderID, password string) error {
	//Get stored progress
	currentHolderSignature, firstDataHash, err := service.repo.GetStoredDoubleSignProgress(continueID)
	if err != nil {
		return err
	}

	//New holder sign
	_, _, newHolderSignature, err := service.sign([]byte(firstDataHash), newHolderID, password)
	if err != nil {
		return err
	}

	//Public event DoubleSignDone
	service.eventHandler.DoubleSignDone(&pb.DoneEvent{
		EventType:              pb.DoneEventType_DOUBLE_SIGN_DONE,
		CurrentHolderSignature: currentHolderSignature,
		NewHolderSignature:     string(newHolderSignature),
	})

	return nil
}

func (service *service) SingleSign(data []byte, userID, password string) error {
	//sign data
	_, _, signature, err := service.sign(data, userID, password)
	if err != nil {
		return err
	}

	//Public event SingleSignDone
	service.eventHandler.SingleSignDone(&pb.DoneEvent{
		EventType:              pb.DoneEventType_SINGLE_SIGN_DONE,
		CurrentHolderSignature: string(signature),
	})

	return nil
}

func (service *service) SingleVerify(userID string, data, signature []byte) (bool, error) {
	//Hash data
	dataHash, _ := hashData(data)

	//Verify
	ok, err := service.verify(userID, dataHash, signature)
	if !ok || err != nil {
		return false, err
	}
	return true, nil
}

func (service *service) DoubleVerify(input *pb.DoubleVerifyRequest) (bool, error) {
	//Hash data
	dataHash, _ := hashData([]byte(input.Data))

	//Verify currentHolder
	ok, err := service.verify(input.CurrentHolderID, dataHash, []byte(input.CurrentHolderSignature))
	if !ok || err != nil {
		return false, err
	}

	//Hash data
	newDataHash, _ := hashData(dataHash)

	//Verify nextHolder
	ok, err = service.verify(input.NewHolderID, newDataHash, []byte(input.NewHolderSignature))
	if !ok || err != nil {
		return false, err
	}
	return true, nil
}

func (service *service) sign(data []byte, id, password string) (dataHash []byte, privateKey *rsa.PrivateKey, signature []byte, err error) {
	//Hash data
	dataHash, hashType := hashData(data)

	//Get currentHolder private key
	privateKey, err = service.keyService.GetPrivateKey(id, password)
	if err != nil {
		return nil, nil, nil, err
	}

	//The currentHolder signs the data hash
	signature, err = generateSignature(dataHash, privateKey, hashType)
	return
}

func hashData(data []byte) (dataHash []byte, hashType crypto.Hash) {
	hash := sha256.New()
	hash.Write(data)

	PSSdata := data
	hashType = crypto.SHA256
	pssh := hashType.New()
	pssh.Write(PSSdata)
	dataHash = pssh.Sum(nil)
	return dataHash, hashType
}

func generateSignature(data []byte, privateKey *rsa.PrivateKey, hashType crypto.Hash) (signature []byte, err error) {
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

func (service *service) verify(userID string, dataHash, signature []byte) (bool, error) {
	//Get public key
	publicKey, err := service.keyService.GetPublicKey(userID)
	if err != nil {
		return false, err
	}

	//Verify hash with the signature
	opts := rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto}
	err = rsa.VerifyPSS(
		publicKey,
		crypto.SHA256,
		dataHash,
		signature,
		&opts,
	)
	if err != nil {
		return false, err
	}
	return true, nil
}
