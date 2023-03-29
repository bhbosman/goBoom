package definedNode

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

func (self *KeyValueExpr) DetermineValue(container IContainer) reflect.Value {
	//TODO implement me
	panic("implement me")
}

func (self *KeyValueExpr) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
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

func (self *KeyValueExpr) AssignExpression(IDefinedNode) {
}

func NewKeyValueExpr(indent int, position token.Position, pos token.Pos, end token.Pos, keyValueExpr *ast.KeyValueExpr) *KeyValueExpr {
	return &KeyValueExpr{
		Location:     NewLocation(indent, position, pos, end),
		keyValueExpr: keyValueExpr,
	}
}
