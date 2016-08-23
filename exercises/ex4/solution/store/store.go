// Package store provides an in-memory key-value store.
package store

import "sync"

// Keyer is an object that can be kept in an InMemory store
type Keyer interface {
	Key() int
}

// An InMemory store handles any Keyer
type InMemory struct {
	sync.RWMutex
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
	db.Lock()
	defer db.Unlock()

	k := v.Key()
	db.data[k] = v
	return k
}

// Get retrieves a value
func (db InMemory) Get(k int) Keyer {
	db.RLock()
	defer db.RUnlock()

	return db.data[k]
}
