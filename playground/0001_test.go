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
		sut := builder.NewAstContextBuilder()
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
	})

}
