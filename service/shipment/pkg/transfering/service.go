package transfering

import (
	"context"
	"time"

	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
	encode "github.com/dueruen/WasteChain/service/shipment/pkg/encode"
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

func NewService(transferingRepo Repository, signClient pb.SignatureServiceClient) Service {
	return &service{transferingRepo, signClient}
}

func (srv *service) TransferShipment(transferRequest *pb.TransferShipmentRequest) error {
	historyItem, error := srv.transferingRepo.TransferShipment(transferRequest, time.Now().String())
	if error != nil {
		return error
	}
	byteEvent := encode.ToByte(historyItem)

	srv.signClient.StartDoubleSign(context.Background(), &pb.StartDoubleSignRequest{
		Data:                  byteEvent,
		CurrentHolderID:       transferRequest.OwnerID,
		CurrentHolderPassword: transferRequest.Password,
		ShipmentID:            transferRequest.ShipmentID,
	})
	return nil
}
