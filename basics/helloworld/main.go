/*
Based on the type of program we are developing, we might generate
executable ( executable is created for the whole application ) or reusable ( Code dependencies, helpers, reused code )
The current program creates an executable
Name of the package determines executable or reusable or dependency type package, main is used for executable here.
*/
package main

// Use code from the standard library of go so that main package will have access to fmt features, (https://pkg.go.dev/std)
import "fmt"

// For each executable package created there must be a func main()
// func - How we declare a function
// main - the name of the function
// () - List of arguments to pass the function
func main() {
	// function body
	fmt.Println("Hello GO programming!")
	// Calling the function
	helloworld()
}

/* Notes:
How is a go program organized ?
1. Package declaration
2. Import other packages that we need
3. Declare functions, write the logic. (in the above order)
*/

// Writing a custom function
func helloworld() {
	fmt.Println("From helloworld! function")
}
