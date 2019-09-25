// Package memstore implements the Store interface with an in-memory data store.
package memstore

import "github.com/xavier268/go-ticket/common"

// MemStore is the Store implementation.
// You should usually have only one.
type MemStore struct {
	m map[string]interface{}
}

// Compiler check
var _ common.Store = new(MemStore)

// New creates a new MemStore.
func New() *MemStore {
	s := new(MemStore)
	s.m = make(map[string]interface{})
	return s
}

// Get the value stored.
func (s *MemStore) Get(k string) interface{} {
	return s.m[k]
}

// Set a value.
func (s *MemStore) Set(k string, v interface{}) {
	s.m[k] = v
}

// Ping check availability
func (s *MemStore) Ping() error {
	return nil
}
