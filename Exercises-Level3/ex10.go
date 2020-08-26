package main

import (
	"fmt"
)

func main() {
	// 'AND' operator
	fmt.Printf("true && true : %v\n", true && true)
	fmt.Printf("true && false : %v\n", true && false)

	// 'OR' operator
	fmt.Printf("true || true : %v\n", true || true)
	fmt.Printf("true || false : %v\n", true || false)

	// 'NOT' operator
	fmt.Printf("!true : %v\n", !true)
}
