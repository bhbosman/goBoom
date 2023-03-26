package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type ValueSpec struct {
	Location
	expressions []ast.Expr
	valueSpec   *ast.ValueSpec
}

func (self *ValueSpec) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *ValueSpec) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *ValueSpec) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

//func (self *ValueSpec) GetIdent() string {
//	panic("figure out")
//	//return self.ident[len(self.ident)-1]
//}

func (self *ValueSpec) AssignExpression(expression ast.Expr) {
	self.expressions = append(self.expressions, expression)
}

func NewValueSpec(indent int, position token.Position, pos token.Pos, end token.Pos, valueSpec *ast.ValueSpec) *ValueSpec {
	return &ValueSpec{
		Location:  NewLocation(indent, position, pos, end),
		valueSpec: valueSpec,
	}
}
