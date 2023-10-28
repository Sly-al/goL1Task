package main

import (
	"fmt"
	"log"
)

/*
Удалить i-ый элемент из слайса.
*/

// удаление элемента при помощи функции copy
func remove1(i int, slice []int) []int {
	if i >= len(slice) || i < 0 {
		log.Fatalf("Slice index out of range")
	}
	copy(slice[i:], slice[i+1:])
	// не забываем укорочавить слайс
	return slice[:len(slice)-1]
}

// удаление при помощи функции append
func remove2(i int, slice []int) []int {
	if i >= len(slice) || i < 0 {
		log.Fatalf("Slice index out of range")
	}
	// берём за основу слайс до i-го элемента, не включая и добавляем в него всё что после i-го
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{802, 804, 805, 806}
	slice = remove1(3, slice)
	fmt.Println(slice, cap(slice))
	slice = remove2(0, slice)
	fmt.Println(slice, cap(slice))
}
