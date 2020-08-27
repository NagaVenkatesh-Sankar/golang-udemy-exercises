package main

import (
	"fmt"
)

func main() {
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	g := foo()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())

	fmt.Println(f())
	fmt.Println(g())
}

// function returing a function
// x is scoped to the function
func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
