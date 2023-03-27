package playground

import (
	"github.com/bhbosman/goBoom/builder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test00001(t *testing.T) {
	t.Run("Test Structs", func(t *testing.T) {
		// Test Case: Test if we can read the definition with primitive field types
		filesToOpen := []string{"0001.go"}
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
		sut.Validate()
		assert.Len(t, sut.TypeSpecs, 1)
		ts := sut.TypeSpecs[0]
		assert.NotNil(t, ts)
		assert.IsType(t, ts, &builder.TypeSpec{})
		assert.Len(t, ts.StructType.Fields, 3)
		for i, field := range ts.StructType.Fields {
			assert.IsType(t, field, &builder.DeclaredField{})
			switch i {
			case 0:
				assert.Equal(t, field.Name(), "SomeInt001")
			case 1:
				assert.Equal(t, field.Name(), "SomeInt002")
			case 2:
				assert.Equal(t, field.Name(), "SomeInt003")
			}
			intType := sut.ValidTypes["int"]
			assert.Same(t, field.ReflectedType(), intType)

		}

	})

}
