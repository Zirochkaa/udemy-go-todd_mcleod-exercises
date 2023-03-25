package main

import "fmt"

type person struct {
	Name string
}

type human interface {
	speak() string
}

func main() {
	p := person{Name: "Oleh"}
	fmt.Println(p.speak())

	// saySomething(p) // Will not work
	saySomething(&p)
}

func (p *person) speak() string {
	return "My name is " + p.Name
}

func saySomething(h human) {
	fmt.Println("Human said -", h.speak())
}
