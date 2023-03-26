package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type ExprStmt struct {
	Location
	ExprStmt *ast.ExprStmt
}

func (self *ExprStmt) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *ExprStmt) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *ExprStmt) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func NewExprStmt(indent int, position token.Position, pos token.Pos, end token.Pos, exprStmt *ast.ExprStmt) *ExprStmt {
	return &ExprStmt{
		Location: NewLocation(indent, position, pos, end),
		ExprStmt: exprStmt}
}
