package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/jharlap/go-101/exercises/ex4/solution/store"
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
	ppl := []Person{
		Person{"Alice", 1},
		Person{"Bob", 2},
		Person{"Claire", 3},
		Person{"Damian", 4},
		Person{"Elaine", 5},
	}

	db := new(store.InMemory)
	var wg sync.WaitGroup
	for _, p := range ppl {
		wg.Add(1)
		go func(p Person) {
			time.Sleep(time.Millisecond)
			db.Put(p)
			fmt.Println(p)
			wg.Done()
		}(p)
	}
	wg.Wait()
}
