package builder

import (
	"fmt"
	"github.com/bhbosman/goBoom/builder/definedNode"
	"github.com/bhbosman/goBoom/builder/functions"
	"github.com/emirpasic/gods/stacks/arraystack"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"reflect"
)

type Scope struct {
	Parent *Scope
	Values map[string]definedNode.IScopedItem
}

type AstContextBuilder struct {
	ValidTypes     map[string]reflect.Type
	ValidFunctions map[functions.FuncKey]reflect.Value
	TypeSpecs      []*definedNode.TypeSpec
	ValueSpec      []*definedNode.ValueSpec
	FuncDecl       []*definedNode.FuncDecl
	RootScope      Scope
	CurrentScope   *Scope
}

func (cb *AstContextBuilder) AddFuncDecl(decl *definedNode.FuncDecl) {
	cb.FuncDecl = append(cb.FuncDecl, decl)
}

func (cb *AstContextBuilder) AddToScope(scopeItem definedNode.IScopedItem) {
	cb.CurrentScope.Values[scopeItem.Name()] = scopeItem
}

func (cb *AstContextBuilder) AddValueSpec(valueSpec *definedNode.ValueSpec) {
	cb.ValueSpec = append(cb.ValueSpec, valueSpec)
}

func (cb *AstContextBuilder) ValidTypeFromKind(kind token.Token) reflect.Type {
	switch kind {
	case token.INT:
		return cb.ValidTypes["int"]

	case token.FLOAT:
		return cb.ValidTypes["float64"]
	}
	panic("implement")
}

func (cb *AstContextBuilder) ValidType(expr definedNode.IDefinedNode) reflect.Type {
	switch dt := expr.(type) {
	case *definedNode.Ident:
		if v, ok := cb.ValidTypes[dt.AstIdent.Name]; ok {
			return v
		}
		panic(fmt.Errorf("not supporting type %v", dt.AstIdent.Name))
	case *definedNode.BasicLit:
		return cb.ValidTypeFromKind(dt.BasicLit.Kind)
	case *definedNode.CallExpr:
		return dt.DetermineType(cb)
	case *definedNode.BinaryExpr:
		return dt.DetermineType(cb)

	}
	panic("implement")
}

func (cb *AstContextBuilder) AddTypeSpec(typeSpec *definedNode.TypeSpec) {
	cb.TypeSpecs = append(cb.TypeSpecs, typeSpec)
}

func (cb *AstContextBuilder) namespace() string {
	return "github.com/bhbosman/goBoom/model"
}

func (cb *AstContextBuilder) ReadFiles(fileName string, readerCloser io.Reader) {
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, fileName, readerCloser, 0)
	if err != nil {
		panic(err)
	}
	indent := 0
	st := arraystack.New()
	clearStack := func(st *arraystack.Stack) func(currentPos ast.Node) {
		return func(currentPos ast.Node) {
			for {
				if value, ok := st.Peek(); ok {

					if currentPos.Pos() >= value.(ast.Node).End() {
						indent--
						if popValue, ok := st.Pop(); ok {
							if complete, ok := popValue.(definedNode.IDefinedNode); ok {
								complete.Complete(cb)
							}
						}
					} else {
						break
					}
				} else {
					break
				}
			}
		}
	}(st)
	ast.Inspect(
		astFile,
		func(stack *arraystack.Stack) func(n ast.Node) bool {
			return func(n ast.Node) bool {
				if n == nil {
					//clearStack(nil)
					return true
				}
				clearStack(n)
				indent++
				createLocalPos := func(n ast.Node) definedNode.IDefinedNode {
					switch v := n.(type) {
					case *ast.ArrayType:

						return definedNode.NewArrayType(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.AssignStmt:
						return definedNode.NewAssignStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.BinaryExpr:
						binaryExpr := definedNode.NewBinaryExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
						if stackValue, ok := st.Peek(); ok {
							if assignIdent, ok := stackValue.(definedNode.IAssignIdent); ok {
								assignIdent.AssignExpression(binaryExpr)
							} else {
								panic("Ident must be assigned")
							}
						}
						return binaryExpr
					case *ast.BlockStmt:
						blockStmt := definedNode.NewBlockStmt(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
						if stackValue, ok := st.Peek(); ok {
							if assignBlockStatement, ok := stackValue.(definedNode.IAssignBlockStatement); ok {
								assignBlockStatement.SetBlockStatement(blockStmt)
								return blockStmt
							} else {
								panic("BlockStatement must be assigned")
							}
						}
						return blockStmt
					case *ast.CallExpr:
						callExpr := definedNode.NewCallExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
						if stackValue, ok := st.Peek(); ok {
							if assignIdent, ok := stackValue.(definedNode.IAssignIdent); ok {
								assignIdent.AssignExpression(callExpr)
							} else {
								panic("Ident must be assigned")
							}
						}
						return callExpr
					case *ast.CompositeLit:
						return definedNode.NewCompositeLit(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.Field:
						if stackValue, ok := st.Peek(); ok {
							if fieldList, ok := stackValue.(definedNode.IAssignIdent); ok {
								field := definedNode.NewMultiFieldDeclaration(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
								fieldList.AssignExpression(field)
								return field
							}
						}
						panic("needs to be figured out. Should be a parent for ast.MultiFieldDeclaration")
					case *ast.FieldList:
						// temp object
						fieldList := definedNode.NewFieldList(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
						return fieldList
					case *ast.File:
						return definedNode.NewFile(indent, fileSet.Position(v.Pos()), v.Pos(), v.End())
					case *ast.ForStmt:
						return definedNode.NewForStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)

					case *ast.FuncType:
						if stackValue, ok := st.Peek(); ok {
							funcType := func(stackValue interface{}) *definedNode.FuncType {
								pos := n.Pos()
								switch parent := stackValue.(type) {
								case definedNode.IIdent:
									return definedNode.NewFuncType(indent, fileSet.Position(v.Pos()), pos, n.End(), parent.GetIdent(), v)
								default:
									return definedNode.NewFuncType(indent, fileSet.Position(v.Pos()), pos, n.End(), "(lambda)", v)
								}
							}(stackValue)

							switch parent := stackValue.(type) {
							case definedNode.IAssignIdent:
								parent.AssignExpression(funcType)
							default:
								panic("no ownership")
							}

							return funcType
						}
						panic("needs to be figured out. Should be a parent for functype")
					case *ast.FuncDecl:
						result := definedNode.NewFuncDecl(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
						return result
					case *ast.FuncLit:
						result := definedNode.NewFuncLit(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
						if stackValue, ok := st.Peek(); ok {
							if assignIdent, ok := stackValue.(definedNode.IAssignIdent); ok {
								assignIdent.AssignExpression(result)
							} else {
								panic("Ident must be assigned")
							}
						}
						return result
					case *ast.Ident:
						result := definedNode.NewIdent(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
						if stackValue, ok := st.Peek(); ok {
							if assignIdent, ok := stackValue.(definedNode.IAssignIdent); ok {
								assignIdent.AssignExpression(result)
							} else {
								panic("Ident must be assigned")
							}
						}
						return result
					case *ast.IfStmt:
						return definedNode.NewIfStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.IncDecStmt:
						return definedNode.NewIncDecStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.IndexExpr:
						return definedNode.NewIndexExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.IndexListExpr:
						return definedNode.NewIndexListExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.KeyValueExpr:
						return definedNode.NewKeyValueExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.RangeStmt:
						return definedNode.NewRangeStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.ReturnStmt:
						return definedNode.NewReturnStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.SelectorExpr:
						selectorExpr := definedNode.NewSelectorExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
						if stackValue, ok := st.Peek(); ok {
							if assignIdent, ok := stackValue.(definedNode.IAssignIdent); ok {
								assignIdent.AssignExpression(selectorExpr)
							} else {
								panic("Ident must be assigned")
							}
						}

						return selectorExpr
					case *ast.StructType:
						if stackValue, ok := st.Peek(); ok {
							structType := definedNode.NewStructType(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
							if assignStructType, ok := stackValue.(definedNode.IAssignIdent); ok {
								assignStructType.AssignExpression(structType)
								return structType
							}
						}
						panic("fix this issue")
					case *ast.TypeSpec:
						return definedNode.NewTypeSpec(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
					case *ast.ValueSpec:
						valueSpec := definedNode.NewValueSpec(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
						return valueSpec
					case *ast.ImportSpec:
						return definedNode.NewImportSpec(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.GenDecl:
						return definedNode.NewGenDecl(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.BasicLit:
						basicLit := definedNode.NewBasicLit(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
						if stackValue, ok := st.Peek(); ok {
							if assignIdent, ok := stackValue.(definedNode.IAssignIdent); ok {
								assignIdent.AssignExpression(basicLit)
							} else {
								panic("Ident must be assigned")
							}
						}
						return basicLit
					case *ast.ExprStmt:
						return definedNode.NewExprStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.BranchStmt:
						return definedNode.NewBranchStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.SwitchStmt:
						return definedNode.NewSwitchStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.CaseClause:
						return definedNode.NewCaseClause(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.DeclStmt:
						return definedNode.NewDeclStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.InterfaceType:
						return definedNode.NewInterfaceType(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.UnaryExpr:
						unaryExpr := definedNode.NewUnaryExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
						if stackValue, ok := st.Peek(); ok {
							if assignIdent, ok := stackValue.(definedNode.IAssignIdent); ok {
								assignIdent.AssignExpression(unaryExpr)
							} else {
								panic("Ident must be assigned")
							}
						}
						return unaryExpr
					case *ast.TypeSwitchStmt:
						return definedNode.NewTypeSwitchStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.TypeAssertExpr:
						return definedNode.NewTypeAssertExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.MapType:
						return definedNode.NewMapType(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.ParenExpr:
						parenExpr := definedNode.NewParenExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
						if stackValue, ok := st.Peek(); ok {
							if assignIdent, ok := stackValue.(definedNode.IAssignIdent); ok {
								assignIdent.AssignExpression(parenExpr)
							} else {
								panic("Ident must be assigned")
							}
						}
						return parenExpr
					default:
						panic("implemnt")
					}
				}
				lp := createLocalPos(n)
				doPush := func(unk interface{}) {
					if complete, ok := unk.(definedNode.IDefinedNode); ok {
						complete.Start(cb)
					}
					st.Push(unk)
				}
				switch v := lp.(type) {
				case definedNode.IRemoveNode:
					if !v.RemoveNode() {
						doPush(v)

					} else {
						if complete, ok := v.(definedNode.IDefinedNode); ok {
							complete.Start(cb)
							complete.Complete(cb)
						}
						indent--
					}
				default:
					doPush(v)
				}

				return true
			}
		}(st),
	)
	clearStack(definedNode.NewLocationP(0, token.Position{}, astFile.End()+1, 0))
	if st.Values() == nil {

	}
}

func (cb *AstContextBuilder) Validate() {
	cb.validateTypeSpecs()
	cb.validateValueSpec()
	cb.validateFuncDecl()

}

func (cb *AstContextBuilder) Generate() {

}

func (cb *AstContextBuilder) validateTypeSpecs() {
	for _, typeSpec := range cb.TypeSpecs {
		typeSpec.Validate(cb)
	}
}

func (cb *AstContextBuilder) validateValueSpec() {
	for _, valueSpec := range cb.ValueSpec {
		valueSpec.Validate(cb)
	}
}

func (cb *AstContextBuilder) Init() {
	cb.CurrentScope = &cb.RootScope
	cb.ValidTypes["int"] = reflect.TypeOf(int(0))
	cb.ValidTypes["float64"] = reflect.TypeOf(float64(0))
	//cb.validTypes["uint"] = reflect.TypeOf(uint(0))
	//cb.validTypes["int8"] = reflect.TypeOf(int8(0))
	//cb.validTypes["int16"] = reflect.TypeOf(int16(0))
	//cb.validTypes["int32"] = reflect.TypeOf(int32(0))
	//cb.validTypes["int64"] = reflect.TypeOf(int64(0))
	//cb.validTypes["uint8"] = reflect.TypeOf(uint8(0))
	//cb.validTypes["uint16"] = reflect.TypeOf(uint16(0))
	//cb.validTypes["uint32"] = reflect.TypeOf(uint32(0))
	//cb.validTypes["uint64"] = reflect.TypeOf(uint64(0))

	//
	functions.RegisterTranslateToInt(cb.ValidFunctions)
}

func (cb *AstContextBuilder) validateFuncDecl() {
	for _, funcDecl := range cb.FuncDecl {
		funcDecl.Validate(cb)
	}

}

func NewAstContextBuilder() *AstContextBuilder {
	return &AstContextBuilder{
		ValidTypes:     make(map[string]reflect.Type),
		ValidFunctions: make(map[functions.FuncKey]reflect.Value),
		TypeSpecs:      nil,
		RootScope: Scope{
			Parent: nil,
			Values: make(map[string]definedNode.IScopedItem),
		},
		CurrentScope: nil,
	}
}
