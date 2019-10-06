package common

// Storer is the minimum required interface for data storage.
type Storer interface {
	Pinger
	Close() error

	// Device management
	GetRole(deviceID string) Role
	SetRole(deviceID string, role Role)
	UnsetRole(deviceID string)

	// CreateRequestID generates a request ID to activate for a given role.
	CreateRequestID(role Role) (requestID string)
	// Activate using the generated requestID
	// You may implemet it so that activation can only happens within certain timeframe,
	// or only once, or only for a single device, ...
	Activate(deviceID string, requestID string) (Role, error)

	GetTicket(tid string) Ticket
	SaveTicket(t Ticket) error
}
