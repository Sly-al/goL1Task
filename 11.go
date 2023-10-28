package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func intersect(arr1, arr2 []int) []int {
	set1 := map[int]struct{}{}
	ans := []int{}
	// за линию строим "множество" из первого
	for _, num := range arr1 {
		set1[num] = struct{}{}
	}
	// проходим по элементам второго
	for _, num := range arr2 {
		// если такое ключ есть во множестве из первого, то нашли пересечение
		if _, ok := set1[num]; ok {
			ans = append(ans, num)
		}
	}
	return ans
}

func main() {
	arr1 := []int{1, 1, 2, 2, 3, 10, 11, 12, 13, 14, 15, 164, 5, 6, 7, 8, 9}
	arr2 := []int{100, 1, 52, 15, 164, 3}
	ans := intersect(arr1, arr2)
	fmt.Println(ans)
}
