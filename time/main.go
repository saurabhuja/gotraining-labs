package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	t := time.Date(2009, 12, 1, 12, 23, 12, 21, time.UTC)
	fmt.Println(t)
	fmt.Println(t.Format(time.ANSIC))
	parsedTime, _ := time.Parse(time.ANSIC, "Tue Dec  1 12:23:12 2009")
	fmt.Println(parsedTime)
}
