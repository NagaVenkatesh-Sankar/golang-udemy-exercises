package main

import (
	"fmt"
)

func main() {
	method1()
	method2()
}

func method1() {
	//create a channel
	c := make(chan int)

	// value is sent
	go func() {
		c <- 42
	}()

	// below is blocked, until the channel receive the value
	fmt.Println(<-c)
}

func method2() {
	// create a buffered channel for 1 element
	c := make(chan int, 1)

	// fill the buffer with value, so it is not blocked
	c <- 42

	// value is printed
	fmt.Println(<-c)
}
