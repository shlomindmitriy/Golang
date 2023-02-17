package main

import (
	"fmt"

	"github.com/shlomindmitriy/Golang/base_server/math"
)

func main() {

	xs := []float64{}
	fmt.Println("Среднее значение:")
	avg := math.Average(xs)
	fmt.Println(avg)

	fmt.Println("Максимальное число:")
	max := Max(xs)
	fmt.Println(max)

	fmt.Println("Минемальное число:")
	min := Min(xs)
	fmt.Println(min)
}

func Max(xs []float64) float64 {

	max := xs[0]

	for i := 0; i < len(xs); i++ {

		if max < xs[i] {
			max = xs[i]

		}

	}

	return max

}

func Min(xs []float64) float64 {

	min := xs[0]

	for i := 0; i < len(xs); i++ {
		if min > xs[i] {

			min = xs[i]
		}
	}
	return min
}
