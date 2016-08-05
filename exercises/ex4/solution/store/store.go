// Package store provides an in-memory key-value store.
package store

import "sync"

// StoreKeyer is an object that can be kept in an InMemory store
type StoreKeyer interface {
	StoreKey() uint64
}

// An InMemory store is thread-safe and handles any StoreKeyer
type InMemory struct {
	sync.RWMutex
	data map[uint64]StoreKeyer
}

// Put stores a value
func (db *InMemory) Put(v StoreKeyer) {
	db.Lock()
	defer db.Unlock()

	if db.data == nil {
		db.data = make(map[uint64]StoreKeyer)
	}
	db.data[v.StoreKey()] = v
}

// Get retrieves a value
func (db *InMemory) Get(k uint64) StoreKeyer {
	db.RLock()
	defer db.RUnlock()

	return db.data[k]
}
