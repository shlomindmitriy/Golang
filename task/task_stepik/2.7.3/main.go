package main

import (
	"fmt"
	"strings"
)
/*
Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.
Входные данные
Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только арабские цифры.

Выходные данные
Выведите максимальную цифру, которая встречается во введенной строке.

Sample Input:
	1112221112

Sample Output:
	2
*/

func main() {
	var a string
	fmt.Scan(&a)

	n := strings.Split(a, "")
	c := n[0]

	for i := 0; i < len(n); i++ {
		if c < n[i] {
			c = n[i]
		}

	}
	fmt.Println(string(c))
}



