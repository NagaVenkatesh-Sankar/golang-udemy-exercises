package main

import (
	"fmt"
)

func main() {
	// ASCII number for A to Z
	for i := 65; i <= 90; i++ {
		fmt.Println(i)

		// inner for loop
		for j := 0; j < 3; j++ {

			// unicode format: #U
			fmt.Printf("\t%#U\n", i)
		}
	}
}
