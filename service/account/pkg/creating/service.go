package creating

import (
	"context"

	pb "github.com/dueruen/WasteChain/service/account/gen/proto"
)

type Service interface {
	CreateCompany(ctx context.Context, in *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error)
	CreateEmployee(ctx context.Context, in *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error)
}

type Repository interface {
	CreateNewCompany(company *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error)
	CreateEmployee(*pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error)
}

type service struct {
	createRepo Repository
}

func NewService(createRepo Repository) Service {
	return &service{createRepo}
}

func (srv *service) CreateCompany(ctx context.Context, in *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error) {
	// TODO make validation
	return srv.createRepo.CreateNewCompany(in)
}

func (srv *service) CreateEmployee(ctx context.Context, in *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {
	// TODO make validation
	return srv.createRepo.CreateEmployee(in)
}
