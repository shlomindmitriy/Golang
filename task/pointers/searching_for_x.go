package main

import "fmt"

// находим значение" x"
func main() {
	x := 1.5
	a := square(&x)
	fmt.Println(a)

	var g, y int
	g = 1
	y = 2

	g1, y1 := swap(&g, &y)
	fmt.Println(*g1, *y1)
	fmt.Println(g, y)

}

func square(x *float64) float64 {
	*x = *x * *x
	return *x
}

// меняет  местами два числа !!
func swap(b *int, c *int) (*int, *int) {
	sw := *b
	*b = *c
	*c = sw
	return b, c
}
