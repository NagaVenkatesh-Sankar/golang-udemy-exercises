package main

import (
	"fmt"
)

func main() {
	// slice of 10 values
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// append 1 value
	x = append(x, 52)
	fmt.Println(x)

	// append more than 1 values
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	// another slice of capacity 5
	y := []int{56, 57, 58, 59, 60}

	for _, v := range y {
		x = append(x, v)
	}
	// expand the slice and append
	// x = append(x, y...)

	fmt.Println(x)

}
