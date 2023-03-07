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

	x := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range x {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, v2 := range v.favIceCreams {
			fmt.Println(i, v2)
		}
		fmt.Println("---------------")
	}
}
