Устные вопросы:


1) С помощью strings.builder, но также можно использовать bytes.Buffer


2) Интерфейсы в go - это определение и описание каких-то конкретных методов которые должны быть у другого типа


Пример интерфейсов:

// Определяем интерфейс


type Shape interface {
    Area() float64
}

// Реализуем интерфейс для типа Circle


type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// Реализуем интерфейс для типа Rectangle


type Rectangle struct {
    Width, Height float64
}


func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}



func printArea(s Shape) {
    fmt.Println("Площадь:", s.Area())
}



func main() {


    // Создаем экземпляры типов Circle и Rectangle


    circle := Circle{Radius: 5.0}


    rectangle := Rectangle{Width: 4.0, Height: 6.0}


}



