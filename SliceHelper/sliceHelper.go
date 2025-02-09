package slicehelper

import "reflect"

type Slice []Element

type Element int

func DeleteLast[T any](slice *[]T) {
	if len(*slice) != 0 {
		*slice = (*slice)[:len(*slice)-1]
	}
}

func DeleteFirst[T any](slice *[]T) {
	if len(*slice) != 0 {
		*slice = (*slice)[1:]
	}
}

func DeleteByIndex[T any](slice *[]T, index int) {
	if len(*slice) != 0 && index < len(*slice) {
		*slice = append((*slice)[:index], (*slice)[index+1:]...)
	}
}

func Compare[T any](slice1 []T, slice2 []T) bool {
	return reflect.DeepEqual(slice1, slice2)
}

// SumSlice — возвращает сумму элементов
func SumSlice(slice Slice) (res Element) {
	for _, s := range slice {
		res += s
	}
	return
}

// MapSlice — применяет функцию op к каждому элементу
func MapSlice(slice Slice, op func(Element) Element) {
	for i, s := range slice {
		slice[i] = op(s)
	}
}

// FolвSlice — сворачивает слайс.
func FoldSlice(slice Slice, op func(Element, Element) Element, init Element) (res Element) {
	res = op(init, slice[0])
	for i := 1; i < len(slice); i++ {
		res = op(res, slice[i])
	}
	return
}
