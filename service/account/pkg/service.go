package pkg

import (
	"github.com/dueruen/WasteChain/service/account/pkg/creating"
	"github.com/dueruen/WasteChain/service/account/pkg/listing"
)

type listingRepository = listing.Repository
type creatingRepository = creating.Repository

type listingService = listing.Service
type creatingService = creating.Service

func Run() {

}

type Repository interface {
	listingRepository
	creatingRepository
}

type Service struct {
	listingService
	creatingService
}

// func NewService(repo Repository) *Service {
// 	return &Service{repo}
// }

// func (srv *Service) CreateCompany(ctx context.Context, in *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error) {
// 	// TODO make validation
// 	return srv.repo.CreateNewCompany(in)
// }

// func (srv *Service) CreateEmployee(ctx context.Context, in *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {
// 	// TODO make validation
// 	return srv.repo.CreateEmployee(in)
// }

// func (srv *Service) ListAllCompanies(ctx context.Context, in *pb.ListAllCompaniesRequest) (*pb.ListAllCompaniesResponse, error) {
// 	return srv.repo.ListAllCompanies(in)
// }

// func (srv *Service) GetCompany(ctx context.Context, in *pb.GetCompanyRequest) (*pb.GetCompanyResponse, error) {
// 	return srv.repo.GetCompany(in)
// }

// func (srv *Service) ListAllEmployeesInCompany(ctx context.Context, in *pb.ListAllEmployeesInCompanyRequest) (*pb.ListAllEmployeesInCompanyResponse, error) {
// 	return srv.repo.ListAllEmployeesInCompany(in)
// }

// func (srv *Service) GetEmployee(ctx context.Context, in *pb.GetEmployeeRequest) (*pb.GetEmployeeResponse, error) {
// 	return srv.repo.GetEmployee(in)
// }
