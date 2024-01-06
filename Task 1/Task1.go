package main

import "fmt"

// Human структура представляет человека
type Human struct {
	Name string
	Age  int
}

// Speak метод для Human, выводящий приветствие
func (h *Human) Speak() {
	fmt.Println("Hello, my name is", h.Name)
}

// Action структура встраивает Human
type Action struct {
	Human // встраивание Human
}

// Move метод для Action, выводящий информацию о движении
func (a *Action) Move() {
	fmt.Println(a.Name, "is moving.")
}

// Jump метод для Action, выводящий информацию о прыжке
func (a *Action) Jump() {
	fmt.Println(a.Name, "is jumping.")
}

func main() {
	// Создаем экземпляр структуры Action
	person := Action{
		Human: Human{
			Name: "John",
			Age:  25,
		},
	}

	// Используем(типо) методы у Action и Human
	person.Speak()
	person.Move()
	person.Jump()
}
