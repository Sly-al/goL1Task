package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

// передаём слайс строк
func reverseArr(arr []string) {
	l := 0
	r := len(arr) - 1
	// реверс слов
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func main() {
	text := "supper dog show online"
	words := strings.Split(text, " ")
	reverseArr(words)
	// джойним слова из слайса в единую строку
	text = strings.Join(words, " ")
	fmt.Println(text)
}
