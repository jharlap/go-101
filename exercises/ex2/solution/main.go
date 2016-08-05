package main

import "fmt"

type Person struct {
	Name     string
	AgeYears int
}

func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.AgeYears)
}

func main() {
	ppl := []Person{
		Person{Name: "Alice", AgeYears: 11},
		Person{Name: "Bob", AgeYears: 9},
	}

	for _, p := range ppl {
		fmt.Println(p)
	}
}
