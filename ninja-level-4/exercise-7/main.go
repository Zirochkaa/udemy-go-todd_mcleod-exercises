package main

import "fmt"

func main() {
	x := make([][]string, 2)

	x[0] = []string{"James", "Bond", "Shaken, not stirred"}
	x[1] = []string{"Miss", "Moneypenny", "Helloooooo, James."}

	for i, v := range x {
		for j, k := range v {
			fmt.Println(i, j, k)
		}
	}
}
