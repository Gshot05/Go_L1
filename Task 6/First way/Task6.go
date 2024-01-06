// Способ 1: Использование канала для сигнала завершения
package main

import (
	"fmt"
	"time"
)

func main() {
	// Канал для остановки горутины
	stopChan := make(chan bool)

	// Запускаем горутину
	go func() {
		//Бесконечный цикл в горутине
		for {
			select {
			//Если получен сигнал из канала stopChan
			case <-stopChan:
				fmt.Println("Горутина завершена.")
				return
			//Если нет данных в канале stopChan, продолжаем работу
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(time.Second)
			}
		}
	}()

	//Главная горутина ждет 3 секунды
	time.Sleep(3 * time.Second)
	fmt.Println("Главная горутина отправляет сигнал завершения в другую горутину.")
	//Отправляем сигнал в канал stopChan
	stopChan <- true

	//Даем время для завершения другой горутины (ждем 2 секунды)
	time.Sleep(2 * time.Second)
	//Выводим сообщение о завершении программы
	fmt.Println("Программа завершена.")
}
