package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type Ident struct {
	Location
	AstIdent *ast.Ident
}

func (self *Ident) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *Ident) Validate(container IContainer) {
	if self.AstIdent.Name == "" {
		panic("ident must have a name")
	}
}

func (self *Ident) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *Ident) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *Ident) RemoveNode() bool {
	return true
}

func NewIdent(indent int, position token.Position, pos token.Pos, end token.Pos, ident *ast.Ident) *Ident {
	result := &Ident{
		Location: NewLocation(indent, position, pos, end),
		AstIdent: ident,
	}
	return result
}
