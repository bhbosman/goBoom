package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type BasicLit struct {
	Location
	basicLit *ast.BasicLit
}

func (self *BasicLit) DetermineType(container IContainer) reflect.Type {
	return container.ValidTypeFromKind(self.basicLit.Kind)
}

func (self *BasicLit) Kind() token.Token { // token.INT, token.FLOAT, token.IMAG, token.CHAR, or token.STRING
	return self.basicLit.Kind
}

func (self *BasicLit) Value() string { // literal string; e.g. 42, 0x7f, 3.14, 1e-9, 2.4i, 'a', '\x7f', "foo" or `\m\n\o`
	return self.basicLit.Value
}

func (self *BasicLit) Validate(container IContainer) {
	_ = container.ValidTypeFromKind(self.basicLit.Kind)
}

func (self *BasicLit) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *BasicLit) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func NewBasicLit(indent int, position token.Position, pos token.Pos, end token.Pos, basicLit *ast.BasicLit) *BasicLit {
	return &BasicLit{
		Location: NewLocation(indent, position, pos, end),
		basicLit: basicLit,
	}
}
