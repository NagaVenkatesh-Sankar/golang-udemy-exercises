package main

import (
	"fmt"
	"golang-udemy-exercises/Exercises-Level12/dog"
)

// General canine animal type
type canine struct {
	name string
	age  int
}

// main function
func main() {
	// create an object of type
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	// print the object values
	fmt.Println(fido)

	// print the dog years
	fmt.Println(dog.YearsTwo(20))
}
