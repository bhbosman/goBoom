package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type TypeSwitchStmt struct {
	Location
	typeSwitchStmt *ast.TypeSwitchStmt
}

func (self *TypeSwitchStmt) DetermineValue(container IContainer) reflect.Value {
	//TODO implement me
	panic("implement me")
}

func (self *TypeSwitchStmt) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *TypeSwitchStmt) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *TypeSwitchStmt) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *TypeSwitchStmt) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *TypeSwitchStmt) SetBlockStatement(stmt *BlockStmt) {
}

func NewTypeSwitchStmt(indent int, position token.Position, pos token.Pos, end token.Pos, typeSwitchStmt *ast.TypeSwitchStmt) *TypeSwitchStmt {
	return &TypeSwitchStmt{
		Location:       NewLocation(indent, position, pos, end),
		typeSwitchStmt: typeSwitchStmt,
	}
}
