package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	jim := person{
		firstName: "jim",
		lastName: "jimov",
		contactInfo: contactInfo{
			email: "pesho@peshov.bg",
			zipCode: 6969,
		},
	}

	jim.updateName("peshkata")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}