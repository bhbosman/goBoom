package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strings"
)

type MultiFieldDeclarationDirection int

const (
	MfdUnknown MultiFieldDeclarationDirection = iota
	MfdInDirection
	MfdOutDirection
)

type MultiFieldDeclaration struct {
	Field
	astField *ast.Field
	// input
	expressions []IDefinedNode
	direction   MultiFieldDeclarationDirection

	// oncomplete of owner as direction is required
	idents       []IDefinedNode
	declaredType IDefinedNode
}

func (self *MultiFieldDeclaration) DetermineValue(container IContainer) reflect.Value {
	//TODO implement me
	panic("implement me")
}

func (self *MultiFieldDeclaration) someIFieldDecl() {
	// should never be called
	panic("should never be called")
}

func (self *MultiFieldDeclaration) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *MultiFieldDeclaration) ReflectedType() reflect.Type {
	panic("this should never be called, as this type is temp object that will be replaced with DeclaredField")
}

func (self *MultiFieldDeclaration) Name() string {
	var ss []string
	for _, s := range self.idents {
		ss = append(ss, s.(*Ident).AstIdent.Name)
	}
	return strings.Join(ss, ",")
}

func (self *MultiFieldDeclaration) Validate(container IContainer) {
	for _, ident := range self.idents {
		ident.Validate(container)
	}
	self.declaredType.Validate(container)
}

func (self *MultiFieldDeclaration) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *MultiFieldDeclaration) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
	if len(self.expressions) == 0 {
		panic("there should at least be one expression")
	}

	self.declaredType = self.expressions[len(self.expressions)-1]
	for i, expression := range self.expressions {
		if i == len(self.expressions)-1 {
			break
		}
		if ident, ok := expression.(*Ident); ok {
			self.idents = append(self.idents, ident)
			continue
		}
		panic("the identifiers should be of type *ast.Ident")
	}
}

func (self *MultiFieldDeclaration) AssignExpression(node IDefinedNode) {
	self.expressions = append(self.expressions, node)
}

func NewMultiFieldDeclaration(indent int, position token.Position, pos token.Pos, end token.Pos, astField *ast.Field) *MultiFieldDeclaration {
	return &MultiFieldDeclaration{
		Field:    NewField(indent, position, pos, end),
		astField: astField,
		//declaredType: nil,
	}
}
