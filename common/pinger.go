package common

// Pinger interface to verify health status.
type Pinger interface {
	Ping() error
}
