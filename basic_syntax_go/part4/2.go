package main

import "fmt"

func main() {
	fmt.Print("введите фаренгейты: ")
	var c float64
	fmt.Scanf("%f", &c)
	output := ((c - 32) * 5 / 9)
	fmt.Println(output)
}
