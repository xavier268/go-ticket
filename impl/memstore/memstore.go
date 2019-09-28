// Package memstore implements the Store interface with an in-memory data store.
package memstore

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/xavier268/go-ticket/common"
)

// MemStore is the Store implementation.
// You should usually have only one.
type MemStore struct {
	did  map[string]common.Role // deviceID to roles
	act  map[string]common.Role // pending activation requests
	rand *rand.Rand             // random generator
}

// Compiler check
var _ common.Store = new(MemStore)

// New creates a new MemStore.
func New() *MemStore {
	s := new(MemStore)
	s.did = make(map[string]common.Role)
	s.act = make(map[string]common.Role)
	// initialize random gen
	s.rand = rand.New(rand.NewSource(time.Now().UnixNano() + 9999999))
	return s
}

// GetRole  return the Device, or a zero-value (RoleNone) if no existent.
func (s *MemStore) GetRole(k string) common.Role {
	return s.did[k]
}

// SetRole set a Role.
func (s *MemStore) SetRole(k string, r common.Role) {
	s.did[k] = r
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

// Close the store - here, does nothing.
func (s *MemStore) Close() error {
	return nil
}

// Activate (only once !) using the generated requestID
func (s *MemStore) Activate(deviceID string, requestID string) (common.Role, error) {
	role, ok := s.act[requestID]
	if !ok {
		return common.RoleNone, common.ErrorInvalidActivationRequest
	}
	delete(s.act, requestID) // prevent reuse !
	return role, nil
}

// CreateRequestID generates a one-time request ID to activate any device for a given role.
func (s *MemStore) CreateRequestID(role common.Role) (rq string) {
	rq = strconv.FormatInt(s.rand.Int63(), 36)
	s.act[rq] = role
	return rq
}
