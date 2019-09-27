package common

// Store is the minimum required interface for data storage.
type Store interface {
	Pinger
	Close() error
	Get(key string) (value interface{})
	Set(key string, value interface{})
}
