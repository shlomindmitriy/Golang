package main

import (
	"fmt"
)

// инициализация структур
type person struct {
	age  int
	name string
}

// Возвращаем несколько значений функции

func main() {
	age, name := add(25, 45, "Петр", "Петров")

	fmt.Println(age)
	fmt.Println(name)

	tom := person{name: "Юра", age: 25}
	fmt.Println(tom)

}

func add(x, y int, fristName, lastName string) (int, string) {
	a := x + y
	fullName := fristName + " " + lastName
	return a, fullName
}
