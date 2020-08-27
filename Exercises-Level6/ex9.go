package main

import (
	"fmt"
)

func main() {
	math(2, 3, add)
	math(2, 3, multiply)
}

func add(x, y int) {
	fmt.Println(x + y)
}

func multiply(x, y int) {
	fmt.Println(x * y)
}

func math(x int, y int, f func(a, b int)) {
	f(x, y)
}
