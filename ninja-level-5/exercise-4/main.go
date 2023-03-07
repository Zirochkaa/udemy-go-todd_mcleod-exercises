package main

import (
	"fmt"
)

func main() {
	x := struct {
		map_example   map[string]string
		slice_example []int
	}{
		map_example: map[string]string{
			"lol":  "wut",
			"lol2": "wut2",
		},
		slice_example: []int{13, 23, 69},
	}

	fmt.Println(x.map_example)
	fmt.Println(x.slice_example)
}
