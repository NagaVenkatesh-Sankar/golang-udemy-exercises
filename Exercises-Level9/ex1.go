package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	// create a waitgroup with delta 2
	var wg sync.WaitGroup
	wg.Add(2)
	// wg.Add(1)

	//go routine 1
	go func() {
		fmt.Println("inside go routine 1")
		fmt.Println("num of go routine :", runtime.NumGoroutine())
		wg.Done()
	}()

	// go routine 2
	go func() {
		fmt.Println("inside go routine 2")
		fmt.Println("num of go routine :", runtime.NumGoroutine())
		wg.Done()
	}()

	fmt.Println("inside main")
	fmt.Println("num of go routine :", runtime.NumGoroutine())
	//wait till all the go routines are done.
	wg.Wait()

	fmt.Println("Ending...")
	fmt.Println("end gs", runtime.NumGoroutine())

}
