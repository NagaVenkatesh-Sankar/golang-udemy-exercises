package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		// close the connection
		close(c)
	}()

	//print all the connection values
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Exit")
}
