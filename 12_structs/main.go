package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	// Shorten with the same data type
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName +
		" " + person.lastName + " and I am " + strconv.Itoa(person.age)
}

// hasBirthday method (pointer receiver) --> change age
func (person *Person) hasBirthday() {
	person.age++
}

// getMarried method (pointer) --> change lastName
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "M" {
		return
	} else {
		person.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{
		firstName: "Samantha",
		lastName:  "Smith",
		city:      "Boston",
		gender:    "F",
		age:       27,
	}

	// Alternative
	person2 := Person{
		"Bob",
		"Johnson",
		"New York",
		"M",
		30,
	}

	fmt.Println(person1.firstName)

	// person1.age++
	// fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Williams")
	person2.getMarried("Thompson")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
