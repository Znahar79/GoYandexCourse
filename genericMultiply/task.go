package main

import "fmt"

func main() {
	fmt.Println(Mul("a", 5))
}

func Mul(a interface{}, b int) interface{} {
	switch a2 := a.(type) {
	case int:
		return a2 * b
	case string:
		result := ""
		for i := 0; i < b; i++ {
			result += a2 + "\n"
		}
		return result
	default:
		return "Неизвестный тип"
	}
}
