// Package Interfaces defines the interfaces and shared structures for executable commands.
package Interfaces

import (
	env "backend/Classes/Env"
	"backend/Classes/Utils"
)

// Command defines the interface for all executable commands.
// Each command must provide methods for retrieving its source position,
// determining its type, and executing its logic.
type Command interface {
	// GetLine returns the line number where the command is defined.
	GetLine() int

	// GetColumn returns the column number where the command is defined.
	GetColumn() int

	// GetType returns the type of the command.
	GetType() Utils.Type

	// Exec executes the command's logic.
	Exec()
}

// CommandStruct provides shared behavior for commands.
// It serves as a minimal base structure that includes the command type,
// and the line and column numbers where the command is defined.
type CommandStruct struct {
	Type   Utils.Type // Type is the type of the command.
	Line   int        // Line is the line number where the command is defined.
	Column int        // Column is the column number where the command is defined.
}

// GetLine returns the line number where the command is defined.
func (c *CommandStruct) GetLine() int {
	return c.Line
}

// GetColumn returns the column number where the command is defined.
func (c *CommandStruct) GetColumn() int {
	return c.Column
}

// GetType returns the type of the command.
func (c *CommandStruct) GetType() Utils.Type {
	return c.Type
}

// Exec provides a default execution method for commands.
// This method should be overridden by specific command implementations.
// Calling this default method results in a runtime panic.
func (c *CommandStruct) Exec() {
	panic("Exec() is not implemented")
}

// AppendError logs a runtime error associated with this command.
// It creates a RuntimeError entry with the command's source position and type.
func (c *CommandStruct) AppendError(msg string) {
	env.AddError(env.RuntimeError{
		Line:    c.GetLine(),
		Column:  c.GetColumn(),
		Command: c.GetType(),
		Msg:     msg,
	})
}

// LogConsole logs a message to the command log console.
func (c *CommandStruct) LogConsole(msg string) {
	env.AddCommandLog(msg)
}
