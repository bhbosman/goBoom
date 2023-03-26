package playground

// Declaration file for tests
import "embed"

//go:embed *.go
var content embed.FS
var namespace = "github.com/bhbosman/goBoom/model"
