package main

import (
	"fmt"
)

const seconds = 3600

func main() {
	fmt.Println("Hello Go World!")
	distance := 60.8

	fmt.Printf("The distance in miles is %f \n", distance*0.62137)
}
