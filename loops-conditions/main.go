package main

import "fmt"

func main() {
	i := 76

	if i > 60 {
		fmt.Println("Passed Exams with First Division")
	}
	switch {
	case i > 60:
		fmt.Println("Passed Exams with First Division")
		fallthrough
	case i > 75:
		fmt.Println("Passed Exams with Distinction")
	case i < 33:
		fmt.Println("Failed in Exams")
	}

	for j := 0; j < i; j++ {
		if j%5 == 0 {
			fmt.Println("Dvisible by 5 ", j)
		}
	}
	j := 0
	for j < i {
		fmt.Println(j)
		j = j + 10
		if j == 50 {
			goto theEnd
		}
	}

theEnd:
	fmt.Println("End of Program")

	var cart []cartItem
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 1.50, 8}
	var bananas = cartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	result := calculateTotal(cart)
	fmt.Printf("Total Cart Value is %.2f ", result)
}

type cartItem struct {
	name     string
	price    float64
	quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []cartItem) float64 {
	var result float64
	for i := range cart {
		fmt.Println(i, cart[i])
		result = result + cart[i].price*float64(cart[i].quantity)
	}
	return result
}
