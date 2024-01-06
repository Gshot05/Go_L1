package main

import (
	"fmt"
)

// ReverseString переворачивает строку
func ReverseString(s string) string {
	// Преобразуем строку в unicode символы
	runes := []rune(s)
	length := len(runes)

	// Индексы для переворачивания
	start, end := 0, length-1

	// Переворачиваем символы в месте
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}

	// Возвращаем перевернутую строку
	return string(runes)
}

func main() {
	// Пример строки с символами Unicode
	inputString := "Шаман - я Русский"

	// Переворачиваем строку
	reversedString := ReverseString(inputString)

	// Выводим результат
	fmt.Printf("Исходная строка: %s\n", inputString)
	fmt.Printf("Перевернутая строка: %s\n", reversedString)
}
