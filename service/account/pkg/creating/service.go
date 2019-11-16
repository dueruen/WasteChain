package creating

import (
	pb "github.com/dueruen/WasteChain/service/account/gen/proto"
)

type Service interface {
	CreateCompany(company *pb.CreateCompany) (*pb.Company, error)
	CreateEmployee(employee *pb.CreateEmployee) (*pb.Employee, error)
}

type Repository interface {
	CreateNewCompany(company *pb.CreateCompany) (*pb.Company, error)
	CreateEmployee(employee *pb.CreateEmployee) (*pb.Employee, error)
}

type service struct {
	createRepo Repository
}

func NewService(createRepo Repository) Service {
	return &service{createRepo}
}

func (srv *service) CreateCompany(company *pb.CreateCompany) (*pb.Company, error) {
	// TODO make validation
	return srv.createRepo.CreateNewCompany(company)
}

func (srv *service) CreateEmployee(employee *pb.CreateEmployee) (*pb.Employee, error) {
	// TODO make validation
	return srv.createRepo.CreateEmployee(employee)
}
