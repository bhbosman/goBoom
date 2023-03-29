package definedNode

import (
	"fmt"
	"go/token"
	"reflect"
)

type File struct {
	Location
	expression IDefinedNode
}

func (self *File) DetermineValue(container IContainer) reflect.Value {
	//TODO implement me
	panic("implement me")
}

func (self *File) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
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
	return self.expression.(*Ident).AstIdent.Name
}

func (self *File) AssignExpression(node IDefinedNode) {
	self.expression = node
}
