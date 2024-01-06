package main

import (
	"fmt"
	"strings"
)

// CharactersUnique возвращает true, если все символы в строке уникальны, иначе false.
func CharactersUnique(s string) bool {
	// Приводим все символы к нижнему регистру для регистронезависимой проверки
	lowercaseStr := strings.ToLower(s)

	// Создаем мапу для отслеживания уникальных символов
	charMap := make(map[rune]bool)

	// Перебираем символы в строке
	for _, char := range lowercaseStr {
		// Если символ уже встречался, то строка не уникальна
		if charMap[char] {
			return false
		}
		// Иначе отмечаем символ как встреченный
		charMap[char] = true
	}

	// Если дошли до конца строки без повторений, то строка уникальна
	return true
}

func main() {
	// Примеры использования функции
	fmt.Println(CharactersUnique("abcd"))      // true
	fmt.Println(CharactersUnique("abCdefAaf")) // false
	fmt.Println(CharactersUnique("aabcd"))     // false
}
