package main

import "fmt"

func main() {
	// сортируем элементы массива ( в обратной последовательности )
	a := []int{1, 2, 3}

	reverse(a)

	fmt.Println(a)

	// сравнение слайсов
	c := [2]int{1, 2}
	b := [2]int{1, 3}
	fmt.Println(c == b)
	//  определение типа Вы можете использовать флаг %T в пакете fmt, чтобы получить представление типа согласно синтаксису Go.

	var x interface{} = []int{1, 2, 3}
	xTap := fmt.Sprintf("%T", x)
	fmt.Println(xTap)

	// Переключатель типов позволяет выбирать между типами
	// Используйте переключатель типа, чтобы сделать несколько утверждений типа последовательно
	var m interface{} = 2.3
	switch v := m.(type) {
	case int:
		fmt.Println("int", v)
	case float64:
		fmt.Println("float64", v)
	default:
		fmt.Println("unknown")
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

}

func Egual(a, b []int) bool {
	if len(a) != len(b) {

		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true

}
