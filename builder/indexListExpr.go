package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type IndexListExpr struct {
	Location
	indexListExpr *ast.IndexListExpr
}

func (self *IndexListExpr) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *IndexListExpr) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *IndexListExpr) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *IndexListExpr) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *IndexListExpr) AssignExpression(IDefinedNode) {

}

func NewIndexListExpr(indent int, position token.Position, pos token.Pos, end token.Pos, indexListExpr *ast.IndexListExpr) *IndexListExpr {
	return &IndexListExpr{
		Location:      NewLocation(indent, position, pos, end),
		indexListExpr: indexListExpr,
	}
}
