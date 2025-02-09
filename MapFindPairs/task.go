package main

import (
	"fmt"
)

func main() {
	A := []int{2, 7, 11, 15}
	k := 18
	result := findPairWithSum(A, k)

	if len(result) > 0 {
		fmt.Println("Pair: ", result)
	} else {
		fmt.Println("Pair not found")
	}
}

func findPairWithSum(nums []int, checkSum int) []int {
	indexMap := make(map[int]int)
	for i, num := range nums {
		diff := checkSum - num

		if idx, exists := indexMap[diff]; exists {
			return []int{idx, i}
		}
		indexMap[num] = i
	}

	return []int{}
}
