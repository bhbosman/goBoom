package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type BranchStmt struct {
	Location
	BranchStmt *ast.BranchStmt
}

func (self *BranchStmt) DetermineValue(container IContainer) reflect.Value {
	//TODO implement me
	panic("implement me")
}

func (self *BranchStmt) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *BranchStmt) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *BranchStmt) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *BranchStmt) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func NewBranchStmt(indent int, position token.Position, pos token.Pos, end token.Pos, branchStmt *ast.BranchStmt) *BranchStmt {
	return &BranchStmt{
		Location:   NewLocation(indent, position, pos, end),
		BranchStmt: branchStmt,
	}
}
