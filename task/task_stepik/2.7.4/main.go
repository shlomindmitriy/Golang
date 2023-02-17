package main

import (
	"fmt"
	"strconv"
)

/*
На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число. 
Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1. В итоге получаем 811181

Sample Input:
	9119

Sample Output:
	811181
*/

func main() {
	var a int
	fmt.Scan(&a)
	rez := make([]int, 0)
		
	for a > 0 {
		new := 0
		new += a % 10
		a /= 10

		rez = append(rez, new)
		if a == 0 {
			break
		}
	}

	var reverse []int

	for i := len(rez) - 1; i >= 0; i-- {
		reverse = append(reverse, rez[i])
	}

	var sum string
	for _, v := range reverse {
		sum += strconv.Itoa(v * v)

	}
	fmt.Println(sum)

}
