package main

import "fmt"

func main() {
	var yearOfBirth int = 1998
	fmt.Println(getText(yearOfBirth))
}

func getText(yearOfBirth int) string {
	switch {
	case yearOfBirth >= 1946 && yearOfBirth < 1965:
		return "Привет, бумер!"
	case yearOfBirth >= 1965 && yearOfBirth < 1981:
		return "Привет, представитель X!"
	case yearOfBirth >= 1981 && yearOfBirth < 1997:
		return "Привет, миллениал!"
	case yearOfBirth >= 1997 && yearOfBirth < 2013:
		return "Привет, зумер!"
	case yearOfBirth >= 2013:
		return "Привет, альфа!"
	default:
		return ""
	}
}
