package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"red", "yellow", "blue"}
	colors = append(colors, "green")
	fmt.Println(colors, &colors)
	fmt.Println(colors[1 : len(colors)-1])
	colors = append(colors[1 : len(colors)-1])
	fmt.Println(colors, &colors)

	colorSlices := make([]string, 1)
	colorSlices[0] = "red"
	fmt.Println(colorSlices, &colorSlices)

	p1 := &colorSlices
	fmt.Println(p1)
	*p1 = append(*p1, "blue")
	fmt.Println(p1, *p1)

	i1 := ([]int{23, 45, 12, 99, 9})
	p2 := &i1
	fmt.Println(*p2)
	sort.Ints(*p2)
	fmt.Println(p2)
	fmt.Println(*p2)

}
