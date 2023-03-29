package definedNode

import (
	"go/ast"
	"go/token"
	"reflect"
)

type ParenExpr struct {
	Location
	Expressions []IDefinedNode
	parenExpr   *ast.ParenExpr
}

func (self *ParenExpr) AssignExpression(node IDefinedNode) {
	self.Expressions = append(self.Expressions, node)
}

func (self *ParenExpr) Start(container IContainer) {
}

func (self *ParenExpr) Complete(container IContainer) {
}

func (self *ParenExpr) Validate(container IContainer) {
	if len(self.Expressions) != 1 {
		panic("count must be one")
	}
	for _, expression := range self.Expressions {
		expression.Validate(container)
	}
}

func (self *ParenExpr) DetermineType(container IContainer) reflect.Type {
	return self.Expressions[0].DetermineType(container)
}

func (self *ParenExpr) DetermineValue(container IContainer) reflect.Value {
	return self.Expressions[0].DetermineValue(container)
}

func NewParenExpr(indent int, position token.Position, pos token.Pos, end token.Pos, parenExpr *ast.ParenExpr) *ParenExpr {
	return &ParenExpr{
		Location:  NewLocation(indent, position, pos, end),
		parenExpr: parenExpr,
	}
}
