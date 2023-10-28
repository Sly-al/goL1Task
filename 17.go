package main

import "fmt"

/*
Реализовать бинарный поиск встроенными методами языка
*/

// бинарный поиск, который ищет первое подходящее значение
func lBinSearch(arr []int, x int) int {
	l := 0
	r := len(arr)
	for l < r {
		m := (l + r) / 2
		if arr[m] >= x {
			r = m
		} else {
			l = m + 1
		}
		//fmt.Println(arr[m])
	}
	return l
}

// бинарный поиск, которые ищет последнее подходящее значение
func rBinSearch(arr []int, x int) int {
	l := 0
	r := len(arr) - 1
	for l < r {
		m := (l + r + 1) / 2
		if arr[m] <= x {
			l = m
		} else {
			r = m - 1
		}
		//fmt.Println(arr[m])
	}
	return l
}

func main() {
	arr := []int{2, 3, 4, 10, 12, 12, 12, 13, 14, 14, 100, 150}
	ans := lBinSearch(arr, 12)
	fmt.Println(" idx lBinSearch", ans)
	ans = rBinSearch(arr, 12)
	fmt.Println("idx rBinSearch", ans)
}
