package grpc

import (
	"context"

	"github.com/dueruen/WasteChain/service/account/pkg/creating"

	kitgrpc "github.com/go-kit/kit/transport/grpc"

	pb "github.com/dueruen/WasteChain/service/account/gen/proto"
	"github.com/dueruen/WasteChain/service/account/pkg/transport"
)

type server struct {
	createCompany             kitgrpc.Handler
	createEmployee            kitgrpc.Handler
	listAllCompanies          kitgrpc.Handler
	getCompany                kitgrpc.Handler
	listAllEmployeesInCompany kitgrpc.Handler
	getEmployee               kitgrpc.Handler
}

func NewGrpcServer(endpoins transport.Endpoints, options []kitgrpc.ServerOption) pb.AccountServiceServer {
	return &server{
		createCompany:             kitgrpc.NewServer(endpoins.CreateCompany, decodeCreateCompanyRequest, encodeCreateCompanyResponse),
		createEmployee:            kitgrpc.NewServer(endpoins.CreateEmployee, decodeCreateEmployeeRequest, encodeCreateEmployeeResponse),
		listAllCompanies:          kitgrpc.NewServer(endpoins.ListAllCompanies, decodeListAllCompaniesRequest, encodeListAllCompaniesResponse),
		getCompany:                kitgrpc.NewServer(endpoins.GetCompany, decodeGetCompanyRequest, encodeGetCompanyResponse),
		listAllEmployeesInCompany: kitgrpc.NewServer(endpoins.ListAllEmployeesInCompany, decodeListAllEmployeesInCompanyRequest, encodeListAllEmployeesInCompanyResponse),
		getEmployee:               kitgrpc.NewServer(endpoins.GetEmployee, decodeGetEmployeeRequest, encodeGetEmployeeResponse),
	}
}

func (server *server) CreateCompany(ctx context.Context, req *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error) {
	_, rep, err := server.createCompany.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateCompanyResponse), nil
}

func decodeCreateCompanyRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateCompanyRequest)
	return transport.CreateCompanyRequest{
		Company: &creating.Company{
			Name:        req.Name,
			Address:     mapPbCreateAddressToCreating(req.Address),
			ContactInfo: []*creating.ContactInfo{mapPbCreateContactInfoToCreating(req.ContactInfo)},
		},
	}, nil
}

func encodeCreateCompanyResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(transport.CreateCompanyResponse)
	return &pb.CreateCompanyResponse{
		Company: mapListCompanyToPb(res.Company),
	}, nil
}

func (server *server) CreateEmployee(ctx context.Context, req *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {
	_, rep, err := server.createEmployee.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateEmployeeResponse), nil
}

func decodeCreateEmployeeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateEmployeeRequest)
	return transport.CreateEmployeeRequest{
		Employee: &creating.Employee{
			Name:      req.Name,
			CompanyID: req.CompanyID,
		},
	}, nil
}

func encodeCreateEmployeeResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(transport.CreateEmployeeResponse)
	return &pb.CreateEmployeeResponse{
		Employee: mapListEmployeeToPb(res.Employee),
	}, nil
}

func (server *server) ListAllCompanies(ctx context.Context, req *pb.ListAllCompaniesRequest) (*pb.ListAllCompaniesResponse, error) {
	_, rep, err := server.listAllCompanies.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListAllCompaniesResponse), nil
}

func decodeListAllCompaniesRequest(_ context.Context, request interface{}) (interface{}, error) {
	//req := request.(*pb.ListAllCompaniesRequest)
	return transport.ListAllCompaniesRequest{}, nil
}

func encodeListAllCompaniesResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(transport.ListAllCompaniesResponse)
	return &pb.ListAllCompaniesResponse{
		Companies: mapAllListCompaniesToPb(res.Companies),
	}, nil
}

func (server *server) GetCompany(ctx context.Context, req *pb.GetCompanyRequest) (*pb.GetCompanyResponse, error) {
	_, rep, err := server.getCompany.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetCompanyResponse), nil
}

func decodeGetCompanyRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetCompanyRequest)
	return transport.GetCompanyRequest{
		ID: req.ID,
	}, nil
}

func encodeGetCompanyResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(transport.GetCompanyResponse)
	return &pb.GetCompanyResponse{
		Company: mapListCompanyToPb(res.Company),
	}, nil
}

func (server *server) ListAllEmployeesInCompany(ctx context.Context, req *pb.ListAllEmployeesInCompanyRequest) (*pb.ListAllEmployeesInCompanyResponse, error) {
	_, rep, err := server.listAllEmployeesInCompany.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListAllEmployeesInCompanyResponse), nil
}

func decodeListAllEmployeesInCompanyRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ListAllEmployeesInCompanyRequest)
	return transport.ListAllEmployeesInCompanyRequest{
		ID: req.ID,
	}, nil
}

func encodeListAllEmployeesInCompanyResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(transport.ListAllEmployeesInCompanyResponse)
	return &pb.ListAllEmployeesInCompanyResponse{
		Employees: mapAllListEmployeeToPb(res.Employees),
	}, nil
}

func (server *server) GetEmployee(ctx context.Context, req *pb.GetEmployeeRequest) (*pb.GetEmployeeResponse, error) {
	_, rep, err := server.getEmployee.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetEmployeeResponse), nil
}

func decodeGetEmployeeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetEmployeeRequest)
	return transport.GetEmployeeRequest{
		ID: req.ID,
	}, nil
}

func encodeGetEmployeeResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(transport.GetEmployeeResponse)
	return &pb.GetEmployeeResponse{
		Employee: mapListEmployeeToPb(res.Employee),
	}, nil
}
