package test

// Declaration file for tests
import "embed"

//go:embed *.go
var content embed.FS
var namespace = "github.com/bhbosman/goBoom/model"

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
