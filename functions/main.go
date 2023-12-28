package main

import (
	"fmt"
)

func main() {
	doSomething()
	fmt.Println(addValues(3, 5))
	fmt.Println(addAllValues(4, 6, 4, 7, 8, 9, 90, 1, 2, 3, 4, 9))

	e := &Employee{"Saurabh", 23, 10}
	e.ChangeName("Ahuja")
	e.IncrementSalary(10)
	e.Print()
	e.PrintMultipleTimes(3)
}

func doSomething() {
	fmt.Println("doSomething")
}

func addValues(i int, j int) int {
	return i + j
}

func addAllValues(ijk ...int) (int, int) {
	result := 0
	for i := range ijk {
		// fmt.Println(i, ijk[i])
		result = result + ijk[i]
	}
	result2 := 0
	for _, v := range ijk {
		result2 = result2 + v
	}
	if result == result2 {
		fmt.Println("Both values are same")
		return result2, len(ijk)
	}
	return result, len(ijk)
}

type Employee struct {
	Name   string
	Id     int
	Salary float64
}

func (e *Employee) Print() {
	fmt.Println(e)
}

func (e *Employee) IncrementSalary(i int) {
	e.Salary = e.Salary + float64(i)*e.Salary/100
}

func (e *Employee) ChangeName(newname string) {
	e.Name = newname
}

func (e *Employee) PrintMultipleTimes(n int) {
	printType := fmt.Sprintln(e)
	for i := 0; i < n; i++ {
		fmt.Println(printType)
	}

}
