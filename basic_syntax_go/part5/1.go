package main

import (
	"fmt"
	"strings"
)

func main() {

	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
			switch i {
			case 0:
				fmt.Println("zero")
			case 1:
				fmt.Println("one")
			case 2:
				fmt.Println("two")
			case 3:
				fmt.Println("three")
			case 4:
				fmt.Println("four")
			case 5:
				fmt.Println("five")
			default:
				fmt.Println("unknown number")
			}
		}

	}
	fmt.Println(strings.Contains("test", "es"),
		strings.Count("test", "t"),
		strings.HasPrefix("test", "te"),
		strings.HasSuffix("test", "st"),
		strings.Index("test", "e"),
		strings.Join([]string{"a", "b"}, "-"),
		strings.Repeat("a", 5),
		strings.Replace("aaaa", "a", "b", 2),
		strings.Split("a-b-c-d-e", "-"),
		strings.ToLower("TEST"),
		strings.ToUpper("test"))

}
