package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scanln(&x)
	fmt.Println(division(x))
}

func division(x int) (int, bool) {

	
	if x%2 == 0 {
		return x, true
	} else {
		return x, false
	}

}
