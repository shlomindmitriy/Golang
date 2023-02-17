package main

import "fmt"

/*
Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку

Sample Input:
	zaabcbd
Sample Output:
	zcd
*/

func main() {
	var a string
	fmt.Scan(&a)

	res := []rune(a)

	rez := make([]rune, 0)

	for i := 0; i < len(res); i++ {
		var flag bool = false
		
		for j := 0; j < len(res); j++ {
			if res[i] == res[j] && i != j {
				flag = true
				break
			}

		}
		if !flag {
			rez = append(rez, res[i])

		}

	}
	fmt.Println(string(rez))

}
