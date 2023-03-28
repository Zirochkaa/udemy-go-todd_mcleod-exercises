package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func main() {
	c := customErr{
		info: "avocado",
	}

	foo(c)
}

func (c customErr) Error() string {
	return fmt.Sprintf("Additional info - %v", c.info)
}

func foo(e error) {
	fmt.Println(e)
}
