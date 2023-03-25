package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	counter := 0
	go_amount := 100

	wg.Add(go_amount)

	for i := 0; i < go_amount; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("Updated", counter)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("End", counter)
}
