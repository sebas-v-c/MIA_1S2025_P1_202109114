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
	Interfaces.CommandStruct // Embeds CommandStruct to inherit common behavior.
}

// NewMount creates a new Mount command instance.
// It initializes the command with its parameters and source location.
//
// Parameters:
//   - line: The line number where the command appears.
//   - column: The column number where the command appears.
//   - params: A map containing command parameters.
//
// Returns:
//   - A pointer to the newly created Mount instance.
func NewMount(line, column int, params map[string]string) *Mount {
	return &Mount{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.MOUNT,
			Params: params,
			Line:   line,
			Column: column,
		},
	}
}

// Exec executes the MOUNT command.
// It first validates the parameters before proceeding. If any required
// parameter is missing, an error is logged.
func (m *Mount) Exec() {
	if err := m.ValidateParams(); err != nil {
		m.AppendError(err.Error()) // Logs error if validation fails.
		return
	}

}

// ValidateParams checks if the required parameters for the MOUNT command are provided.
//
// Returns:
//   - error: Returns an error if a required parameter is missing.
func (m *Mount) ValidateParams() error {
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
