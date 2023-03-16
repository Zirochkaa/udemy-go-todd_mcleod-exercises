package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p := person{
		first: "Oleh",
		last:  "Me",
		age:   42,
	}

	p.speak()
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last, "and my age is", p.age)
}
