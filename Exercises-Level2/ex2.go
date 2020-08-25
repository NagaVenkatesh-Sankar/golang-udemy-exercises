package main

import (
	"fmt"
)

func main() {
	// using the logical operators
	equal := 5 == 5
	lessEqual := 2 <= 5
	greatEqual := 5 >= 2
	notEqual := 2 != 5
	less := 2 < 5
	great := 5 > 2

	// print the variable results
	fmt.Println(equal)
	fmt.Println(lessEqual)
	fmt.Println(greatEqual)
	fmt.Println(notEqual)
	fmt.Println(less)
	fmt.Println(great)
}
