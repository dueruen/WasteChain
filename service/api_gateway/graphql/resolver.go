package graphql

import (
	"context"
	b64 "encoding/base64"
	"errors"

	pb "github.com/dueruen/WasteChain/service/api_gateway/gen/proto"
)

type Resolver struct {
	AccountClient        pb.AccountServiceClient
	SignatureClient      pb.SignatureServiceClient
	AuthenticationClient pb.AuthenticationServiceClient
	ShipmentClient       pb.ShipmentServiceClient
	QRClient             pb.QRServiceClient
}

func (r *Resolver) HistoryItem() HistoryItemResolver {
	return &historyItemResolver{r}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) TransferShipmentResponse() TransferShipmentResponseResolver {
	return &transferShipmentResponseResolver{r}
}

type historyItemResolver struct{ *Resolver }

//HistoryItem Resolver
func (r *historyItemResolver) Event(ctx context.Context, obj *pb.HistoryItem) (int, error) {
	if obj.Event == pb.ShipmentEvent_CREATED {
		return 0, nil
	}
	if obj.Event == pb.ShipmentEvent_TRANSFERED {
		return 1, nil
	}
	if obj.Event == pb.ShipmentEvent_PROCESSED {
		return 2, nil
	}

	return -1, errors.New("Invalid ShipmentEvent given")
}

type mutationResolver struct{ *Resolver }

//Mutation Resolvers for Account service
func (r *mutationResolver) CreateCompany(ctx context.Context, company pb.CreateCompany) (*pb.Company, error) {
	res, err := r.AccountClient.CreateCompany(ctx, &pb.CreateCompanyRequest{
		Company: &company,
	})
	if err != nil {
		return nil, err
	}
	return res.Company, nil
}
func (r *mutationResolver) CreateEmployee(ctx context.Context, employee pb.CreateEmployee) (*pb.Employee, error) {
	res, err := r.AccountClient.CreateEmployee(ctx, &pb.CreateEmployeeRequest{
		Employee: &employee,
	})
	if err != nil {
		return nil, err
	}
	return res.Employee, nil
}

//Mutation Resolvers for Signature service
func (r *mutationResolver) ContinueDoubleSign(ctx context.Context, request pb.ContinueDoubleSignRequest) (string, error) {
	_, err := r.SignatureClient.ContinueDoubleSign(ctx, &request)
	if err != nil {
		return err.Error(), err
	}
	return "", nil
}

//Mutation Resolvers for Authetication service
func (r *mutationResolver) Login(ctx context.Context, request pb.LoginRequest) (*pb.LoginResponse, error) {
	res, err := r.AuthenticationClient.Login(ctx, &request)
	if err != nil {
		return &pb.LoginResponse{}, err
	}
	return &pb.LoginResponse{Token: res.Token, Id: res.Id}, nil
}

//Mutation Resolvers for Shipment service
func (r *mutationResolver) CreateShipment(ctx context.Context, request pb.CreateShipmentRequest) (string, error) {
	res, err := r.ShipmentClient.CreateShipment(ctx, &request)
	if err != nil {
		return res.ID, err
	}
	return res.ID, nil
}

func (r *mutationResolver) TransferShipment(ctx context.Context, request pb.TransferShipmentRequest) (*pb.TransferShipmentResponse, error) {
	res, err := r.ShipmentClient.TransferShipment(ctx, &request)
	if err != nil {
		return &pb.TransferShipmentResponse{Error: err.Error()}, err
	}
	return res, nil
}

func (r *mutationResolver) ProcessShipment(ctx context.Context, request pb.ProcessShipmentRequest) (string, error) {
	_, err := r.ShipmentClient.ProcessShipment(ctx, &request)
	if err != nil {
		return err.Error(), err
	}
	return "", nil
}

type queryResolver struct{ *Resolver }

//Query Resolvers for Account service
func (r *queryResolver) ListAllCompanies(ctx context.Context) ([]*pb.Company, error) {
	res, err := r.AccountClient.ListAllCompanies(ctx, &pb.ListAllCompaniesRequest{})
	if err != nil {
		return nil, err
	}
	return res.Companies, nil
}
func (r *queryResolver) GetCompany(ctx context.Context, companyID string) (*pb.Company, error) {
	res, err := r.AccountClient.GetCompany(ctx, &pb.GetCompanyRequest{
		ID: companyID,
	})
	if err != nil {
		return nil, err
	}
	return res.Company, nil
}
func (r *queryResolver) ListAllEmployeesInCompany(ctx context.Context, companyID string) ([]*pb.Employee, error) {
	res, err := r.AccountClient.ListAllEmployeesInCompany(ctx, &pb.ListAllEmployeesInCompanyRequest{
		ID: companyID,
	})
	if err != nil {
		return nil, err
	}
	return res.Employees, nil
}
func (r *queryResolver) GetEmployee(ctx context.Context, employeeID string) (*pb.Employee, error) {
	res, err := r.AccountClient.GetEmployee(ctx, &pb.GetEmployeeRequest{
		ID: employeeID,
	})
	if err != nil {
		return nil, err
	}
	return res.Employee, nil
}

//Query Resolvers for Shipment service
func (r *queryResolver) GetShipmentDetails(ctx context.Context, request string) (*pb.Shipment, error) {
	res, err := r.ShipmentClient.GetShipmentDetails(ctx, &pb.GetShipmentDetailsRequest{ID: request})
	if err != nil {
		return nil, err
	}
	return res.Shipment, nil
}

func (r *queryResolver) ListAllShipments(ctx context.Context) ([]*pb.Shipment, error) {
	res, err := r.ShipmentClient.ListAllShipments(ctx, &pb.ListAllShipmentsRequest{})
	if err != nil {
		return nil, err
	}
	return res.ShipmentList, nil
}

func (r *queryResolver) ToQr(ctx context.Context, data string) (string, error) {
	res, err := r.QRClient.CreateQRCode(ctx, &pb.CreateQRRequest{DataString: data})
	if err != nil {
		return "", err
	}
	base64 := b64.StdEncoding.EncodeToString(res.QRCode)
	return base64, nil
}

type transferShipmentResponseResolver struct{ *Resolver }

func (r *transferShipmentResponseResolver) QRCode(ctx context.Context, obj *pb.TransferShipmentResponse) (string, error) {
	base64 := b64.StdEncoding.EncodeToString(obj.QRCode)
	return base64, nil
}
