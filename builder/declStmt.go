package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type DeclStmt struct {
	Location
	DeclStmt *ast.DeclStmt
}

func (self *DeclStmt) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *DeclStmt) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *DeclStmt) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *DeclStmt) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func NewDeclStmt(indent int, position token.Position, pos token.Pos, end token.Pos, declStmt *ast.DeclStmt) *DeclStmt {
	return &DeclStmt{
		Location: NewLocation(indent, position, pos, end),
		DeclStmt: declStmt,
	}
}
