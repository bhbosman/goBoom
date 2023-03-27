package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type RangeStmt struct {
	Location
	rangeStmt *ast.RangeStmt
}

func (self *RangeStmt) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *RangeStmt) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *RangeStmt) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *RangeStmt) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *RangeStmt) SetBlockStatement(stmt *BlockStmt) {
}

func (self *RangeStmt) AssignExpression(IDefinedNode) {
}

func NewRangeStmt(indent int, position token.Position, pos token.Pos, end token.Pos, rangeStmt *ast.RangeStmt) *RangeStmt {
	return &RangeStmt{
		Location:  NewLocation(indent, position, pos, end),
		rangeStmt: rangeStmt,
	}
}
