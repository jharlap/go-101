// Package store provides an in-memory key-value store.
package store

// Keyer is a type that can be kept in an InMemory store
type Keyer interface {
	Key() int
}

// An InMemory store handles any Keyer
type InMemory struct {
	data map[int]Keyer
}

// New creates an InMemory store
func New() InMemory {
	return InMemory{
		data: make(map[int]Keyer),
	}
}

// Put stores a value
func (db *InMemory) Put(v Keyer) int {
	k := v.Key()
	db.data[k] = v
	return k
}

// Get retrieves a value
func (db InMemory) Get(k int) Keyer {
	return db.data[k]
}
