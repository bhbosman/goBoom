package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type FuncLit struct {
	Location
	funcType *FuncType
	funcLit  *ast.FuncLit
}

func (self *FuncLit) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *FuncLit) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *FuncLit) SetBlockStatement(stmt *BlockStmt) {
}

func (self *FuncLit) SetFuncType(funcType *FuncType) {
	self.funcType = funcType
}

func (self *FuncLit) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func NewFuncLit(indent int, position token.Position, pos token.Pos, end token.Pos, funcLit *ast.FuncLit) *FuncLit {
	result := &FuncLit{
		Location: NewLocation(indent, position, pos, end),
		funcLit:  funcLit,
	}
	return result
}
