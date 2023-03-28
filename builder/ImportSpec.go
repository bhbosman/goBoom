package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type ImportSpec struct {
	Location
	importSpec *ast.ImportSpec
}

func (self *ImportSpec) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *ImportSpec) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *ImportSpec) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *ImportSpec) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func NewImportSpec(indent int, position token.Position, pos token.Pos, end token.Pos, importSpec *ast.ImportSpec) *ImportSpec {
	return &ImportSpec{
		Location:   NewLocation(indent, position, pos, end),
		importSpec: importSpec,
	}
}
