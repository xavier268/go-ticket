// Package memstore implements the Store interface with an in-memory data store.
package memstore

import "github.com/xavier268/go-ticket/common"

// MemStore is the Store implementation.
// You should usually have only one.
type MemStore struct {
	m map[string]common.Role
}

// Compiler check
var _ common.Store = new(MemStore)

// New creates a new MemStore.
func New() *MemStore {
	s := new(MemStore)
	s.m = make(map[string]common.Role)
	return s
}

// GetRole  return the Device, or a zero-value (RoleNone) if no existent.
func (s *MemStore) GetRole(k string) common.Role {
	return s.m[k]
}

// SetRole set a Role.
func (s *MemStore) SetRole(k string, r common.Role) {
	s.m[k] = r
}

// UnsetRole set the role to RoleNone.
func (s *MemStore) UnsetRole(k string) {
	// delete(s.m, k)
	s.SetRole(k, common.RoleNone)
}

// Ping check availability
func (s *MemStore) Ping() error {
	return nil
}

// Close close the store - here, does nothing.
func (s *MemStore) Close() error {
	return nil
}
