package main

import (
	"fmt"
)

const year = 2020

// iota will start with 0 inside a const block
const (
	year0 = year + iota
	year1 = year + iota
	year2 = year + iota
)

// iota will be reset again to 0 with the 'const'
const year3 = year + iota

func main() {
	fmt.Printf("year0 : %d\n", year0)
	fmt.Printf("year1 : %d\n", year1)
	fmt.Printf("year2 : %d\n", year2)
	fmt.Printf("year3 : %d\n", year3)
}
