package main

import "fmt"

/*
Разработать конвейер чисел. Даны два канала:
в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	arr := [...]int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	in := make(chan int)
	out := make(chan int)
	// асинхронная запись чисел в канал in
	go func() {
		for _, num := range arr {
			in <- num
		}
		close(in)
	}()
	// асинхроннное чтение из канала in и запись произведения в канал out
	go func() {
		for num := range in {
			out <- num * 2
		}
		close(out)
	}()
	// вывод
	for num := range out {
		fmt.Println(num)
	}

}
