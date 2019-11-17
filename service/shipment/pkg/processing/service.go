package processing

import (
	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
)

type Service interface {
	ProcessShipment(processingRequest *pb.ProcessShipmentRequest) error
}

type Repository interface {
	ProcessShipment(processingRequest *pb.ProcessShipmentRequest) error
}

type service struct {
	processingRepo Repository
}

func NewService(processingRepo Repository) Service {
	return &service{processingRepo}
}

func (srv *service) ProcessShipment(processingRequest *pb.ProcessShipmentRequest) error {
	return srv.processingRepo.ProcessShipment(processingRequest)
}
