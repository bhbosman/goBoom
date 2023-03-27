package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strings"
)

type MultiFieldDeclaration struct {
	Field
	astField *ast.Field
	// input
	expressions []ast.Expr

	// oncomplete
	idents       []string
	declaredType ast.Expr
}

func (self *MultiFieldDeclaration) ReflectedType() reflect.Type {
	panic("this should never be called, as this type is temp object that will be replaced with DeclaredField")
}

func (self *MultiFieldDeclaration) DeclaredTypeExpression() ast.Expr {
	// if self.declaredType is nil, then the Complete have not executed
	return self.declaredType
}

func (self *MultiFieldDeclaration) Name() string {
	return strings.Join(self.idents, ",")
}

func (self *MultiFieldDeclaration) Validate(container IContainer) {

}

func (self *MultiFieldDeclaration) SetFuncType(funcType *FuncType) {
}

func (self *MultiFieldDeclaration) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *MultiFieldDeclaration) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
	if len(self.expressions) < 2 {
		panic("there should always be a at least one ast.ident, followed by an ast.Expr")
	}
	self.declaredType = self.expressions[len(self.expressions)-1]
	for i, expression := range self.expressions {
		if i == len(self.expressions)-1 {
			break
		}
		if ident, ok := expression.(*ast.Ident); ok {
			self.idents = append(self.idents, ident.Name)
			continue
		}
		panic("the identifiers should be of type *ast.Ident")
	}
}

func (self *MultiFieldDeclaration) AssignExpression(expression ast.Expr) {
	self.expressions = append(self.expressions, expression)
}

func NewMultiFieldDeclaration(indent int, position token.Position, pos token.Pos, end token.Pos, astField *ast.Field) *MultiFieldDeclaration {
	return &MultiFieldDeclaration{
		Field:        NewField(indent, position, pos, end),
		astField:     astField,
		declaredType: nil,
	}
}
