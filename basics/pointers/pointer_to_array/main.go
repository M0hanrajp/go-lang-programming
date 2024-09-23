package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := [5]int32{1, 2, 3, 4, 5}
	fmt.Printf("Address of 'a' in main function = %d\n", uintptr(unsafe.Pointer(&a)))
	fmt.Printf("Address of 'a' in main function without unsafe package = %d\n", &a)
	// Decalare a variable pa that points to an array of 5 integers
	// other accepted declarations var pa = &a, var pa *[5]int32 = &a
	pa := &a
	ModifyArray(pa)
	DisplayArray(pa)
	fmt.Printf("Modified array in main = %d\n", a)
}

func DisplayArray(array_input *[5]int32) {
	fmt.Printf("Address of array_input = %d\nWhat is at *array_input = %d\n", uintptr(unsafe.Pointer(array_input)), *array_input)
	for i := 0; i < len(array_input); i++ {
		fmt.Printf("Element[%d] = %2d & address = %d\n", i, array_input[i], uintptr(unsafe.Pointer(&array_input[i])))
	}
}

func ModifyArray(array_input *[5]int32) {
	for i := 0; i < len(array_input); i++ {
		array_input[i] = 2 * array_input[i]
	}
}

// Note: https://go.dev/tour/moretypes/1, Unlike C, Go has no pointer arithmetic.
