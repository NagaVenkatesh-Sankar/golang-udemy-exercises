package main

import (
	"fmt"
)

func main() {
	//slice of 10 values
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	// first argument is a slice and then the values are expanded to append
	y := append(x[:3], x[6:]...) // [42, 43, 44, 48, 49, 50, 51]
	fmt.Println(y)
}
