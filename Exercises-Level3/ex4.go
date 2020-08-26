package main

import (
	"fmt"
)

func main() {
	year := 1988
	for {
		// break the for loop on the condition
		if year > 2020 {
			break
		}
		fmt.Println(year)
		year++
	}
}
