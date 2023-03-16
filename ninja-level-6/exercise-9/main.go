package main

import "fmt"

func main() {
	g := func(x int) int {
		fmt.Println("I am a passed function with value", x)
		return x
	}
	x := foo(g, 42)
	fmt.Println(x)
}

func foo(f func(x int) int, x int) int {
	f(x)
	x++
	return x
}
