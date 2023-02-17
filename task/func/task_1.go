package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scanln(&x)
	fmt.Println("Функция деления ")
	fmt.Println(division(x))

	fmt.Println("Наименьшее число ")
	var y [4]int
	y[0] = 9
	y[1] = 5
	y[2] = 8
	y[3] = 2
	a := largestNumber(y[0], y[1], y[2], y[3])

	fmt.Println(a)

	fmt.Println("генерация нечетных чисел")
	nextEven := makeOddGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println("число фибоначи")

	c := fib(10) //"10" число может быть любым
	fmt.Println(c)
}

// принимает число и делит его пополам , возвращает tru или false
func division(x int) (int, bool) {

	if x%2 == 0 {
		return x, true
	} else {
		return x, false
	}

}
