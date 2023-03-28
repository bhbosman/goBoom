package builder

import (
	"go/ast"
	"go/token"
	"reflect"
)

type IIdent interface {
	GetIdent() string
}
type Node interface {
	ast.Node
	Indent() int
	Position() token.Position
}

type IDefinedNode interface {
	Node
	Start(IContainer)
	Complete(IContainer)
	Validate(IContainer)
	DetermineType(container IContainer) reflect.Type
}

type IAssignIdent interface {
	AssignExpression(node IDefinedNode)
}

type IField interface {
	IDefinedNode
	Name() string
	someIFieldDecl()
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
	AddTypeSpec(*TypeSpec)
	AddValueSpec(*ValueSpec)
	ValidType(expr IDefinedNode) reflect.Type
	ValidTypeFromKind(kind token.Token) reflect.Type
}
