package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type SelectorExpr struct {
	Location
	expression   ast.Expr
	selectorExpr *ast.SelectorExpr
}

func (self *SelectorExpr) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *SelectorExpr) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *SelectorExpr) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *SelectorExpr) GetIdent() string {
	return self.expression.(*ast.Ident).Name
}

func (self *SelectorExpr) AssignExpression(expression ast.Expr) {
	self.expression = expression
}

func NewSelectorExpr(indent int, position token.Position, pos token.Pos, end token.Pos, selectorExpr *ast.SelectorExpr) *SelectorExpr {
	return &SelectorExpr{
		Location:     NewLocation(indent, position, pos, end),
		selectorExpr: selectorExpr,
	}
}
