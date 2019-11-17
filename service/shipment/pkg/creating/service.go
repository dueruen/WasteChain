package creating

import (
	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
)

type Service interface {
	CreateShipment(creationRequest *pb.CreateShipmentRequest) (string, error)
}

type Repository interface {
	CreateNewShipment(creationRequest *pb.CreateShipmentRequest) (string, error)
}

type service struct {
	createRepo Repository
}

func NewService(createRepo Repository) Service {
	return &service{createRepo}
}

func (srv *service) CreateShipment(creationRequest *pb.CreateShipmentRequest) (string, error) {
	return srv.createRepo.CreateNewShipment(creationRequest)
}
