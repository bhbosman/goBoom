package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type FuncDecl struct {
	Location
	expression IDefinedNode
	funcType   *FuncType
	funcDecl   *ast.FuncDecl
}

func (self *FuncDecl) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *FuncDecl) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *FuncDecl) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *FuncDecl) SetBlockStatement(stmt *BlockStmt) {
}

func (self *FuncDecl) SetFuncType(funcType *FuncType) {
	self.funcType = funcType
}

func (self *FuncDecl) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *FuncDecl) GetIdent() string {
	return self.expression.(*Ident).AstIdent.Name
}

func (self *FuncDecl) AssignExpression(expression IDefinedNode) {
	self.expression = expression
}

func NewFuncDecl(indent int, position token.Position, pos token.Pos, end token.Pos, funcDecl *ast.FuncDecl,
) *FuncDecl {
	result := &FuncDecl{
		Location: NewLocation(indent, position, pos, end),
		funcDecl: funcDecl,
	}
	return result
}
