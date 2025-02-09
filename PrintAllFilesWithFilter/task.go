package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	PrintAllFilesWithFilter("D:/Work/GO_Projects/YandexCourse", "Map")
}

func PrintAllFilesWithFilter(pathDir string, filter string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := os.ReadDir(pathDir)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}
	//  проходим по списку
	for _, f := range files {
		// получаем имя элемента
		// filepath.Join — функция, которая собирает путь к элементу с разделителями
		filename := filepath.Join(pathDir, f.Name())
		// печатаем имя элемента
		if strings.Contains(filename, filter) {
			fmt.Println(filename)
		}

		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			PrintAllFilesWithFilter(filename, filter)
		}
	}
}
