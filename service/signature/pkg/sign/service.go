package sign

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"

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
	VerifyHistory(req *pb.VerifyHistoryRequest) *pb.VerifyHistoryResponse
}

type service struct {
	repo         Repository
	keyService   key.Service
	eventHandler EventHandler
	qrClient     pb.QRServiceClient
	blockClient  pb.BlockchainServiceClient
}

func NewService(repo Repository, keyService key.Service, eventHandler EventHandler, qrClient pb.QRServiceClient, blockClient pb.BlockchainServiceClient) Service {
	return &service{repo, keyService, eventHandler, qrClient, blockClient}
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

	doneEvent := &pb.DoneEvent{
		EventType:              pb.DoneEventType_DOUBLE_SIGN_DONE,
		CurrentHolderSignature: currentHolderSignature,
		NewHolderSignature:     newHolderSignature,
		ShipmentID:             shipmentID,
	}

	data, _ := json.Marshal(doneEvent)
	doneEvent.Data = data

	//Public event DoubleSignDone
	service.eventHandler.DoubleSignDone(doneEvent)
	return nil
}

func (service *service) SingleSign(req *pb.SingleSignRequest) error {
	//sign data
	_, _, signature, err := service.sign(req.Data, req.UserID, req.Password)
	if err != nil {
		return err
	}

	fmt.Println("SIGNATURE########: ", signature)
	doneEvent := &pb.DoneEvent{
		EventType:              pb.DoneEventType_SINGLE_SIGN_DONE,
		CurrentHolderSignature: signature,
		ShipmentID:             req.ShipmentID,
	}

	data, _ := json.Marshal(doneEvent)
	doneEvent.Data = data
	fmt.Println("¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤¤")
	itemData := pb.DoneEvent{}
	json.Unmarshal(data, &itemData)
	fmt.Println("DATA string: ", string(data))
	fmt.Println("dATAA unmarshal: ", itemData)

	//Public event SingleSignDone
	service.eventHandler.SingleSignDone(doneEvent)

	return nil
}

func (service *service) VerifyHistory(req *pb.VerifyHistoryRequest) *pb.VerifyHistoryResponse {
	res, err := service.blockClient.GetShipmentData(context.Background(), &pb.GetShipmentDataRequest{
		ShipmentID: req.ShipmentID,
	})

	if err != nil {
		fmt.Println("SIGN GetShipmentData ERR: ", err)
		return &pb.VerifyHistoryResponse{Ok: false, Error: err.Error()}
	}
	if len(res.History) == 0 {
		fmt.Println("SIGN NO HISTORY on IOTA")
		return &pb.VerifyHistoryResponse{Ok: false, Error: errors.New("No history").Error()}
	}

	for i, item := range res.History {
		itemData := pb.DoneEvent{}
		json.Unmarshal(item, &itemData)

		reqItem := req.History[i]
		fmt.Println("&&&&&&&&&&&&&&&&&&& DATA: ", string(reqItem.Data))

		if (i == 0 && reqItem.CurrentHolderID != "" && reqItem.NewHolderID == "") || (i != 0 && reqItem.CurrentHolderID != "" && reqItem.NewHolderID == "") {
			err := service.singleVerify(reqItem.CurrentHolderID, itemData.CurrentHolderSignature, reqItem.Data)
			if err != nil {
				fmt.Println("SIGN singleVerify ERR: ", err)
				return &pb.VerifyHistoryResponse{Ok: false, Error: err.Error()}
			}
		} else if i != 0 && reqItem.CurrentHolderID != "" && reqItem.NewHolderID != "" {
			err := service.doubleVerify(reqItem.CurrentHolderID, reqItem.NewHolderID, reqItem.Data, itemData.CurrentHolderSignature, itemData.NewHolderSignature)
			if err != nil {
				fmt.Println("SIGN doubleVerify ERR: ", err)
				return &pb.VerifyHistoryResponse{Ok: false, Error: err.Error()}
			}
		} else {
			fmt.Println("SIGN Error verifying history")
			return &pb.VerifyHistoryResponse{Ok: false, Error: errors.New("Error verifying history").Error()}
		}
	}

	fmt.Println("VERIFY DATA OK!!!!")
	return &pb.VerifyHistoryResponse{Ok: true}
}

func (service *service) singleVerify(id string, signature, data []byte) error {
	//Hash data
	dataHash, _ := hashData(data)

	//Verify
	ok, err := service.verify(id, dataHash, signature)
	if !ok || err != nil {
		return err
	}
	return nil
}

func (service *service) doubleVerify(currentHolderID, newHolderID string, data, currentHolderSignature, newHolderSignature []byte) error {
	//Hash data
	dataHash, _ := hashData(data)

	//Verify currentHolder
	_, err := service.verify(currentHolderID, dataHash, currentHolderSignature)
	if err != nil {
		return err
	}

	//Hash data
	newDataHash, _ := hashData(dataHash)

	//Verify nextHolder
	_, err = service.verify(newHolderID, newDataHash, newHolderSignature)
	if err != nil {
		return err
	}
	return nil
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
		fmt.Println("GetPublicKey ERR: ", err)
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
		fmt.Println("VerifyPSS Err: ", err)
		return false, err
	}
	return true, nil
}
