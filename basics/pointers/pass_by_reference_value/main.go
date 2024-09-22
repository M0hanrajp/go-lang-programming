package main

import "fmt"

func main() {
	// Pass by reference example using integer
	mainDecimalValue := 10
	fmt.Printf("From main funciton:: Value = %d, address = %d\n", mainDecimalValue, &mainDecimalValue)
	PassByReference(&mainDecimalValue)
	fmt.Printf("After modifying in main funciton:: Value of variable in main = %d, address that variable holds = %d\n", mainDecimalValue, &mainDecimalValue)

	// Passs by value using string
	mainSentence := "This string is from main"
	fmt.Printf("From main funciton:: Value = %s, address = %d\n", mainSentence, &mainSentence)
	PassByValue(mainSentence)
	fmt.Printf("After modifying in main funciton:: Value = %s, address of variable = %d\n", mainSentence, &mainSentence)
}

func PassByReference(insert *int) {
	fmt.Printf("From PassByReference funciton:: Value using dereference operator = %d, address that variable holds = %d\n", *insert, insert)
	// modifying the value from the main function
	*insert = *insert * 10
	fmt.Printf("After modifying in PassByReference funciton:: Value using dereference operator = %d, address that variable holds = %d\n", *insert, insert)
}

func PassByValue(insert string) {
	fmt.Printf("From PassByValue funciton:: Value = %s, address of variable = %d\n", insert, &insert)
	// Modify the string
	insert = "This string is from PassByValue"
	fmt.Printf("After modifying in PassByValue funciton:: Value = %s, address of variable = %d\n", insert, &insert)
}
