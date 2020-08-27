package main

import (
	"fmt"
)

func main() {
	fmt.Println("in main function")
	foo()
}

func foo() {
	// defer will be executed at the end
	defer func() {
		fmt.Println("end of foo function")
	}()

	fmt.Println("inside foo function")
}
