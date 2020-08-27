package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

// embed a type within another type
type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	// set values without the identifiers
	t := truck{
		vehicle{2, "black"},
		true,
	}

	// set values with the identifiers
	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: false,
	}

	// print all of the values
	fmt.Println(t)
	fmt.Println(s)

	// print a single value
	fmt.Println(t.doors)
	fmt.Println(s.doors)
}
