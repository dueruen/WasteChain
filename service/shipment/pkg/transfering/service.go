package transfering

import (
	"time"

	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
)

type Service interface {
	TransferShipment(*pb.TransferShipmentRequest) error
}

type Repository interface {
	TransferShipment(*pb.TransferShipmentRequest, string) error
}

type service struct {
	transferingRepo Repository
}

func NewService(transferingRepo Repository) Service {
	return &service{transferingRepo}
}

func (srv *service) TransferShipment(transferRequest *pb.TransferShipmentRequest) error {
	return srv.transferingRepo.TransferShipment(transferRequest, time.Now().String())
}
