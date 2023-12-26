package main

import "fmt"

func main() {
	var colors [3]string
	colors[0] = "red"
	colors[1] = "yellow"
	colors[2] = "blue"
	fmt.Println(colors)
	var numbers = [5]int{4, 5, 6, 3, 2}
	fmt.Println(numbers, len(numbers))
}
