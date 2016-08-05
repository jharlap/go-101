package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type User struct {
	Person   // User is a Person
	Username string
}

func (u User) String() string {
	return fmt.Sprintf("%s [ %s ]", u.Username, u.Name) // notice u.Name
}

func (u User) EquivalentString() string {
	return fmt.Sprintf("%s [ %s ]", u.Username, u.Person.Name) // also works
}

func main() {
	u := User{
		Person:   Person{Name: "Cthulu"}, // Why?
		Username: "snugglebug",
	}

	fmt.Println(u.String())
	// snugglebug [ Cthulu ]
}
