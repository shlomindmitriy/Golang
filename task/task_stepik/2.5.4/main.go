package main

import "fmt"

/*
 На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

Sample Input:
	ihgewlqlkot

Sample Output:
	hello

*/

func main() {
	var a string
	fmt.Scan(&a)

	for i, v := range a {

		if i%2 != 0 {
			
			
			fmt.Println(string(v))
		}
	}
}
