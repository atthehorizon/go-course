package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Flanagan",
		contactInfo: contactInfo{
			email: "alex@flanagan.com",
			zip:   48433,
		},
	}
	alex.updateFirstName("Alexander")
	alex.print()
}

func (pp *person) updateFirstName(newName string) {
	(*pp).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
