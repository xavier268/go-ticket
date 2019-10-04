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

	// Process ticketID, with a given role.
	// If found and valid, error is nil.
	// Html fragment is a human readable feed back, always available.
	Process(tktID string, role Role) (htmlFragment string, validity error)
	// DisplayTkt display public information such as ticket holder
	// Do not check validity.
	// Only error if ticket does not exists.
	String() string
}