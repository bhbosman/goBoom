package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type Ident struct {
	Location
	ident *ast.Ident
}

func (self *Ident) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
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
		ident:    ident,
	}
	return result
}
