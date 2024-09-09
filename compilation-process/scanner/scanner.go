package main

import (
	"fmt"
	// Package scanner implements a scanner for Go source text.
	"go/scanner"
	/* Package token defines constants representing the lexical
		tokens of the Go programming language and basic operations
	    on tokens (printing, predicates). */
	"go/token"
)

func main() {
	// a variable that holds the string (enclosed in backticks) into a slice of bytes.
	// ensures compatibility with scanner.Scanner, as input is src in s.Init
	src := (`package main

	import "fmt"

	func main() {
	  fmt.Println("Hello, world!")
	}
	`)

	// tokenize the code
	var s scanner.Scanner
	// variable fset that will hold a new file set,
	// soruce for fset and file: https://pkg.go.dev/go/token@go1.23.0#NewFileSet
	fset := token.NewFileSet()
	// holds, string, base offset, size of the file
	file := fset.AddFile("scanner", fset.Base(), len(src))

	// Initialize the scanner with file information, source code, error handling method
	// & mode
	s.Init(file, []byte(src), nil, 0)

	for {
		// position, token, string literal(if any) are assgined to s.Scan() which will return the values.
		pos, tok, lit := s.Scan()
		// print the data till tok equals token.EOF, the loop breaks, stopping the scanning process.
		fmt.Printf("%-6s  %-8s  %q\n", fset.Position(pos), tok, lit)
		if tok == token.EOF {
			break
		}
	}
}
