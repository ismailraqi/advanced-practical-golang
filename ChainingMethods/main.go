package main

import (
	"fmt"
	"strings"
)

// Person holds the data structure for the current person info
type Person struct {
	Name string
	Age  int
}

// PR initializes the 'Person' struct with an empty values
var PR = Person{}

// PersonName is the first chain layer
func PersonName(personName string) *Person {
	if len(strings.TrimSpace(personName)) == 0 {
		return &Person{}
	}
	newPerson := Person{
		Name: personName,
	}
	fmt.Println("1st chain")
	return &newPerson
}

// AgePerson is the 2nd chain layer
func (p *Person) AgePerson(age int) *Person {
	p.Age = age
	fmt.Println("2nd chain")
	return p
}

// AddPerson is the 3rd chain layer
func (p *Person) AddPerson() *Person {
	fmt.Println("3rd chain")
	fmt.Println("person name: ", p.Name, " age: ", p.Age)
	return p
}

func main() {
	n := PersonName("Ismail").AgePerson(21).AddPerson()
	fmt.Println("n:", n.Name)
}
