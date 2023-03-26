package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type KeyValueExpr struct {
	Location
	keyValueExpr *ast.KeyValueExpr
}

func (self *KeyValueExpr) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *KeyValueExpr) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *KeyValueExpr) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *KeyValueExpr) AssignExpression(expression ast.Expr) {
}

func NewKeyValueExpr(indent int, position token.Position, pos token.Pos, end token.Pos, keyValueExpr *ast.KeyValueExpr) *KeyValueExpr {
	return &KeyValueExpr{
		Location:     NewLocation(indent, position, pos, end),
		keyValueExpr: keyValueExpr,
	}
}
