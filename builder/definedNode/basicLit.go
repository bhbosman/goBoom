package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
)

type BasicLit struct {
	Location
	BasicLit *ast.BasicLit
}

func (self *BasicLit) DetermineValue(container IContainer) reflect.Value {
	switch self.BasicLit.Kind {
	case token.INT:
		parseInt, err := strconv.ParseInt(self.BasicLit.Value, 0, 64)
		if err != nil {
			panic(err)
		}
		return reflect.ValueOf(int(parseInt))
	case token.FLOAT:
	case token.IMAG:
	case token.CHAR:
	case token.STRING:
	}

	//TODO implement me
	panic("implement me")
}

func (self *BasicLit) DetermineType(container IContainer) reflect.Type {
	return container.ValidTypeFromKind(self.BasicLit.Kind)
}

func (self *BasicLit) Kind() token.Token { // token.INT, token.FLOAT, token.IMAG, token.CHAR, or token.STRING
	return self.BasicLit.Kind
}

func (self *BasicLit) Value() string { // literal string; e.g. 42, 0x7f, 3.14, 1e-9, 2.4i, 'a', '\x7f', "foo" or `\m\n\o`
	return self.BasicLit.Value
}

func (self *BasicLit) Validate(container IContainer) {
	_ = container.ValidTypeFromKind(self.BasicLit.Kind)
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
		BasicLit: basicLit,
	}
}
