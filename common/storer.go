package common

// Storer is the minimum required interface for data storage.
type Storer interface {
	Pinger
	Close() error

	// Device management
	GetRole(deviceID string) Role
	SetRole(deviceID string, role Role)
	UnsetRole(deviceID string)

	// CreateRequestID generates a one-time request ID to activate for a given role.
	CreateRequestID(role Role) (requestID string)
	// Activate (once) using the generated requestID
	Activate(deviceID string, requestID string) (Role, error)

	GetTicket(tid string) Ticket
	SaveTicket(t Ticket) error
}
