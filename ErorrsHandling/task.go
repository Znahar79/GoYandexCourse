package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// 1) вставьте определение типа для []error
// 2) определите метод Error для вашего типа, который будет выводить
//    все ошибки слайса
// 3) реализуйте функцию MyCheck
//
// ...

type MyErrorStruct struct {
	Errors []error
}

func main() {
	for {
		fmt.Printf("Укажите строку (q для выхода): ")
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		ret = strings.TrimRight(ret, "\n")
		if ret == `q` {
			break
		}
		if err = MyCheck(ret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(`Строка прошла проверку`)
		}
	}
}

func (errs MyErrorStruct) Error() string {
	var out string
	for _, err := range errs.Errors {
		out += err.Error() + ";"
	}
	return strings.TrimRight(out, `; `)
}

func MyCheck(checkString string) error {
	errorStruct := MyErrorStruct{}
	for _, r := range checkString {
		if unicode.IsDigit(r) {
			errorStruct.Errors = append(errorStruct.Errors, errors.New("found numbers"))
		}
	}

	if len(checkString) >= 20 {
		errorStruct.Errors = append(errorStruct.Errors, errors.New("line is too long"))
	}

	if strings.Count(checkString, " ") != 2 {
		errorStruct.Errors = append(errorStruct.Errors, errors.New("no two spaces"))
	}

	if len(errorStruct.Errors) != 0 {
		return errorStruct
	}
	return nil
}
