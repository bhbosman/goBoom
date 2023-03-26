package builder

import (
	"go/ast"
	"go/token"
)

type DeclaredField struct {
	Field
	name         string
	declaredType ast.Expr
}

func NewDeclaredField(indent int, position token.Position, pos token.Pos, end token.Pos, name string, declaredType ast.Expr) *DeclaredField {
	return &DeclaredField{
		Field:        NewField(indent, position, pos, end),
		name:         name,
		declaredType: declaredType,
	}
}
