package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func helloGo() {
	fmt.Println("Execute a Goroutine")
	wg.Done()
}

func main() {
	wg.Add(1)
	go helloGo()

	wg.Add(1)
	go func() {
		helloGo()
	}()

	for i := 0; i < 5; i++ {
		// wg.Add(1) that's be executed in only one goroutine
		go helloGo()
	}

	wg.Wait()
	fmt.Println(runtime.NumGoroutine())

	fmt.Println("main function has been ended here")

	fmt.Println(runtime.NumGoroutine())
}
