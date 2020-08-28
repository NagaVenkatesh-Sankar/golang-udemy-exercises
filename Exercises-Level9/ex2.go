package main

import (
	"fmt"
)

type person struct {
	name string
}
type human interface {
	speak()
}

func saySomthing(h human) {
	h.speak()
}

func (p *person) speak() {
	fmt.Println("Hello, I am", p.name)
}

func main() {
	p1 := person{
		name: "Naga",
	}
	saySomthing(&p1)
}
