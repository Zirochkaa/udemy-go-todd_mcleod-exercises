package main

import "fmt"

func main() {
	year := 1997
	for {
		if year > 2023 {
			break
		}

		fmt.Println(year)
		year++
	}
}
