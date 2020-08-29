package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			//send the data in 'c' channel
			c <- i
		}
		// send the data in 'q' channel
		q <- 1
		// close(c)
	}()

	return c
}

func receive(c, q <-chan int) {
	// even if the channel isnt closed, we get the quit signal from another channel
	for {
		// select case if the data is received from channel 'c' or 'q'
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			{
				fmt.Println("Quit signal received")
				// out of loop in 'q' channel data
				return
			}
		}
	}
}
