package main

import "fmt"

func main() {
	// array declaration & definition

	// Type 1: explicit declaration
	var arrayDeclarationExplicit [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("arrayDeclaration:", arrayDeclarationExplicit)

	// var type is not mentioned it will be infered by the compiler
	// better for big code bases
	var arrayDeclarationInfered = [5]int{1, 2, 3, 4, 5}
	fmt.Println("arrayDec:", arrayDeclarationInfered)

	// Type 2:
	arrayShortDeclaration := [5]float32{2.2, 3.3, 4.5, 6.7, 5.3}
	fmt.Println("float type array short declaration:", arrayShortDeclaration)

	// Type 3: accessing array elements using index operator
	var a [5]int
	fmt.Println("Whole array contents:", a)
	a[4] = 100
	a[3] = 200
	fmt.Println("set:", a)
	fmt.Println("get:", a[3], a[4])
	fmt.Println("len:", len(a))

	// Type 4: 2 dimensional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("Contents of the 2d array: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Contents of the 2d array: ", twoD)

	// Accessing array elements using a for loop
	sentence := [4]string{"I'm", "learning", "go", "language"}
	for i := 0; i < 4; i++ {
		fmt.Printf("%s ", sentence[i])
	}
}
