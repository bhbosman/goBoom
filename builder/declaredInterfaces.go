package builder

import (
	"go/ast"
	"go/token"
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
	someIFieldDecl()
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
	ValidType(expr ast.Expr)
}

type IDefinedNode interface {
	ast.Node
	Start(IContainer)
	Complete(IContainer)
	Validate(IContainer)
}
