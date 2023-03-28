package playground

import (
	"github.com/bhbosman/goBoom/builder"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test00002(t *testing.T) {
	filesToOpen := []string{"0002.go"}
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
		assert.Len(t, sut.ValueSpec, 2)
		t.Run("Test first ValueSpec", func(t *testing.T) {
			valueSpec := sut.ValueSpec[0]
			assert.Len(t, valueSpec.Expressions, 6)
		})
		t.Run("Test first ValueSpec", func(t *testing.T) {
			valueSpec := sut.ValueSpec[1]
			assert.Len(t, valueSpec.Expressions, 6)
		})
	})
}

func TestName(t *testing.T) {

	var a any
	reflect.TypeOf(a)
	//funcOf := reflect.TypeOf(math.Abs)

	funcOf := reflect.FuncOf(
		[]reflect.Type{reflect.TypeOf(int8(0))},
		[]reflect.Type{reflect.TypeOf(0)},
		false,
	)

	fn := reflect.MakeFunc(
		funcOf,
		func(args []reflect.Value) (results []reflect.Value) {
			reflect.ValueOf(0).SetInt()
			typeOfInt := reflect.TypeOf((0))
			s := args[0].Type().String()
			println(s)
			if args[0].Type().AssignableTo(typeOfInt) {
				return nil
			}

			return nil
		},
	)
	s := int(8)

	fn.Call(
		[]reflect.Value{
			reflect.ValueOf(int8(s)),
		},
	)
}
