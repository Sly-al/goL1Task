package main

import (
	"fmt"
	"math/rand"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

*/

/* 1. если есть возможность, то стоит избегать использования глобальных перменных
   2. зачем выделять дополнительную память под переменную, если сразу можно сделать срез от возвращаемого значения
*/

func createHugeString(num int) string {
	str := []rune("0123456789абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	res := ""
	for i := 0; i < num; i++ {
		res += string(str[rand.Intn(len(str))])
	}
	return res
}

func someFunc() string {
	return createHugeString(1 << 10)[:100]
}

func main() {
	fmt.Print(someFunc())
}
