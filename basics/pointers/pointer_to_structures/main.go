package main

import "fmt"

type EdgexServices struct {
	NameOfService   string
	IsServiceActive bool
}

func main() {
	sessionA := EdgexServices{NameOfService: "Core Data", IsServiceActive: true}
	fmt.Printf("Default structure in Main function:\n%v\n", sessionA)
	ChangeServiceState(&sessionA)
	fmt.Printf("Modified structure after function call:\n%v\n", sessionA)
}

func ChangeServiceState(session *EdgexServices) {
	fmt.Printf("The type of variable in this function is : %T\n", session)
	if session.IsServiceActive == true {
		session.IsServiceActive = false
	}
}
