package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type InterfaceType struct {
	Location
	interfaceType *ast.InterfaceType
	fields        []IField
}

func (self *InterfaceType) AssignExpression(node IDefinedNode) {
}

func (self *InterfaceType) DetermineValue(container IContainer) reflect.Value {
	//TODO implement me
	panic("implement me")
}

func (self *InterfaceType) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *InterfaceType) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

//func (self *InterfaceType) SetFuncType(funcType *FuncType) {
//
//}

//func (self *InterfaceType) AddField(field IField) {
//	self.fields = append(self.fields, field)
//}

func (self *InterfaceType) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *InterfaceType) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func NewInterfaceType(indent int, position token.Position, pos token.Pos, end token.Pos, interfaceType *ast.InterfaceType) *InterfaceType {
	return &InterfaceType{
		Location:      NewLocation(indent, position, pos, end),
		interfaceType: interfaceType,
	}
}
