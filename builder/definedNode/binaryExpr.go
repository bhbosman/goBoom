package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type BinaryExpr struct {
	Location
	expressions []IDefinedNode
	binaryExpr  *ast.BinaryExpr
}

func (self *BinaryExpr) DetermineValue(container IContainer) reflect.Value {
	valueA := self.expressions[0].DetermineValue(container)
	valueB := self.expressions[1].DetermineValue(container)
	switch {
	case
		func() bool {
			return valueA.Type() == reflect.TypeOf(int(0))
		}():
		switch self.binaryExpr.Op {
		case token.ADD:
			v := int(valueA.Int()) + int(valueB.Int())
			return reflect.ValueOf(v)
		case token.SUB: // -
			v := int(valueA.Int()) - int(valueB.Int())
			return reflect.ValueOf(v)
		case token.MUL: // *
			v := int(valueA.Int()) * int(valueB.Int())
			return reflect.ValueOf(v)
		case token.QUO: // /
			v := int(valueA.Int()) / int(valueB.Int())
			return reflect.ValueOf(v)
		case token.REM: // %
			v := int(valueA.Int()) % int(valueB.Int())
			return reflect.ValueOf(v)
		case token.AND: // &
			v := int(valueA.Int()) & int(valueB.Int())
			return reflect.ValueOf(v)
		case token.OR: // |
			v := int(valueA.Int()) | int(valueB.Int())
			return reflect.ValueOf(v)
		case token.XOR: // ^
			v := int(valueA.Int()) ^ int(valueB.Int())
			return reflect.ValueOf(v)
		case token.SHL: // <<
			v := int(valueA.Int()) << int(valueB.Int())
			return reflect.ValueOf(v)
		case token.SHR: // >>
			v := int(valueA.Int()) >> int(valueB.Int())
			return reflect.ValueOf(v)
		case token.AND_NOT: // &^
			v := int(valueA.Int()) &^ int(valueB.Int())
			return reflect.ValueOf(v)

		default:
		}
	}
	panic("implement")

}

func (self *BinaryExpr) DetermineType(container IContainer) reflect.Type {
	type00 := self.expressions[0].DetermineType(container)
	type01 := self.expressions[1].DetermineType(container)
	if type00 != type01 {
		panic("types must be the same")
	}
	return type00

}

func (self *BinaryExpr) Validate(container IContainer) {
	if len(self.expressions) != 2 {
		panic("binary expression need 2 expresssions")
	}
	for _, expression := range self.expressions {
		expression.Validate(container)
	}
}

func (self *BinaryExpr) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *BinaryExpr) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *BinaryExpr) AssignExpression(node IDefinedNode) {
	self.expressions = append(self.expressions, node)
}

func NewBinaryExpr(indent int, position token.Position, pos token.Pos, end token.Pos, binaryExpr *ast.BinaryExpr) *BinaryExpr {
	return &BinaryExpr{
		Location:   NewLocation(indent, position, pos, end),
		binaryExpr: binaryExpr,
	}
}
