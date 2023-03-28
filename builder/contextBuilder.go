package builder

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"reflect"
)

type AstContextBuilder struct {
	ValidTypes map[string]reflect.Type
	TypeSpecs  []*TypeSpec
	ValueSpec  []*ValueSpec
}

func (cb *AstContextBuilder) AddValueSpec(valueSpec *ValueSpec) {
	cb.ValueSpec = append(cb.ValueSpec, valueSpec)
}

func (cb *AstContextBuilder) ValidTypeFromKind(kind token.Token) reflect.Type {
	switch kind {
	case token.INT:
		return cb.ValidTypes["int"]
	}
	panic("implement")
}

func (cb *AstContextBuilder) ValidType(expr IDefinedNode) reflect.Type {
	switch dt := expr.(type) {
	case *Ident:
		if v, ok := cb.ValidTypes[dt.AstIdent.Name]; ok {
			return v
		}
		panic(fmt.Errorf("not supporting type %v", dt.AstIdent.Name))
	case *BasicLit:
		return cb.ValidTypeFromKind(dt.basicLit.Kind)
	}
	panic("implement")
}

func (cb *AstContextBuilder) AddTypeSpec(typeSpec *TypeSpec) {
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
							if complete, ok := popValue.(IDefinedNode); ok {
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
				createLocalPos := func(n ast.Node) IDefinedNode {
					switch v := n.(type) {
					case *ast.ArrayType:

						return NewArrayType(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.AssignStmt:
						return NewAssignStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.BinaryExpr:
						return NewBinaryExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.BlockStmt:
						blockStmt := NewBlockStmt(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
						if stackValue, ok := st.Peek(); ok {
							if assignBlockStatement, ok := stackValue.(IAssignBlockStatement); ok {
								assignBlockStatement.SetBlockStatement(blockStmt)
								return blockStmt
							} else {
								panic("BlockStatement must be assigned")
							}
						}
						return blockStmt
					case *ast.CallExpr:
						callExpr := NewCallExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
						if stackValue, ok := st.Peek(); ok {
							if assignIdent, ok := stackValue.(IAssignIdent); ok {
								assignIdent.AssignExpression(callExpr)
							} else {
								panic("Ident must be assigned")
							}
						}
						return callExpr
					case *ast.CompositeLit:
						return NewCompositeLit(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.Field:
						if stackValue, ok := st.Peek(); ok {
							if fieldList, ok := stackValue.(IAddField); ok {
								field := NewMultiFieldDeclaration(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
								fieldList.AddField(field)
								return field
							}
						}
						panic("needs to be figured out. Should be a parent for ast.MultiFieldDeclaration")
					case *ast.FieldList:
						// temp object
						fieldList := NewFieldList(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
						return fieldList
					case *ast.File:
						return NewFile(indent, fileSet.Position(v.Pos()), v.Pos(), v.End())
					case *ast.ForStmt:
						return NewForStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)

					case *ast.FuncType:
						if stackValue, ok := st.Peek(); ok {
							funcType := func(stackValue interface{}) *FuncType {
								pos := n.Pos()
								switch parent := stackValue.(type) {
								case IIdent:
									return NewFuncType(indent, fileSet.Position(v.Pos()), pos, n.End(), parent.GetIdent(), v)
								default:
									return NewFuncType(indent, fileSet.Position(v.Pos()), pos, n.End(), "(lambda)", v)
								}
							}(stackValue)

							switch parent := stackValue.(type) {
							case IAssignFuncType:
								parent.SetFuncType(funcType)
							default:
								panic("no ownership")
							}

							return funcType
						}
						panic("needs to be figured out. Should be a parent for functype")
					case *ast.FuncDecl:
						result := NewFuncDecl(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
						return result
					case *ast.FuncLit:
						result := NewFuncLit(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
						return result
					case *ast.Ident:
						result := NewIdent(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
						if stackValue, ok := st.Peek(); ok {
							if assignIdent, ok := stackValue.(IAssignIdent); ok {
								assignIdent.AssignExpression(result)
							} else {
								panic("Ident must be assigned")
							}
						}
						return result
					case *ast.IfStmt:
						return NewIfStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.IncDecStmt:
						return NewIncDecStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.IndexExpr:
						return NewIndexExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.IndexListExpr:
						return NewIndexListExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.KeyValueExpr:
						return NewKeyValueExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.RangeStmt:
						return NewRangeStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.ReturnStmt:
						return NewReturnStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.SelectorExpr:
						return NewSelectorExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.StructType:
						if stackValue, ok := st.Peek(); ok {
							structType := NewStructType(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
							if assignStructType, ok := stackValue.(IAssignIdent); ok {
								assignStructType.AssignExpression(structType)
								return structType
							}
						}
						panic("fix this issue")
					case *ast.TypeSpec:
						return NewTypeSpec(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
					case *ast.ValueSpec:
						valueSpec := NewValueSpec(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
						return valueSpec
					case *ast.ImportSpec:
						return NewImportSpec(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.GenDecl:
						return NewGenDecl(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.BasicLit:
						basicLit := NewBasicLit(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
						if stackValue, ok := st.Peek(); ok {
							if assignIdent, ok := stackValue.(IAssignIdent); ok {
								assignIdent.AssignExpression(basicLit)
							} else {
								panic("Ident must be assigned")
							}
						}
						return basicLit
					case *ast.ExprStmt:
						return NewExprStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.BranchStmt:
						return NewBranchStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.SwitchStmt:
						return NewSwitchStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.CaseClause:
						return NewCaseClause(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.DeclStmt:
						return NewDeclStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.InterfaceType:
						return NewInterfaceType(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.UnaryExpr:
						unaryExpr := NewUnaryExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
						if stackValue, ok := st.Peek(); ok {
							if assignIdent, ok := stackValue.(IAssignIdent); ok {
								assignIdent.AssignExpression(unaryExpr)
							} else {
								panic("Ident must be assigned")
							}
						}
						return unaryExpr
					case *ast.TypeSwitchStmt:
						return NewTypeSwitchStmt(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.TypeAssertExpr:
						return NewTypeAssertExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.MapType:
						return NewMapType(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					default:
						panic("implemnt")
					}
				}
				lp := createLocalPos(n)
				doPush := func(unk interface{}) {
					if complete, ok := unk.(IDefinedNode); ok {
						complete.Start(cb)
					}
					st.Push(unk)
				}
				switch v := lp.(type) {
				case IRemoveNode:
					if !v.RemoveNode() {
						doPush(v)

					} else {
						if complete, ok := v.(IDefinedNode); ok {
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
	clearStack(
		&Location{
			pos: astFile.End() + 1,
			end: 0,
		},
	)
	if st.Values() == nil {

	}
}

func (cb *AstContextBuilder) Validate() {
	cb.validateTypeSpecs()
	cb.validateValueSpec()
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
	cb.ValidTypes["int"] = reflect.TypeOf(int(0))
	//cb.validTypes["uint"] = reflect.TypeOf(uint(0))
	//cb.validTypes["int8"] = reflect.TypeOf(int8(0))
	//cb.validTypes["int16"] = reflect.TypeOf(int16(0))
	//cb.validTypes["int32"] = reflect.TypeOf(int32(0))
	//cb.validTypes["int64"] = reflect.TypeOf(int64(0))
	//cb.validTypes["uint8"] = reflect.TypeOf(uint8(0))
	//cb.validTypes["uint16"] = reflect.TypeOf(uint16(0))
	//cb.validTypes["uint32"] = reflect.TypeOf(uint32(0))
	//cb.validTypes["uint64"] = reflect.TypeOf(uint64(0))
}

func NewAstContextBuilder() *AstContextBuilder {
	return &AstContextBuilder{
		ValidTypes: make(map[string]reflect.Type),
		TypeSpecs:  nil,
	}
}
