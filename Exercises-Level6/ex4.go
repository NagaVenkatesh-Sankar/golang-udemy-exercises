package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

// receiver function of the person type
func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

func main() {
	p := person{
		first: "Naga",
		last:  "Venkatesh",
		age:   32,
	}
	// call the associated function
	p.speak()
}
