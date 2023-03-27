package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type IfStmt struct {
	Location
	ifStmt *ast.IfStmt
}

func (self *IfStmt) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *IfStmt) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *IfStmt) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *IfStmt) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *IfStmt) AssignExpression(IDefinedNode) {
}

func (self *IfStmt) SetBlockStatement(stmt *BlockStmt) {
}

func NewIfStmt(indent int, position token.Position, pos token.Pos, end token.Pos, ifStmt *ast.IfStmt) *IfStmt {
	return &IfStmt{
		Location: NewLocation(indent, position, pos, end),
		ifStmt:   ifStmt,
	}
}
