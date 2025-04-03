// Package Commands provides implementations for various filesystem commands,
// such as disk removal, user management, and more.
package Commands

import (
	"backend/Classes/Interfaces"
	"backend/Classes/Utils"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// Rmdisk represents the command to remove a disk file.
// It embeds the common CommandStruct for basic command metadata and includes a map for command parameters.
type Rmdisk struct {
	Interfaces.CommandStruct                   // Embedded command structure providing metadata (type, line, column, etc.)
	Params                   map[string]string // Params holds the command parameters (e.g., file path for removal).
}

// NewRmdisk creates a new instance of the Rmdisk command with the specified source position and parameters.
func NewRmdisk(line, column int, params map[string]string) *Rmdisk {
	return &Rmdisk{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.RMDISK, // Command type for removing a disk.
			Line:   line,         // The line number where the command is defined.
			Column: column,       // The column number where the command is defined.
		},
		Params: params, // Command parameters, including the "path" of the disk file to remove.
	}
}

// Exec executes the Rmdisk command logic.
// It validates parameters, then attempts to remove the specified disk file.
// If the removal is successful, it logs a success message; otherwise, it appends an error.
func (r *Rmdisk) Exec() {
	// Validate that the required parameters are present.
	if err := r.validateParams(); err != nil {
		r.AppendError(err.Error())
		return
	}

	// Attempt to remove the disk file at the specified path.
	// os.Remove returns an error if the file could not be deleted.
	file := os.Remove(r.Params["path"])
	if file != nil {
		r.AppendError("Error deleting file: " + r.Params["path"])
		return
	}

	// Log a success message indicating that the disk file was removed.
	r.LogConsole("=================RMDISK=================\n" +
		"File succesfully removed" +
		"\n=================END RMDISK=================\n")
}

// validateParams ensures that the required parameters for the Rmdisk command are present.
// If the "path" parameter exists but does not have the ".mia" extension (case-insensitive),
// the extension is appended. It returns an error if the "path" parameter is missing.
func (r *Rmdisk) validateParams() error {
	// Check if the "path" parameter exists.
	if _, ok := r.Params["path"]; ok {
		// If the file extension is not ".mia" (case-insensitive), append ".mia" to the path.
		if !strings.EqualFold(filepath.Ext(r.Params["path"]), ".mia") {
			r.Params["path"] = r.Params["path"] + ".mia"
		}
		return nil
	}
	// Return an error if the obligatory "path" parameter does not exist.
	return errors.New("obligatory parameter not exist")
}
