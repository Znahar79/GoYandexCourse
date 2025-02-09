package main

import "fmt"

func main() {
	slice := make([]int, 100)

	number := 1
	for index := range slice {
		slice[index] = number
		number++
	}

	slice = append(slice[:10], slice[89:]...)

	slice = reverseInts(slice)

	fmt.Println(slice)
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
