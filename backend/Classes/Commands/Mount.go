package Commands

import (
	"backend/Classes/Interfaces"
	"backend/Classes/Utils"
	"errors"
	"path/filepath"
	"strings"
)

// Mount represents the MOUNT command, which is used to mount a disk.
type Mount struct {
	Interfaces.CommandStruct // Inherits shared behavior.
	Params                   map[string]string
}

// NewMount creates a new Mount command instance.
// It initializes the command with its parameters and source location.
func NewMount(line, column int, params map[string]string) *Mount {
	return &Mount{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.MOUNT,
			Line:   line,
			Column: column,
		},
		Params: params, // Params field now belongs to Mount specifically
	}
}

// Exec executes the MOUNT command.
func (m *Mount) Exec() {
	if err := m.validateParams(); err != nil {
		m.AppendError(err.Error()) // Logs error if validation fails.
		return
	}

	// TODO: Implement actual mount logic.
}

// validateParams checks if the required parameters for the MOUNT command are provided.
func (m *Mount) validateParams() error {
	if _, ok := m.Params["path"]; !ok {
		return errors.New("missing parameter -path")
	}
	if _, ok := m.Params["name"]; !ok {
		return errors.New("missing parameter -name")
	}

	// Ensure the file extension is ".mia"
	if !strings.EqualFold(filepath.Ext(m.Params["path"]), ".mia") {
		m.Params["path"] = m.Params["path"] + ".mia"
	}

	return nil
}
