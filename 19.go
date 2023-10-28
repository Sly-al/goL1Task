package main

import (
	"fmt"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.

*/

// передаём строку по указателю, чтобы её же и поменять
func reverse(inpStr *string) {
	// переводим символы в слайс рун(так мы сможем работать с любым символом из unicode)
	arrLetters := []rune(*inpStr)
	l := 0
	r := len(arrLetters) - 1
	// реверс
	for l < r {
		arrLetters[l], arrLetters[r] = arrLetters[r], arrLetters[l]
		l++
		r--
	}
	*inpStr = string(arrLetters)
}

func main() {
	inpStr := "ПРиветю . / world"
	fmt.Println(inpStr)
	reverse(&inpStr)
	fmt.Println(inpStr)
}
