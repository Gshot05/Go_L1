package main

import (
	"fmt"
	"sync"
)

// Counter представляет структуру счетчика
type Counter struct {
	mu    sync.Mutex
	value int
	wg    sync.WaitGroup
}

// Increment инкрементирует значение счетчика на 1
func (c *Counter) Increment() {
	defer c.wg.Done()

	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

// GetValue возвращает текущее значение счетчика
func (c *Counter) GetValue() int {
	return c.value
}

func main() {
	// Создаем экземпляр счетчика
	counter := Counter{}

	// Количество горутин для инкрементации
	numGoroutines := 10

	// Добавляем количество горутин в WaitGroup
	counter.wg.Add(numGoroutines)

	// Запускаем горутины для инкрементации
	for i := 0; i < numGoroutines; i++ {
		go counter.Increment()
	}

	// Ожидаем завершения всех горутин
	counter.wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.GetValue())
}
