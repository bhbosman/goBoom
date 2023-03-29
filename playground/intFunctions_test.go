package playground

import (
	"github.com/bhbosman/goBoom/builder"
	_ "github.com/golang/mock/gomock"
	"reflect"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntFunctions(t *testing.T) {

	//gomock.Controller
	filesToOpen := []string{"intFunctions.go"}
	var sut = builder.NewAstContextBuilder()
	sut.Init()
	for _, fileToOpen := range filesToOpen {
		readCloser, err := content.Open(fileToOpen)
		if !assert.NoError(t, err) {
			return
		}
		sut.ReadFiles(fileToOpen, readCloser)
		err = readCloser.Close()
		if !assert.NoError(t, err) {
			return
		}
	}
	t.Run("Check Validation", func(t *testing.T) {
		sut.Validate()
		t.Run("Check validation of Nothing", func(t *testing.T) {
			v1 := sut.RootScope.Values["Nothing"]
			assert.NotNil(t, v1)
			funcOf := reflect.FuncOf([]reflect.Type{}, []reflect.Type{}, false)
			assert.Same(t, funcOf, v1.DetermineType(sut))
			//v1.DetermineValue(sut)
		})
		t.Run("Check validation of OnlyInputWithNames", func(t *testing.T) {
			v1 := sut.RootScope.Values["OnlyInputWithNames"]
			assert.NotNil(t, v1)
			funcOf := reflect.FuncOf(
				[]reflect.Type{
					reflect.TypeOf(0),
				},
				[]reflect.Type{},
				false,
			)
			assert.Same(t, funcOf, v1.DetermineType(sut))
			//v1.DetermineValue(sut)
		})
		t.Run("Check validation of OnlyInputWithNoNames", func(t *testing.T) {
			v1 := sut.RootScope.Values["OnlyInputWithNoNames"]
			assert.NotNil(t, v1)
			funcOf := reflect.FuncOf(
				[]reflect.Type{
					reflect.TypeOf(0),
				},
				[]reflect.Type{},
				false,
			)
			assert.Same(t, funcOf, v1.DetermineType(sut))
			//v1.DetermineValue(sut)
		})
		t.Run("Check validation of OnlyOutput", func(t *testing.T) {
			v1 := sut.RootScope.Values["OnlyOutput"]
			assert.NotNil(t, v1)
			funcOf := reflect.FuncOf(
				[]reflect.Type{},
				[]reflect.Type{
					reflect.TypeOf(0),
				},
				false,
			)
			assert.Same(t, funcOf, v1.DetermineType(sut))
			//v1.DetermineValue(sut)
		})
		t.Run("Check validation of Both", func(t *testing.T) {
			v1 := sut.RootScope.Values["Both"]
			assert.NotNil(t, v1)
			funcOf := reflect.FuncOf(
				[]reflect.Type{
					reflect.TypeOf(0),
				},
				[]reflect.Type{
					reflect.TypeOf(0),
				},
				false,
			)
			assert.Same(t, funcOf, v1.DetermineType(sut))
			//v1.DetermineValue(sut)
		})
		t.Run("Check validation of MethodMinus", func(t *testing.T) {
			v1 := sut.RootScope.Values["MethodMinus"]
			assert.NotNil(t, v1)
			funcOf := reflect.FuncOf(
				[]reflect.Type{
					reflect.TypeOf(0),
					reflect.TypeOf(0),
				},
				[]reflect.Type{
					reflect.TypeOf(0),
				},
				false,
			)
			assert.Same(t, funcOf, v1.DetermineType(sut))
			//v1.DetermineValue(sut)
		})
		t.Run("Check validation of MethodMod", func(t *testing.T) {
			v1 := sut.RootScope.Values["MethodMod"]
			assert.NotNil(t, v1)
			funcOf := reflect.FuncOf(
				[]reflect.Type{
					reflect.TypeOf(0),
					reflect.TypeOf(0),
				},
				[]reflect.Type{
					reflect.TypeOf(0),
					reflect.TypeOf(0),
				},
				false,
			)
			assert.Same(t, funcOf, v1.DetermineType(sut))
			//v1.DetermineValue(sut)
		})
		t.Run("Check validation of MethodModWithReturnVariables", func(t *testing.T) {
			v1 := sut.RootScope.Values["MethodModWithReturnVariables"]
			assert.NotNil(t, v1)
			funcOf := reflect.FuncOf(
				[]reflect.Type{
					reflect.TypeOf(0),
					reflect.TypeOf(0),
				},
				[]reflect.Type{
					reflect.TypeOf(0),
					reflect.TypeOf(0),
				},
				false,
			)
			assert.Same(t, funcOf, v1.DetermineType(sut))
			//v1.DetermineValue(sut)
		})

	})
}
