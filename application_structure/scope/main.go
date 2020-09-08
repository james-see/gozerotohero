package main

import "fmt"

const done = false // package scope

func main() {
	x := 10 // local scope
	fmt.Printf("Hello %d\n", x)
	fmt.Println(done)
}
