package definedNode

import (
	"go/ast"
	"go/token"
)

type DeclaredPosition struct {
	FileName string
	Pos      token.Pos
}

type EntityField struct {
	Name string
	Type string
	Expr ast.Expr
}

type Entity struct {
	Name             string
	DeclaredPosition DeclaredPosition
	EntityField      []EntityField
}
