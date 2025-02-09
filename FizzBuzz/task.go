package main

import "fmt"

func main() {
	fizzBuzzSwitch()
}

func fizzBuzzIf() {
	for i := 1; i < 101; i++ {
		result := ""

		if i%3 == 0 {
			result += "Fizz"
		}

		if i%5 == 0 {
			result += "Buzz"
		}

		if result == "" {
			result = fmt.Sprint(i)
		}

		fmt.Println(result)
	}
}

func fizzBuzzSwitch() {
	for i := 1; i < 101; i++ {
		result := ""

		switch {
		case i%3 == 0 && i%5 == 0:
			result += "FizzBuzz"
		case i%3 == 0:
			result += "Fizz"
		case i%5 == 0:
			result += "Buzz"
		default:
			result = fmt.Sprint(i)
		}

		fmt.Println(result)
	}
}
