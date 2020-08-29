package main

import (
	"fmt"
)

func main() {
	// sender channel
	// cs := make(chan<- int)

	// receiver channel
	// cr := make(<-chan int)

	// bidirectional channel
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
