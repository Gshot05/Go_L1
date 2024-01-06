package main

import (
	"fmt"
)

func main() {
	// Предположим, что a и b - числовые переменные, значения которых > 2^20.
	a := float64(2<<20) + 1
	b := float64(2<<20) + 3

	// Умножение
	multiplicationResult := a * b
	fmt.Printf("Умножение: %.2f * %.2f = %.2f\n", a, b, multiplicationResult)

	// Деление
	divisionResult := a / b
	fmt.Printf("Деление: %.2f / %.2f = %.2f\n", a, b, divisionResult)

	// Сложение
	additionResult := a + b
	fmt.Printf("Сложение: %.2f + %.2f = %.2f\n", a, b, additionResult)

	// Вычитание
	subtractionResult := a - b
	fmt.Printf("Вычитание: %.2f - %.2f = %.2f\n", a, b, subtractionResult)
}
