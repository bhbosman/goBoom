package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type AssignStmt struct {
	Location
	expressions []IDefinedNode
	assignStmt  *ast.AssignStmt
}

func (self *AssignStmt) DetermineValue(container IContainer) reflect.Value {
	//TODO implement me
	panic("implement me")
}

func (self *AssignStmt) DetermineType(container IContainer) reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (self *AssignStmt) Validate(container IContainer) {
	//TODO implement me
	panic("implement me")
}

func (self *AssignStmt) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *AssignStmt) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
}

func (self *AssignStmt) AssignExpression(node IDefinedNode) {
	self.expressions = append(self.expressions, node)
}

func NewAssignStmt(indent int, position token.Position, pos token.Pos, end token.Pos, assignStmt *ast.AssignStmt) *AssignStmt {
	return &AssignStmt{
		Location:   NewLocation(indent, position, pos, end),
		assignStmt: assignStmt,
	}
}
