package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	src := []byte(`
	package main
	import "fmt"
	import "go/scanner"
	import "unsafe"
	func main() {
		fmt.Println("Hello, world!")
	}
	`)

	// create a new file set
	fset := token.NewFileSet()
	// ParseFile parses the source code of a single Go source file and returns the corresponding ast.File node.
	file, err := parser.ParseFile(fset, "parser", src, 0)
	if err != nil {
		log.Fatal(err)
	}
	// print in ast format
	ast.Print(fset, file)
}
