package common

// Role defines the role a device can be granted.
type Role byte

// Predfined Roles for devices.
const (
	RoleNone Role = iota // Zero-value, does nothing.
	RoleEntry
	RoleExit
	RoleInside
	RoleOutside
	RoleAdmin
	RoleSuper
)

// Device defines the status of a given device.
type Device struct {
	ID   string
	Role Role
}
