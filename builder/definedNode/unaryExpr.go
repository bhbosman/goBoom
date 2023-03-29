package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type UnaryExpr struct {
	Location
	Expressions []IDefinedNode
	UnaryExpr   *ast.UnaryExpr
}

func (self *UnaryExpr) AssignExpression(node IDefinedNode) {
	self.Expressions = append(self.Expressions, node)
}

func (self *UnaryExpr) DetermineValue(container IContainer) reflect.Value {
	//TODO implement me
	panic("implement me")
}

func (self *UnaryExpr) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *UnaryExpr) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *UnaryExpr) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *UnaryExpr) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func NewUnaryExpr(indent int, position token.Position, pos token.Pos, end token.Pos, unaryExpr *ast.UnaryExpr) *UnaryExpr {
	return &UnaryExpr{
		Location:  NewLocation(indent, position, pos, end),
		UnaryExpr: unaryExpr,
	}
}
