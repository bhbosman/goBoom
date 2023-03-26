package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type StructType struct {
	Location
	fields     []IField
	structType *ast.StructType
}

func (self *StructType) Validate(container IContainer) {
	for _, field := range self.fields {
		if declaredField, ok := field.(*DeclaredField); ok {
			container.ValidType(declaredField.declaredType)
		}
	}
}

func (self *StructType) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *StructType) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
	var newFields []IField
	for _, field := range self.fields {
		if multiFieldDeclaration, ok := field.(*MultiFieldDeclaration); ok {
			for _, ident := range multiFieldDeclaration.idents {
				declaredField := NewDeclaredField(field.Indent(), field.Position(), field.Pos(), field.End(), ident, multiFieldDeclaration.declaredType)
				newFields = append(newFields, declaredField)
			}
			continue
		}
		panic("expect *MultiFieldDeclaration")
	}
	self.fields = newFields

}

func (self *StructType) AddField(field IField) {
	self.fields = append(self.fields, field)
}

func NewStructType(indent int, position token.Position, pos token.Pos, end token.Pos, structType *ast.StructType) *StructType {
	result := &StructType{
		Location:   NewLocation(indent, position, pos, end),
		structType: structType,
	}
	return result
}
