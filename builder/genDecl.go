package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type GenDecl struct {
	Location
	genDecl *ast.GenDecl
}

func (self *GenDecl) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *GenDecl) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *GenDecl) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func NewGenDecl(indent int, position token.Position, pos token.Pos, end token.Pos, genDecl *ast.GenDecl) *GenDecl {
	return &GenDecl{
		Location: NewLocation(indent, position, pos, end),
		genDecl:  genDecl,
	}
}
