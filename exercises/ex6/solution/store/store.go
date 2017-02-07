// Package store provides an in-memory key-value store.
package store

import (
	"errors"
	"math/rand"
	"sync"
)

// Keyer is a type that can be kept in an InMemory store
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

var errUnexpected = errors.New("an unexpected error")

// Put stores a value
func (db *InMemory) Put(v Keyer) (int, error) {
	db.Lock()
	defer db.Unlock()

	if rand.Intn(10) < 5 {
		return 0, errUnexpected
	}

	k := v.Key()
	db.data[k] = v
	return k, nil
}

// Get retrieves a value
func (db *InMemory) Get(k int) (Keyer, error) {
	db.RLock()
	defer db.RUnlock()

	if rand.Intn(10) < 5 {
		return nil, errUnexpected
	}

	return db.data[k], nil
}
