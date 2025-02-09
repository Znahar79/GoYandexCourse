package main

import (
	"fmt"
	"math"
)

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

func main() {
	myFigure := square
	fmt.Println("Figure: ", myFigure)
	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("Error: unknown figure")
		return
	}

	if ar != nil {
		x := 5.0
		myArea := ar(x)

		fmt.Println("Area: ", myArea)
	}
}

func area(figure figures) (func(float64) float64, bool) {
	switch figure {
	case square:
		return areaSquare, true
	case circle:
		return areaCircle, true
	case triangle:
		return areaTriangle, true
	default:
		return nil, true

	}
}

func areaSquare(value float64) float64 {
	return value * value
}

func areaCircle(value float64) float64 {
	return math.Pi * value * value
}

func areaTriangle(value float64) float64 {
	return (math.Sqrt(3) / 4) * value * value
}
