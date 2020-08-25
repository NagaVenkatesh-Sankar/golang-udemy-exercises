package main

import (
	"fmt"
)

// untyped and typed constants
const (
	x     = 5
	y int = 10
)

func main() {
	fmt.Println(x, y)

	fmt.Printf("%T\t%T\n", x, y)

	// untyped number can be assigned without conversion
	var z float64
	z = x

	// cannot assign typed int const to another number type: float64 without conversion
	// z = y
	// z = float64(y)

	// print the type and the value
	fmt.Printf("%T\t%v\n", z, z)
}
