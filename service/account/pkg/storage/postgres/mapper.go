package postgres

import (
	"github.com/dueruen/WasteChain/service/account/pkg/creating"
	"github.com/dueruen/WasteChain/service/account/pkg/listing"
)

func mapCreateCompany(create *creating.Company) *Company {
	return &Company{
		Name:        create.Name,
		Address:     mapCreateAddress(create.Address),
		ContactInfo: mapAllCreateContactInfo(create.ContactInfo),
	}
}

func mapToListCompany(company *Company) *listing.Company {
	return &listing.Company{
		ID:          company.ID,
		Name:        company.Name,
		Address:     mapToListAddress(company.Address),
		ContactInfo: mapAllToListContactInfo(company.ContactInfo),
		Employees:   mapAllToListEmployee(company.Employees),
	}
}

func mapAllToListCompany(companies []*Company) []*listing.Company {
	mapped := make([]*listing.Company, 0)
	for _, value := range companies {
		mapped = append(mapped, mapToListCompany(value))
	}
	return mapped
}

func mapCreateAddress(create *creating.Address) *Address {
	return &Address{
		RoadName: create.RoadName,
		Number:   create.Number,
		ZipCode:  create.ZipCode,
	}
}

func mapToListAddress(address *Address) *listing.Address {
	return &listing.Address{
		ID:       address.ID,
		RoadName: address.RoadName,
		Number:   address.Number,
		ZipCode:  address.ZipCode,
	}
}

func mapCreateContactInfo(create *creating.ContactInfo) *ContactInfo {
	return &ContactInfo{
		Title:       create.Title,
		Mail:        create.Mail,
		PhoneNumber: create.PhoneNumber,
		CompanyID:   create.CompanyID,
	}
}

func mapAllCreateContactInfo(contactInfo []*creating.ContactInfo) []*ContactInfo {
	mapped := make([]*ContactInfo, 0)
	for _, value := range contactInfo {
		mapped = append(mapped, mapCreateContactInfo(value))
	}
	return mapped
}

func mapToListContactInfo(contactInfo *ContactInfo) *listing.ContactInfo {
	return &listing.ContactInfo{
		ID:          contactInfo.ID,
		Title:       contactInfo.Title,
		Mail:        contactInfo.Mail,
		PhoneNumber: contactInfo.PhoneNumber,
		CompanyID:   contactInfo.CompanyID,
	}
}

func mapAllToListContactInfo(contactInfo []*ContactInfo) []*listing.ContactInfo {
	mapped := make([]*listing.ContactInfo, 0)
	for _, value := range contactInfo {
		mapped = append(mapped, mapToListContactInfo(value))
	}
	return mapped
}

func mapCreateEmployee(create *creating.Employee) *Employee {
	return &Employee{
		Name:      create.Name,
		CompanyID: create.CompanyID,
	}
}

func mapAllCreateEmployee(employees []*creating.Employee) []*Employee {
	mapped := make([]*Employee, 0)
	for _, value := range employees {
		mapped = append(mapped, mapCreateEmployee(value))
	}
	return mapped
}

func mapToListEmployee(employee *Employee) *listing.Employee {
	return &listing.Employee{
		ID:        employee.ID,
		Name:      employee.Name,
		CompanyID: employee.CompanyID,
	}
}

func mapAllToListEmployee(employees []*Employee) []*listing.Employee {
	mapped := make([]*listing.Employee, 0)
	for _, value := range employees {
		mapped = append(mapped, mapToListEmployee(value))
	}
	return mapped
}
