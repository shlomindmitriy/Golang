package main

import (
	"fmt"
)

// наименьшее число

func largestNumber(x ...int) int {

	min := x[0]
	for i := 0; i < len(x); i++ {
		if min > x[i] {
			min = x[i]
		}
		fmt.Println(x[i])

	}

	return min

}
