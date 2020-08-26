package main

import (
	"fmt"
)

func main() {
	const n = 10
	// complete for loop
	for i := 1; i <= n; i++ {
		fmt.Printf("%d\t", i)
	}

	// only with condition
	j := 0
	for j <= n {
		j++
		fmt.Printf("%d\t", j)
	}

	// empty for loop
	k := 0
	for {
		if k >= n {
			break
		}
		k++
		fmt.Printf("%d\t", k)
	}

}
