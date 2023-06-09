package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type TypeAssertExpr struct {
	Location
	typeAssertExpr *ast.TypeAssertExpr
}

func (self *TypeAssertExpr) DetermineValue(container IContainer) reflect.Value {
	//TODO implement me
	panic("implement me")
}

func (self *TypeAssertExpr) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *TypeAssertExpr) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *TypeAssertExpr) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *TypeAssertExpr) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *TypeAssertExpr) AssignExpression(IDefinedNode) {
}

func NewTypeAssertExpr(indent int, position token.Position, pos token.Pos, end token.Pos, typeAssertExpr *ast.TypeAssertExpr) *TypeAssertExpr {
	return &TypeAssertExpr{
		Location:       NewLocation(indent, position, pos, end),
		typeAssertExpr: typeAssertExpr,
	}
}
