package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 12, 45, 23
	intSum := i1 + i2 + i3
	fmt.Println(intSum)
	f1, f2, f3 := 23.23, 45.32, 67.43
	floatSum := f1 + f2 + f3
	fmt.Println(floatSum)
	floatSum = math.Round(floatSum*100) / 100
	fmt.Println(floatSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("%.2f", circumference)
}
