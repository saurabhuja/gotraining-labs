package main

import "fmt"

func main() {
	// m := new(map[string]int) // using new instead of make doesnt allocate memory
	m := make(map[string]int)
	m["i"] = 23
	fmt.Println(m)
}
