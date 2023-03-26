package builder

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type IncDecStmt struct {
	Location
	incDecStmt *ast.IncDecStmt
}

func (self *IncDecStmt) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *IncDecStmt) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *IncDecStmt) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *IncDecStmt) AssignExpression(expression ast.Expr) {
}

func NewIncDecStmt(indent int, position token.Position, pos token.Pos, end token.Pos, incDecStmt *ast.IncDecStmt) *IncDecStmt {
	return &IncDecStmt{
		Location:   NewLocation(indent, position, pos, end),
		incDecStmt: incDecStmt,
	}
}
