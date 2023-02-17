package main

/*
	На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").

*/
import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)

	fmt.Println(CountLetters(str))

}
func CountLetters(str string) (string) {
	var g int

	res := make([]rune, 0)

	for _, v := range str {
		res = append(res, v)
	}

	g = len(res) - 1

	for i := 0; i < len(res); i++ {

		if res[i] != res[g] {
			return "Нет"
		}
		if res[i] == res[g] {
			break
		}

		g--
	}

	return "Палиндром"
}
