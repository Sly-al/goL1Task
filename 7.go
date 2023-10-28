package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map
*/

// запись в мапу
func writer(i int, m map[int]int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock() // перед каждой записью лочим мьютекс для избежания состояния гонки
	defer mutex.Unlock()
	m[i] = i * i
}

func main() {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)
	m := make(map[int]int)
	wg.Add(100)
	// создаём 100 потоков
	for i := 0; i < 100; i++ {
		go writer(i, m, &wg, &mutex)
	}
	wg.Wait()
	fmt.Println(m)
}
