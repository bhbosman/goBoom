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

func (self *BasicLit) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
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
