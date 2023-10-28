package main

import (
	"fmt"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func main() {
	var (
		n int64
		i int
		b int
	)
	fmt.Print("Введите число и номер бита через пробел: ")
	fmt.Scan(&n, &i)
	fmt.Print("Введите значение битика ")
	fmt.Scan(&b)
	fmt.Printf("Before: %b\n", n)
	// устанавливаем 1 при помощи логического или
	if b == 1 {
		mask := n | (1 << i) // or
		fmt.Printf("After setting 1: %b\n", mask)
		// устанавливаем бит ноль при помощи логического И, при этом XOR делает необходимый бит(1 << i) нулевым
	} else {
		mask := n &^ (1 << i) // and not(bit clear)
		fmt.Printf("After setting 0: %b\n", mask, 2)
	}

}
