package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type ReturnStmt struct {
	Location
	returnStmt *ast.ReturnStmt
}

func (self *ReturnStmt) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *ReturnStmt) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *ReturnStmt) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *ReturnStmt) AssignExpression(expression ast.Expr) {
}

func NewReturnStmt(indent int, position token.Position, pos token.Pos, end token.Pos, returnStmt *ast.ReturnStmt) *ReturnStmt {
	return &ReturnStmt{
		Location:   NewLocation(indent, position, pos, end),
		returnStmt: returnStmt,
	}
}
