package playground

import (
	"github.com/bhbosman/goBoom/builder"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test00002(t *testing.T) {
	filesToOpen := []string{"intAssignments.go"}
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
		assert.Len(t, sut.ValueSpec, 8)
		t.Run("Test first ValueSpec", func(t *testing.T) {
			valueSpec := sut.ValueSpec[0]
			assert.Len(t, valueSpec.Expressions, 6)
		})
		t.Run("Test second ValueSpec", func(t *testing.T) {
			valueSpec := sut.ValueSpec[1]
			assert.Len(t, valueSpec.Expressions, 6)
		})
		assert.Len(t, sut.RootScope.Values, 12)
		t.Run("Check types of scoped values", func(t *testing.T) {
			v1 := sut.RootScope.Values["one"]
			assert.Same(t, reflect.TypeOf(0), v1.DetermineType(nil))
			v2 := sut.RootScope.Values["two"]
			assert.Same(t, reflect.TypeOf(0), v2.DetermineType(nil))
			v3 := sut.RootScope.Values["three"]
			assert.Same(t, reflect.TypeOf(0), v3.DetermineType(nil))
			v4 := sut.RootScope.Values["four"]
			assert.Same(t, reflect.TypeOf(0), v4.DetermineType(nil))
			v5 := sut.RootScope.Values["five"]
			assert.Same(t, reflect.TypeOf(0), v5.DetermineType(nil))
			v6 := sut.RootScope.Values["six"]
			assert.Same(t, reflect.TypeOf(0), v6.DetermineType(nil))
			v7 := sut.RootScope.Values["six"]
			assert.Same(t, reflect.TypeOf(0), v7.DetermineType(nil))
			v8 := sut.RootScope.Values["six"]
			assert.Same(t, reflect.TypeOf(0), v8.DetermineType(nil))
			v9 := sut.RootScope.Values["six"]
			assert.Same(t, reflect.TypeOf(0), v9.DetermineType(nil))
			v10 := sut.RootScope.Values["six"]
			assert.Same(t, reflect.TypeOf(0), v10.DetermineType(nil))
			v11 := sut.RootScope.Values["six"]
			assert.Same(t, reflect.TypeOf(0), v11.DetermineType(nil))
			v12 := sut.RootScope.Values["six"]
			assert.Same(t, reflect.TypeOf(0), v12.DetermineType(nil))
		})
		t.Run("Determine Value", func(t *testing.T) {
			v1 := sut.RootScope.Values["one"]
			assert.Same(t, reflect.TypeOf(0), v1.DetermineType(nil))
			value1 := v1.DetermineValue(sut)
			assert.Same(t, reflect.TypeOf(0), value1.Type())
			assert.Equal(t, int64(1), value1.Int())

			v2 := sut.RootScope.Values["two"]
			assert.Same(t, reflect.TypeOf(0), v2.DetermineType(nil))
			value2 := v2.DetermineValue(sut)
			assert.Same(t, reflect.TypeOf(0), value2.Type())
			assert.Equal(t, int64(2), value2.Int())

			v3 := sut.RootScope.Values["three"]
			assert.Same(t, reflect.TypeOf(0), v3.DetermineType(nil))
			value3 := v3.DetermineValue(sut)
			assert.Same(t, reflect.TypeOf(0), value3.Type())
			assert.Equal(t, int64(3), value3.Int())

			v4 := sut.RootScope.Values["four"]
			assert.Same(t, reflect.TypeOf(0), v4.DetermineType(nil))
			value4 := v4.DetermineValue(sut)
			assert.Same(t, reflect.TypeOf(0), value4.Type())
			assert.Equal(t, int64(4), value4.Int())

			v5 := sut.RootScope.Values["five"]
			assert.Same(t, reflect.TypeOf(0), v5.DetermineType(nil))
			value5 := v5.DetermineValue(sut)
			assert.Same(t, reflect.TypeOf(0), value5.Type())
			assert.Equal(t, int64(5), value5.Int())

			v6 := sut.RootScope.Values["six"]
			assert.Same(t, reflect.TypeOf(0), v6.DetermineType(nil))
			value6 := v6.DetermineValue(sut)
			assert.Same(t, reflect.TypeOf(0), value6.Type())
			assert.Equal(t, int64(6), value6.Int())

			v7 := sut.RootScope.Values["seven"]
			assert.Same(t, reflect.TypeOf(0), v6.DetermineType(nil))
			value7 := v7.DetermineValue(sut)
			assert.Same(t, reflect.TypeOf(0), value7.Type())
			assert.Equal(t, int64(7), value7.Int())

			v8 := sut.RootScope.Values["eight"]
			assert.Same(t, reflect.TypeOf(0), v8.DetermineType(nil))
			value8 := v8.DetermineValue(sut)
			assert.Same(t, reflect.TypeOf(0), value8.Type())
			assert.Equal(t, int64(8), value8.Int())

			v9 := sut.RootScope.Values["nine"]
			assert.Same(t, reflect.TypeOf(0), v9.DetermineType(nil))
			value9 := v9.DetermineValue(sut)
			assert.Same(t, reflect.TypeOf(0), value9.Type())
			assert.Equal(t, int64(9), value9.Int())

			v10 := sut.RootScope.Values["ten"]
			assert.Same(t, reflect.TypeOf(0), v10.DetermineType(nil))
			value10 := v10.DetermineValue(sut)
			assert.Same(t, reflect.TypeOf(0), value10.Type())
			assert.Equal(t, int64(10), value10.Int())

			v11 := sut.RootScope.Values["eleven"]
			assert.Same(t, reflect.TypeOf(0), v11.DetermineType(nil))
			value11 := v11.DetermineValue(sut)
			assert.Same(t, reflect.TypeOf(0), value11.Type())
			assert.Equal(t, int64(11), value11.Int())

			v12 := sut.RootScope.Values["twelve"]
			assert.Same(t, reflect.TypeOf(0), v12.DetermineType(nil))
			value12 := v12.DetermineValue(sut)
			assert.Same(t, reflect.TypeOf(0), value12.Type())
			assert.Equal(t, int64(12), value12.Int())
		})
	})
}
