package common

import "fmt"

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

// String convert role to string
func (r *Role) String() string {
	switch *r {
	case RoleNone:
		return "None"
	case RoleEntry:
		return "Entry"
	case RoleExit:
		return "Exit"
	case RoleInside:
		return "Inside"
	case RoleOutside:
		return "Outside"
	case RoleAdmin:
		return "Admin"
	case RoleSuper:
		return "SuperUser"
	default:
		return fmt.Sprintf("Unknown role N° %d", int(*r))
	}
}
