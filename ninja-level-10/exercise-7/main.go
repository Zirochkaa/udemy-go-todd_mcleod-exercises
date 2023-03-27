package main

import "fmt"

func main() {
	c := make(chan int)
	goroutines_amount := 10
	numbers_amount := 10

	fmt.Println("start")

	for i := 0; i < goroutines_amount; i++ {
		go func() {
			for i := 0; i < numbers_amount; i++ {
				c <- i
			}
		}()
	}

	for i := 0; i < numbers_amount*goroutines_amount; i++ {
		fmt.Println(i, <-c)
	}

	fmt.Println("end")
}
