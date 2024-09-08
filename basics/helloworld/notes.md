## Section 2 a simple start

**Lecture 7 & 8**

1. **Five Key Questions**:
   - How to run the code inside the project.
   - The meaning of the first line (`package main`).
   - The purpose of the `import "fmt"` line.
   - Understanding the `func` keyword.
   - The general pattern for organizing Go code.

2. **Running the Code**:
   - Navigate to the project directory in the terminal.
   - Use the `go run main.go` command to compile and execute the program.
   - The `go` command is a portal to working with Go on the local machine.

3. **Go Command Line Interface (CLI)**:
   - **`go run`**: Compiles and immediately executes the program.
   - **`go build`**: Compiles the program without executing it, creating an executable file.
   - **`go fmt`**: Automatically formats the code.
   - **`go install` and `go get`**: Handle dependencies in the project.
   - **`go test`**: Runs and executes test files associated with the project.

4. **Difference Between `go run` and `go build`**:
   - `go run` compiles and executes the program immediately.
   - `go build` only compiles the program, creating an executable file that can be run later.

5. **Practical Example**:
   - Running `go build main.go` creates an executable file (`main` or `main.exe`).
   - The executable can be run using `./main` on Mac/Linux or `main.exe` on Windows.

```bash
mpunix@LIN-MP22QN1X:~/go-lang-programming/basics/hellworld$ go version
go version go1.23.0 linux/amd64
mpunix@LIN-MP22QN1X:~/go-lang-programming/basics/hellworld$ go run main.go
Hello GO programming!
what does this do ?
mpunix@LIN-MP22QN1X:~/go-lang-programming/basics/hellworld$ go build main.go
mpunix@LIN-MP22QN1X:~/go-lang-programming/basics/hellworld$ ls -l
total 2084
-rwxr-xr-x 1 mpunix mpunix 2129445 Sep  6 00:08 main
-rw-r--r-- 1 mpunix mpunix     118 Sep  6 00:02 main.go
mpunix@LIN-MP22QN1X:~/go-lang-programming/basics/hellworld$ ./main
Hello GO programming!
what does this do ?
```

**Lecture 9**

1. **Understanding `package main`**:
   - **Package Concept**: A package in Go is like a project or workspace, a collection of related source code files.
   - **Declaration Requirement**: Every file in a package must declare the package it belongs to at the top.

2. **Types of Packages**:
   - **Executable Packages**: These are compiled into runnable files (e.g., `main` or `main.exe`). They are used for creating applications that perform tasks.
   - **Reusable Packages**: These contain reusable code, such as libraries or dependencies, and are not directly executable.

3. **Naming the Package**:
   - **`main` Package**: Using `package main` indicates that the package is executable. This is why the `go build` command creates an executable file when the package is named `main`.
   - **Other Names**: Any other package name (e.g., `package apple`) will not produce an executable file when built.

4. **Function `main`**:
   - An executable package must contain a `main` function, which serves as the entry point for the application.

5. **Practical Demonstration**:
   - Changing the package name to something other than `main` and running `go build` will not produce an executable file.
   - Reverting the package name back to `main` and building the project again will produce the executable.

```bash
~/go-lang-programming/basics/hellworld$ cat main.go
/*
Based on the type of program we are developing, we might generate
executable ( executable is created for the whole application ) or reusable ( Code dependencies, helpers, reused code )
The current program creates an executable
Name of the package determines executable or reusable or dependency type package, main is used for executable here.
*/
package hellworld

import "fmt"

// For each executable package created there must be a func main()
func main() {
        fmt.Println("Hello GO programming!")
        fmt.Print("what does this do ?\n")
}
mpunix@LIN-MP22QN1X:~/go-lang-programming/basics/hellworld$ go build main.go
mpunix@LIN-MP22QN1X:~/go-lang-programming/basics/hellworld$ ls -al
total 16
drwxr-xr-x 2 mpunix mpunix 4096 Sep  6 00:42 .
drwxr-xr-x 3 mpunix mpunix 4096 Sep  5 23:45 ..
-rw-r--r-- 1 mpunix mpunix 3570 Sep  6 00:42 Worklog.md
-rw-r--r-- 1 mpunix mpunix  539 Sep  6 00:42 main.go
## Changed the main.go at this point
mpunix@LIN-MP22QN1X:~/go-lang-programming/basics/hellworld$ go build main.go
mpunix@LIN-MP22QN1X:~/go-lang-programming/basics/hellworld$ ls -al
total 2096
drwxr-xr-x 2 mpunix mpunix    4096 Sep  6 00:43 .
drwxr-xr-x 3 mpunix mpunix    4096 Sep  5 23:45 ..
-rw-r--r-- 1 mpunix mpunix    3570 Sep  6 00:42 Worklog.md
-rwxr-xr-x 1 mpunix mpunix 2129453 Sep  6 00:43 main
-rw-r--r-- 1 mpunix mpunix     534 Sep  6 00:43 main.go
```
6. **Summary**:
   - **`package main`**: Indicates an executable package.
   - **Other Package Names**: Indicate reusable or dependency packages.
   - **Executable Package Requirement**: Must have a `main` function.

**Lecture 10 & 11**

1. **Understanding `import fmt`**:
   - **Purpose of `import`**: The `import` statement is used to give the current package access to code from another package.
   - **`fmt` Package**: `fmt` is a standard library package in Go, short for "format." It provides functions for formatting and printing text to the terminal, useful for debugging and output.

2. **How Packages Work Together**:
   - **Main Package**: The central package in a project.
   - **Standard Library Packages**: Other packages included with Go by default (e.g., `fmt`, `math`).
   - **Importing Packages**: To use code from another package, you must explicitly import it using the `import` statement (e.g., `import fmt`).

3. **Using External Packages**:
   - **Beyond Standard Library**: You can also import packages authored by other engineers (e.g., `calculator`, `uploader`).
   - **Reusable Packages**: These are packages that provide reusable code or libraries.

4. **Documentation**:
   - **Go Standard Library Documentation**: Available at [golang.org/pkg](https://golang.org/pkg).
   - **Importance of Documentation**: The instructor emphasizes the importance of referring to the official documentation to understand and use standard library packages effectively.

5. **Practical Example**:
   - **Importing `fmt`**: Forms a link from the main package to the `fmt` package, allowing the use of its functions.
   - **Other Imports**: Similarly, you can import other standard or external packages as needed.

6. **Summary**:
   - **`import` Statement**: Used to include code from other packages.
   - **`fmt` Package**: Part of the standard library, used for formatting and printing text.
   - **Documentation**: Essential for learning and using Go's standard packages.

7. **Understanding `func main`**:
   - **Function Declaration**: Functions in Go are declared using the `func` keyword, followed by the function name, an argument list in parentheses, and a body enclosed in curly braces.
   - **Syntax**:
     ```go
     func functionName(arguments) {
         // function body
     }
     ```
   - **Example**:
     ```go
     func main() {
         fmt.Println("Hello, World!")
     }
     ```

8. **File Organization in Go**:
   - **Standard Pattern**: Every Go file follows a consistent pattern:
     1. **Package Declaration**: Declares the package the file belongs to.
        ```go
        package main
        ```
     2. **Import Statements**: Lists the packages needed for the file.
        ```go
        import "fmt"
        ```
     3. **Function and Logic**: Contains the main logic, functions, and variable declarations.
        ```go
        func main() {
            fmt.Println("Hello, World!")
        }
        ```

9. **Summary of Main Concepts**:
   - **Packages**: Group related files and must be declared at the top of each file.
   - **Imports**: Used to include code from other packages.
   - **Functions**: Declared with `func`, they encapsulate code to perform specific tasks.
   - **File Structure**: Consistently follows the pattern of package declaration, imports, and function definitions.
