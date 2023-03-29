package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type ValueSpec struct {
	Location
	Expressions []IDefinedNode
	valueSpec   *ast.ValueSpec
}

func (valueSpec *ValueSpec) DetermineValue(container IContainer) reflect.Value {
	//TODO implement me
	panic("implement me")
}

func (valueSpec *ValueSpec) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (valueSpec *ValueSpec) Validate(container IContainer) {
	if len(valueSpec.Expressions)%2 != 0 {
		panic("there should be pairs of ident/expression")
	}
	half := len(valueSpec.Expressions) / 2
	for i := 0; i < half; i++ {
		ident := valueSpec.Expressions[i]
		ident.Validate(container)
		expression := valueSpec.Expressions[i+half]
		expression.Validate(container)

		declaredField := NewDeclaredField(ident, expression)
		declaredField.Validate(container)

		container.AddToScope(declaredField)

	}
}

func (valueSpec *ValueSpec) Start(IContainer) {
	valueSpec.Print(fmt.Sprintf("Start %v", reflect.TypeOf(valueSpec).String()))
}

func (valueSpec *ValueSpec) Complete(container IContainer) {
	valueSpec.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(valueSpec).String()))
	container.AddValueSpec(valueSpec)
}

func (valueSpec *ValueSpec) AssignExpression(node IDefinedNode) {
	valueSpec.Expressions = append(valueSpec.Expressions, node)
}

func NewValueSpec(indent int, position token.Position, pos token.Pos, end token.Pos, valueSpec *ast.ValueSpec) *ValueSpec {
	return &ValueSpec{
		Location:  NewLocation(indent, position, pos, end),
		valueSpec: valueSpec,
	}
}
