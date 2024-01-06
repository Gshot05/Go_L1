package main

import (
	"fmt"
	"strings"
)

// ReverseWords переворачивает слова в строке
func ReverseWords(s string) string {
	// Разделяем строку на слова
	words := strings.Fields(s)

	// Индексы для переворачивания слов
	start, end := 0, len(words)-1

	// Переворачиваем слова в месте
	for start < end {
		words[start], words[end] = words[end], words[start]
		start++
		end--
	}

	// Склеиваем перевернутые слова в строку с пробелами
	reversedString := strings.Join(words, " ")

	return reversedString
}

func main() {
	// Пример строки со словами
	inputString := "snow dog sun"

	// Переворачиваем слова
	reversedWords := ReverseWords(inputString)

	// Выводим результат
	fmt.Printf("Исходная строка: %s\n", inputString)
	fmt.Printf("Перевернутые слова: %s\n", reversedWords)
}
