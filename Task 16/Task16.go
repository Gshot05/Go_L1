package main

import "fmt"

// Функция partition выполняет разделение среза на две части относительно опорного элемента.
// Возвращает отсортированный срез и индекс опорного элемента.
func partition(arr []int, low, high int) ([]int, int) {
	// Опорный элемент
	pivot := arr[high]

	// Индекс для элемента, который должен стать на свое место
	i := low

	// Проходим по элементам от low до high-1
	for j := low; j < high; j++ {
		// Если текущий элемент меньше опорного
		if arr[j] < pivot {
			// Меняем местами элементы и увеличиваем индекс i
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// Меняем опорный элемент со следующим элементом после i
	arr[i], arr[high] = arr[high], arr[i]

	// Возвращаем отсортированный срез и индекс опорного элемента
	return arr, i
}

// Функция quickSort рекурсивно сортирует срез от low до high.
// Возвращает отсортированный срез.
func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)

		// Рекурсивно сортируем левую и правую части относительно опорного элемента
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}

	return arr
}

// Функция quickSortStart является оберткой для начала быстрой сортировки.
// Возвращает отсортированный срез.
func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func main() {
	// Пример использования быстрой сортировки
	fmt.Println(quickSortStart([]int{5, 6, 7, 2, 1, 0}))
}
