package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type MapType struct {
	Location
	mapType *ast.MapType
}

func (self *MapType) DetermineValue(container IContainer) reflect.Value {
	//TODO implement me
	panic("implement me")
}

func (self *MapType) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *MapType) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *MapType) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *MapType) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *MapType) AssignExpression(IDefinedNode) {
}

func NewMapType(indent int, position token.Position, pos token.Pos, end token.Pos, mapType *ast.MapType) *MapType {
	return &MapType{
		Location: NewLocation(indent, position, pos, end),
		mapType:  mapType,
	}
}
