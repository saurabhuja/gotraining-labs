package main

type EmployeeService interface {
	print()
	incrementSalary()
	toJson() string
}
