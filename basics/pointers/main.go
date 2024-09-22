/* Pointers */
// Variables are the names given to a memory location where the actual data is stored.
// To access the stored data we need the address of that particular memory location.
/* To remember all the memory addresses(Hexadecimal Format) manually is an overhead that’s
   why we use variables to store data and variables can be accessed just by using their name. */

/* A pointer is a special kind of variable that is not only used to store the memory addresses of other variables but
   also points where the memory is located and provides ways to find out the value stored at that memory location. */

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	BasicPointerUnderstanding()
}

func BasicPointerUnderstanding() {
	// Pointer declaration, creating a variable, dereference operator
	var x int8 = 20
	var px *int8 = &x
	fmt.Printf("Type of x = %T, value of x = %d & address of x = %d\n", x, x, &x)
	fmt.Printf("Type of px = %T, value of px (what px holds) = %d, & address of px = %d\n", px, px, &px)
	fmt.Printf("Accessing value of x using dereference operator on px (where *px points to ?): %d\n", *px)
	fmt.Printf("Size of x = %d & size of px = %d\n", unsafe.Sizeof(x), unsafe.Sizeof(px))
}

func UninitializedPointerVaraible() {
	// an uninitialized pointer will always have a nil value.
	var test *int
	fmt.Printf("Variable 'test' is initialized to : %d, %v\n", test, test)
}

func TypeInferenceOfPointerVariableInteger(input int) {
	// Type inference in declaring variables
	// compiler auto detects the type of variable based on assignment
	var sp = &input
	fmt.Printf("What is the type of variable 'input' ? : %T\nWhat does input hold ?: %d\naddress of input ?: %d\n", input, input, &input)
	fmt.Printf("What is the type of variable 'sp' ? : %T\nWhat does sp hold ?: %d\nWhat sp points to ?: %d\n", sp, sp, *sp)
	// pen := "ink"
	// Uncommenting the following line will cause a type mismatch error
	// sp = &pen
	// fmt.Printf("What is the type of variable 'normal' ? : %T\nWhat does normal holds ?: %d\nWhat normal points to ?: %d\n", normal, normal, *normal)
}

func TypeInferenceOfPointerVariableString(insert *string) {
	// Type inference in declaring variables
	// compiler auto detects the type of variable based on assignment
	var dp = &insert
	fmt.Printf("What is the type of variable 'insert' ? : %T\nWhat does insert hold ?: %d\naddress of insert ?: %d\n", insert, insert, &insert)
	fmt.Printf("What is the type of variable 'dp' ? : %T\nWhat does dp hold ?: %d\nWhat dp points to using *dp ?: %d\nWhat dp points to using **dp ?: %s\n", dp, dp, *dp, **dp)
}

func ShortDecalarationOfPointer() {
	// Shorthand declaration
	p := "this is a string"
	q := &p
	fmt.Printf("What is the type of variable 'p' ? : %T\nWhat does p hold ?: %s\nWhat is p's address ?: %d\n", p, p, &p)
	fmt.Printf("What is the type of variable 'q' ? : %T\nWhat does q hold ?: %d\nWhat q points to ?: %s\n", q, q, *q)
}

func CallDecimal() int {
	decimal := 9090
	fmt.Printf("What is the type of decimal: %T\nvalue of decimal: %d\naddress of decimal: %d\n", decimal, decimal, &decimal)
	return decimal
}

func CallSentence() *string {
	sentence := "How does double pointer work ?"
	fmt.Printf("What is the type of sentence: %T\nvalue of sentence: %s\naddress of sentence: %d\n", sentence, sentence, &sentence)
	return &sentence
}
