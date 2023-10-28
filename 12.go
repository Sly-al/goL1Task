package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	words := [...]string{"cat", "cat", "dog", "cat", "tree"}
	// для создания множества используем мапу с пустыми значениями
	set := make(map[string]struct{})
	for _, now := range words {
		set[now] = struct{}{}
	}
	fmt.Println(set)
}
