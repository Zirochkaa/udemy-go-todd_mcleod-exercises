package main

import (
	"fmt"
)

func main() {
	first()
	second()
}

func first() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func second() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
