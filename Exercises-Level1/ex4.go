package main

import (
	"fmt"
)

//own type with underlying type as 'int'
type mynum int

var x mynum

func main() {
	//print default value
	fmt.Println(x)

	//print the type
	fmt.Printf("%T\n", x)

	//assign the int value
	x = 42

	//print the vaue
	fmt.Println(x)
}
