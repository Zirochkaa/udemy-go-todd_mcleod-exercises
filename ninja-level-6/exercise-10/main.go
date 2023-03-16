package main

import "fmt"

func main() {
	f1 := foo()
	f1()
	f1()
	f1()
	f2 := foo()
	f2()
	f2()
}

func foo() func() {
	x := 10
	return func() {
		fmt.Println(x)
		x++
		fmt.Println(x)
		fmt.Println("----------")
	}
}
