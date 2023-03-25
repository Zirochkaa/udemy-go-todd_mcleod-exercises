package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("avocado")
		wg.Done()
	}()

	go func() {
		fmt.Println("mango")
		wg.Done()
	}()

	wg.Wait()
}
