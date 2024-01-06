package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем каналы для передачи данных между этапами конвейера
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Этап 1: Горутина, генерирующая числа
	wg.Add(1)
	go generateNumbers(ch1, &wg)

	// Этап 2: Горутина, умножающая числа на 2
	wg.Add(1)
	go multiplyByTwo(ch1, ch2, &wg)

	// Этап 3: Горутина, выводящая числа
	wg.Add(1)
	go printNumbers(ch2, &wg)

	// Завершаем конвейер после завершения всех горутин
	wg.Wait()
}

// Этап 1: Генерация чисел
func generateNumbers(out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Пример чисел
	numbers := []int{1, 2, 3, 4, 5}

	// Отправляем числа в первый канал
	for _, num := range numbers {
		out <- num
	}

	// Закрываем первый канал, чтобы сообщить, что данные завершены
	close(out)
}

// Этап 2: Умножение чисел на 2
func multiplyByTwo(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Умножаем числа на 2 и отправляем результат во второй канал
	for num := range in {
		result := num * 2
		out <- result
	}

	// Закрываем второй канал после обработки всех данных
	close(out)
}

// Этап 3: Вывод чисел в stdout
func printNumbers(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Выводим числа из второго канала в стандартный вывод (stdout)
	for num := range in {
		fmt.Println(num)
	}
}
