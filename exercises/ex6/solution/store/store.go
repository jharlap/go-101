// Package store provides an in-memory key-value store.
package store

import (
	"errors"
	"math/rand"
	"sync"
)

// StoreKeyer is an object that can be kept in an InMemory store
type StoreKeyer interface {
	StoreKey() uint64
}

// An InMemory store is thread-safe and handles any StoreKeyer
type InMemory struct {
	sync.RWMutex
	data map[uint64]StoreKeyer
}

var errUnexpected = errors.New("an unexpected error")

// Put stores a value
func (db *InMemory) Put(v StoreKeyer) error {
	db.Lock()
	defer db.Unlock()

	if db.data == nil {
		db.data = make(map[uint64]StoreKeyer)
	}

	if rand.Intn(10) < 5 {
		return errUnexpected
	}

	db.data[v.StoreKey()] = v
	return nil
}

// Get retrieves a value
func (db *InMemory) Get(k uint64) (StoreKeyer, error) {
	db.RLock()
	defer db.RUnlock()

	if rand.Intn(10) < 5 {
		return nil, errUnexpected
	}

	return db.data[k], nil
}
