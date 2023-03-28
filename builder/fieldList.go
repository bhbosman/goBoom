package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type FieldList struct {
	Location
	fields    []*MultiFieldDeclaration
	fieldList *ast.FieldList
}

func (self *FieldList) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *FieldList) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *FieldList) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *FieldList) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *FieldList) RemoveNode() bool {
	return true
}

func NewFieldList(indent int, position token.Position, pos token.Pos, end token.Pos, fieldList *ast.FieldList) *FieldList {
	result := &FieldList{
		Location:  NewLocation(indent, position, pos, end),
		fieldList: fieldList,
	}
	return result
}
