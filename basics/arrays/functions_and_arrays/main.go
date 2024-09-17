package main

import "fmt"

// an array as an argument
func myfun(array [6]int, size int) int {
	// the size of the array expects a constant, cannot use variable here
	funcitonArray := [6]int{}
	for k := 0; k < size; k++ {
		funcitonArray[k] = array[k]
		fmt.Printf("Element [%d] of array[%d] = %d\n", k, k, array[k])
	}
	fmt.Println("From Function, the array passed: ", funcitonArray)
	return 0
}

// Main function
func main() {

	// Creating and initializing an array
	var arr = [6]int{67, 59, 29, 35, 42, 34}
	myfun(arr, 6)
}
