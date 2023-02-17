package main

import "fmt"

func main() {
	m := map[string]string{ //проверяем наличие ключа

		"port": "ship",
		"rot":  "got",
		"bot":  "nil",
	}
	_, v := m["boty"]

	fmt.Println(v)

	a := []int{1, 2, 3} //объединяем срезы в один
	b := []int{4, 5, 6}
	c := append(a, b...)
	fmt.Println(c)

}
