package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name   string
	Id     int
	Salary float64
}

func (e *Employee) print() {
	fmt.Println(e)
}

func (e *Employee) incrementSalary() {
	e.Salary = e.Salary + 1
}

func (e *Employee) changeName(newname string) {
	e.Name = newname
}

func (e *Employee) toJson() string {
	data, _ := json.Marshal(e)
	return string(data)
}
