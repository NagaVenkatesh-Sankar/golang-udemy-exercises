package main

import (
	"fmt"
)

func main() {
	// create a slice of 10 values
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	//slicing the slice
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(x)

	//runtime error : out of range, as the capacity is only 10
	// fmt.Println(x[5:11])
}
