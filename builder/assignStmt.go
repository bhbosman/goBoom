package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type AssignStmt struct {
	Location
	expressions []ast.Expr
	assignStmt  *ast.AssignStmt
}

func (self *AssignStmt) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *AssignStmt) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *AssignStmt) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *AssignStmt) AssignExpression(expression ast.Expr) {
	self.expressions = append(self.expressions, expression)
}

func NewAssignStmt(indent int, position token.Position, pos token.Pos, end token.Pos, assignStmt *ast.AssignStmt) *AssignStmt {
	return &AssignStmt{
		Location:   NewLocation(indent, position, pos, end),
		assignStmt: assignStmt,
	}
}
