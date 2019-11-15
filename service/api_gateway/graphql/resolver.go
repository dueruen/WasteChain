package graphql

// import (
// 	"context"
// )

// // type Resolver struct {
// // 	accountClient pb.AccountServiceClient
// // }

// // func (r *Resolver) Mutation() MutationResolver {
// // 	return &mutationResolver{r}
// // }
// // func (r *Resolver) Query() QueryResolver {
// // 	return &queryResolver{r}
// // }

// // type mutationResolver struct{ *Resolver }

// // func (r *mutationResolver) CreateCompany(ctx context.Context, company *CreateCompany) (*Company, error) {
// // 	res, err := r.accountClient.CreateCompany(ctx, &pb.CreateCompanyRequest{company})
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return nil, res.Company
// // }

// // func (r *mutationResolver) CreateEmployee(ctx context.Context, employee CreateEmployee) (*Employee, error) {
// // 	panic("not implemented")
// // }

// // type queryResolver struct{ *Resolver }

// // func (r *queryResolver) ListAllCompanies(ctx context.Context) ([]*Company, error) {
// // 	panic("not implemented")
// // }
// // func (r *queryResolver) GetCompany(ctx context.Context, companyID string) (*Company, error) {
// // 	panic("not implemented")
// // }
// // func (r *queryResolver) ListAllEmployeesInCompany(ctx context.Context, companyID string) ([]*Employee, error) {
// // 	panic("not implemented")
// // }
// // func (r *queryResolver) GetEmployee(ctx context.Context, employeeID string) (*Employee, error) {
// // 	panic("not implemented")
// // }

// type Resolver struct{}

// func (r *Resolver) Mutation() MutationResolver {
// 	return &mutationResolver{r}
// }
// func (r *Resolver) Query() QueryResolver {
// 	return &queryResolver{r}
// }

// type mutationResolver struct{ *Resolver }

// func (r *mutationResolver) CreateCompany(ctx context.Context, company *CreateCompany) (*Company, error) {
// 	panic("not implemented")
// }
// func (r *mutationResolver) CreateEmployee(ctx context.Context, employee CreateEmployee) (*Employee, error) {
// 	panic("not implemented")
// }

// type queryResolver struct{ *Resolver }

// func (r *queryResolver) ListAllCompanies(ctx context.Context) ([]*Company, error) {
// 	panic("not implemented")
// }
// func (r *queryResolver) GetCompany(ctx context.Context, companyID string) (*Company, error) {
// 	panic("not implemented")
// }
// func (r *queryResolver) ListAllEmployeesInCompany(ctx context.Context, companyID string) ([]*Employee, error) {
// 	panic("not implemented")
// }
// func (r *queryResolver) GetEmployee(ctx context.Context, employeeID string) (*Employee, error) {
// 	panic("not implemented")
// }

//package tmp

import (
	"context"

	pb "github.com/dueruen/WasteChain/service/api_gateway/gen/proto"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	AccountClient pb.AccountServiceClient
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

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

type queryResolver struct{ *Resolver }

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
