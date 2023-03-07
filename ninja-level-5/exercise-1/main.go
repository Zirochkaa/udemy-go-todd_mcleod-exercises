package main

import "fmt"

type person struct {
	firstName    string
	lastName     string
	favIceCreams []string
}

func main() {
	p1 := person{
		firstName:    "James",
		lastName:     "Bond",
		favIceCreams: []string{"chocolate", "martini", "rum and coke"},
	}
	p2 := person{
		firstName:    "Oleh",
		lastName:     "Me",
		favIceCreams: []string{"strawberry", "vanilla"},
	}

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for i, v := range p1.favIceCreams {
		fmt.Println(i, v)
	}

	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	for i, v := range p2.favIceCreams {
		fmt.Println(i, v)
	}
}
