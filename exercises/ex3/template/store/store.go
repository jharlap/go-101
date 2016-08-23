// Package store provides an in-memory key-value store.
package store

// Keyer is an object that can be kept in an InMemory store
type Keyer interface {
	Key() int
}

// An InMemory store handles any Keyer
type InMemory struct {
	data // add your map here
}

// New creates an InMemory store
func New() InMemory {
	// make and return a store
}

// Put stores a value
func (db *InMemory) Put(v Keyer) int {
	// insert the value into the store, in the slot determined by Key()
	// return the key used
}

// Get retrieves a value
func (db InMemory) Get(k int) Keyer {
	// return the value with the given key
}
