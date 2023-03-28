package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type IndexExpr struct {
	Location
	indexExpr *ast.IndexExpr
}

func (self *IndexExpr) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *IndexExpr) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *IndexExpr) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *IndexExpr) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *IndexExpr) AssignExpression(IDefinedNode) {

}

func (self *IndexExpr) GetIdent() string {
	panic("dddd")
}

func NewIndexExpr(indent int, position token.Position, pos token.Pos, end token.Pos, indexExpr *ast.IndexExpr) *IndexExpr {
	return &IndexExpr{
		Location:  NewLocation(indent, position, pos, end),
		indexExpr: indexExpr,
	}
}
