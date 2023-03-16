package main

import (
	"fmt"
)

func main() {
	x := func(x ...int) {
		fmt.Println("You passed following arguments :", x)
	}

	x([]int{14, 42, 69}...)
}
