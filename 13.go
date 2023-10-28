package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func swapTuple(a *int, b *int) {
	*a, *b = *b, *a
}

func swapSum(a *int, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func swapXOR(a *int, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func main() {
	var a, b int
	fmt.Print("Введите два числа: ")
	fmt.Scan(&a, &b)
	//свап при помощи кортежа
	swapTuple(&a, &b)
	fmt.Println(a, b)

	// свап при помощи суммы и разности
	swapSum(&a, &b)
	fmt.Println(a, b)

	// свап при помощи XOR
	swapXOR(&a, &b)
	fmt.Println(a, b)

}
