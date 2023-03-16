package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func main() {
	s := square{
		length: 3,
	}
	c := circle{
		radius: 1,
	}

	info(s)
	info(c)
}

func info(s Shape) {
	fmt.Println(s.Area())
}

func (s square) Area() float64 {
	return math.Pow(s.length, 2)
}

func (c circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
