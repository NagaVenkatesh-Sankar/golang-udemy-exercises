package main

import (
	"fmt"
)

func main() {
	var x int
	x = 5

	fmt.Println("Main function")

	// print the address location of the variable
	fmt.Println(&x)
}
