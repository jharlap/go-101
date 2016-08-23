package main

import (
	"fmt"

	"github.com/jharlap/go-101/exercises/ex6/solution/store"
)

type Person struct {
	Name     string
	AgeYears int
}

func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.AgeYears)
}

func (p Person) Key() int {
	return len(p.Name) + p.AgeYears
}

func storeWorker(n int, db *store.InMemory, ppl <-chan Person, ids chan<- int) {
	for p := range ppl {
		fmt.Printf("Store worker %d storing %s\n", n, p.Name)
		k, err := db.Put(p)
		if err != nil {
			fmt.Printf("Error storing %s: %s\n", p.Name, err)
		}
		ids <- k
	}
}

func fetchWorker(n int, db *store.InMemory, ids <-chan int) {
	for id := range ids {
		v, err := db.Get(id)
		if err != nil {
			fmt.Printf("Error fetching %d: %s\n", id, err)
		}
		if p, ok := v.(Person); ok {
			fmt.Printf("Fetch worker %d fetched %s\n", n, p.Name)
		}
	}
}

func main() {
	db := store.New()
	ppl := make(chan Person)
	ids := make(chan int)

	go storeWorker(1, &db, ppl, ids)
	go storeWorker(2, &db, ppl, ids)
	go fetchWorker(1, &db, ids)
	go fetchWorker(2, &db, ids)
	go fetchWorker(3, &db, ids)

	for i, name := range []string{"Alice", "Bob", "Claire", "Damian", "Elaine", "Francesc"} {
		ppl <- Person{Name: name, AgeYears: i}
	}
	close(ppl)

	<-make(chan struct{}) // never end (don't really do this)
}

// Output:
//
// Store worker 1 storing Alice
// Error storing Alice: an unexpected error
// Store worker 1 storing Claire
// Store worker 1 storing Damian
// Error fetching 9: an unexpected error
// Store worker 2 storing Bob
// Fetch worker 3 fetched Bob
// Error fetching 8: an unexpected error
// Store worker 2 storing Francesc
// Error fetching 13: an unexpected error
// Store worker 1 storing Elaine
// Error storing Elaine: an unexpected error
// Error fetching 10: an unexpected error
// fatal error: all goroutines are asleep - deadlock!
