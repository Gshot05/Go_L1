package main

import "strings"

var justString string

// Функция someFunc выполняет следующие действия:
// Создает огромную строку v с использованием функции createHugeString.
// Создает срез байтов нужного размера (100 байт) с помощью make.
// Приводит срез байтов к строке и присваивает justString этой строкой.
// Копирует данные из строки v в justString, обрезая до 100 байт.
func someFunc() {
	v := createHugeString(1 << 10)
	justString = string(make([]byte, 100))
	copy([]byte(justString), []byte(v))
}

// Функция createHugeString создает строку, повторяя символ "a" указанное количество раз (size).
// В данном случае, создается строка из 2^10 символов "a".
func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

// Вызывает функцию someFunc для выполнения операций по обработке строк.
func main() {
	someFunc()
}
