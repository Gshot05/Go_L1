package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Создаем мьютекс для обеспечения конкурентного доступа к map
	var mu sync.Mutex

	// Создаем map для хранения данных
	data := make(map[string]int)

	// Создаем WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Количество горутин для записи данных
	numWorkers := 5

	// Запускаем горутины для записи данных
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go writeData(i, &wg, &mu, data)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим данные из map после завершения горутин
	fmt.Println("Итоговые данные:", data)
}

func writeData(workerID int, wg *sync.WaitGroup, mu *sync.Mutex, data map[string]int) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		// Генерируем уникальный ключ для каждой горутины
		key := fmt.Sprintf("Key%d_%d", workerID, i)

		// Генерируем уникальное значение для каждой горутины
		value := workerID*10 + i

		// Блокируем мьютекс перед записью в map
		mu.Lock()
		// Записываем данные в map
		data[key] = value
		// Разблокируем мьютекс после записи
		mu.Unlock()

		// Имитируем задержку для наглядности
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("Горутина %d записала данные: %s=%d\n", workerID, key, value)
	}
}
