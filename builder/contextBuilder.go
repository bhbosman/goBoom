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
	validTypes map[string]reflect.Type
	typeSpecs  []*TypeSpec
}

func (self *AstContextBuilder) ValidType(expr ast.Expr) {
	switch dt := expr.(type) {
	case *ast.Ident:
		if _, ok := self.validTypes[dt.Name]; !ok {
			panic(fmt.Errorf("not supporting type %v", dt.Name))
		}
	default:
		panic("implement")
	}

}

func (self *AstContextBuilder) AddTypeSpec(typeSpec *TypeSpec) {
	self.typeSpecs = append(self.typeSpecs, typeSpec)
}

func (self *AstContextBuilder) namespace() string {
	return "github.com/bhbosman/goBoom/model"
}

func (self *AstContextBuilder) ReadFiles(fileName string, readerCloser io.Reader) {
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
								complete.Complete(self)
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
						return NewCallExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
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
						if stackValue, ok := st.Peek(); ok {
							if assignIdent, ok := stackValue.(IAssignIdent); ok {
								assignIdent.AssignExpression(v)
							} else {
								panic("Ident must be assigned")
							}
						}
						result := NewIdent(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
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
							if assignStructType, ok := stackValue.(IAssignStructType); ok {
								assignStructType.AssignStructType(structType)
								return structType
							}
						}
						panic("fix this issue")
					case *ast.TypeSpec:
						return NewTypeSpec(indent, fileSet.Position(v.Pos()), n.Pos(), n.End(), v)
					case *ast.ValueSpec:
						return NewValueSpec(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.ImportSpec:
						return NewImportSpec(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.GenDecl:
						return NewGenDecl(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
					case *ast.BasicLit:
						return NewBasicLit(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
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
						return NewUnaryExpr(indent, fileSet.Position(v.Pos()), v.Pos(), v.End(), v)
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
						complete.Start(self)
					}
					st.Push(unk)
				}
				switch v := lp.(type) {
				case IRemoveNode:
					if !v.RemoveNode() {
						doPush(v)

					} else {
						if complete, ok := v.(IDefinedNode); ok {
							complete.Start(self)
							complete.Complete(self)
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

func (self *AstContextBuilder) Validate() {
	self.validateTypeSpecs()
}

func (self *AstContextBuilder) Generate() {

}

func (self *AstContextBuilder) validateTypeSpecs() {
	for _, typeSpec := range self.typeSpecs {
		typeSpec.Validate(self)

	}
}

func (self *AstContextBuilder) Init() {
	self.validTypes["int"] = reflect.TypeOf(int(0))
	//self.validTypes["uint"] = reflect.TypeOf(uint(0))
	//self.validTypes["int8"] = reflect.TypeOf(int8(0))
	//self.validTypes["int16"] = reflect.TypeOf(int16(0))
	//self.validTypes["int32"] = reflect.TypeOf(int32(0))
	//self.validTypes["int64"] = reflect.TypeOf(int64(0))
	//self.validTypes["uint8"] = reflect.TypeOf(uint8(0))
	//self.validTypes["uint16"] = reflect.TypeOf(uint16(0))
	//self.validTypes["uint32"] = reflect.TypeOf(uint32(0))
	//self.validTypes["uint64"] = reflect.TypeOf(uint64(0))
}

func NewAstContextBuilder() *AstContextBuilder {
	return &AstContextBuilder{
		validTypes: make(map[string]reflect.Type),
		typeSpecs:  nil,
	}
}
