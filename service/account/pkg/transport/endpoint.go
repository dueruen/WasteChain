package transport

import (
	"context"

	pb "github.com/dueruen/WasteChain/service/account/gen/proto"
	"github.com/dueruen/WasteChain/service/account/pkg/creating"
	"github.com/dueruen/WasteChain/service/account/pkg/listing"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateCompany             endpoint.Endpoint
	CreateEmployee            endpoint.Endpoint
	ListAllCompanies          endpoint.Endpoint
	GetCompany                endpoint.Endpoint
	ListAllEmployeesInCompany endpoint.Endpoint
	GetEmployee               endpoint.Endpoint
}

func MakeEndpoints(createService creating.Service, listService listing.Service) Endpoints {
	return Endpoints{
		CreateCompany:             makeCreateCompanyEndpoint(createService),
		CreateEmployee:            makeCreateEmployeeEndpoint(createService),
		ListAllCompanies:          makeListAllCompaniesEndpoint(listService),
		GetCompany:                makeGetCompanyEndpoint(listService),
		ListAllEmployeesInCompany: makeListAllEmployeesInCompanyEndpoint(listService),
		GetEmployee:               makeGetEmployeeEndpoint(listService),
	}
}

func makeCreateCompanyEndpoint(service creating.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.CreateCompanyRequest)
		res, _ := service.CreateCompany(req.Company)
		return &pb.CreateCompanyResponse{Company: res}, nil
	}
}

func makeCreateEmployeeEndpoint(service creating.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.CreateEmployeeRequest)
		res, _ := service.CreateEmployee(req.Employee)
		return &pb.CreateEmployeeResponse{Employee: res}, nil
	}
}

func makeListAllCompaniesEndpoint(service listing.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, _ := service.ListAllCompanies()
		return &pb.ListAllCompaniesResponse{Companies: res}, nil
	}
}

func makeGetCompanyEndpoint(service listing.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.GetCompanyRequest)
		res, _ := service.GetCompany(req.ID)
		return &pb.GetCompanyResponse{Company: res}, nil
	}
}

func makeListAllEmployeesInCompanyEndpoint(service listing.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ListAllEmployeesInCompanyRequest)
		res, _ := service.ListAllEmployeesInCompany(req.ID)
		return &pb.ListAllEmployeesInCompanyResponse{Employees: res}, nil
	}
}

func makeGetEmployeeEndpoint(service listing.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pb.GetEmployeeRequest)
		res, _ := service.GetEmployee(req.ID)
		return &pb.GetEmployeeResponse{Employee: res}, nil
	}
}
