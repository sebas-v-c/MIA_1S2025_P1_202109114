// Package Structs provides basic data structures used throughout the application.
package Structs

import "fmt"

// User represents a user in the system.
// It contains the user's ID, group, name, and password.
type User struct {
	Id       int32  // Id is the unique identifier for the user.
	Group    string // Group represents the user's group or role.
	Name     string // Name is the user's full name.
	Password string // Password stores the user's password.
}

// ToString returns a string representation of the User.
// The output format is: "<Id>,U,<Group>,<Name>,<Password>"
func (u *User) ToString() string {
	return fmt.Sprintf("%d,U,%s,%s,%s", u.Id, u.Group, u.Name, u.Password)
}
