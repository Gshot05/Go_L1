package main

import "fmt"

// removeElement удаляет i-й элемент из слайса и возвращает новый слайс.
func removeElement(slice []int, index int) []int {
	// Проверка на корректный индекс
	if index < 0 || index >= len(slice) {
		fmt.Println("Некорректный индекс для удаления.")
		return slice
	}

	// Удаление элемента из слайса
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	// Пример использования
	numbers := []int{1, 2, 3, 4, 5}
	indexToRemove := 2

	// Удаление элемента
	resultSlice := removeElement(numbers, indexToRemove)

	// Вывод результата
	fmt.Printf("Исходный слайс: %v\n", numbers)
	fmt.Printf("Слайс после удаления элемента %d: %v\n", indexToRemove, resultSlice)
}
