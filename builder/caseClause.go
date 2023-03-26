package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type CaseClause struct {
	Location
	caseClause *ast.CaseClause
}

func (self *CaseClause) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *CaseClause) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *CaseClause) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *CaseClause) AssignExpression(expression ast.Expr) {
}

func NewCaseClause(indent int, position token.Position, pos token.Pos, end token.Pos, caseClause *ast.CaseClause) *CaseClause {
	return &CaseClause{
		Location:   NewLocation(indent, position, pos, end),
		caseClause: caseClause,
	}
}