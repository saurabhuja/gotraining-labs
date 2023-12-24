package main

import "fmt"

func main() {
	e := &Employee{"Saurabh", 1, 54}
	e.print()
	e.changeName("Ahuja")
	e.incrementSalary()
	jsonData := e.toJson()
	fmt.Println(jsonData)

	service1 := EmployeeService(e)
	service1.incrementSalary()
	service1.print()
	fmt.Println(service1.toJson())
}
