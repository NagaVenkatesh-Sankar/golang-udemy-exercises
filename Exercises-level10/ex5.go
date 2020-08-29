package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	// close the connection
	close(c)

	// 'ok' will be false on closed connection
	v, ok = <-c
	fmt.Println(v, ok)
}
