package main

import (
	"fmt"
)

// setBit устанавливает i-й бит в переменной value в значение bitValue.
func setBit(value int64, position uint, bitValue uint) int64 {
	// Создаем маску с установленным битом в заданной позиции.
	mask := int64(1) << position

	if bitValue == 1 {
		// Если bitValue равен 1, устанавливаем бит в 1 с использованием операции "ИЛИ".
		return value | mask
	} else {
		// Если bitValue не равен 1, устанавливаем бит в 0 с использованием операции "И с инверсией".
		return value &^ mask
	}
}

func main() {
	// Исходное число, в котором мы хотим установить бит.
	var number int64 = 42

	// Позиция бита, который мы хотим установить (начиная с 0).
	position := uint(2)

	// Значение, которое мы хотим установить для бита (0 или 1).
	bitValue := uint(1)

	// Вызываем функцию setBit, чтобы установить бит в переменной number.
	result := setBit(number, position, bitValue)

	// Выводим результат.
	fmt.Printf("Исходное число: %d\n", number)
	fmt.Printf("Установленный бит в позиции %d равен %d\n", position, bitValue)
	fmt.Printf("Результат: %d\n", result)
}
