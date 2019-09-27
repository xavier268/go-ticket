package common

// Store is the minimum required interface for data storage.
type Store interface {
	Pinger
	Close() error
	GetRole(deviceID string) Role
	SetRole(deviceID string, role Role)
	UnsetRole(deviceID string)
}
