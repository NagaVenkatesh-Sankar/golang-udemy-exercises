package main

import (
	"fmt"
)

func main() {
	// array of capacity 5
	x := [5]int{1, 2, 3, 4, 5}

	// looping through all of the array
	for i, v := range x {
		fmt.Println(i, v)
	}

	// print the type
	fmt.Printf("%T\n", x)
}
