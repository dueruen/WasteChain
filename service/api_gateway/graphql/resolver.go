package graphql

import (
	"context"

	pb "github.com/dueruen/WasteChain/service/api_gateway/gen/proto"
)

type Resolver struct {
	AccountClient        pb.AccountServiceClient
	QRClient             pb.QRServiceClient
	SignatureClient      pb.SignatureServiceClient
	AuthenticationClient pb.AuthenticationServiceClient
	ShipmentClient       pb.ShipmentServiceClient
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
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

//Mutation Resolvers for QR service
func (r *mutationResolver) CreateQRCode(ctx context.Context, dataString string) (*[]byte, error) {
	res, err := r.QRClient.CreateQRCode(ctx, &pb.CreateQRRequest{
		DataString: dataString,
	})
	if err != nil {
		return nil, err
	}
	return &res.QRCode, nil
}

//Mutation Resolvers for Signature service
func (r *mutationResolver) ContinueDoubleSign(ctx context.Context, request *pb.ContinueDoubleSignRequest) error {
	_, err := r.SignatureClient.ContinueDoubleSign(ctx, request)
	if err != nil {
		return err
	}
	return nil
}

//Mutation Resolvers for Authetication service
func (r *mutationResolver) CreateCredentials(ctx context.Context, request *pb.CreateCredentialsRequest) error {
	_, err := r.AuthenticationClient.CreateCredentials(ctx, request)
	if err != nil {
		return err
	}
	return nil
}

func (r *mutationResolver) Login(ctx context.Context, request *pb.LoginRequest) (string, error) {
	res, err := r.AuthenticationClient.Login(ctx, request)
	if err != nil {
		return "", err
	}
	return res.Token, nil
}

func (r *mutationResolver) Validate(ctx context.Context, request *pb.ValidateRequest) (bool, error) {
	res, err := r.AuthenticationClient.Validate(ctx, request)
	if err != nil {
		return res.Valid, err
	}
	return res.Valid, nil
}

//Mutation Resolvers for Shipment service
func (r *mutationResolver) CreateShipment(ctx context.Context, request *pb.CreateShipmentRequest) (string, error) {
	res, err := r.ShipmentClient.CreateShipment(ctx, request)
	if err != nil {
		return res.ID, err
	}
	return res.ID, nil
}

func (r *mutationResolver) TransferShipment(ctx context.Context, request *pb.TransferShipmentRequest) error {
	_, err := r.ShipmentClient.TransferShipment(ctx, request)
	if err != nil {
		return err
	}
	return nil
}

func (r *mutationResolver) ProcessShipment(ctx context.Context, request *pb.ProcessShipmentRequest) error {
	_, err := r.ShipmentClient.ProcessShipment(ctx, request)
	if err != nil {
		return err
	}
	return nil
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
func (r *queryResolver) GetShipmentDetails(ctx context.Context, request *pb.GetShipmentDetailsRequest) (*pb.Shipment, error) {
	res, err := r.ShipmentClient.GetShipmentDetails(ctx, request)
	if err != nil {
		return nil, err
	}
	return res.Shipment, nil
}

func (r *queryResolver) ListAllShipments(ctx context.Context, request *pb.ListAllShipmentsRequest) ([]*pb.Shipment, error) {
	res, err := r.ShipmentClient.ListAllShipments(ctx, request)
	if err != nil {
		return nil, err
	}
	return res.ShipmentList, nil
}
