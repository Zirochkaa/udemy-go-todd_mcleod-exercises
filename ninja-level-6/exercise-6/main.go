package main

import (
	"fmt"
)

func main() {
	func(x ...int) {
		fmt.Println("You passed following arguments :", x)
	}([]int{14, 42, 69}...)
}
