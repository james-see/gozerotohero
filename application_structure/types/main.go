package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i1 int8 = 100
	fmt.Printf("%T\n", i1)

	var numbers = [4]int{4, 5, 6, 7}
	fmt.Println(numbers)

	// like lists in python
	var cities = []string{"London", "Tokyo", "New York"}
	fmt.Println(cities)

	// maps like dicts
	balances := map[string]float64{
		"USD": 435.23,
		"EUR": 511.11,
	}
	fmt.Println(balances)

	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T", you)

	//Pointer TYPE
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is type %T and value is %v", ptr, ptr)

	// string to int and int to string
	i, _ := strconv.Atoi("-50")
	s := strconv.Itoa(20)
	fmt.Printf("\n%v is %v\n", s, i)

	type age int
	type oldAge age // int is underlying type
	type veryOldAge age

}
