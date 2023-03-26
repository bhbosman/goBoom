package test

import (
	"github.com/bhbosman/goBoom/builder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test00004(t *testing.T) {
	t.Run("Validate", func(t *testing.T) {
		filesToOpen := []string{"00000.go", "00001.go"}
		sut := builder.NewAstContextBuilder()
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
	})
}
