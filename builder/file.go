package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type File struct {
	Location
	expression ast.Expr
}

func (self *File) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *File) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *File) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func NewFile(indent int, position token.Position, pos token.Pos, end token.Pos) *File {
	return &File{
		Location: NewLocation(indent, position, pos, end),
	}
}

func (self *File) GetIdent() string {
	return self.expression.(*ast.Ident).Name
}

func (self *File) AssignExpression(expression ast.Expr) {
	self.expression = expression
}
