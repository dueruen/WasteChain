package listing

import (
	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
)

type Service interface {
	GetShipmentDetails(*pb.GetShipmentDetailsRequest) (error, *pb.Shipment)
	ListAllShipments() (error, []*pb.Shipment)
}

type Repository interface {
	GetShipmentDetails(*pb.GetShipmentDetailsRequest) (error, *pb.Shipment)
	ListAllShipments() (error, []*pb.Shipment)
}

type service struct {
	listRepo Repository
}

func NewService(listRepo Repository) Service {
	return &service{listRepo}
}

func (srv *service) GetShipmentDetails(request *pb.GetShipmentDetailsRequest) (error, *pb.Shipment) {
	return srv.listRepo.GetShipmentDetails(request)
}

func (srv *service) ListAllShipments() (error, []*pb.Shipment) {
	return srv.listRepo.ListAllShipments()
}
