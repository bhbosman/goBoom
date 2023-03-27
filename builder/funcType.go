package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type FuncType struct {
	Location
	ident     string
	hasParams bool
	hasResult bool
	fields    []IField
	funcType  *ast.FuncType
}

func (self *FuncType) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *FuncType) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *FuncType) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *FuncType) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *FuncType) AddField(field IField) {
	self.fields = append(self.fields, field)
}

func NewFuncType(indent int, position token.Position, pos token.Pos, end token.Pos, ident string, funcType *ast.FuncType) *FuncType {
	result := &FuncType{
		Location:  NewLocation(indent, position, pos, end),
		ident:     ident,
		hasParams: funcType.Params != nil,
		hasResult: funcType.Results != nil,
		funcType:  funcType,
	}
	return result
}
