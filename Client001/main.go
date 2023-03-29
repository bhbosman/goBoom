package main

import (
	"embed"
	_ "github.com/bhbosman/goBoom/Client001/internal"
)

//go:embed internal/*.go
var content embed.FS

func main() {
	//dirEntries, err := content.ReadDir("internal")
	//if err != nil {
	//	return
	//}
	//contextBuilder := builder.NewAstContextBuilder()
	//for _, entry := range dirEntries {
	//	fileName := filepath.Join("internal", entry.Name())
	//	file, err := content.Open(fileName)
	//	if err != nil {
	//		return
	//	}
	//	contextBuilder.Build(fileName, file)
	//}
	//contextBuilder.Generate()
}
