package main

import (
	"fmt"
)

func main() {
	markPercent := 60

	// if condition to check the remainder value
	if markPercent >= 90 {
		fmt.Println("Got 'A' grade")
	} else if markPercent >= 70 {
		fmt.Println("Got 'B' grade")
	} else if markPercent >= 40 {
		fmt.Println("Got 'C' grade")
	} else {
		fmt.Println("No grade. Need to improve")
	}
}
