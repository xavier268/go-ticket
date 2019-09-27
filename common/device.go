package common

// Role defines the role a device can be granted.
type Role byte

// Predfined Roles for devices.
const (
	Default Role = iota // Zero-value, does nothing.
	Entry
	Exit
	Inside
	Outside
	Admin
)

// Device defines the status of a given device.
type Device struct {
	ID   string
	Role Role
}
