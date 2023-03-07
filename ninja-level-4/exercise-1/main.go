package main

import "fmt"

func main() {
	x := [5]int{14, 33, 42, 69, 99}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
