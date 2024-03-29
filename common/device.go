package common

import "fmt"

// Role defines the role a device can be granted.
type Role int

// Predefined Roles for devices.
const (
	RoleNone Role = iota // Zero-value, does nothing.
	RoleEntry
	RoleExit
	RoleReview
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
	case RoleReview:
		return "Review"
	case RoleAdmin:
		return "Admin"
	case RoleSuper:
		return "SuperUser"
	default:
		return fmt.Sprintf("Unknown role N° %d", int(*r))
	}
}
