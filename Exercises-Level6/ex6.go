package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("inside an ananymous function")
	}()

	fmt.Println("in outer main")

}
