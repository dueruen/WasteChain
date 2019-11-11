package listing

import (
	"context"

	pb "github.com/dueruen/WasteChain/service/account/gen/proto"
)

type Service interface {
	ListAllCompanies(ctx context.Context, in *pb.ListAllCompaniesRequest) (*pb.ListAllCompaniesResponse, error)
	GetCompany(ctx context.Context, in *pb.GetCompanyRequest) (*pb.GetCompanyResponse, error)
	ListAllEmployeesInCompany(ctx context.Context, in *pb.ListAllEmployeesInCompanyRequest) (*pb.ListAllEmployeesInCompanyResponse, error)
	GetEmployee(ctx context.Context, in *pb.GetEmployeeRequest) (*pb.GetEmployeeResponse, error)
}

type Repository interface {
	ListAllCompanies(in *pb.ListAllCompaniesRequest) (*pb.ListAllCompaniesResponse, error)
	GetCompany(in *pb.GetCompanyRequest) (*pb.GetCompanyResponse, error)
	ListAllEmployeesInCompany(in *pb.ListAllEmployeesInCompanyRequest) (*pb.ListAllEmployeesInCompanyResponse, error)
	GetEmployee(in *pb.GetEmployeeRequest) (*pb.GetEmployeeResponse, error)
}

type service struct {
	listRepo Repository
}

func NewService(listRepo Repository) Service {
	return &service{listRepo}
}

func (srv *service) ListAllCompanies(ctx context.Context, in *pb.ListAllCompaniesRequest) (*pb.ListAllCompaniesResponse, error) {
	return srv.listRepo.ListAllCompanies(in)
}

func (srv *service) GetCompany(ctx context.Context, in *pb.GetCompanyRequest) (*pb.GetCompanyResponse, error) {
	return srv.listRepo.GetCompany(in)
}

func (srv *service) ListAllEmployeesInCompany(ctx context.Context, in *pb.ListAllEmployeesInCompanyRequest) (*pb.ListAllEmployeesInCompanyResponse, error) {
	return srv.listRepo.ListAllEmployeesInCompany(in)
}

func (srv *service) GetEmployee(ctx context.Context, in *pb.GetEmployeeRequest) (*pb.GetEmployeeResponse, error) {
	return srv.listRepo.GetEmployee(in)
}
