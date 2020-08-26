package main

import (
	"fmt"
)

func main() {
	// create a slice
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// loop through all the slice values
	for i, v := range x {
		fmt.Println(i, v)
	}

	// print the type
	fmt.Printf("%T\n", x)
}
