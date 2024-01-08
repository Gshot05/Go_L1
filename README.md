Ответы на Устные Вопросы по Go

Первый вопрос
 
Q: Как создать строку в Go с использованием strings.builder?

A: В Go для создания строки с использованием strings.Builder:

```sh
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

```sh
slice := make([]int, 5, 10)
```
 
С использованием литерала среза:

```sh
slice := []int{1, 2, 3, 4, 5}
```


Срез среза (подсрез):

```sh
originalSlice := []int{1, 2, 3, 4, 5}
subSlice := originalSlice[1:3] 
```

Для map:

С использованием make:

```sh
myMap1 := make(map[string]int) // Создание пустой мапы
```

С использованием литерала мапы:

```sh
myMap2 := map[string]int{
    "one":   1,
    "two":   2,
    "three": 3,
}
```

Использование функции new:

```sh
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
