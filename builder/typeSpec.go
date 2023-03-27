package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type TypeSpec struct {
	Location
	DefinedNode []IDefinedNode
	typeSpec    *ast.TypeSpec
}

func (self *TypeSpec) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *TypeSpec) Validate(container IContainer) {
	if len(self.DefinedNode) != 2 {
		panic("need two elements")
	}
	if self.DefinedNode[0].(*Ident).AstIdent.Name == "" {
		panic("there should be a name")
	}
	self.DefinedNode[1].Validate(container)
}

func (self *TypeSpec) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *TypeSpec) Complete(container IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
	container.AddTypeSpec(self)
}

func (self *TypeSpec) GetIdent() string {
	return self.DefinedNode[0].(*Ident).AstIdent.Name
}

func (self *TypeSpec) AssignExpression(node IDefinedNode) {
	self.DefinedNode = append(self.DefinedNode, node)
}

func NewTypeSpec(indent int, position token.Position, pos token.Pos, end token.Pos, typeSpec *ast.TypeSpec) *TypeSpec {
	result := &TypeSpec{
		Location: NewLocation(indent, position, pos, end),
		typeSpec: typeSpec,
	}
	return result
}
