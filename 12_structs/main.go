package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Value Receiver Method ()
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " year old."
}

// Pointer Receiver Method ()

func (p *Person) hasBirthday() {
	p.age++
}

// Pointer Receiver example 2 Method ()
func (p *Person) isMarried(spouseName string) {
	if p.gender == "Female" {
		p.lastName = p.lastName + " " + spouseName
	} else {
		return
	}
}

func main() {
	person1 := Person{firstName: "Arnab Kumar", lastName: "Nandy", city: "Kolkata", gender: "Male", age: 27}
	person2 := Person{firstName: "Sonia", lastName: "Das", city: "Kolkata", gender: "Female", age: 26}
	fmt.Println(person1)
	person2.city = "Bengaluru"
	fmt.Println(person2)

	person1.hasBirthday()
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
	person1.isMarried("Arnab")
	person2.isMarried("Arnab")
	fmt.Println(person2.greet())
	fmt.Println(person1.greet())
}
