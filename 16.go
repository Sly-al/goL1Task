package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)-1]
	var left []int
	var right []int

	// развбиение элементов на большие и меньше опорного
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	// рекурсивная сортировка элементов больших и меньших опорного
	sortedArr := append(QuickSort(left), pivot)
	sortedArr = append(sortedArr, QuickSort(right)...)

	return sortedArr
}

func main() {
	arr := []int{5, -2, 9, 1, 7, 6, 3, 8, 100}
	fmt.Println("Before sorting:", arr)
	sortedArr := QuickSort(arr)
	fmt.Println("After sorting:", sortedArr)
}
