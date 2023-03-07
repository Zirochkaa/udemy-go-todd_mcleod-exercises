package main

import "fmt"

func main() {
	x := []int{14, 33, 42, 69, 99, 114, 133, 142, 169, 199}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
