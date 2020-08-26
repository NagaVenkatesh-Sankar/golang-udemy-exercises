package main

import (
	"fmt"
	"strconv"
)

func main() {
	// making a slice with length and capacity
	x := make([]string, 50, 150)
	// x := make([]string, 0, 150)
	// x := make([]string, 0)
	// x := []string{}

	fmt.Println(len(x))
	fmt.Println(cap(x))

	// slice grows the length of the slice
	// the underlying array capacity is same as 150
	x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// convert int to string and append to the slice
	for i := 1; i <= 100; i++ {
		x = append(x, strconv.Itoa(i))
	}
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// print out all the slice values
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
