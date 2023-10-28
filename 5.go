package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

func main() {
	var (
		n  int
		wg sync.WaitGroup
	)
	fmt.Print("Введите количество секунд ")
	fmt.Scan(&n)
	// создаем timer для отсчёта секунд
	timer := time.After(time.Duration(n) * time.Second)
	// канал для отправки данных
	intChan := make(chan int)
	wg.Add(1)
	go func(intChan <-chan int) {
		defer wg.Done()
		for mes := range intChan {
			fmt.Println(mes)
		}
	}(intChan)
	num := 0
	for {
		// мультиплексирование для ожидания таймера и отправки сообщений
		select {
		case <-timer:
			close(intChan)
			wg.Done()
			return
		default:
			intChan <- num
			num++
		}
	}

}
