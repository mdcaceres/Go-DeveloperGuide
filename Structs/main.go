package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	migue := person{"Miguel",
		"Caceres",
		contactInfo{
			"mdcaceres95@gmail.com",
			5000,
		},
	}

	miguePointer := &migue
	miguePointer.updateName("Michael")
	migue.print()
}

func (p *person) updateName(name string) {
	(*p).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
