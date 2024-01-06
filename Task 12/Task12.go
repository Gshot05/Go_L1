package main

import "fmt"

func main() {
	// Создаем пустое множество строк
	mySet := make(map[string]bool)

	// Заданная последовательность строк
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Добавляем каждую строку в множество
	for _, word := range sequence {
		mySet[word] = true
	}

	// Выводим уникальные элементы множества
	fmt.Println("Собственное множество:", getKeys(mySet))
}

func getKeys(set map[string]bool) []string {
	// Создаем срез строк для хранения ключей
	keys := make([]string, 0, len(set))

	// Итерируем по всем ключам в множестве
	for key := range set {
		// Добавляем текущий ключ в срез
		keys = append(keys, key)
	}

	// Возвращаем срез, содержащий все ключи из множества
	return keys
}
