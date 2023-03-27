package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type BinaryExpr struct {
	Location
	binaryExpr *ast.BinaryExpr
}

func (self *BinaryExpr) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *BinaryExpr) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *BinaryExpr) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *BinaryExpr) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *BinaryExpr) AssignExpression(IDefinedNode) {
}

func NewBinaryExpr(indent int, position token.Position, pos token.Pos, end token.Pos, binaryExpr *ast.BinaryExpr) *BinaryExpr {
	return &BinaryExpr{
		Location:   NewLocation(indent, position, pos, end),
		binaryExpr: binaryExpr,
	}
}
