package main

import "fmt"

// intersect возвращает пересечение двух множеств
func intersect(set1, set2 []int) []int {
	// Создаем слайс (срез) для хранения результата пересечения
	result := make([]int, 0)

	// Создаем карту (map) для быстрого поиска элементов во втором множестве
	set2Map := make(map[int]bool)
	for _, elem := range set2 {
		set2Map[elem] = true
	}

	// Проверяем каждый элемент первого множества
	for _, elem := range set1 {
		// Если элемент также присутствует во втором множестве, добавляем его в результат
		if set2Map[elem] {
			result = append(result, elem)
		}
	}

	// Возвращаем результат пересечения
	return result
}

func main() {
	// Пример использования
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	// Вызываем функцию intersect для нахождения пересечения
	intersection := intersect(set1, set2)

	// Выводим результат
	fmt.Println("Пересечение:", intersection)
}
