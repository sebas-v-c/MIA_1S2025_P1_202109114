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

	// GetResult retrieves the result of the command execution.
	GetResult() string

	// validateParams checks the parameters sent to the command.
	// This ensures required parameters exist before execution.
	// If no validation is required, it returns nil.
	validateParams() error
}

// CommandStruct is a base struct providing shared behavior for commands.
type CommandStruct struct {
	Result string            // The result of executing the command.
	Type   Utils.Type        // The type of command.
	Params map[string]string // Parameters passed to the command.
	Line   int               // The line number of the command in the source.
	Column int               // The column number of the command in the source.
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

// GetResult retrieves the result of the command execution.
// This method should be overridden by specific command implementations.
func (c *CommandStruct) GetResult() string {
	panic("GetResult() is not implemented")
}

// validateParams checks the parameters sent to the command.
// By default, it assumes no validation is needed and returns nil.
// Commands that require specific parameters must override this method.
func (c *CommandStruct) validateParams() error {
	return nil
}

// AppendError logs a runtime error related to this command.
// This method is used to report errors that occur during execution.
func (c *CommandStruct) AppendError(msg string) {
	env.AddError(env.RuntimeError{
		Line:    c.GetLine(),
		Column:  c.GetColumn(),
		Command: c.GetType(),
		Msg:     msg,
	})
}
