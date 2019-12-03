package processing

import (
	"context"
	"time"

	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
	encode "github.com/dueruen/WasteChain/service/shipment/pkg/encode"
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

func NewService(processingRepo Repository, signClient pb.SignatureServiceClient) Service {
	return &service{processingRepo, signClient}
}

func (srv *service) ProcessShipment(processingRequest *pb.ProcessShipmentRequest) error {
	historyItem, error := srv.processingRepo.ProcessShipment(processingRequest, time.Now().String())
	if error != nil {
		return error
	}
	byteEvent := encode.ToByte(historyItem)

	srv.signClient.SingleSign(context.Background(), &pb.SingleSignRequest{
		Data:       byteEvent,
		UserID:     processingRequest.OwnerID,
		Password:   processingRequest.Password,
		ShipmentID: processingRequest.ShipmentID,
	})
	return nil
}
