package main

import (
	"fmt"

	"github.com/Golang/task/base_server/fail/math"
)

func main() {

	xs := []float64{1, 2, 3, 4}

	avg := math.Average(xs)

	fmt.Println(avg)
}
