package main

import (
	"fmt"
	"ninjalevel13/exercise1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	lusia := canine{
		name: "Lusia",
		age:  dog.Years(2),
	}
	fmt.Println(lusia)
	fmt.Println(dog.YearsTwo(8))
}
