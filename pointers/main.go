package main

import "fmt"

func main() {
	intI := 12
	var p *int
	p = &intI
	fmt.Println(p, *p, &intI)

	floatI := 43.54
	p1 := &floatI
	fmt.Println(p1, *p1)
	*p1 = *p1 / 31
	fmt.Println(floatI)
	fmt.Println(*p1)
}
