package main

import (
	"fmt"
	"strconv"
)

/*
Используем анонимные функции на практике.

Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint, которую внутри функции main в дальнейшем можно будет вызвать по имени fn.

Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют нечетные цифры и цифра 0. Если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.

Пакет main объявлять не нужно!
Вводить и выводить что-либо не нужно!
Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.

Sample Input:

727178
Sample Output:

28
*/
func main() {
	fn := func(a uint) uint {

		fmt.Scan(&a)

		n := int(a)

		m := strconv.Itoa(n)
		strNew := ""

		for i := 0; i < len(m); i++ {

			s, _ := strconv.Atoi(string(m[i]))

			if s%2 == 0 && s != 0 {

				strNew += string(m[i])
			}

		}

		result, _ := strconv.Atoi(strNew)

		if result == 0 {
			result = 100
		}

		return uint(result)
	}

	x := fn(10)
	fmt.Println(x)

}
