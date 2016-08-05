package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type User struct {
	Person   Person // User has a Person
	Username string
}

func (u User) String() string {
	return fmt.Sprintf("%s [ %s ]", u.Username, u.Person.Name)
}

func main() {
	u := User{
		Person: Person{
			Name: "Cthulu",
		},
		Username: "snugglebug",
	}

	fmt.Println(u.String())
	// snugglebug [ Cthulu ]
}
