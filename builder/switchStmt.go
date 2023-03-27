package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type SwitchStmt struct {
	Location
	switchStmt *ast.SwitchStmt
}

func (self *SwitchStmt) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *SwitchStmt) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *SwitchStmt) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *SwitchStmt) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *SwitchStmt) SetBlockStatement(stmt *BlockStmt) {
}

func (self *SwitchStmt) AssignExpression(IDefinedNode) {
}

func NewSwitchStmt(indent int, position token.Position, pos token.Pos, end token.Pos, switchStmt *ast.SwitchStmt) *SwitchStmt {
	return &SwitchStmt{
		Location:   NewLocation(indent, position, pos, end),
		switchStmt: switchStmt,
	}
}
