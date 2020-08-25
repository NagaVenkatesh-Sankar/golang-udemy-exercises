package main

import (
	"fmt"
)

//package scoped variables declared and initialized
var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	//print formatted to a single string
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)

	//print to the terminal
	fmt.Println(s)

}
