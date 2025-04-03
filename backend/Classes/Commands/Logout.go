// Package Commands provides implementations for various filesystem commands,
// including user logout functionality.
package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Utils"
)

// Logout represents the LOGOUT command, which is used to log out the currently logged-in user.
// It embeds the common CommandStruct to inherit shared metadata (such as Type, Line, and Column).
type Logout struct {
	Interfaces.CommandStruct // Embeds common command behavior and metadata.
}

// NewLogout creates a new instance of the Logout command with the specified source location.
// It initializes the command with the LOGOUT type, and sets the line and column where the command is defined.
func NewLogout(line, column int) *Logout {
	return &Logout{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.LOGOUT, // Sets the command type to LOGOUT.
			Line:   line,         // Source code line where the command is defined.
			Column: column,       // Source code column where the command is defined.
		},
	}
}

// Exec executes the LOGOUT command.
// It checks whether a user is currently logged in; if not, it logs an error.
// If a user is logged in, it sets the global current user to nil (logging out),
// then logs a success message to the console.
func (l *Logout) Exec() {
	// Check if a user is currently logged in.
	if env.CurrentUser == nil {
		l.AppendError("You are not logged in")
		return
	}

	// Build the console output.
	var consoleString string
	consoleString += "=================LOGOUT=================\n"

	// Set the global CurrentUser to nil, effectively logging out.
	env.CurrentUser = nil

	consoleString += "Logout success!\n"
	consoleString += "=================END LOGOUT=================\n"

	// Log the console output.
	l.LogConsole(consoleString)
}
