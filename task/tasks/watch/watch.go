package main

import "fmt"

func main() {
	// вариант первый
	var a int
	fmt.Scan(&a)
	h := (a / 30)
	m := (a * 2) % 60
	fmt.Println("It is", h, "hours", m, "minutes.")
}

//  вариант два

// func main(){
// var a int
// 	fmt.Scan(&a)
// 	h := (a*2)/60
// 	m :=(a*2)%60
// 	fmt.Println("It is",h,"hours" ,m,"minutes.")
// }
