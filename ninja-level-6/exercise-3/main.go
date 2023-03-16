package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello, avocado")
}

func foo() {
	defer func() {
		fmt.Println("Thanks")
	}()
	fmt.Println("Hello to you too")
}
