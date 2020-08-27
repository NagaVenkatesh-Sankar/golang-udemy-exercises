package main

import (
	"fmt"
)

func main() {
	// assign function to a variable
	nextThree := func(i int) []int {
		return []int{1 + i, 2 + i, 3 + i}
	}
	//call the function
	n := nextThree(10)

	//print
	fmt.Println(n)

}
