package main

import "fmt"
import "unsafe"

func main() {

	// Finding the length of the array
	arrayOne := [7]int{1, 2}
	fmt.Printf("Length of arrayOne is: %2d\n", len(arrayOne))
	fmt.Printf("size of arrayOne is: %2d\n", unsafe.Sizeof(arrayOne))

	// Duplicate values and total length of the array is 10
	arrayTwo := [...]float32{3.4, 5.6, 6.9, 7.8, 6.7, 4.5, 5.6, 6.6, 7.7, 5.6}
	fmt.Printf("Length of arrayTwo is: %d\n", len(arrayTwo))
	fmt.Printf("size of arrayTwo is: %2d\n", unsafe.Sizeof(arrayTwo))

	// How to access the above elements using for loop
	for i := 0; i < len(arrayTwo); i++ {
		fmt.Printf("element[%d] of arrayTwo: %.2f\n", i, arrayTwo[i])
	}

	// Array is of value type not of reference type
	arrayDefault := [3]int{1, 1, 1}
	fmt.Println("Default array: ", arrayDefault)
	// array can be copied using the short declaration
	arrayCopy := arrayDefault
	arrayCopy[1] = 3
	fmt.Println("Modified array: ", arrayCopy)
	fmt.Println("(After modification)Default array: ", arrayDefault)

	// Comparison of arrays
	// They need to be of same size & same data type
	arrayCompare := [7]int{1, 2, 3, 4, 5, 6, 8}
	fmt.Println("Is arrayCompare equals arrayOne ?:", arrayCompare == arrayOne)
	arrayOneDuplicate := [7]int{}
	fmt.Println("Is arrayOneDuplicate equals arrayOne ?:", arrayOneDuplicate == arrayOne)
	arrayWithoutSize := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Is arrayWithoutSize equals arrayCompare ?:", arrayCompare == arrayWithoutSize)

	// copy an array by reference
	my_arr1 := [6]int{12, 45, 67, 65, 34, 34}
	// Here, the elements are passed by reference
	my_arr2 := &my_arr1
	fmt.Println("Array_1: ", my_arr1)
	fmt.Println("Array_2: ", *my_arr2)
	my_arr1[5] = 44
	fmt.Println("Array_1: ", my_arr1)
	fmt.Println("Array_2: ", *my_arr2)

}
