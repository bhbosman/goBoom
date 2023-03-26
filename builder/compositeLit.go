package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type CompositeLit struct {
	Location
	compositeLit *ast.CompositeLit
}

func (self *CompositeLit) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *CompositeLit) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *CompositeLit) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *CompositeLit) AssignExpression(expression ast.Expr) {
}

func NewCompositeLit(indent int, position token.Position, pos token.Pos, end token.Pos, compositeLit *ast.CompositeLit) *CompositeLit {
	return &CompositeLit{
		Location:     NewLocation(indent, position, pos, end),
		compositeLit: compositeLit,
	}
}
