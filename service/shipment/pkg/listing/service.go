package listing

import (
	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
)

type Service interface {
	GetShipmentDetails(string) (error, *pb.Shipment)
	ListAllShipments() (error, []*pb.Shipment)
}

type Repository interface {
	GetShipmentDetails(string) (error, *pb.Shipment)
	ListAllShipments() (error, []*pb.Shipment)
}

type service struct {
	listRepo Repository
}

func NewService(listRepo Repository) Service {
	return &service{listRepo}
}

func (srv *service) GetShipmentDetails(id string) (error, *pb.Shipment) {
	return srv.listRepo.GetShipmentDetails(id)
}

func (srv *service) ListAllShipments() (error, []*pb.Shipment) {
	return srv.listRepo.ListAllShipments()
}
