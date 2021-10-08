package main

import (
	"fmt"
	root "github.com/virtengine/virtengine/cmd/virtengine/cmd"
	"github.com/spf13/cobra/doc"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage is:\n\tvirtengine_docgen <output path>\n")
		os.Exit(1)
	}
	outputPath := os.Args[1]
	cmd, _ := root.NewRootCmd()
	err := doc.GenMarkdownTree(cmd, outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed generating markdown into %q:%v\n", outputPath, err)
		os.Exit(1)
	}
}
