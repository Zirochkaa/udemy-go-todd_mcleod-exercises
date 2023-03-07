package main

import (
	"fmt"
	"ninjalevel12/exercise2/quote"
	"ninjalevel12/exercise2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
