package definedNode

import (
	"go/token"
	"reflect"
)

type Operation struct {
	declaredField DeclaredField
	expression    IDefinedNode
}

func (self *Operation) Pos() token.Pos {
	return self.declaredField.Pos()
}

func (self *Operation) End() token.Pos {
	return self.declaredField.End()
}

func (self *Operation) Indent() int {
	return self.declaredField.Indent()
}

func (self *Operation) Position() token.Position {
	return self.declaredField.Position()
}

func (self *Operation) Start(container IContainer) {
}

func (self *Operation) Complete(container IContainer) {
}

func (self *Operation) Validate(container IContainer) {
	self.declaredField.Validate(container)
	self.expression.Validate(container)
}

func (self *Operation) DetermineType(container IContainer) reflect.Type {
	return self.declaredField.DetermineType(container)
}
