package main

import "fmt"

type contact struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func main() {
	alex := person{firstName: "Abhinav", lastName: "Singh", contact: contact{email: "abhinav@gmail.com", zipCode: 136118}}
	updateName(alex)
	fmt.Print("%+v", alex)
}

func updateName(p *person) {
	p.firstName = "BLA Bla"
}
