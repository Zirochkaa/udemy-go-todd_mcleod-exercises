package main

import (
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	// How to import `dog` package???
	lusia := canine{name: "Lusia", age: dog.Years(2)}
	fmt.Println(lusia)
}
