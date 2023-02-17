package main

import "fmt"

type greeter interface {
	greet(string) string
}

type russia struct{}
type american struct{}

func (r *russia) greet(name string) string {
	return fmt.Sprintf("Привет  %s", name)
}

func (a *american) greet(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func sayHello(g greeter, name string) {
	fmt.Println(g.greet(name))
}

func main() {
	sayHello(&russia{}, "Петя ")
	sayHello(&american{}, "Bill")
}
