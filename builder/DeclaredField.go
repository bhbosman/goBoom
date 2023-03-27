package builder

import (
	"go/ast"
	"go/token"
	"reflect"
)

type DeclaredField struct {
	Field
	name          string
	declaredType  ast.Expr
	reflectedType reflect.Type
}

func (df *DeclaredField) ReflectedType() reflect.Type {
	return df.reflectedType
}

func (df *DeclaredField) DeclaredTypeExpression() ast.Expr {
	return df.declaredType
}

func (df *DeclaredField) Name() string {
	return df.name
}

func NewDeclaredField(indent int, position token.Position, pos token.Pos, end token.Pos, name string, declaredType ast.Expr) *DeclaredField {
	return &DeclaredField{
		Field:        NewField(indent, position, pos, end),
		name:         name,
		declaredType: declaredType,
	}
}
