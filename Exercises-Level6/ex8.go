package main

import (
	"fmt"
)

// function to return a function, that returns a string
func greet(msg string) func() string {
	return func() string {
		return "Hello, " + msg
	}
}

func main() {
	g := greet("World")
	fmt.Println(g())

	g = greet("Golang User")
	fmt.Println(g())
}
