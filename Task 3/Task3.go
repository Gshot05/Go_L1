package main

import (
	"fmt"
	"sync"
)

// Функция для нахождения квадратов
// В функции передаем указаетель wg на пакет sync
// Nuber int показывает что мы передаём числа
// resultChan поток в который мы складываем данные типа int
func squareWorker(wg *sync.WaitGroup, number int, resultChan chan<- int) {
	//Отклдываем wg.Done чтобы уменьшить счётчик горутин
	defer wg.Done()
	//Находим квадрат числа
	result := number * number
	//Записываем полученное в resultChan
	resultChan <- result
}

// Главная функцияя
func main() {
	//Создали срез чисел
	numbers := []int{2, 4, 6, 8, 10}
	//Создали канал ограниченного размера и типа int
	resultChan := make(chan int, len(numbers))
	//Переменна wg ссылается на пакет sync
	var wg sync.WaitGroup

	//Создали цикл размером с наш срез
	for _, num := range numbers {
		//Увеличиваем счётчик горутин
		wg.Add(1)
		//С помощью горутины используем нашу функцию
		go squareWorker(&wg, num, resultChan)
	}

	// Закрываем канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var sum int

	// Читаем результаты из канала и суммируем их
	for result := range resultChan {
		sum += result
	}

	fmt.Printf("Сумма квадратов: %d\n", sum)
}
