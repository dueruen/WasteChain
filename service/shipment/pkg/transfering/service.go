package transfering

import (
	"context"
	"time"

	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
	encode "github.com/dueruen/WasteChain/service/shipment/pkg/encode"
)

type Service interface {
	TransferShipment(*pb.TransferShipmentRequest) (*pb.TransferShipmentResponse, error)
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

func (srv *service) TransferShipment(transferRequest *pb.TransferShipmentRequest) (*pb.TransferShipmentResponse, error) {
	historyItem, err := srv.transferingRepo.TransferShipment(transferRequest, time.Now().String())

	if err != nil {
		return &pb.TransferShipmentResponse{Error: err.Error()}, err
	}
	byteEvent := encode.ToByte(historyItem)

	res, err := srv.signClient.StartDoubleSign(context.Background(), &pb.StartDoubleSignRequest{
		Data:                  byteEvent,
		CurrentHolderID:       transferRequest.OwnerID,
		CurrentHolderPassword: transferRequest.Password,
		ShipmentID:            transferRequest.ShipmentID,
	})
	if err != nil {
		return &pb.TransferShipmentResponse{Error: err.Error()}, err
	}

	return &pb.TransferShipmentResponse{
		ContinueID: res.ContinueID,
		QRCode:     res.QRCode,
	}, nil
}
