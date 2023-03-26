package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type TypeSpec struct {
	Location
	name       string
	structType *StructType
	typeSpec   *ast.TypeSpec
}

func (self *TypeSpec) AssignStructType(structType *StructType) {
	self.structType = structType
}

func (self *TypeSpec) Validate(container IContainer) {
	if self.name == "" {
		panic("there should be a name")
	}
	self.structType.Validate(container)
}

func (self *TypeSpec) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *TypeSpec) Complete(container IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
	container.AddTypeSpec(self)
}

func (self *TypeSpec) GetIdent() string {
	return self.name
}

func (self *TypeSpec) AssignExpression(expression ast.Expr) {
	self.name = expression.(*ast.Ident).Name
}

func NewTypeSpec(indent int, position token.Position, pos token.Pos, end token.Pos, typeSpec *ast.TypeSpec) *TypeSpec {
	result := &TypeSpec{
		Location: NewLocation(indent, position, pos, end),
		typeSpec: typeSpec,
	}
	return result
}
