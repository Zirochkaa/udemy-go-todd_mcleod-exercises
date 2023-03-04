package main

import "fmt"

func main() {
	favFood := "avocado"
	if favFood == "apple" {
		fmt.Println("Your fav food is not avocado")
	} else if favFood == "avocado" {
		fmt.Printf("Your fav food is %v\n", favFood)
	} else {
		fmt.Println("You have no fav food")
	}
}
