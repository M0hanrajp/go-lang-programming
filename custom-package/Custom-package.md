### Making your own packages

```bash
mpunix@LIN-MP22QN1X:~/go-lang-programming/custom-package$ ls
main.go  mypackage
mpunix@LIN-MP22QN1X:~/go-lang-programming/custom-package$ go run main.go
main.go:3:8: package mypackage is not in std (/snap/go/10698/src/mypackage)
...
~/go-lang-programming/custom-package$ go mod init myproject
go: creating new go.mod: module myproject
go: to add module requirements and sums:
        go mod tidy
~/go-lang-programming/custom-package$ go run main.go
Hello from mypackage!
```

The `go.mod` file is crucial in Go projects for several reasons. It helps manage dependencies, define module paths, and ensure consistent builds. Here's how it contributes to your project:

### 1. **Module Path Definition**

The `go.mod` file defines the module path, which is used in import statements. For example:

```go
module myproject
```

This line tells Go that the root of your module is `myproject`. When you import `"myproject/mypackage"`, Go knows to look for the `mypackage` directory within the `myproject` module.

### 2. **Dependency Management**

The `go.mod` file lists all the dependencies your project needs. When you run `go build` or `go run`, Go will automatically download and manage these dependencies. For example:

```go
require (
    github.com/some/dependency v1.2.3
)
```

### 3. **Versioning**

The `go.mod` file specifies the Go version your project is using:

```go
go 1.23
```

This ensures that your project is built with the correct version of Go, which helps avoid compatibility issues.

### 4. **Reproducible Builds**

By locking dependencies to specific versions, the `go.mod` file ensures that your project builds consistently across different environments. This is crucial for collaboration and deployment.

### Example

Here's a complete example of how `go.mod` integrates with your project:

**go.mod**:
```go
module myproject

go 1.23
```

**main.go**:
```go
package main

import (
    "myproject/mypackage"
)

func main() {
    mypackage.Hello()
}
```

**mypackage/mypackage.go**:
```go
package mypackage

import "fmt"

func Hello() {
    fmt.Println("Hello from mypackage!")
}
```

### Summary

- **Defines Module Path**: Helps Go locate your custom packages.
- **Manages Dependencies**: Automatically handles downloading and updating dependencies.
- **Ensures Version Consistency**: Specifies the Go version to avoid compatibility issues.
- **Supports Reproducible Builds**: Locks dependencies to specific versions for consistent builds.

### How to run from different directory

1. **Set the `GO111MODULE` environment variable to `on`:** This ensures that Go uses the module system.

   ```sh
   export GO111MODULE=on
   ```

2. **Run the program using the `go run` command with the module-aware mode:** This allows you to specify the module root directory.

   ```sh
   go run custom-package/main.go
   ```

Alternatively, you can use the `go build` command to build the executable and then run it:

1. **Build the executable:**

   ```sh
   go build -o myprogram custom-package/main.go
   ```

2. **Run the executable:**

   ```sh
   ./myprogram
   ```

The `GO111MODULE` environment variable is used to control the module support in Go. Go modules are the dependency management system introduced in Go 1.11, which provide a way to declare, manage, and use dependencies in your Go projects.

The `GO111MODULE` variable can have three possible values:

1. **`off`**: Disables module support. Go will look for dependencies in the `GOPATH` directory.
2. **`on`**: Enables module support, regardless of whether a `go.mod` file is present.
3. **`auto`**: Enables module support only when a `go.mod` file is present in the current directory or any parent directory. This is the default setting in Go 1.13 and later¹.

By setting `GO111MODULE=on`, you are explicitly enabling module support, ensuring that Go uses the module system to manage dependencies, even if you are outside of the module's root directory.

This helps in scenarios where you want to run Go commands from directories that are not within the module's root, like in your case where you want to execute the program from the `go-lang-programming` directory.

### Running from a test directory

The error message indicates that Go is treating `myproject/mypackage` as a program rather than an importable package. This is likely because there is a `main.go` file in the `custom-package` directory, which Go interprets as a standalone program.

To resolve this, you can:

1. **Remove or rename the `main.go` file in the `custom-package` directory:** This will prevent Go from treating it as a program.

2. **Ensure the `mypackage` is correctly structured as a package:**

   Here’s how you can do it:

   1. **Rename or move the `main.go` file in the `custom-package` directory:**

      ```sh
      mv ~/go-lang-programming/custom-package/main.go ~/go-lang-programming/custom-package/main_program.go
      ```

   2. **Update the `go.mod` file in the `test-directory` to correctly reference the `mypackage`:**

      ```go
      // ~/go-lang-programming/test-directory/go.mod
      module test-directory

      go 1.23.0

      require myproject/mypackage v0.0.0

      replace myproject/mypackage => ../custom-package/mypackage
      ```

   3. **Ensure the import path in your `main.go` file in the `test-directory` is correct:**

      ```go
      // ~/go-lang-programming/test-directory/main.go
      package main

      import (
          "myproject/mypackage"
      )

      func main() {
          mypackage.Hello()
      }
      ```

   4. **Run the `main.go` file from the `test-directory`:**

      ```sh
      cd ~/go-lang-programming/test-directory
      go run main.go
      ```

By renaming or moving the `main.go` file in the `custom-package` directory, you ensure that Go treats `mypackage` as an importable package rather than a standalone program. This should resolve the import cycle error and allow you to run your `main.go` file from the `test-directory`.
