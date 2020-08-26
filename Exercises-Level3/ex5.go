package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		// modulo operator(%) gives the remainder value
		fmt.Printf("When %v is divided by 4, the remainder is %v\n", i, i%4)
	}
}
