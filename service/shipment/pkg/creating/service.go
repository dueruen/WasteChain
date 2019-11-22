package creating

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
)

type Service interface {
	CreateShipment(creationRequest *pb.CreateShipmentRequest) (string, error)
}

type Repository interface {
	CreateNewShipment(creationRequest *pb.CreateShipmentRequest, timestamp string) (string, *pb.HistoryItem, error)
}

type service struct {
	createRepo Repository
	signClient pb.SignatureServiceClient
}

type dataEvent struct {
	Event      pb.ShipmentEvent
	OwnerID    string
	ReceiverID string
	TimeStamp  string
	Location   string
}

func NewService(createRepo Repository, signClient pb.SignatureServiceClient) Service {
	return &service{createRepo, signClient}
}

func (srv *service) CreateShipment(creationRequest *pb.CreateShipmentRequest) (string, error) {
	id, historyItem, error := srv.createRepo.CreateNewShipment(creationRequest, time.Now().String())
	dataEvent := mapHistoryItemToDataEvent(historyItem)
	byteEvent := dataEventToByteArray(dataEvent)

	srv.signClient.SingleSign(context.Background(), &pb.SingleSignRequest{
		Data:       byteEvent,
		UserID:     creationRequest.CurrentHolderID,
		Password:   creationRequest.Password,
		ShipmentID: id,
	})
	return id, error
}

func mapHistoryItemToDataEvent(historyItem *pb.HistoryItem) *dataEvent {
	newDataEvent := &dataEvent{
		Event:      historyItem.Event,
		OwnerID:    historyItem.OwnerID,
		ReceiverID: historyItem.ReceiverID,
		TimeStamp:  historyItem.TimeStamp,
		Location:   historyItem.Location,
	}

	return newDataEvent
}

func dataEventToByteArray(event *dataEvent) []byte {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	enc.Encode(event)
	return buf.Bytes()
}
