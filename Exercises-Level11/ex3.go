package main

import (
	"fmt"
)

type customErr struct {
}

// implement the Error Interface for CustomErr type
func (er customErr) Error() string {
	return "custom error print"
}

//accepts error as a parameter
func foo(e error) {
	fmt.Println(e.Error())
}

func main() {
	c1 := customErr{}
	foo(c1)
}
