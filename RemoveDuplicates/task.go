package main

import (
	"fmt"
)

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	fmt.Println("Original slice: ", input)
	fmt.Println("Removed dublicates: ", RemoveDuplicates(input))
}

func RemoveDuplicates(input []string) []string {
	inputMap := make(map[string]int)
	for _, element := range input {
		inputMap[element]++
	}

	var result = make([]string, len(inputMap))
	i := 0
	for k := range inputMap {
		result[i] = k
		i++
	}

	return result
}
