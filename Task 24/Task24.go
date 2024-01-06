package main

import (
	"fmt"
	"math"
)

// Point - структура, представляющая точку в двумерном пространстве.
type Point struct {
	x, y float64
}

// newPoint - конструктор для создания новой точки.
func newPoint(x, y float64) Point {
	return Point{x, y}
}

// Distance - метод для вычисления расстояния между текущей точкой и другой точкой.
func (p Point) Distance(q Point) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки
	point1 := newPoint(1.0, 2.0)
	point2 := newPoint(4.0, 6.0)

	// Вычисляем расстояние между точками
	distance := point1.Distance(point2)

	// Вывод результата
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
