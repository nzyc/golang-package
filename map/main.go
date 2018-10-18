package main

import "fmt"

func main () {
	var numInt map[string]int
	numInt = make(map[string]int)
	numInt["one"] = 1
	numInt["two"] = 1
	fmt.Println(numInt)

	rating := make(map[string]int)
	rating["one"] = 123
	rating["two"] = 123
	fmt.Println(rating)
	delete(rating,"two")
	fmt.Println(rating)
}
