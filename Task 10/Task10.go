package main

import (
	"fmt"
	"sort"
)

func main() {
	// Заданные температурные значения
	temperatures := []float64{
		-25.4,
		-27.0,
		13.0,
		19.0,
		15.5,
		24.5,
		-21.0,
		32.5}

	// Создаем мапу для группировки температур
	groups := make(map[int][]float64)

	// Сортируем температурные значения для удобства
	sort.Float64s(temperatures)

	// Группируем температурные значения по диапазонам с шагом в 10 градусов
	for _, temp := range temperatures {
		groupKey := int(temp/10) * 10
		groups[groupKey] = append(groups[groupKey], temp)
	}

	// Выводим результаты группировки
	for key, values := range groups {
		fmt.Printf("%d: %v\n", key, values)
	}
}
