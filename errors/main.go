package main

import "fmt"

var scope float64
var something float32

func main() {
	scope := 45
	fmt.Println(scope, something)
	fmt.Printf("Value of scope = %d\nValue of something = %f\n", scope, something)
	fmt.Println("Some value")
}
