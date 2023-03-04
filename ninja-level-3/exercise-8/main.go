package main

import "fmt"

func main() {
	switch {
	case 2 == 2:
		fmt.Println("2 == 2")
		fallthrough
	case 3 == 3:
		fmt.Println("3 == 3")
	}
}
