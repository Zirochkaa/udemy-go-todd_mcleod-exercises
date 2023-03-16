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
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}

func changeMe(p *person) {
	(*p).first = "Avocado"
	p.age = 1
}
