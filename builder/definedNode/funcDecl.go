package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type FuncDecl struct {
	Location
	Expressions []IDefinedNode

	name     *Ident
	funcType *FuncType

	funcDecl *ast.FuncDecl
}

func (self *FuncDecl) Name() string {
	return self.name.AstIdent.Name
}

func (self *FuncDecl) DetermineValue(container IContainer) reflect.Value {
	//TODO implement me
	panic("implement me")
}

func (self *FuncDecl) DetermineType(container IContainer) reflect.Type {
	return self.funcType.DetermineType(container)
}

func (self *FuncDecl) Validate(container IContainer) {
	self.name.Validate(container)
	self.funcType.Validate(container)
	container.AddToScope(self)
}

func (self *FuncDecl) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *FuncDecl) SetBlockStatement(stmt *BlockStmt) {
}

func (self *FuncDecl) Complete(container IContainer) {
	if len(self.Expressions) != 2 {
		panic("need name and func type")
	}
	if param, ok := self.Expressions[0].(*Ident); ok {
		self.name = param
	} else {
		panic("require *Ident")
	}
	if param, ok := self.Expressions[1].(*FuncType); ok {
		self.funcType = param
	} else {
		panic("require *FuncType")
	}
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
	container.AddFuncDecl(self)
}

func (self *FuncDecl) GetIdent() string {
	return self.Expressions[0].(*Ident).AstIdent.Name
}

func (self *FuncDecl) AssignExpression(expression IDefinedNode) {
	self.Expressions = append(self.Expressions, expression)
}

func NewFuncDecl(indent int, position token.Position, pos token.Pos, end token.Pos, funcDecl *ast.FuncDecl,
) *FuncDecl {
	result := &FuncDecl{
		Location: NewLocation(indent, position, pos, end),
		funcDecl: funcDecl,
	}
	return result
}
