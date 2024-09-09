## Go compilation process:

### 3 phase of compiler
* The scanner, which converts the source code into a list of tokens, for use by the parser.
* The parser, which converts the tokens into an Abstract Syntax Tree to be used by code generation.
* The code generation, which converts the Abstract Syntax Tree to machine code.

Note:

> 1. **Packages for Tools**: The packages like `go/scanner`, `go/parser`, `go/token`, and `go/ast` are designed to help developers create tools that work with Go source code. These tools might include code editors, linters, or other utilities that analyze or manipulate Go code.

> 2. **Not Used by the Compiler**: The actual Go compiler, which translates Go code into executable programs, doesn't use these packages. Instead, it has its own way of handling Go code.

> 3. **Historical Reason**: The reason the Go compiler doesn't use these packages is historical. The Go compiler was originally written in the C programming language. When it was later converted to Go, it retained much of its original structure and didn't switch to using these newer packages.


### Scanner

### Notes on Go Compiler and Scanner

#### **Compiler Basics**
- **First Step**: Breaks raw source code into tokens (lexical analysis).
- **Tokens**: Keywords, strings, variable names, function names, etc.
  - Examples: "package", "main", "func".
- **Token Representation**: Each token has a position, type, and raw text.

#### **Go Specifics**
- **Packages**: `go/scanner` and `go/token` allow manual execution of the scanner.
- **Purpose**: Inspect what the Go compiler sees after scanning the source code.

#### **Example Program**
- **Goal**: Print all tokens of a "Hello World" program.
- **Steps**:
  1. Create source code string.
  2. Initialize `scanner.Scanner` struct.
  3. Call `Scan()` repeatedly.
  4. Print token’s position, type, and literal string.
  5. Stop at End of File (EOF) marker.

```go
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
	src := []byte(`
	package main

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
	s.Init(file, src, nil, 0)

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
```

**Why Convert to a Byte Slice?**
* Compatibility: Many functions in Go, especially those dealing with I/O and low-level operations, work with byte slices rather than strings.
* Efficiency: Byte slices are mutable, meaning you can modify them in place without creating new copies, which can be more efficient for certain operations.

source: https://github.com/golang/go/blob/master/src/go/scanner/scanner.go


**Role of `FileSet` with Byte Slice**
* Position Management:
  * Byte Slice: The byte slice (src) contains the raw source code.
  * FileSet: The FileSet manages the positions of tokens within this source code. It translates byte offsets into more meaningful positions like line and column numbers.
* Contextual Information:
  * Raw Positions: Without a FileSet, the scanner would only provide raw byte offsets (e.g., “position 10”).
  * Human-Readable Positions: The FileSet converts these raw positions into human-readable formats (e.g., “line 1, column 10”), which are much more useful for debugging and error reporting.
* Multiple Files:
  * Single File: Even if you’re working with a single byte slice, the FileSet helps in managing its positions.
  * Multiple Files: If you later decide to work with multiple source files, the FileSet can handle all of them together, ensuring that positions are unique and correctly mapped.

source: https://pkg.go.dev/go/token@go1.23.0#NewFileSet

* Example: If the base offset is 1000, the first token in the new file might have a position of 1001, the second token 1002, and so on.

```go
	s.Init(file, src, nil, 0)
```
* s.init https://pkg.go.dev/go/scanner#Scanner.Init
* The mode parameter determines how comments are handled.
* Passing nil means no custom error handling function is provided. The scanner will use its default error handling.
* Mode controls scanner behavior, how comments are handeled, Passing 0 means no special scanning modes are enabled. The scanner will operate in its default mode.

```go
	for {
		pos, tok, lit := s.Scan()
		fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)

		if tok == token.EOF {
			break
		}
	}
}
```
* https://pkg.go.dev/go/scanner#Scanner.Scan
  * Scan scans the next token and returns the token position, the token, and its literal string if applicable.
  * It will return values based on the code such that:
    * pos: Holds the position of the current token.
    * tok: Holds the type of the current token.
    * lit: Holds the literal value of the current token, if any.
```go
		fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)
```
Sure, let's break down the `fmt.Printf` statement in detail:

### `fmt.Printf` Function
- **Purpose**: `fmt.Printf` is a function from the `fmt` package in Go that formats and prints data according to a specified format string.

### Format String
```go
"%-6s%-8s%q\n"
```
- **`%-6s`**:
  - **`%s`**: This is a verb that formats the argument as a string.
  - **`-6`**: This specifies a minimum width of 6 characters and left-aligns the string within this width.
- **`%-8s`**:
  - **`%s`**: Formats the argument as a string.
  - **`-8`**: Specifies a minimum width of 8 characters and left-aligns the string within this width.
- **`%q`**:
  - Formats the argument as a double-quoted string.
- **`\n`**:
  - Newline character, which moves the cursor to the next line after printing.

### Arguments
```go
fset.Position(pos), tok, lit
```
- **`fset.Position(pos)`**:
  - **`fset.Position`**: This method converts the raw position (`pos`) into a human-readable format, typically showing the line and column number.
  - **Example**: If `pos` corresponds to the start of the file, `fset.Position(pos)` might return `1:1` (line 1, column 1).
- **`tok`**:
  - This is the token type, such as `token.PACKAGE`, `token.FUNC`, etc.
  - **Example**: `token.PACKAGE` might be printed as `package`.
- **`lit`**:
  - This is the literal value of the token, if applicable.
  - **Example**: For a string literal, `lit` might be `"Hello, world!"`.

### Putting It All Together
The `fmt.Printf` statement formats and prints the position, token type, and literal value of each token in a structured way.

### Example Output
Assuming the source code is:
```go
package main
```
The output might look like:
```
1:1    package "package"
1:9    main    "main"
```
- **`1:1`**: Position of the `package` keyword (line 1, column 1).
- **`package`**: Token type.
- **`"package"`**: Literal value.
- **`1:9`**: Position of the `main` identifier (line 1, column 9).
- **`main`**: Token type.
- **`"main"`**: Literal value.

### Summary
- **`fmt.Printf`**: Formats and prints the token information.
- **Format String**: Specifies how each piece of information should be formatted.
- **Arguments**: Provide the position, token type, and literal value to be printed.

#### **Key Points**
- **Manual Scanning**: Allows inspection of tokens as seen by the Go compiler.
- **Semicolon Insertion**: Explains why Go does not require explicit semicolons.

### Output

```bash
:~/go-lang-programming/compilation-process$ go run scanner.go
scanner:2:2  package   "package"
scanner:2:10  IDENT     "main"
scanner:2:14  ;         "\n"
scanner:4:2  import    "import"
scanner:4:9  STRING    "\"fmt\""
scanner:4:14  ;         "\n"
scanner:6:2  func      "func"
scanner:6:7  IDENT     "main"
scanner:6:11  (         ""
scanner:6:12  )         ""
scanner:6:14  {         ""
scanner:7:4  IDENT     "fmt"
scanner:7:7  .         ""
scanner:7:8  IDENT     "Println"
scanner:7:15  (         ""
scanner:7:16  STRING    "\"Hello, world!\""
scanner:7:31  )         ""
scanner:7:32  ;         "\n"
scanner:8:2  }         ""
scanner:8:3  ;         "\n"
scanner:9:2  EOF       ""
```

### Parser

* The parser is a phase of the compiler that converts the tokens into an Abstract Syntax Tree (AST).
* The AST is a structured representation of the source code. In the AST we will be able to see the program structure, such as functions and constant declarations.
* Go has again provided us with packages to parse the program and view the AST: go/parser and go/ast.

```go
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
	func main() {
		fmt.Println("Hello, world!")
	}
	`)

	// create a new file set
	fset := token.NewFileSet()


	file, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		log.Fatal(err)
	}

	ast.Print(fset, file)
}
```
Program remains same, pareser.ParseFile is used for converting token to AST

```go
file, err := parser.ParseFile(fset, "", src, 0)
```
source: https://pkg.go.dev/go/parser@go1.23.0#ParseFile

* ParseFile parses the source code of a single Go source file and returns the corresponding ast.File node.
* file: This is a pointer to an ast.File structure, which represents the parsed Go source file.
* err: This is an error value that will be nil if the parsing was successful, or contain an error message if something went wrong.
* For printing we are using ast package, https://pkg.go.dev/go/ast@go1.23.0#Print

**AST output**
```bash
~/go-lang-programming/compilation-process$ go run parser/parser.go
     0  *ast.File {
     1  .  Package: 2:2
     2  .  Name: *ast.Ident {
     3  .  .  NamePos: 2:10
     4  .  .  Name: "main"
     5  .  }
     6  .  Decls: []ast.Decl (len = 2) {
     7  .  .  0: *ast.GenDecl {
     8  .  .  .  TokPos: 3:2
     9  .  .  .  Tok: import
    10  .  .  .  Lparen: -
    11  .  .  .  Specs: []ast.Spec (len = 1) {
    12  .  .  .  .  0: *ast.ImportSpec {
    13  .  .  .  .  .  Path: *ast.BasicLit {
    14  .  .  .  .  .  .  ValuePos: 3:9
    15  .  .  .  .  .  .  Kind: STRING
    16  .  .  .  .  .  .  Value: "\"fmt\""
    17  .  .  .  .  .  }
    18  .  .  .  .  .  EndPos: -
    19  .  .  .  .  }
    20  .  .  .  }
    21  .  .  .  Rparen: -
    22  .  .  }
    23  .  .  1: *ast.FuncDecl {
    24  .  .  .  Name: *ast.Ident {
    25  .  .  .  .  NamePos: 4:7
    26  .  .  .  .  Name: "main"
    27  .  .  .  .  Obj: *ast.Object {
    28  .  .  .  .  .  Kind: func
    29  .  .  .  .  .  Name: "main"
    30  .  .  .  .  .  Decl: *(obj @ 23)
    31  .  .  .  .  }
    32  .  .  .  }
    33  .  .  .  Type: *ast.FuncType {
    34  .  .  .  .  Func: 4:2
    35  .  .  .  .  Params: *ast.FieldList {
    36  .  .  .  .  .  Opening: 4:11
    37  .  .  .  .  .  Closing: 4:12
    38  .  .  .  .  }
    39  .  .  .  }
    40  .  .  .  Body: *ast.BlockStmt {
    41  .  .  .  .  Lbrace: 4:14
    42  .  .  .  .  List: []ast.Stmt (len = 1) {
    43  .  .  .  .  .  0: *ast.ExprStmt {
    44  .  .  .  .  .  .  X: *ast.CallExpr {
    45  .  .  .  .  .  .  .  Fun: *ast.SelectorExpr {
    46  .  .  .  .  .  .  .  .  X: *ast.Ident {
    47  .  .  .  .  .  .  .  .  .  NamePos: 5:3
    48  .  .  .  .  .  .  .  .  .  Name: "fmt"
    49  .  .  .  .  .  .  .  .  }
    50  .  .  .  .  .  .  .  .  Sel: *ast.Ident {
    51  .  .  .  .  .  .  .  .  .  NamePos: 5:7
    52  .  .  .  .  .  .  .  .  .  Name: "Println"
    53  .  .  .  .  .  .  .  .  }
    54  .  .  .  .  .  .  .  }
    55  .  .  .  .  .  .  .  Lparen: 5:14
    56  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
    57  .  .  .  .  .  .  .  .  0: *ast.BasicLit {
    58  .  .  .  .  .  .  .  .  .  ValuePos: 5:15
    59  .  .  .  .  .  .  .  .  .  Kind: STRING
    60  .  .  .  .  .  .  .  .  .  Value: "\"Hello, world!\""
    61  .  .  .  .  .  .  .  .  }
    62  .  .  .  .  .  .  .  }
    63  .  .  .  .  .  .  .  Ellipsis: -
    64  .  .  .  .  .  .  .  Rparen: 5:30
    65  .  .  .  .  .  .  }
    66  .  .  .  .  .  }
    67  .  .  .  .  }
    68  .  .  .  .  Rbrace: 6:2
    69  .  .  .  }
    70  .  .  }
    71  .  }
    72  .  FileStart: 1:1
    73  .  FileEnd: 7:2
    74  .  Scope: *ast.Scope {
    75  .  .  Objects: map[string]*ast.Object (len = 1) {
    76  .  .  .  "main": *(obj @ 27)
    77  .  .  }
    78  .  }
    79  .  Imports: []*ast.ImportSpec (len = 1) {
    80  .  .  0: *(obj @ 12)
    81  .  }
    82  .  Unresolved: []*ast.Ident (len = 1) {
    83  .  .  0: *(obj @ 46)
    84  .  }
    85  .  GoVersion: ""
    86  }
```

### Notes on AST Structure

- **Decls Field**: Contains all declarations in the file (imports, constants, variables, functions).
  - **Current Declarations**:
    - Import of the `fmt` package.
    - `main` function.

- **Main Function**:
  - **Name**: Represented as an identifier with the value `main`.
  - **Declaration (Type Field)**: Would contain parameters and return type if specified.
  - **Body**: List of statements (only one in this case).

- **fmt.Println Statement**:
  - **ExprStmt**: Represents an expression (function call, literal, binary operation, unary operation, etc.).
    - **CallExpr**: The actual function call.
      - **Fun**: Reference to the function call (a `SelectorExpr` selecting `Println` from `fmt`).
      - **Args**: List of expressions (arguments to the function).
        - **BasicLit**: Represents the literal string `"Hello, world!"` with type `STRING`.

To extract all function calls, we are going to use the following code: https://gist.github.com/koesie10/ba6af59e0dd8213260e5944c1464b0b1

### Machine code generation

1. **Validation**:
   - After resolving imports and checking types, the program is confirmed as valid Go code.

2. **Conversion to SSA**:
   - **AST to SSA**: The Abstract Syntax Tree (AST) is converted to Static Single Assignment (SSA) form.
   - **SSA Characteristics**:
     - Variables are always defined before use.
     - Each variable is assigned exactly once.

3. **Optimization Passes**:
   - **Purpose**: Simplify or speed up code execution.
   - **Examples**:
     - **Dead Code Elimination**: Remove code that will never execute (e.g., `if (false) { fmt.Println("test") }`).
     - **Nil Check Removal**: Remove unnecessary nil checks that the compiler can prove will never fail.

4. **Intermediate Representation**:
   - SSA is an intermediate representation closer to the final machine code, facilitating easier optimization.

To show the generated SSA,
* set the GOSSAFUNC environment variable to the function we would like to view the SSA of.
  * in this case main.
  * We will also need to pass the -S flag to the compiler, so it will print the code and create an HTML file.
  * We will also compile the file for Linux 64-bit,
  * to compile the file we will run:
  ```bash
  $ GOSSAFUNC=main GOOS=linux GOARCH=amd64 go build -gcflags "-S" <filename>.go
  ```
It will print all SSA, but it will also generate a ssa.html file which is interactive so we will use that.
* Lines that are greyed out means they might be removed in different phases.

**MAchine code stdout**
```bash
go-lang-programming/compilation-process/machine_code$ GOSSAFUNC=main GOOS=linux GOARCH=amd64 go build -gcflags "-S" machine.go
# runtime
dumped SSA for main,1 to /home/mpunix/go-lang-programming/compilation-process/machine_code/ssa.html
# command-line-arguments
dumped SSA for main,1 to ./ssa.html
main.main STEXT size=83 args=0x0 locals=0x40 funcid=0x0 align=0x0
        0x0000 00000 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:5)   TEXT    main.main(SB), ABIInternal, $64-0
        0x0000 00000 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:5)   CMPQ    SP, 16(R14)
        0x0004 00004 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:5)   PCDATA  $0, $-2
        0x0004 00004 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:5)   JLS     76
        0x0006 00006 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:5)   PCDATA  $0, $-1
        0x0006 00006 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:5)   PUSHQ   BP
        0x0007 00007 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:5)   MOVQ    SP, BP
        0x000a 00010 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:5)   SUBQ    $56, SP
        0x000e 00014 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:5)   FUNCDATA        $0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
        0x000e 00014 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:5)   FUNCDATA        $1, gclocals·EaPwxsZ75yY1hHMVZLmk6g==(SB)
        0x000e 00014 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:5)   FUNCDATA        $2, main.main.stkobj(SB)
        0x000e 00014 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:6)   LEAQ    type:string(SB), DX
        0x0015 00021 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:6)   MOVQ    DX, main..autotmp_8+40(SP)
        0x001a 00026 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:6)   LEAQ    main..stmp_0(SB), DX
        0x0021 00033 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:6)   MOVQ    DX, main..autotmp_8+48(SP)
        0x0026 00038 (/snap/go/10698/src/fmt/print.go:314)      MOVQ    os.Stdout(SB), BX
        0x002d 00045 (<unknown line number>)    NOP
        0x002d 00045 (/snap/go/10698/src/fmt/print.go:314)      LEAQ    go:itab.*os.File,io.Writer(SB), AX
        0x0034 00052 (/snap/go/10698/src/fmt/print.go:314)      LEAQ    main..autotmp_8+40(SP), CX
        0x0039 00057 (/snap/go/10698/src/fmt/print.go:314)      MOVL    $1, DI
        0x003e 00062 (/snap/go/10698/src/fmt/print.go:314)      MOVQ    DI, SI
        0x0041 00065 (/snap/go/10698/src/fmt/print.go:314)      PCDATA  $1, $0
        0x0041 00065 (/snap/go/10698/src/fmt/print.go:314)      CALL    fmt.Fprintln(SB)
        0x0046 00070 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:7)   ADDQ    $56, SP
        0x004a 00074 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:7)   POPQ    BP
        0x004b 00075 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:7)   RET
        0x004c 00076 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:7)   NOP
        0x004c 00076 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:5)   PCDATA  $1, $-1
        0x004c 00076 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:5)   PCDATA  $0, $-2
        0x004c 00076 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:5)   CALL    runtime.morestack_noctxt(SB)
        0x0051 00081 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:5)   PCDATA  $0, $-1
        0x0051 00081 (/home/mpunix/go-lang-programming/compilation-process/machine_code/machine.go:5)   JMP     0
        0x0000 49 3b 66 10 76 46 55 48 89 e5 48 83 ec 38 48 8d  I;f.vFUH..H..8H.
        0x0010 15 00 00 00 00 48 89 54 24 28 48 8d 15 00 00 00  .....H.T$(H.....
        0x0020 00 48 89 54 24 30 48 8b 1d 00 00 00 00 48 8d 05  .H.T$0H......H..
        0x0030 00 00 00 00 48 8d 4c 24 28 bf 01 00 00 00 48 89  ....H.L$(.....H.
        0x0040 fe e8 00 00 00 00 48 83 c4 38 5d c3 e8 00 00 00  ......H..8].....
        0x0050 00 eb ad                                         ...
        rel 2+0 t=R_USEIFACE type:string+0
        rel 2+0 t=R_USEIFACE type:*os.File+0
        rel 17+4 t=R_PCREL type:string+0
        rel 29+4 t=R_PCREL main..stmp_0+0
        rel 41+4 t=R_PCREL os.Stdout+0
        rel 48+4 t=R_PCREL go:itab.*os.File,io.Writer+0
        rel 66+4 t=R_CALL fmt.Fprintln+0
        rel 77+4 t=R_CALL runtime.morestack_noctxt+0
type:.eq.sync/atomic.Pointer[os.dirInfo] STEXT dupok nosplit size=10 args=0x10 locals=0x0 funcid=0x0 align=0x0
        0x0000 00000 (<autogenerated>:1)        TEXT    type:.eq.sync/atomic.Pointer[os.dirInfo](SB), DUPOK|NOSPLIT|NOFRAME|ABIInternal, $0-16
        0x0000 00000 (<autogenerated>:1)        FUNCDATA        $0, gclocals·TjPuuCwdlCpTaRQGRKTrYw==(SB)
        0x0000 00000 (<autogenerated>:1)        FUNCDATA        $1, gclocals·J5F+7Qw7O7ve2QcWC7DpeQ==(SB)
        0x0000 00000 (<autogenerated>:1)        FUNCDATA        $5, type:.eq.sync/atomic.Pointer[os.dirInfo].arginfo1(SB)
        0x0000 00000 (<autogenerated>:1)        FUNCDATA        $6, type:.eq.sync/atomic.Pointer[os.dirInfo].argliveinfo(SB)
        0x0000 00000 (<autogenerated>:1)        PCDATA  $3, $1
        0x0000 00000 (<autogenerated>:1)        MOVQ    (AX), CX
        0x0003 00003 (<autogenerated>:1)        CMPQ    (BX), CX
        0x0006 00006 (<autogenerated>:1)        SETEQ   AL
        0x0009 00009 (<autogenerated>:1)        RET
        0x0000 48 8b 08 48 39 0b 0f 94 c0 c3                    H..H9.....
go:cuinfo.producer.main SDWARFCUINFO dupok size=0
        0x0000 72 65 67 61 62 69                                regabi
go:cuinfo.packagename.main SDWARFCUINFO dupok size=0
        0x0000 6d 61 69 6e                                      main
go:info.fmt.Println$abstract SDWARFABSFCN dupok size=44
        0x0000 05 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 b9 02  .fmt.Println....
        0x0010 01 21 61 00 00 00 00 00 00 21 6e 00 01 00 00 00  .!a......!n.....
        0x0020 00 21 65 72 72 00 01 00 00 00 00 00              .!err.......
        rel 0+0 t=R_USETYPE type:[]interface {}+0
        rel 0+0 t=R_USETYPE type:error+0
        rel 0+0 t=R_USETYPE type:int+0
        rel 21+4 t=R_DWARFSECREF go:info.[]interface {}+0
        rel 29+4 t=R_DWARFSECREF go:info.int+0
        rel 39+4 t=R_DWARFSECREF go:info.error+0
go:itab.*os.File,io.Writer SRODATA dupok size=32
        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 5a 22 ee 60 00 00 00 00 00 00 00 00 00 00 00 00  Z".`............
        rel 0+8 t=R_ADDR type:io.Writer+0
        rel 8+8 t=R_ADDR type:*os.File+0
        rel 24+8 t=RelocType(-32767) os.(*File).Write+0
sync/atomic..dict.Pointer[os.dirInfo] SRODATA dupok size=128
        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0050 00 00 00 00 00 00 00 00                          ........
        rel 0+0 t=R_USEIFACE type:*os.dirInfo+0
        rel 0+0 t=R_USEIFACE type:*sync/atomic.Pointer[os.dirInfo]+0
        rel 0+0 t=R_USEIFACE type:*os.dirInfo+0
        rel 0+0 t=R_USEIFACE type:*os.dirInfo+0
        rel 0+0 t=R_USEIFACE type:*os.dirInfo+0
        rel 0+0 t=R_USEIFACE type:*os.dirInfo+0
        rel 0+0 t=R_USEIFACE type:*os.dirInfo+0
        rel 0+0 t=R_USEIFACE type:unsafe.Pointer+0
        rel 0+8 t=R_ADDR type:*sync/atomic.Pointer[os.dirInfo]+0
        rel 0+0 t=R_USEIFACE type:*os.dirInfo+0
        rel 0+0 t=R_USEIFACE type:unsafe.Pointer+0
        rel 0+0 t=R_USEIFACE type:*sync/atomic.Pointer[os.dirInfo]+0
        rel 0+0 t=R_USEIFACE type:*os.dirInfo+0
        rel 0+0 t=R_USEIFACE type:unsafe.Pointer+0
        rel 0+0 t=R_USEIFACE type:*sync/atomic.Pointer[os.dirInfo]+0
        rel 0+0 t=R_USEIFACE type:*os.dirInfo+0
        rel 0+0 t=R_USEIFACE type:*os.dirInfo+0
        rel 0+0 t=R_USEIFACE type:*os.dirInfo+0
        rel 0+0 t=R_USEIFACE type:unsafe.Pointer+0
        rel 0+0 t=R_USEIFACE type:*sync/atomic.Pointer[os.dirInfo]+0
        rel 0+0 t=R_USEIFACE type:unsafe.Pointer+0
        rel 0+0 t=R_USEIFACE type:*os.dirInfo+0
        rel 8+8 t=R_ADDR type:*os.dirInfo+0
        rel 16+8 t=R_ADDR type:*os.dirInfo+0
        rel 24+8 t=R_ADDR type:*sync/atomic.Pointer[os.dirInfo]+0
        rel 32+8 t=R_ADDR type:*os.dirInfo+0
        rel 40+8 t=R_ADDR type:*sync/atomic.Pointer[os.dirInfo]+0
        rel 48+8 t=R_ADDR type:*os.dirInfo+0
        rel 56+8 t=R_ADDR type:*os.dirInfo+0
        rel 64+8 t=R_ADDR type:*os.dirInfo+0
        rel 72+8 t=R_ADDR type:*sync/atomic.Pointer[os.dirInfo]+0
        rel 80+8 t=R_ADDR type:*os.dirInfo+0
main..inittask SNOPTRDATA size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+0 t=R_INITORDER fmt..inittask+0
runtime.memequal64·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=R_ADDR runtime.memequal64+0
runtime.gcbits.0100000000000000 SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
type:.namedata.*atomic.Pointer[os.dirInfo]. SRODATA dupok size=29
        0x0000 01 1b 2a 61 74 6f 6d 69 63 2e 50 6f 69 6e 74 65  ..*atomic.Pointe
        0x0010 72 5b 6f 73 2e 64 69 72 49 6e 66 6f 5d           r[os.dirInfo]
type:.eqfunc.sync/atomic.Pointer[os.dirInfo] SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=R_ADDR type:.eq.sync/atomic.Pointer[os.dirInfo]+0
runtime.memequal0·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=R_ADDR runtime.memequal0+0
type:.namedata.*[0]*os.dirInfo- SRODATA dupok size=17
        0x0000 00 0f 2a 5b 30 5d 2a 6f 73 2e 64 69 72 49 6e 66  ..*[0]*os.dirInf
        0x0010 6f                                               o
type:*[0]*os.dirInfo SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 c6 0a ea a1 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=R_ADDR runtime.memequal64·f+0
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*[0]*os.dirInfo-+0
        rel 48+8 t=R_ADDR type:[0]*os.dirInfo+0
runtime.gcbits. SRODATA dupok size=0
type:.namedata.*[]*os.dirInfo- SRODATA dupok size=16
        0x0000 00 0e 2a 5b 5d 2a 6f 73 2e 64 69 72 49 6e 66 6f  ..*[]*os.dirInfo
type:*[]*os.dirInfo SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 31 df 2b 6e 08 08 08 36 00 00 00 00 00 00 00 00  1.+n...6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=R_ADDR runtime.memequal64·f+0
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*[]*os.dirInfo-+0
        rel 48+8 t=R_ADDR type:[]*os.dirInfo+0
type:[]*os.dirInfo SRODATA dupok size=56
        0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 16 77 13 b1 02 08 08 17 00 00 00 00 00 00 00 00  .w..............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*[]*os.dirInfo-+0
        rel 44+4 t=RelocType(-32763) type:*[]*os.dirInfo+0
        rel 48+8 t=R_ADDR type:*os.dirInfo+0
type:[0]*os.dirInfo SRODATA dupok size=72
        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 e5 80 e9 79 0a 08 08 11 00 00 00 00 00 00 00 00  ...y............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=R_ADDR runtime.memequal0·f+0
        rel 32+8 t=R_ADDR runtime.gcbits.+0
        rel 40+4 t=R_ADDROFF type:.namedata.*[0]*os.dirInfo-+0
        rel 44+4 t=RelocType(-32763) type:*[0]*os.dirInfo+0
        rel 48+8 t=R_ADDR type:*os.dirInfo+0
        rel 56+8 t=R_ADDR type:[]*os.dirInfo+0
type:.importpath.sync/atomic. SRODATA dupok size=13
        0x0000 00 0b 73 79 6e 63 2f 61 74 6f 6d 69 63           ..sync/atomic
type:.namedata._- SRODATA dupok size=3
        0x0000 00 01 5f                                         .._
type:.namedata.v- SRODATA dupok size=3
        0x0000 00 01 76                                         ..v
type:sync/atomic.Pointer[os.dirInfo] SRODATA dupok size=168
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 76 69 31 3d 07 08 08 19 00 00 00 00 00 00 00 00  vi1=............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 03 00 00 00 00 00 00 00 03 00 00 00 00 00 00 00  ................
        0x0050 00 00 00 00 00 00 00 00 58 00 00 00 00 00 00 00  ........X.......
        0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0090 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x00a0 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=R_ADDR type:.eqfunc.sync/atomic.Pointer[os.dirInfo]+0
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*atomic.Pointer[os.dirInfo].+0
        rel 44+4 t=R_ADDROFF type:*sync/atomic.Pointer[os.dirInfo]+0
        rel 48+8 t=R_ADDR type:.importpath.sync/atomic.+0
        rel 56+8 t=R_ADDR type:sync/atomic.Pointer[os.dirInfo]+96
        rel 80+4 t=R_ADDROFF type:.importpath.sync/atomic.+0
        rel 96+8 t=R_ADDR type:.namedata._-+0
        rel 104+8 t=R_ADDR type:[0]*os.dirInfo+0
        rel 120+8 t=R_ADDR type:.namedata._-+0
        rel 128+8 t=R_ADDR type:sync/atomic.noCopy+0
        rel 144+8 t=R_ADDR type:.namedata.v-+0
        rel 152+8 t=R_ADDR type:unsafe.Pointer+0
type:.namedata.*func(*atomic.Pointer[os.dirInfo], *os.dirInfo, *os.dirInfo) bool- SRODATA dupok size=67
        0x0000 00 41 2a 66 75 6e 63 28 2a 61 74 6f 6d 69 63 2e  .A*func(*atomic.
        0x0010 50 6f 69 6e 74 65 72 5b 6f 73 2e 64 69 72 49 6e  Pointer[os.dirIn
        0x0020 66 6f 5d 2c 20 2a 6f 73 2e 64 69 72 49 6e 66 6f  fo], *os.dirInfo
        0x0030 2c 20 2a 6f 73 2e 64 69 72 49 6e 66 6f 29 20 62  , *os.dirInfo) b
        0x0040 6f 6f 6c                                         ool
type:*func(*sync/atomic.Pointer[os.dirInfo], *os.dirInfo, *os.dirInfo) bool SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 63 08 c0 ac 08 08 08 36 00 00 00 00 00 00 00 00  c......6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=R_ADDR runtime.memequal64·f+0
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*func(*atomic.Pointer[os.dirInfo], *os.dirInfo, *os.dirInfo) bool-+0
        rel 48+8 t=R_ADDR type:func(*sync/atomic.Pointer[os.dirInfo], *os.dirInfo, *os.dirInfo) bool+0
type:func(*sync/atomic.Pointer[os.dirInfo], *os.dirInfo, *os.dirInfo) bool SRODATA dupok size=88
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 0d 10 62 e1 02 08 08 33 00 00 00 00 00 00 00 00  ..b....3........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 03 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0050 00 00 00 00 00 00 00 00                          ........
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*func(*atomic.Pointer[os.dirInfo], *os.dirInfo, *os.dirInfo) bool-+0
        rel 44+4 t=RelocType(-32763) type:*func(*sync/atomic.Pointer[os.dirInfo], *os.dirInfo, *os.dirInfo) bool+0
        rel 56+8 t=R_ADDR type:*sync/atomic.Pointer[os.dirInfo]+0
        rel 64+8 t=R_ADDR type:*os.dirInfo+0
        rel 72+8 t=R_ADDR type:*os.dirInfo+0
        rel 80+8 t=R_ADDR type:bool+0
type:.namedata.*func(*atomic.Pointer[os.dirInfo]) *os.dirInfo- SRODATA dupok size=48
        0x0000 00 2e 2a 66 75 6e 63 28 2a 61 74 6f 6d 69 63 2e  ..*func(*atomic.
        0x0010 50 6f 69 6e 74 65 72 5b 6f 73 2e 64 69 72 49 6e  Pointer[os.dirIn
        0x0020 66 6f 5d 29 20 2a 6f 73 2e 64 69 72 49 6e 66 6f  fo]) *os.dirInfo
type:*func(*sync/atomic.Pointer[os.dirInfo]) *os.dirInfo SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 8f 1f c8 72 08 08 08 36 00 00 00 00 00 00 00 00  ...r...6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=R_ADDR runtime.memequal64·f+0
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*func(*atomic.Pointer[os.dirInfo]) *os.dirInfo-+0
        rel 48+8 t=R_ADDR type:func(*sync/atomic.Pointer[os.dirInfo]) *os.dirInfo+0
type:func(*sync/atomic.Pointer[os.dirInfo]) *os.dirInfo SRODATA dupok size=72
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 8b 8e aa d0 02 08 08 33 00 00 00 00 00 00 00 00  .......3........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 01 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 00 00 00 00 00 00 00 00                          ........
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*func(*atomic.Pointer[os.dirInfo]) *os.dirInfo-+0
        rel 44+4 t=RelocType(-32763) type:*func(*sync/atomic.Pointer[os.dirInfo]) *os.dirInfo+0
        rel 56+8 t=R_ADDR type:*sync/atomic.Pointer[os.dirInfo]+0
        rel 64+8 t=R_ADDR type:*os.dirInfo+0
type:.namedata.*func(*atomic.Pointer[os.dirInfo], *os.dirInfo)- SRODATA dupok size=49
        0x0000 00 2f 2a 66 75 6e 63 28 2a 61 74 6f 6d 69 63 2e  ./*func(*atomic.
        0x0010 50 6f 69 6e 74 65 72 5b 6f 73 2e 64 69 72 49 6e  Pointer[os.dirIn
        0x0020 66 6f 5d 2c 20 2a 6f 73 2e 64 69 72 49 6e 66 6f  fo], *os.dirInfo
        0x0030 29                                               )
type:*func(*sync/atomic.Pointer[os.dirInfo], *os.dirInfo) SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 24 28 03 fb 08 08 08 36 00 00 00 00 00 00 00 00  $(.....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=R_ADDR runtime.memequal64·f+0
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*func(*atomic.Pointer[os.dirInfo], *os.dirInfo)-+0
        rel 48+8 t=R_ADDR type:func(*sync/atomic.Pointer[os.dirInfo], *os.dirInfo)+0
type:func(*sync/atomic.Pointer[os.dirInfo], *os.dirInfo) SRODATA dupok size=72
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 15 bf 46 19 02 08 08 33 00 00 00 00 00 00 00 00  ..F....3........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 02 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 00 00 00 00 00 00 00 00                          ........
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*func(*atomic.Pointer[os.dirInfo], *os.dirInfo)-+0
        rel 44+4 t=RelocType(-32763) type:*func(*sync/atomic.Pointer[os.dirInfo], *os.dirInfo)+0
        rel 56+8 t=R_ADDR type:*sync/atomic.Pointer[os.dirInfo]+0
        rel 64+8 t=R_ADDR type:*os.dirInfo+0
type:.namedata.*func(*atomic.Pointer[os.dirInfo], *os.dirInfo) *os.dirInfo- SRODATA dupok size=61
        0x0000 00 3b 2a 66 75 6e 63 28 2a 61 74 6f 6d 69 63 2e  .;*func(*atomic.
        0x0010 50 6f 69 6e 74 65 72 5b 6f 73 2e 64 69 72 49 6e  Pointer[os.dirIn
        0x0020 66 6f 5d 2c 20 2a 6f 73 2e 64 69 72 49 6e 66 6f  fo], *os.dirInfo
        0x0030 29 20 2a 6f 73 2e 64 69 72 49 6e 66 6f           ) *os.dirInfo
type:*func(*sync/atomic.Pointer[os.dirInfo], *os.dirInfo) *os.dirInfo SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 85 c7 60 3d 08 08 08 36 00 00 00 00 00 00 00 00  ..`=...6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=R_ADDR runtime.memequal64·f+0
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*func(*atomic.Pointer[os.dirInfo], *os.dirInfo) *os.dirInfo-+0
        rel 48+8 t=R_ADDR type:func(*sync/atomic.Pointer[os.dirInfo], *os.dirInfo) *os.dirInfo+0
type:func(*sync/atomic.Pointer[os.dirInfo], *os.dirInfo) *os.dirInfo SRODATA dupok size=80
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 bc a5 28 74 02 08 08 33 00 00 00 00 00 00 00 00  ..(t...3........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 02 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*func(*atomic.Pointer[os.dirInfo], *os.dirInfo) *os.dirInfo-+0
        rel 44+4 t=RelocType(-32763) type:*func(*sync/atomic.Pointer[os.dirInfo], *os.dirInfo) *os.dirInfo+0
        rel 56+8 t=R_ADDR type:*sync/atomic.Pointer[os.dirInfo]+0
        rel 64+8 t=R_ADDR type:*os.dirInfo+0
        rel 72+8 t=R_ADDR type:*os.dirInfo+0
type:.namedata.CompareAndSwap. SRODATA dupok size=16
        0x0000 01 0e 43 6f 6d 70 61 72 65 41 6e 64 53 77 61 70  ..CompareAndSwap
type:.namedata.*func(*os.dirInfo, *os.dirInfo) bool- SRODATA dupok size=38
        0x0000 00 24 2a 66 75 6e 63 28 2a 6f 73 2e 64 69 72 49  .$*func(*os.dirI
        0x0010 6e 66 6f 2c 20 2a 6f 73 2e 64 69 72 49 6e 66 6f  nfo, *os.dirInfo
        0x0020 29 20 62 6f 6f 6c                                ) bool
type:*func(*os.dirInfo, *os.dirInfo) bool SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 6e 1c 7b 9d 08 08 08 36 00 00 00 00 00 00 00 00  n.{....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=R_ADDR runtime.memequal64·f+0
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*func(*os.dirInfo, *os.dirInfo) bool-+0
        rel 48+8 t=R_ADDR type:func(*os.dirInfo, *os.dirInfo) bool+0
type:func(*os.dirInfo, *os.dirInfo) bool SRODATA dupok size=80
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 34 6e 96 f6 02 08 08 33 00 00 00 00 00 00 00 00  4n.....3........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 02 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*func(*os.dirInfo, *os.dirInfo) bool-+0
        rel 44+4 t=RelocType(-32763) type:*func(*os.dirInfo, *os.dirInfo) bool+0
        rel 56+8 t=R_ADDR type:*os.dirInfo+0
        rel 64+8 t=R_ADDR type:*os.dirInfo+0
        rel 72+8 t=R_ADDR type:bool+0
type:.namedata.Load. SRODATA dupok size=6
        0x0000 01 04 4c 6f 61 64                                ..Load
type:.namedata.*func() *os.dirInfo- SRODATA dupok size=21
        0x0000 00 13 2a 66 75 6e 63 28 29 20 2a 6f 73 2e 64 69  ..*func() *os.di
        0x0010 72 49 6e 66 6f                                   rInfo
type:*func() *os.dirInfo SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 15 d3 0f 9f 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=R_ADDR runtime.memequal64·f+0
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*func() *os.dirInfo-+0
        rel 48+8 t=R_ADDR type:func() *os.dirInfo+0
type:func() *os.dirInfo SRODATA dupok size=64
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 77 0c bb 3a 02 08 08 33 00 00 00 00 00 00 00 00  w..:...3........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*func() *os.dirInfo-+0
        rel 44+4 t=RelocType(-32763) type:*func() *os.dirInfo+0
        rel 56+8 t=R_ADDR type:*os.dirInfo+0
type:.namedata.Store. SRODATA dupok size=7
        0x0000 01 05 53 74 6f 72 65                             ..Store
type:.namedata.*func(*os.dirInfo)- SRODATA dupok size=20
        0x0000 00 12 2a 66 75 6e 63 28 2a 6f 73 2e 64 69 72 49  ..*func(*os.dirI
        0x0010 6e 66 6f 29                                      nfo)
type:*func(*os.dirInfo) SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 0f 30 20 57 08 08 08 36 00 00 00 00 00 00 00 00  .0 W...6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=R_ADDR runtime.memequal64·f+0
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*func(*os.dirInfo)-+0
        rel 48+8 t=R_ADDR type:func(*os.dirInfo)+0
type:func(*os.dirInfo) SRODATA dupok size=64
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 53 df 95 59 02 08 08 33 00 00 00 00 00 00 00 00  S..Y...3........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*func(*os.dirInfo)-+0
        rel 44+4 t=RelocType(-32763) type:*func(*os.dirInfo)+0
        rel 56+8 t=R_ADDR type:*os.dirInfo+0
type:.namedata.Swap. SRODATA dupok size=6
        0x0000 01 04 53 77 61 70                                ..Swap
type:.namedata.*func(*os.dirInfo) *os.dirInfo- SRODATA dupok size=32
        0x0000 00 1e 2a 66 75 6e 63 28 2a 6f 73 2e 64 69 72 49  ..*func(*os.dirI
        0x0010 6e 66 6f 29 20 2a 6f 73 2e 64 69 72 49 6e 66 6f  nfo) *os.dirInfo
type:*func(*os.dirInfo) *os.dirInfo SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 9c 4f e0 c8 08 08 08 36 00 00 00 00 00 00 00 00  .O.....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=R_ADDR runtime.memequal64·f+0
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*func(*os.dirInfo) *os.dirInfo-+0
        rel 48+8 t=R_ADDR type:func(*os.dirInfo) *os.dirInfo+0
type:func(*os.dirInfo) *os.dirInfo SRODATA dupok size=72
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 7a 84 08 95 02 08 08 33 00 00 00 00 00 00 00 00  z......3........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 01 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 00 00 00 00 00 00 00 00                          ........
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*func(*os.dirInfo) *os.dirInfo-+0
        rel 44+4 t=RelocType(-32763) type:*func(*os.dirInfo) *os.dirInfo+0
        rel 56+8 t=R_ADDR type:*os.dirInfo+0
        rel 64+8 t=R_ADDR type:*os.dirInfo+0
type:*sync/atomic.Pointer[os.dirInfo] SRODATA dupok size=136
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 ae 00 26 16 09 08 08 36 00 00 00 00 00 00 00 00  ..&....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 04 00 04 00  ................
        0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0080 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=R_ADDR runtime.memequal64·f+0
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*atomic.Pointer[os.dirInfo].+0
        rel 48+8 t=R_ADDR type:sync/atomic.Pointer[os.dirInfo]+0
        rel 56+4 t=R_ADDROFF type:.importpath.sync/atomic.+0
        rel 72+4 t=R_ADDROFF type:.namedata.CompareAndSwap.+0
        rel 76+4 t=R_METHODOFF type:func(*os.dirInfo, *os.dirInfo) bool+0
        rel 80+4 t=R_METHODOFF sync/atomic.(*Pointer[os.dirInfo]).CompareAndSwap+0
        rel 84+4 t=R_METHODOFF sync/atomic.(*Pointer[os.dirInfo]).CompareAndSwap+0
        rel 88+4 t=R_ADDROFF type:.namedata.Load.+0
        rel 92+4 t=R_METHODOFF type:func() *os.dirInfo+0
        rel 96+4 t=R_METHODOFF sync/atomic.(*Pointer[os.dirInfo]).Load+0
        rel 100+4 t=R_METHODOFF sync/atomic.(*Pointer[os.dirInfo]).Load+0
        rel 104+4 t=R_ADDROFF type:.namedata.Store.+0
        rel 108+4 t=R_METHODOFF type:func(*os.dirInfo)+0
        rel 112+4 t=R_METHODOFF sync/atomic.(*Pointer[os.dirInfo]).Store+0
        rel 116+4 t=R_METHODOFF sync/atomic.(*Pointer[os.dirInfo]).Store+0
        rel 120+4 t=R_ADDROFF type:.namedata.Swap.+0
        rel 124+4 t=R_METHODOFF type:func(*os.dirInfo) *os.dirInfo+0
        rel 128+4 t=R_METHODOFF sync/atomic.(*Pointer[os.dirInfo]).Swap+0
        rel 132+4 t=R_METHODOFF sync/atomic.(*Pointer[os.dirInfo]).Swap+0
go:string."Hello, world!" SRODATA dupok size=13
        0x0000 48 65 6c 6c 6f 2c 20 77 6f 72 6c 64 21           Hello, world!
main..stmp_0 SRODATA static size=16
        0x0000 00 00 00 00 00 00 00 00 0d 00 00 00 00 00 00 00  ................
        rel 0+8 t=R_ADDR go:string."Hello, world!"+0
runtime.nilinterequal·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=R_ADDR runtime.nilinterequal+0
type:.namedata.*[1]interface {}- SRODATA dupok size=18
        0x0000 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65 20  ..*[1]interface
        0x0010 7b 7d                                            {}
type:*[1]interface {} SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 a8 0e 57 36 08 08 08 36 00 00 00 00 00 00 00 00  ..W6...6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=R_ADDR runtime.memequal64·f+0
        rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*[1]interface {}-+0
        rel 48+8 t=R_ADDR type:[1]interface {}+0
runtime.gcbits.0200000000000000 SRODATA dupok size=8
        0x0000 02 00 00 00 00 00 00 00                          ........
type:[1]interface {} SRODATA dupok size=72
        0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
        0x0010 6e 20 6a 3d 02 08 08 11 00 00 00 00 00 00 00 00  n j=............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 01 00 00 00 00 00 00 00                          ........
        rel 24+8 t=R_ADDR runtime.nilinterequal·f+0
        rel 32+8 t=R_ADDR runtime.gcbits.0200000000000000+0
        rel 40+4 t=R_ADDROFF type:.namedata.*[1]interface {}-+0
        rel 44+4 t=RelocType(-32763) type:*[1]interface {}+0
        rel 48+8 t=R_ADDR type:interface {}+0
        rel 56+8 t=R_ADDR type:[]interface {}+0
gclocals·g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·EaPwxsZ75yY1hHMVZLmk6g== SRODATA dupok size=9
        0x0000 01 00 00 00 02 00 00 00 00                       .........
main.main.stkobj SRODATA static size=24
        0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
        0x0010 10 00 00 00 00 00 00 00                          ........
        rel 20+4 t=R_ADDROFF runtime.gcbits.0200000000000000+0
gclocals·TjPuuCwdlCpTaRQGRKTrYw== SRODATA dupok size=10
        0x0000 02 00 00 00 02 00 00 00 03 00                    ..........
gclocals·J5F+7Qw7O7ve2QcWC7DpeQ== SRODATA dupok size=8
        0x0000 02 00 00 00 00 00 00 00                          ........
type:.eq.sync/atomic.Pointer[os.dirInfo].arginfo1 SRODATA static dupok size=3
        0x0000 08 08 ff                                         ...
type:.eq.sync/atomic.Pointer[os.dirInfo].argliveinfo SRODATA static dupok size=2
        0x0000 00 00
```
