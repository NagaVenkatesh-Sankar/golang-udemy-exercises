package main

import (
	"fmt"
)

func main() {
	i1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := foo(i1...)
	fmt.Println(s1)

	s3 := foo(10, 20, 30)
	fmt.Println(s3)

	i2 := []int{1, 2, 3, 4, 5}
	s2 := bar(i2)
	fmt.Println(s2)
}

// variadic function
func foo(x ...int) int {
	fmt.Println("foo", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	fmt.Println("bar", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
