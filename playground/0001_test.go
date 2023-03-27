package playground

import (
	"github.com/bhbosman/goBoom/builder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test00001(t *testing.T) {
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
	t.Run("Test Structs", func(t *testing.T) {
		sut.Validate()
		assert.Len(t, sut.TypeSpecs, 1)
		typeSpecs := sut.TypeSpecs[0]
		assert.NotNil(t, typeSpecs)
		assert.IsType(t, typeSpecs.DefinedNode[1], (*builder.StructType)(nil))
		structType := typeSpecs.DefinedNode[1].(*builder.StructType)
		assert.Len(t, structType.Fields, 3)
		for i, field := range structType.Fields {
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
