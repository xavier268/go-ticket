package common

// Store is the minimum required interface for data storage.
type Store interface {
	Pinger
	Close() error
	GetDevice(deviceID string) Device
	SetDevice(device Device) (deviceID string)
	UnSetDevice(device Device)
}
