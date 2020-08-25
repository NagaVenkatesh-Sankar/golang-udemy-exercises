package main

import (
	"fmt"
)

func main() {
	//assign the interger value
	a := 6

	// print in decimal, binary and hex format
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	// left shift 1 bit and print
	a = a << 1
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
}
