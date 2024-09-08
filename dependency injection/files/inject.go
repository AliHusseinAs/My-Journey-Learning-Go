package main

import "strconv"

// dependency

func (E Employee) getName() string {
	return "Name is " + E.Name
}

func (E Employee) getSalary() string {
	salar := strconv.FormatInt(E.Salary, 10)
	return "Salary is : " + salar
}
