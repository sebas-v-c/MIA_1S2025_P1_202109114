package Interfaces

import (
	env "backend/Classes/Env"
	"backend/Classes/Utils"
)

// Command defines the interface for all executable commands.
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
// It no longer contains `Params` or `Result`, making it truly minimal.
type CommandStruct struct {
	Type   Utils.Type // The type of command.
	Line   int        // The line number of the command in the source.
	Column int        // The column number of the command in the source.
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
func (c *CommandStruct) Exec() {
	panic("Exec() is not implemented")
}

// AppendError logs a runtime error related to this command.
func (c *CommandStruct) AppendError(msg string) {
	env.AddError(env.RuntimeError{
		Line:    c.GetLine(),
		Column:  c.GetColumn(),
		Command: c.GetType(),
		Msg:     msg,
	})
}

func (c *CommandStruct) LogConsole(msg string) {
	env.AddCommandLog(msg)
}
