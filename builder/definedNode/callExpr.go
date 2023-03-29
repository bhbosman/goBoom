package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type CallExpr struct {
	Location
	expressions []IDefinedNode
	callExpr    *ast.CallExpr
}

func (self *CallExpr) DetermineValue(container IContainer) reflect.Value {
	switch {
	case
		func() bool {
			if len(self.expressions) == 2 {
				_, ok01 := self.expressions[0].(*Ident)
				_, ok02 := self.expressions[1].(*BasicLit)
				return ok01 && ok02
			}
			return false
		}():
		{
			bl, _ := self.expressions[1].(*BasicLit)
			return bl.DetermineValue(container)
		}
	}
	//TODO implement me
	panic("implement me")
}

func (self *CallExpr) DetermineType(container IContainer) reflect.Type {
	switch {
	case
		func() bool {
			if len(self.expressions) == 2 {
				_, ok01 := self.expressions[0].(*Ident)
				_, ok02 := self.expressions[1].(*BasicLit)
				return ok01 && ok02
			}
			return false
		}():
		{
			bl, _ := self.expressions[1].(*BasicLit)
			return bl.DetermineType(container)
		}
	}
	//TODO implement me
	panic("implement me")
}

func (self *CallExpr) Validate(container IContainer) {
	for _, expression := range self.expressions {
		expression.Validate(container)
	}
}

func (self *CallExpr) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *CallExpr) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *CallExpr) AssignExpression(node IDefinedNode) {
	self.expressions = append(self.expressions, node)
}

func NewCallExpr(indent int, position token.Position, pos token.Pos, end token.Pos, callExpr *ast.CallExpr) *CallExpr {
	return &CallExpr{
		Location: NewLocation(indent, position, pos, end),
		callExpr: callExpr,
	}
}
