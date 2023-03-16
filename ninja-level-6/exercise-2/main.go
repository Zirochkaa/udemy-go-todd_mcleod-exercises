package main

import "fmt"

func main() {
	a := foo([]int{5, 6, 2}...)
	b := bar([]int{2, 3, 5})
	fmt.Println(a)
	fmt.Println(b)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
