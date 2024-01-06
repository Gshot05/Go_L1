package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Пример переменных разных типов
	var intValue interface{} = 42
	var stringValue interface{} = "Hello, Go!"
	var boolValue interface{} = true
	var channelValue interface{} = make(chan int)

	// Определение типа переменной
	printType(intValue)
	printType(stringValue)
	printType(boolValue)
	printType(channelValue)
}

func printType(value interface{}) {
	// Получаем тип переменной с использованием либы reflect
	t := reflect.TypeOf(value)

	// Выводим тип переменной
	fmt.Printf("Значение: %v, Тип: %v\n", value, t)
}
