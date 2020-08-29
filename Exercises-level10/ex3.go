package main

import (
	"fmt"
	"time"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)
	// go routine function
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		fmt.Println("out of loop")
		// check what happes if the channel is not closed for some time
		time.Sleep(time.Second * 3)
		// close the channel, so the receiver will end the loop
		close(c)
	}()

	// returning the channel
	return c
}

func receive(c <-chan int) {
	fmt.Println("in receiver")

	// loop for all the data, till the channel is closed,
	for v := range c {
		fmt.Println(v)
	}
}
