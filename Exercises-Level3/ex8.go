package main

import (
	"fmt"
)

func main() {
	// no switch expression is given
	switch {
	// case condition is evaluated for first occurance of 'true'
	case false:
		fmt.Println("It will not print")
	case true:
		fmt.Println("It prints")
	}
}
