package main

import "fmt"

func main() {
	price, inStock := 100, true
	_ = inStock
	if price < 100 {
		fmt.Println("Buy it.")
	} else if price > 100 {
		fmt.Println("Too expensive.")
	} else {
		fmt.Println("Right on edge.")
	}
}
