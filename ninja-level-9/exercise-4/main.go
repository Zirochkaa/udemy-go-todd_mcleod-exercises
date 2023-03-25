package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	counter := 0
	go_amount := 100

	wg.Add(go_amount)

	for i := 0; i < go_amount; i++ {
		go func() {
			m.Lock()
			v := counter
			v++
			counter = v
			fmt.Println("Updated", counter)
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("End", counter)
}
