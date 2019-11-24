package processing

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
)

type Service interface {
	ProcessShipment(processingRequest *pb.ProcessShipmentRequest) error
}

type Repository interface {
	ProcessShipment(processingRequest *pb.ProcessShipmentRequest, timestamp string) (*pb.HistoryItem, error)
}

type service struct {
	processingRepo Repository
	signClient     pb.SignatureServiceClient
}

type dataEvent struct {
	Event      pb.ShipmentEvent
	OwnerID    string
	ReceiverID string
	TimeStamp  string
	Location   string
}

func NewService(processingRepo Repository, signClient pb.SignatureServiceClient) Service {
	return &service{processingRepo, signClient}
}

func (srv *service) ProcessShipment(processingRequest *pb.ProcessShipmentRequest) error {
	historyItem, error := srv.processingRepo.ProcessShipment(processingRequest, time.Now().String())
	if error != nil {
		return error
	}
	dataEvent := mapHistoryItemToDataEvent(historyItem)
	byteEvent := dataEventToByteArray(dataEvent)

	srv.signClient.SingleSign(context.Background(), &pb.SingleSignRequest{
		Data:       byteEvent,
		UserID:     processingRequest.OwnerID,
		Password:   processingRequest.Password,
		ShipmentID: processingRequest.ID,
	})
	return nil
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
