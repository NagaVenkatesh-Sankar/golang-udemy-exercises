package main

import (
	"fmt"
)

// composite struct type
type person struct {
	firstName   string
	lastName    string
	favIcecream []string
}

func main() {
	// set a value for person
	p1 := person{
		firstName: "Naga",
		lastName:  "Venkatesh",
		favIcecream: []string{
			"mango",
			"choco",
		},
	}
	p2 := person{
		firstName: "Naga",
		lastName:  "Tanvi",
		favIcecream: []string{
			"strawberry",
			"kiwi",
		},
	}

	// add the person to map with lastname as key
	persons := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	// print with values
	for _, p := range persons {
		fmt.Println(p.firstName)
		fmt.Println(p.lastName)
		for _, s := range p.favIcecream {
			fmt.Println("icecream : ", s)
		}
	}
}
