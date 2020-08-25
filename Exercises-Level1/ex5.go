package main

import (
	"fmt"
)

// own type of underlying int
type hotdog int

var x hotdog
var y int

func main() {
	//print the default value and type of the own type variable
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// assign the value to own type
	x = 42

	//print the value
	fmt.Println(x)

	//convert the own type with underlying type to another underlying type variable
	y = int(x)

	// print the value and the type of the converted variable

	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
