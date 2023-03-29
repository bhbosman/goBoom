package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type ArrayType struct {
	Location
	arrayType *ast.ArrayType
}

func (self *ArrayType) DetermineValue(container IContainer) reflect.Value {
	//TODO implement me
	panic("implement me")
}

func (self *ArrayType) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *ArrayType) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *ArrayType) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *ArrayType) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *ArrayType) AssignExpression(IDefinedNode) {
}

func NewArrayType(indent int, position token.Position, pos token.Pos, end token.Pos, arrayType *ast.ArrayType) *ArrayType {
	return &ArrayType{
		Location:  NewLocation(indent, position, pos, end),
		arrayType: arrayType,
	}
}
