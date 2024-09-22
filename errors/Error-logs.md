## Error logs

`var` is not used with card & is directly assigned
```go
func main() {

        card = "Ace of spades"
```
Error:
```bash
# command-line-arguments
./main.go:7:2: undefined: card
```

---
variable declared once using `:=` need not be used again for assigning new value
```go
        cardOne := "Five of hearts"
        cardOne := "Five of diamonds" // use cardOne = "Five of diamonds"
```
Error:
```bash
# command-line-arguments
./main.go:22:10: no new variables on left side of :=
```
---

Based on the go lang spec: https://go.dev/ref/spec#Short_variable_declarations

>Short variable declarations may appear only inside functions. In some contexts such as the initializers for "if", "for", or "switch" statements, they can be used to declare local temporary variables.

```go
package main

import "fmt"

scope := 10

func main() {
    fmt.Println(scope)
}
```
Error:
```bash
mpunix@LIN-MP22QN1X:~/go-lang-programming$  go run errors/main.go
# command-line-arguments
errors/main.go:5:1: syntax error: non-declaration statement outside function body
```
---

Using type inference for assigning a variable with type pointer. but cannot change the type once assigned.

Error:
```bash
# command-line-arguments
basics/pointers/main.go:32:11: cannot use &pen (value of type *string) as *int value in assignment
```
Code:
```go
	// Type inference in declaring variables
	// compiler auto detects the type of variable based on assignment
	pen := "ink"
	var normal = &x
	fmt.Printf("What is the type of variable 'normal' ? : %T\nWhat does normal holds ?: %d\nWhat normal points to ?: %d\n", normal, normal, *normal)
	normal = &pen
	fmt.Printf("What is the type of variable 'normal' ? : %T\nWhat does normal holds ?: %d\nWhat normal points to ?: %d\n", normal, normal, *normal)
```
