package grpc

import (
	pb "github.com/dueruen/WasteChain/service/account/gen/proto"
	"github.com/dueruen/WasteChain/service/account/pkg/creating"
	"github.com/dueruen/WasteChain/service/account/pkg/listing"
)

func mapListCompanyToPb(company *listing.Company) *pb.Company {
	return &pb.Company{
		ID:          company.ID,
		Name:        company.Name,
		Address:     mapListAddressToPb(company.Address),
		ContactInfo: mapAllListContactInfoToPb(company.ContactInfo),
		Employees:   mapAllListEmployeeToPb(company.Employees),
	}
}

func mapPbCompanyToCreate(company *pb.Company) *creating.Company {
	return &creating.Company{
		Name:    company.Name,
		Address: mapPbAddressToCreating(company.Address),
	}
}

func mapPbCompanyListing(company *pb.Company) *listing.Company {
	return &listing.Company{
		ID:          company.ID,
		Name:        company.Name,
		Address:     mapPbAddressToListing(company.Address),
		ContactInfo: mapAllPbContactInfoToListing(company.ContactInfo),
		Employees:   mapAllPbEmployeeToListing(company.Employees),
	}
}

func mapAllListCompaniesToPb(companies []*listing.Company) []*pb.Company {
	mapped := make([]*pb.Company, 0)
	for _, value := range companies {
		mapped = append(mapped, mapListCompanyToPb(value))
	}
	return mapped
}

func mapAllPbCompaniesToListing(companies []*pb.Company) []*listing.Company {
	mapped := make([]*listing.Company, 0)
	for _, value := range companies {
		mapped = append(mapped, mapPbCompanyListing(value))
	}
	return mapped
}

func mapListAddressToPb(list *listing.Address) *pb.Address {
	return &pb.Address{
		ID:       list.ID,
		RoadName: list.RoadName,
		Number:   list.Number,
		ZipCode:  list.ZipCode,
	}
}

func mapPbAddressToCreating(create *pb.Address) *creating.Address {
	return &creating.Address{
		RoadName: create.RoadName,
		Number:   create.Number,
		ZipCode:  create.ZipCode,
	}
}

func mapPbCreateAddressToCreating(create *pb.CreateAddress) *creating.Address {
	return &creating.Address{
		RoadName: create.RoadName,
		Number:   create.Number,
		ZipCode:  create.ZipCode,
	}
}

func mapPbAddressToListing(address *pb.Address) *listing.Address {
	return &listing.Address{
		ID:       address.ID,
		RoadName: address.RoadName,
		Number:   address.Number,
		ZipCode:  address.ZipCode,
	}
}

func mapPbCreateContactInfoToCreating(contactInfo *pb.CreateContactInfo) *creating.ContactInfo {
	return &creating.ContactInfo{
		Title:       contactInfo.Title,
		Mail:        contactInfo.Mail,
		PhoneNumber: contactInfo.PhoneNumber,
	}
}

func mapPbContactInfoToCreating(contactInfo *pb.ContactInfo) *creating.ContactInfo {
	return &creating.ContactInfo{
		Title:       contactInfo.Title,
		Mail:        contactInfo.Mail,
		PhoneNumber: contactInfo.PhoneNumber,
		CompanyID:   contactInfo.CompanyID,
	}
}

func mapPbContactInfoToListing(contactInfo *pb.ContactInfo) *listing.ContactInfo {
	return &listing.ContactInfo{
		ID:          contactInfo.ID,
		Title:       contactInfo.Title,
		Mail:        contactInfo.Mail,
		PhoneNumber: contactInfo.PhoneNumber,
		CompanyID:   contactInfo.CompanyID,
	}
}

func mapListContactInfoToPb(contactInfo *listing.ContactInfo) *pb.ContactInfo {
	return &pb.ContactInfo{
		ID:          contactInfo.ID,
		Title:       contactInfo.Title,
		Mail:        contactInfo.Mail,
		PhoneNumber: contactInfo.PhoneNumber,
		CompanyID:   contactInfo.CompanyID,
	}
}

func mapAllPbContactInfoToCreating(contactInfo []*pb.ContactInfo) []*creating.ContactInfo {
	mapped := make([]*creating.ContactInfo, 0)
	for _, value := range contactInfo {
		mapped = append(mapped, mapPbContactInfoToCreating(value))
	}
	return mapped
}

func mapAllPbContactInfoToListing(contactInfo []*pb.ContactInfo) []*listing.ContactInfo {
	mapped := make([]*listing.ContactInfo, 0)
	for _, value := range contactInfo {
		mapped = append(mapped, mapPbContactInfoToListing(value))
	}
	return mapped
}

func mapAllListContactInfoToPb(contactInfo []*listing.ContactInfo) []*pb.ContactInfo {
	mapped := make([]*pb.ContactInfo, 0)
	for _, value := range contactInfo {
		mapped = append(mapped, mapListContactInfoToPb(value))
	}
	return mapped
}

func mapPbEmployeeToCreate(employee *pb.Employee) *creating.Employee {
	return &creating.Employee{
		Name:      employee.Name,
		CompanyID: employee.CompanyID,
	}
}

func mapListEmployeeToPb(employee *listing.Employee) *pb.Employee {
	return &pb.Employee{
		ID:        employee.ID,
		Name:      employee.Name,
		CompanyID: employee.CompanyID,
	}
}

func mapPbEmployeeToListing(employee *pb.Employee) *listing.Employee {
	return &listing.Employee{
		ID:        employee.ID,
		Name:      employee.Name,
		CompanyID: employee.CompanyID,
	}
}

func mapAllPbEmployeeToListing(employees []*pb.Employee) []*listing.Employee {
	mapped := make([]*listing.Employee, 0)
	for _, value := range employees {
		mapped = append(mapped, mapPbEmployeeToListing(value))
	}
	return mapped
}

func mapAllListEmployeeToPb(employees []*listing.Employee) []*pb.Employee {
	mapped := make([]*pb.Employee, 0)
	for _, value := range employees {
		mapped = append(mapped, mapListEmployeeToPb(value))
	}
	return mapped
}
