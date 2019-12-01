package sign

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	pb "github.com/dueruen/WasteChain/service/signature/gen/proto"

	"github.com/dueruen/WasteChain/service/signature/pkg/key"
	"github.com/gofrs/uuid"
)

type Repository interface {
	StoreDoubleSignProgress(progressID, currentHolderID, shipmentID string, signature, dataHash []byte) error
	GetStoredDoubleSignProgress(id string) (currentHolderID, shipmentID string, signature, dataHash []byte, err error)
}

type EventHandler interface {
	DoubleSignNeeded(event *pb.DoubleSignNeededEvent)
	DoubleSignDone(event *pb.DoneEvent)
	SingleSignDone(event *pb.DoneEvent)
}

type Service interface {
	StartDoubleSign(req *pb.StartDoubleSignRequest) error
	FinishStartDoubleSign(progressID string, qrCode []byte)
	ContinueDoubleSign(req *pb.ContinueDoubleSignRequest) error
	SingleSign(req *pb.SingleSignRequest) error
	SingleVerify(req *pb.SingleVerifyRequest) (bool, error)
	DoubleVerify(req *pb.DoubleVerifyRequest) (bool, error)
}

type service struct {
	repo         Repository
	keyService   key.Service
	eventHandler EventHandler
	qrClient     pb.QRServiceClient
}

func NewService(repo Repository, keyService key.Service, eventHandler EventHandler, qrClient pb.QRServiceClient) Service {
	return &service{repo, keyService, eventHandler, qrClient}
}

func (service *service) StartDoubleSign(req *pb.StartDoubleSignRequest) error {
	//sign data
	dataHash, _, signature, err := service.sign(req.Data, req.CurrentHolderID, req.CurrentHolderPassword)
	if err != nil {
		return err
	}

	//Store progres for the next to sign
	id, _ := uuid.NewV4()
	err = service.repo.StoreDoubleSignProgress(id.String(), req.CurrentHolderID, req.ShipmentID, signature, dataHash)
	if err != nil {
		return err
	}

	service.qrClient.CreateQRCode(context.Background(), &pb.CreateQRRequest{
		ID:         id.String(),
		DataString: id.String(),
	})

	return nil
}

func (service *service) FinishStartDoubleSign(progressID string, qrCode []byte) {
	//Get stored progress
	currentHolderID, _, _, _, _ := service.repo.GetStoredDoubleSignProgress(progressID)

	//Public event DoubleSignNeeded
	service.eventHandler.DoubleSignNeeded(&pb.DoubleSignNeededEvent{
		CurrentHolderID: currentHolderID,
		ContinueID:      progressID,
		QRCode:          qrCode,
	})
}

func (service *service) ContinueDoubleSign(req *pb.ContinueDoubleSignRequest) error {
	//Get stored progress
	_, shipmentID, currentHolderSignature, firstDataHash, err := service.repo.GetStoredDoubleSignProgress(req.ContinueID)
	if err != nil {
		return err
	}

	//New holder sign
	_, _, newHolderSignature, err := service.sign(firstDataHash, req.NewHolderID, req.NewHolderPassword)
	if err != nil {
		return err
	}

	//Public event DoubleSignDone
	service.eventHandler.DoubleSignDone(&pb.DoneEvent{
		EventType:              pb.DoneEventType_DOUBLE_SIGN_DONE,
		CurrentHolderSignature: currentHolderSignature,
		NewHolderSignature:     newHolderSignature,
		ShipmentID:             shipmentID,
	})
	return nil
}

func (service *service) SingleSign(req *pb.SingleSignRequest) error {
	//sign data
	_, _, signature, err := service.sign(req.Data, req.UserID, req.Password)
	if err != nil {
		return err
	}

	//Public event SingleSignDone
	service.eventHandler.SingleSignDone(&pb.DoneEvent{
		EventType:              pb.DoneEventType_SINGLE_SIGN_DONE,
		CurrentHolderSignature: signature,
		ShipmentID:             req.ShipmentID,
	})
	return nil
}

func (service *service) SingleVerify(req *pb.SingleVerifyRequest) (bool, error) {
	//Hash data
	dataHash, _ := hashData(req.Data)

	//Verify
	ok, err := service.verify(req.UserID, dataHash, req.Signature)
	if !ok || err != nil {
		return false, err
	}
	return true, nil
}

func (service *service) DoubleVerify(req *pb.DoubleVerifyRequest) (bool, error) {
	//Hash data
	dataHash, _ := hashData(req.Data)

	//Verify currentHolder
	_, err := service.verify(req.CurrentHolderID, dataHash, req.CurrentHolderSignature)
	if err != nil {
		return false, err
	}

	//Hash data
	newDataHash, _ := hashData(dataHash)

	//Verify nextHolder
	_, err = service.verify(req.NewHolderID, newDataHash, req.NewHolderSignature)
	if err != nil {
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
