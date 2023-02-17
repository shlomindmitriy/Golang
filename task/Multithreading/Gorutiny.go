package main

import (
	"fmt"
	"math/rand"
	"time"
)

var c chan int

func main() {

	go f(0)
	var input string
	fmt.Scanln(&input)

	// выводит "конец" по истечению 3 секунд
	select {
	case m := <-c:
		haldre(m)
	case <-time.After(3 * time.Second):
		fmt.Println("конец")

	}
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}

}
func haldre(int) {

}
