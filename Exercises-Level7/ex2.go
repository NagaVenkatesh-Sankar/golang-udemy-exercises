package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

// method 1 : pointer parameter
func changeMe(p *person) {
	// either of the below is valid for the variables
	p.last = "Tanvi"
	(*p).last = "Tanviii"
}

// method 2 : pointer receiver
func (p *person) change() {
	p.last = "Sankar"
}

func main() {
	p1 := person{
		first: "Naga",
		last:  "Venkatesh",
	}

	fmt.Println(p1)

	// pass the address as argument
	changeMe(&p1)
	fmt.Println(p1)

	// call the method
	p1.change()
	fmt.Println(p1)
}
