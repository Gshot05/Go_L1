package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Используем глобальную переменную wg для ожидания завершения всех воркеров
var wg sync.WaitGroup

// Создаем канал, который будет использоваться для получения сигнала о завершении программы (Ctrl+C)
var exitChan = make(chan os.Signal, 1)

func main() {
	//Получаем количество воркеров
	var numWorkers int
	fmt.Println("Сколько будет воркеров?")
	_, err := fmt.Scan(&numWorkers)
	if err != nil {
		fmt.Println("Ошибка при вводе числа воркеров:", err)
		return
	}

	//Создаем канал для передачи данных между главной горутиной и воркерами
	dataChan := make(chan int)

	//Регистрируем канал exitChan как обработчик сигналов SIGINT и SIGTERM (Ctrl+C)
	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)

	//Создаем горутины-воркеры и передаем им нужные параметры
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, dataChan, exitChan)
	}

	//Анонимная горутина для отправки данных в канал dataChan
	go func() {
		//Завершаем канал dataChan при завершении работы горутины
		defer close(dataChan)
		for i := 1; ; i++ {
			dataChan <- i
			time.Sleep(time.Second) //Задержка для наглядности
		}
	}()

	//Ожидание сигнала завершения программы (Ctrl+C)
	<-exitChan

	//Ожидание завершения всех воркеров
	wg.Wait()

	//Вывод сообщения о завершении программы
	fmt.Println("Программа завершена.")
}

// Функция worker представляет собой воркера
func worker(id int, wg *sync.WaitGroup, dataChan <-chan int, exitChan chan os.Signal) {
	defer wg.Done() // Уменьшаем счетчик ожидания при завершении работы воркера
	for {
		select {
		// Чтение данных из канала dataChan
		case data := <-dataChan:
			fmt.Printf("Worker %d: Received data %d\n", id, data)
		// Обработка сигнала завершения программы
		case <-exitChan:
			fmt.Printf("Worker %d: Shutting down\n", id)
			return
		}
	}
}
