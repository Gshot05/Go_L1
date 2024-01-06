package main

import (
	"fmt"
	"time"
)

// Sleep - собственная функция "засыпания"
func Sleep(seconds int) {
	// Создаем канал
	done := make(chan bool)

	// Запускаем горутину для задержки
	go func() {
		time.Sleep(time.Duration(seconds) * time.Second)
		done <- true
	}()

	// Ожидаем завершения горутины
	<-done
}

func main() {
	fmt.Println("Начало выполнения программы")

	// Используем свою функцию Sleep для задержки на 2 секунды
	Sleep(2)

	fmt.Println("Программа завершена")
}
