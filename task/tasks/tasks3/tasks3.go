package main

import "fmt"

// ищем пересечения массива

func main() {
	a := []int{37, 5, 1, 2}
	b := []int{6, 2, 4, 37}

	fmt.Printf("%v\n", arrange(a, b))

	x := []int{1, 1, 1}
	y := []int{1, 1, 1, 1}
	fmt.Printf("%v\n", arrange(x, y))

}

func arrange(x, y []int) []int {
	counter := make(map[int]int)
	var result []int

	for _, v := range x {
		if _, c := counter[v]; !c {
			counter[v] = 1
		} else {
			counter[v] += 1
		}
	}
	for _, v := range y {
		if cout, n := counter[v]; n && cout > 0 {
			counter[v] -= 1
			result = append(result, v)
		}

	}
	return result
}
