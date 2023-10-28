package main

import (
	"fmt"
	"unicode"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func checkUnique(s string) bool {
	// самодельное множество
	m := make(map[rune]struct{})
	for _, ch := range s {
		// кастуем руну в нижний регситр
		ch = unicode.ToLower(ch)
		// если такоя руна уже встречлась, то символы в строке не уникальны
		if _, ok := m[ch]; ok {
			return false
		}
		m[ch] = struct{}{}
	}
	// если прошли по всей строки и не нашли повторений, то символы все уникальны
	return true
}

func main() {
	s := "abcde"
	fmt.Println(checkUnique(s))
}
