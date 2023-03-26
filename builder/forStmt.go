package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type ForStmt struct {
	Location
	forStmt *ast.ForStmt
}

func (self *ForStmt) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *ForStmt) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *ForStmt) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *ForStmt) SetBlockStatement(stmt *BlockStmt) {
}

func (self *ForStmt) AssignExpression(expression ast.Expr) {
}

func NewForStmt(indent int, position token.Position, pos token.Pos, end token.Pos, forStmt *ast.ForStmt) *ForStmt {
	return &ForStmt{
		Location: NewLocation(indent, position, pos, end),
		forStmt:  forStmt}
}
