package main

import (
	"fmt"

	"github.com/jharlap/go-101/exercises/ex3/solution/store"
)

type Person struct {
	Name     string
	AgeYears int
}

func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.AgeYears)
}

func (p Person) StoreKey() uint64 {
	return uint64(len(p.Name) + p.AgeYears)
}

func main() {
	a := Person{Name: "Alice", AgeYears: 11}
	b := Person{Name: "Bob", AgeYears: 9}

	db := new(store.InMemory)
	db.Put(a)
	db.Put(b)

	fmt.Println(db)
	// &{map[12:{Bob 9} 16:{Alice 11}]}

	va := db.Get(a.StoreKey())
	fmt.Println(va)
	// Alice is 11 years old
}
