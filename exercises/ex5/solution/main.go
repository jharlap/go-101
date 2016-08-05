package main

import (
	"fmt"

	"github.com/jharlap/go-101/exercises/ex5/solution/store"
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

func storeWorker(n int, db *store.InMemory, ppl <-chan Person, ids chan<- uint64) {
	for p := range ppl {
		fmt.Printf("Store worker %d storing %s\n", n, p.Name)
		db.Put(p)
		ids <- p.StoreKey()
	}
}

func fetchWorker(n int, db *store.InMemory, ids <-chan uint64) {
	for id := range ids {
		v := db.Get(id)
		if v != nil {
			p := v.(Person)
			fmt.Printf("Fetch worker %d fetched %s\n", n, p.Name)
		}
	}
}

func main() {
	db := new(store.InMemory)
	ppl := make(chan Person)
	ids := make(chan uint64)

	go storeWorker(1, db, ppl, ids)
	go storeWorker(2, db, ppl, ids)
	go fetchWorker(1, db, ids)
	go fetchWorker(2, db, ids)
	go fetchWorker(3, db, ids)

	for i, name := range []string{"Alice", "Bob", "Claire", "Damian", "Elaine", "Francesc"} {
		ppl <- Person{Name: name, AgeYears: i}
	}
	close(ppl)

	<-make(chan struct{}) // never end (don't really do this)
}

// Output:
//
// Store worker 1 storing Alice
// Fetch worker 3 fetched Alice
// Store worker 1 storing Claire
// Store worker 1 storing Damian
// Fetch worker 3 fetched Damian
// Store worker 1 storing Elaine
// Store worker 1 storing Francesc
// Fetch worker 3 fetched Francesc
// Store worker 2 storing Bob
// Fetch worker 2 fetched Elaine
// Fetch worker 3 fetched Bob
// Fetch worker 1 fetched Claire
// fatal error: all goroutines are asleep - deadlock!
