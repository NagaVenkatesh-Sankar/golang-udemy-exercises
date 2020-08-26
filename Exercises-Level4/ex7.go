package main

import (
	"fmt"
)

func main() {
	// two slice of strings
	s1 := []string{"James", "Bond", "Shaken, not stirred"}
	s2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	// 2 dimensional slice
	ss := [][]string{s1, s2}
	for _, s := range ss {
		for _, v := range s {
			fmt.Printf("%v\t", v)
		}
		fmt.Printf("\n")
	}

}
