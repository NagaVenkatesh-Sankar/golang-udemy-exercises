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

	// print the values
	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for _, s := range p1.favIcecream {
		fmt.Println("icecream : ", s)
	}

	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	for _, s := range p2.favIcecream {
		fmt.Println("icecream : ", s)
	}

}
