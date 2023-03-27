package builder

import (
	"go/token"
	"reflect"
)

type DeclaredField struct {
	name          IDefinedNode
	declaredType  IDefinedNode
	reflectedType reflect.Type
}

func (self *DeclaredField) Pos() token.Pos {
	return self.name.Pos()
}

func (self *DeclaredField) End() token.Pos {
	return self.name.End()
}

func (self *DeclaredField) Indent() int {
	return self.name.Indent()
}

func (self *DeclaredField) Position() token.Position {
	return self.name.Position()
}

func (self *DeclaredField) someIFieldDecl() {
	// should never be called
	panic("should never be called")
}

func (df *DeclaredField) DetermineType(container IContainer) reflect.Type {
	return df.reflectedType
}

func (df *DeclaredField) Start(container IContainer) {
}

func (df *DeclaredField) Complete(container IContainer) {
}

func (df *DeclaredField) Validate(container IContainer) {
	df.name.Validate(container)
	df.declaredType.Validate(container)
	df.reflectedType = container.ValidType(df.declaredType)
}

func (df *DeclaredField) ReflectedType() reflect.Type {
	return df.reflectedType
}

func (df *DeclaredField) Name() string {
	return df.name.(*Ident).AstIdent.Name
}

func NewDeclaredField(name IDefinedNode, declaredType IDefinedNode) *DeclaredField {
	return &DeclaredField{
		name:         name,
		declaredType: declaredType,
	}
}
