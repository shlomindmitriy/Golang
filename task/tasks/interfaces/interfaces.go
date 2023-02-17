package main

import "fmt"

type animal interface { // animal - какоето обстрактное животное
	makeSound() // makeSaund -издавать звук
}

type cat struct{}

type dog struct{}

func (c *cat) makeSound() {
	fmt.Println("meow")

}

func (d *dog) makeSound() {
	fmt.Println("gaw")
}

func main() {
	var c, d animal = &cat{}, &dog{}
	c.makeSound()
	d.makeSound()

}
