package listing

type Company struct {
	ID          string
	Name        string
	Address     *Address
	ContactInfo []*ContactInfo
	Employees   []*Employee
}

type Address struct {
	ID       string
	RoadName string
	Number   int32
	ZipCode  int32
}

type ContactInfo struct {
	ID          string
	Title       string
	PhoneNumber int32
	Mail        string
	CompanyID   string
}

type Employee struct {
	ID        string
	Name      string
	CompanyID string
}
