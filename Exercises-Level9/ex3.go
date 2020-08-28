package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg1, wg2 sync.WaitGroup
	//race condition method 1
	incrementer1 := 0
	gs1 := 100
	wg1.Add(gs1)

	for i := 0; i < gs1; i++ {
		go func() {
			v := incrementer1
			// yeild the processor which allow other go routines to execute
			runtime.Gosched()
			v++
			incrementer1 = v
			// fmt.Println(incrementer1)
			wg1.Done()
		}()
	}
	wg1.Wait()
	fmt.Println("incrementer 1 value:", incrementer1)

	//race condition method 2
	incrementer2 := 0
	gs2 := 100
	wg2.Add(gs2)

	for i := 0; i < gs2; i++ {
		go func() {
			v := incrementer2
			v++
			incrementer2 = v
			// fmt.Println(incrementer2)
			wg2.Done()
		}()
	}
	wg2.Wait()
	fmt.Println("incrementer 2 value:", incrementer2)

	//race condition method 3
	incrementer3 := 0
	gs3 := 100
	for i := 0; i < gs3; i++ {
		go func() {
			w := incrementer3
			w++
			incrementer3 = w
		}()
	}
	fmt.Println("incrementer 3 value:", incrementer3)
}
