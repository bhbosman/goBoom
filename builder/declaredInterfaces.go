package builder

import (
	"go/ast"
	"go/token"
	"reflect"
)

type IIdent interface {
	GetIdent() string
}

type IAssignStructType interface {
	AssignStructType(*StructType)
}

type IAssignIdent interface {
	AssignExpression(expression ast.Expr)
}
type Node interface {
	ast.Node
	Indent() int
	Position() token.Position
}

type IField interface {
	Node
	Name() string
	someIFieldDecl()
	DeclaredTypeExpression() ast.Expr
	ReflectedType() reflect.Type
}

type IAddField interface {
	AddField(field IField)
}

type IAssignFuncType interface {
	SetFuncType(*FuncType)
}

type IRemoveNode interface {
	RemoveNode() bool
}

type IAssignBlockStatement interface {
	SetBlockStatement(*BlockStmt)
}

type IContainer interface {
	AddTypeSpec(spec *TypeSpec)
	ValidType(expr ast.Expr) reflect.Type
}

type IDefinedNode interface {
	ast.Node
	Start(IContainer)
	Complete(IContainer)
	Validate(IContainer)
}
