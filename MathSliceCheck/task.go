package main

import (
	"fmt"
	slicehelper "yandexcourse/SliceHelper"
)

func main() {

	s := slicehelper.Slice{1, 2, 3}
	fmt.Println(s)
	fmt.Println("Сумма слайса: ", slicehelper.SumSlice(s))

	slicehelper.MapSlice(s, func(i slicehelper.Element) slicehelper.Element {
		return i * 2
	})

	fmt.Println("Слайс, умноженный на два: ", s)

	fmt.Println("Сумма слайса: ", slicehelper.SumSlice(s))

	fmt.Println("Свёртка слайса умножением ",
		slicehelper.FoldSlice(s,
			func(x slicehelper.Element, y slicehelper.Element) slicehelper.Element {
				return x * y
			},
			1))

	fmt.Println("Свёртка слайса сложением ",
		slicehelper.FoldSlice(s,
			func(x slicehelper.Element, y slicehelper.Element) slicehelper.Element {
				return x + y
			},
			0))

}
