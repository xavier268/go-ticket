// Package memstore implements the Store interface with an in-memory data store.
package memstore

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/xavier268/go-ticket/common"
)

// MemStore is the Store implementation.
// You should usually have only one.
type MemStore struct {
	did  map[string]common.Role   // deviceID to roles
	act  map[string]common.Role   // pending activation requests
	tkt  map[string]common.Ticket // ticket database
	rand *rand.Rand               // random generator
}

// Compiler check
var _ common.Storer = new(MemStore)

// New creates a new MemStore.
func New() *MemStore {
	s := new(MemStore)
	s.did = make(map[string]common.Role)
	s.act = make(map[string]common.Role)
	s.tkt = make(map[string]common.Ticket)
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

// Activate using the generated requestID
// It may be called multiple times, but will only activate a single device.
func (s *MemStore) Activate(deviceID string, requestID string) (common.Role, error) {
	role, ok := s.act[requestID]
	if !ok {
		return common.RoleNone, common.ErrorInvalidActivationRequest
	}
	// prevent reuse !
	delete(s.act, requestID)
	// Do the actual activation
	s.did[deviceID] = role
	return role, nil
}

// CreateRequestID generates a one-time request ID to activate any device for a given role.
func (s *MemStore) CreateRequestID(role common.Role) (rq string) {
	rq = strconv.FormatInt(s.rand.Int63(), 36)
	s.act[rq] = role
	return rq
}

// GetTicket retieves a ticket from data store.
func (s *MemStore) GetTicket(tid string) common.Ticket {
	return s.tkt[tid]
}

// SaveTicket in store.
func (s *MemStore) SaveTicket(t common.Ticket) error {
	s.tkt[t.TID] = t
	return nil
}

// String() for debugging.
func (s *MemStore) String() string {
	res := "\n=== MemStore ==\nDevices registred :"
	for k, v := range s.did {
		res += fmt.Sprintf("\n   Device %20.20s => %s", k, v.String())
	}
	res += "\nActivation requests pending :"
	for k, v := range s.act {
		res += fmt.Sprintf("\n   RequestID %20.20s => %s", k, v.String())
	}
	return res + "\n"
}
