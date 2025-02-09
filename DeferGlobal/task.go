package main

import "fmt"

var Global = 5

func main() {
	fmt.Println("Global:", Global)
	defer ReturnGlobal(Global)

	Global = 10
	fmt.Println("Updated Global:", Global)
}

func ReturnGlobal(oldValue int) {
	Global = oldValue
	fmt.Println("Returned Global:", Global)
}
