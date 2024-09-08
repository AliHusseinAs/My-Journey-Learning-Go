package main

import (
	pr "fmt"
)

func main() {
	E := Employee{
		Name:   "Ali",
		Salary: 5000,
	}
	pr.Println(printName(E))
	pr.Println(printSalary(E))
}
