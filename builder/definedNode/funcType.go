package definedNode

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type FuncType struct {
	Location
	ident       string
	hasParams   bool
	hasResult   bool
	paramCount  int
	resultCount int
	expressions []IDefinedNode
	funcType    *ast.FuncType
}

func (self *FuncType) AssignExpression(node IDefinedNode) {
	self.expressions = append(self.expressions, node)
}

func (self *FuncType) DetermineValue(container IContainer) reflect.Value {
	//TODO implement me
	panic("implement me")
}

func (self *FuncType) DetermineType(container IContainer) reflect.Type {
	var in, out []reflect.Type
	for _, expression := range self.expressions {
		switch v := expression.(type) {
		case *MultiFieldDeclaration:
			addToDirection := func(direction MultiFieldDeclarationDirection) {
				switch direction {
				case MfdInDirection:
					in = append(in, v.declaredType.DetermineType(container))
					break
				case MfdOutDirection:
					out = append(out, v.declaredType.DetermineType(container))
					break
				}
			}
			if len(v.idents) > 0 {
				for range v.idents {
					addToDirection(v.direction)
				}
			} else {
				addToDirection(v.direction)
			}
		}
	}

	return reflect.FuncOf(in, out, false)
}

func (self *FuncType) Validate(container IContainer) {
	for _, expression := range self.expressions {
		expression.Validate(container)
	}
}

func (self *FuncType) Start(IContainer) {
	self.Print(fmt.Sprintf("Start %v", reflect.TypeOf(self).String()))
}

func (self *FuncType) Complete(IContainer) {
	self.Print(fmt.Sprintf("Complete %v", reflect.TypeOf(self).String()))
	switch {
	case !self.hasResult && !self.hasParams:
		if len(self.expressions) != 0 {
			panic("needs no fields when there in no params and result")
		}
	case !self.hasResult && self.hasParams:
		if len(self.expressions) != self.paramCount {
			panic("need fields for params")
		}
		for _, expression := range self.expressions {
			if mfd, ok := expression.(*MultiFieldDeclaration); ok {
				mfd.direction = MfdInDirection
			}
		}

	case self.hasResult && !self.hasParams:
		if len(self.expressions) != self.resultCount {
			panic("need fields for params")
		}
		for _, expression := range self.expressions {
			if mfd, ok := expression.(*MultiFieldDeclaration); ok {
				mfd.direction = MfdOutDirection
			}
		}
	case self.hasResult && self.hasParams:
		if len(self.expressions) != self.resultCount+self.paramCount {
			panic("counts does not match up")

		}
		for i, expression := range self.expressions {
			switch {
			case 0 <= i && i < self.paramCount:
				if mfd, ok := expression.(*MultiFieldDeclaration); ok {
					mfd.direction = MfdInDirection
				}
			case self.paramCount <= i && i < self.paramCount+self.resultCount:
				if mfd, ok := expression.(*MultiFieldDeclaration); ok {
					mfd.direction = MfdOutDirection
				}
			}

		}
	}
}

func NewFuncType(indent int, position token.Position, pos token.Pos, end token.Pos, ident string, funcType *ast.FuncType) *FuncType {
	result := &FuncType{
		Location:  NewLocation(indent, position, pos, end),
		ident:     ident,
		hasParams: funcType.Params != nil && len(funcType.Params.List) > 0,
		hasResult: funcType.Results != nil && len(funcType.Results.List) > 0,
		paramCount: func() int {
			if funcType.Params != nil {
				return len(funcType.Params.List)
			}
			return 0
		}(),
		resultCount: func() int {
			if funcType.Results != nil {
				return len(funcType.Results.List)
			}
			return 0
		}(),
		funcType: funcType,
	}
	return result
}
