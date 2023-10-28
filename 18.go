package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

// структура, состоящая из переменное счётчика и мьютекса
type storage struct {
	counter int
	mu      sync.Mutex
}

// безопасный метод для увеличения счётчика
func (s *storage) increment() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counter++
}

// число потоков
const N = 228

func main() {
	var wg sync.WaitGroup
	st := storage{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		// конкуретно увеличиваем счётчик
		go func() {
			defer wg.Done()
			st.increment()
		}()
	}
	wg.Wait()
	fmt.Println(st.counter)
}
