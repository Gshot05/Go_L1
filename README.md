Ответы на Устные Вопросы по Go

Первый вопрос
 
Q: Как создать строку в Go с использованием strings.builder?

A: В Go для создания строки с использованием strings.Builder:

```go
var builder strings.Builder
builder.WriteString("Hello, ")
builder.WriteString("world!")
result := builder.String()
```

Также можно использовать bytes.Buffer аналогичным образом.


Второй вопрос

Q: Что такое интерфейсы в Go?

A: В Go интерфейсы - это определение и описание набора методов, которые должны быть реализованы другими типами. Они широко используются для достижения полиморфизма и абстракции.


Третий вопрос

Q: Какова разница между Mutex и RWMutex в Go?

A: Mutex полностью блокирует канал для записи или чтения, в то время как RWMutex позволяет блокировать отдельно чтение и запись в канал.


Четвертый вопрос

Q: В чем разница между небуферизированным и буферизированным каналом в Go?

A:
Небуферизированный канал не имеет размера (по умолчанию 0).
Буферизированный канал имеет четкий размер, что позволяет буферизовать определенное количество элементов.


Пятый вопрос

Q: Какой размер у структуры struct{}{} в Go?

A: У такой структуры размер - 0, поскольку у структуры вес имеют только её поля.


Шестой вопрос

Q: Поддерживает ли Go перегрузку методов и функций?

A: Нет, в Go отсутствует перегрузка методов и функций. Рекомендуется просто давать разные имена методам и функциям.


Седьмой вопрос

Q: В каком порядке происходит итерация по мапе в Go?

A: Порядок итерации по мапе в Go случайный, и не существует гарантированного порядка элементов. Порядок может меняться при каждой итерации.


Восьмой вопрос

Q: Какова разница между new и make в Go?

A:
new инициализирует нулевое значение для данного типа и возвращает указатель на этот тип.
make используется исключительно для создания и инициализации срезов, отображений и каналов, возвращая ненулевой экземпляр указанного типа.


Девятый вопрос

Q: Как создать срез в Go?

A:

Для slice:

С помощью make:

```go
slice := make([]int, 5, 10)
```
 
С использованием литерала среза:

```go
slice := []int{1, 2, 3, 4, 5}
```


Срез среза (подсрез):

```go
originalSlice := []int{1, 2, 3, 4, 5}
subSlice := originalSlice[1:3] 
```

Для map:

С использованием make:

```go
myMap1 := make(map[string]int) // Создание пустой мапы
```

С использованием литерала мапы:

```go
myMap2 := map[string]int{
    "one":   1,
    "two":   2,
    "three": 3,
}
```

Использование функции new:

```go
myMap3 := new(map[string]int)
*myMap3 = make(map[string]int)
```

С использованием var:
```go
var myMap4 map[string]int 
```


Инициализация мапы и добавление элементов:

```go
myMap5 := map[string]int{}
myMap5["one"] = 1
myMap5["two"] = 2
```


Десятый вопрос:

Q: Что выведет данная программа и почему?

```go
func update(p *int) {
    b := 2
    p = &b
}
func main() {
var (
    a = 1
    p = &a
)
    fmt.Println(*p)
    update(p)
    fmt.Println(*p)
}
```

A:

Ответ будет:

```sh
1
1
```

Так как функция update получает копию указателя на переменную a, поэтому изменения внутри функции не влияют на оригинальный указатель в функции main.


Одиннадцатый вопрос

Q: Что выведет данная программа и почему?

```go
func main() {
    wg := sync.WaitGroup{}
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
    }(wg, i)
}
    wg.Wait()
    fmt.Println("exit")
}
```

A:

В этом коде есть ошибка в том, как используется sync.WaitGroup. Параметр wg внутри анонимной функции передается по значению (копируется), и поэтому каждая горутина получает свою копию sync.WaitGroup, которая не взаимодействует с оригинальным объектом wg.

Чтобы исправить это, надо передавать указатель на sync.WaitGroup вместо значения.


Двенадцатый вопрос:

Q: Что выведет данная программа и почему?

```go
func main() {
n := 0
if true {
    n := 1
    n++
}


fmt.Println(n)
}

```
A:


Данная программа выведет 0, так как в if-casе у нас создается так называемая теневая переменная, которая существует только в данном условии, если мы хотим получить результат, надо сделать "=", а не ":="



Тринадцатый вопрос

Q: Что выведет данная программа и почему?

```go
func someAction(v []int8, b int8) {
    v[0] = 100
    v = append(v, b)
}
func main() {
   var a = []int8{1, 2, 3, 4, 5}
   someAction(a, 6)
   fmt.Println(a)
}
```

A:

Данный код вывыедет 

```sh 
[100 2 3 4 5]
 ```

Функция someAction изменяет значение первого элемента среза v (который фактически ссылается на тот же срез, что и a) на 100. Таким образом, срез a в функции main будет
 ```sh
{100, 2, 3, 4, 5}
```


Четырнадцатый вопрос


Q: Что выведет данный код и почему?

```go
func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)

	fmt.Print(slice)
}
```

A: Данный код вывыдее [b b a] [a a], так как у нас есть слайс с содержимым [a a], а в анонимную функцию мы передаём копию оригинального слайса, где добовляем "a", и заменяем первые два элемента на "b" 
