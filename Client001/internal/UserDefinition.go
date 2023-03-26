package internal

type UserDefinition struct {
	Id      int
	Name    string
	Surname string
}

type UserAddress struct {
	Id               int
	UserDefinitionId int
	Line01           string
	Line02           string
	Line03           string
	Line04           string
	Line05           string
	Province         string
	Country          string
	PostalCode       string
}

type UserDefinitionView struct {
	Id         int
	Name       string
	Surname    string
	Line01     string
	Line02     string
	Line03     string
	Line04     string
	Line05     string
	Province   string
	Country    string
	PostalCode string
}
