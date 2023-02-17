package main

import (
	"fmt"
	"math"
)

func main() {
	var k,  v float64
	fmt.Scan( &v, &k)

	// fmt.Println(T(t))
	// fmt.Println(W(k,m))
W(v,k)
	// fmt.Println(M(p, v))
	fmt.Println(W)
}

// func T(w float64) float64 {

// 	res := 6 / w
// 	return T(res)
// }
func W(k, m float64)  {
	
	t:= k / m
	t = math.Sqrt(t)
	return 
}
// func M(p, v float64) float64{
// 	m:=p*v

// 	return W
// }
