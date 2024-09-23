package main

import "fmt"

func main() {
	minutes := 45
	fmt.Printf("Each half in a football match lasts:: %d minutes, [Debug address: %d]\n", minutes, &minutes)
	// Call the function to multiply
	MultiplyByTow(&minutes)
	fmt.Printf("A Football match lasts:: %d minutes, [Debug address: %d]\n", minutes, &minutes)
}

func MultiplyByTow(insert *int) {
	*insert = *insert * 2
}
