package builder

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

func (self *TypeAssertExpr) AssignExpression(expression ast.Expr) {
}

func NewTypeAssertExpr(indent int, position token.Position, pos token.Pos, end token.Pos, typeAssertExpr *ast.TypeAssertExpr) *TypeAssertExpr {
	return &TypeAssertExpr{
		Location:       NewLocation(indent, position, pos, end),
		typeAssertExpr: typeAssertExpr,
	}
}
