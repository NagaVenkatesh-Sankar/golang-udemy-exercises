package main

import "fmt"

func main() {
	ni := 10
	nj := 10
	c := make(chan int)
	for i := 0; i < ni; i++ {
		// 10 go routines
		go func() {
			for j := 0; j < nj; j++ {
				// pass value to channel
				c <- j
			}

		}()

	}
	//print for all the values
	for i := 0; i < ni*nj; i++ {
		fmt.Println(i, <-c)
	}
}
