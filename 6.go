package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины
*/

// функция в горутине доходит до return, завершается функция и горутина
func foo(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	return
}

// запускаем функцию в горутине, ждём конца выполнения
func v1() {
	fmt.Println("Variant 1:")
	var wg sync.WaitGroup
	wg.Add(1)
	go foo(&wg)
	wg.Wait()
}

// запускаем горутину, которая считываем из канала, завершение горутины происходит после закрытия канала
func v2() {
	fmt.Println("Variant 2:")
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go func(ch <-chan int) {
		defer wg.Done()
		for mes := range ch {
			fmt.Println(mes)
		}
	}(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}

func main() {
	v1()
	time.Sleep(time.Second)
	v2()
}
