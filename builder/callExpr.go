package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type CallExpr struct {
	Location
	callExpr *ast.CallExpr
}

func (self *CallExpr) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *CallExpr) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *CallExpr) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *CallExpr) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *CallExpr) AssignExpression(IDefinedNode) {
}

func NewCallExpr(indent int, position token.Position, pos token.Pos, end token.Pos, callExpr *ast.CallExpr) *CallExpr {
	return &CallExpr{
		Location: NewLocation(indent, position, pos, end),
		callExpr: callExpr,
	}
}
