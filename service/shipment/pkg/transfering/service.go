package transfering

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
)

type Service interface {
	TransferShipment(*pb.TransferShipmentRequest) error
}

type Repository interface {
	TransferShipment(*pb.TransferShipmentRequest, string) (*pb.HistoryItem, error)
}

type service struct {
	transferingRepo Repository
	signClient      pb.SignatureServiceClient
}

type dataEvent struct {
	Event      pb.ShipmentEvent
	OwnerID    string
	ReceiverID string
	TimeStamp  string
	Location   string
}

func NewService(transferingRepo Repository, signClient pb.SignatureServiceClient) Service {
	return &service{transferingRepo, signClient}
}

func (srv *service) TransferShipment(transferRequest *pb.TransferShipmentRequest) error {
	historyItem, error := srv.transferingRepo.TransferShipment(transferRequest, time.Now().String())
	dataEvent := mapHistoryItemToDataEvent(historyItem)
	byteEvent := dataEventToByteArray(dataEvent)

	srv.signClient.StartDoubleSign(context.Background(), &pb.StartDoubleSignRequest{
		Data:                  byteEvent,
		CurrentHolderID:       transferRequest.OwnerID,
		CurrentHolderPassword: transferRequest.Password,
		ShipmentID:            transferRequest.ShipmentID,
	})
	return error
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
