package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var counter int64

	go_amount := 100

	wg.Add(go_amount)

	for i := 0; i < go_amount; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Updated", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("End", counter)
}
