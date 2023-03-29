package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type StructType struct {
	Location
	Fields     []IField
	structType *ast.StructType
}

func (self *StructType) AssignExpression(node IDefinedNode) {
}

func (self *StructType) DetermineValue(container IContainer) reflect.Value {
	//TODO implement me
	panic("implement me")
}

func (self *StructType) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *StructType) Validate(container IContainer) {
	for _, field := range self.Fields {
		field.Validate(container)
	}

}

func (self *StructType) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *StructType) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
	//var newFields []IField
	//for _, field := range self.Fields {
	//	if multiFieldDeclaration, ok := field.(*MultiFieldDeclaration); ok {
	//		for _, ident := range multiFieldDeclaration.idents {
	//			declaredField := NewDeclaredField(ident, multiFieldDeclaration.declaredType)
	//			newFields = append(newFields, declaredField)
	//		}
	//		continue
	//	}
	//	panic("expect *MultiFieldDeclaration")
	//}
	//self.Fields = newFields

}

//func (self *StructType) AddField(field IField) {
//	self.Fields = append(self.Fields, field)
//}

func NewStructType(indent int, position token.Position, pos token.Pos, end token.Pos, structType *ast.StructType) *StructType {
	result := &StructType{
		Location:   NewLocation(indent, position, pos, end),
		structType: structType,
	}
	return result
}
