package creating

type Company struct {
	Name        string
	Address     *Address
	ContactInfo []*ContactInfo
}

type Address struct {
	RoadName string
	Number   int32
	ZipCode  int32
}

type ContactInfo struct {
	Title       string
	PhoneNumber int32
	Mail        string
	CompanyID   string
}

type Employee struct {
	Name      string
	CompanyID string
}
