package main

import "fmt"

func main() {
	fmt.Print("введите футы:")
	var t float64
	fmt.Scanf("%f", &t)
	output := t / 3.281
	fmt.Println(output)
}
