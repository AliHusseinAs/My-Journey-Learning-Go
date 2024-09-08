package main

import (
	pr "fmt"
)

type Person struct {
	Name string
}

type funcs interface {
	Greeting() string
}

func (p Person) Greeting() string {
	return "hello" + p.Name
}

func printGree(f funcs) string {

	// pr.Println(f.Greeting())
	return f.Greeting()
}

func main() {
	p := Person{
		Name: "Ali"}
	pr.Println(printGree(p))
}
