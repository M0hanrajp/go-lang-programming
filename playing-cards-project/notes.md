**Error: undefined: <function name>**

```bash
go run basics/playing-cards-project/main.go
# command-line-arguments
basics/playing-cards-project/main.go:19:2: undefined: OtherMain
```
The issue you're encountering is likely due to how Go handles package imports and file organization within a project. Even though both `main.go` and `project.go` are part of the same package (`main`), Go only compiles files that are explicitly included in the build or run command. When you run `go run basics/playing-cards-project/main.go`, only `main.go` is compiled, and it does not see the `OtherMain` function from `project.go`.

### Solution:
To resolve this, you should run all the Go files in the directory together. You can do this by running the following command:

```bash
go run basics/playing-cards-project/*.go
```

This command will compile and run both `main.go` and `project.go`, allowing the `main.go` file to access the `OtherMain` function from `project.go`.

**How to import the project.go as a pacakge in main.go**

**Using Go Modules**

Go modules allow you to manage dependencies and organize packages in a better way. To set up a Go module for your project:

1. Navigate to the root of your project directory (e.g., `~/go-lang-programming`).
2. Initialize a Go module by running:

   ```bash
   go mod init playing-cards-project
   ```
   This will create a `go.mod` file, defining your project as a module.

3. Then, organize the code in sub-packages if needed. For instance, you could put the `OtherMain` function into a separate package instead of keeping everything in the `main` package.

4. Move `project.go` into a subfolder/package, such as `cards`:
   
   ```
   playing-cards-project/
   ├── main.go
   └── cards/
       └── project.go
   ```

5. Modify `project.go` to be part of the `cards` package:

   ```go
   // basics/playing-cards-project/cards/project.go
   package cards

   import "fmt"

    < program body >
   ```

6. In `main.go`, import the `cards` package:

   ```go
   package main

   import (
       "fmt"
       "playing-cards-project/cards"
   )

   func main() {
       fmt.Println("Commented code execution:")
       cards.OtherMain()
   }
   ```

7. Now, you can run the project by simply executing:

   ```bash
   go run main.go
   ```

### 2. **Using `go build`**
Another way to ensure both files are compiled without using wildcards is to first build the project:

1. Run the `go build` command in the project directory:

   ```bash
   go build basics/playing-cards-project/
   ```

2. After building the project, you can run the generated binary, which includes both files:

   ```bash
   ./playing-cards-project
   ```

This method compiles all `.go` files in the directory, ensuring they are all part of the build process.
