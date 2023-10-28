package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func worker(i int, intChan <-chan int, wg *sync.WaitGroup) {
	for message := range intChan {
		fmt.Printf("Worлer %d recived: %d", i, message)
	}
	wg.Done()
}

func main() {
	var (
		n  int
		wg sync.WaitGroup
	)
	num := 0
	fmt.Print("Enter amount of workers ")
	fmt.Scan(&n)
	intChan := make(chan int)
	wg.Add(n)
	// запускаем воркеров
	for i := 0; i < n; i++ {
		go worker(i, intChan, &wg)
	}
	// канал, который будет "отлавливать" сигнал Ctrl+C с консоли
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	for {
		// мультиплексирование для отправки сообщений и отлавдивания Ctrl+C
		select {
		case <-sigch:
			close(intChan)
			wg.Wait()
			return
		default:
			intChan <- num
			num++
		}
	}
}
