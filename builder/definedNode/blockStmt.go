package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type BlockStmt struct {
	Location
	blockStmt *ast.BlockStmt
}

func (self *BlockStmt) DetermineValue(container IContainer) reflect.Value {
	//TODO implement me
	panic("implement me")
}

func (self *BlockStmt) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *BlockStmt) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *BlockStmt) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *BlockStmt) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *BlockStmt) RemoveNode() bool {
	return true
}

func NewBlockStmt(indent int, position token.Position, pos token.Pos, end token.Pos, blockStmt *ast.BlockStmt) *BlockStmt {
	return &BlockStmt{
		Location:  NewLocation(indent, position, pos, end),
		blockStmt: blockStmt,
	}
}
