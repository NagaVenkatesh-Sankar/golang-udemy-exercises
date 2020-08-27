package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())
	fmt.Println(bar())

	fmt.Println(foo(), foo2())

	// below fails
	// fmt.Println(foo(), bar())

}

// single return value
func foo() int {
	return 10
}

// two return value
func bar() (int, string) {
	return 32, "Naga Venkatesh"
}

func foo2() int {
	return 10
}
