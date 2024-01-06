package main

import (
	"fmt"
	"sync"
)

// Функция для нахождения квадратов
func squareWorker(wg *sync.WaitGroup, number int) {
	//Отклдываем wg.Done чтобы уменьшить счётчик горутин
	defer wg.Done()
	//Находим квадрат
	result := number * number
	//Форматируем и выводим ответ
	fmt.Printf("%d squared is %d\n", number, result)
}

// Функция main, точка входа в нашу прогу
func main() {
	//Создаём срез чисел
	numbers := []int{2, 4, 6, 8, 10}
	//Создали перемнную wg с указанием на пакет sync
	var wg sync.WaitGroup
	//Цикл в котором пробегаемся по срезу
	for _, num := range numbers {
		//Установили значение на 1, чтобы увеличивать значения активных горутин
		wg.Add(1)
		//С помощью горутины выполняем нахождение квадратов
		go squareWorker(&wg, num)
	}

	// Ждем завершения всех горутин
	wg.Wait()
}
