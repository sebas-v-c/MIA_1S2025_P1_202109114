// Package Commands provides implementations for various filesystem commands,
// such as creating disks, files, and managing the filesystem.
package Commands

import (
	"backend/Classes/Interfaces"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"errors"
	"path/filepath"
	"strconv"
	"strings"
)

// Mkdisk represents the MKDISK command, which is used to create a new disk file
// with a specified size and configuration. It embeds the common CommandStruct to
// inherit shared metadata and behavior, and holds a map for command parameters.
type Mkdisk struct {
	Interfaces.CommandStruct                   // Embedded command structure (provides Type, Line, Column, etc.)
	Params                   map[string]string // Params holds command-specific parameters such as "size", "path", "fit", and "unit".
}

// NewMkdisk creates a new Mkdisk command instance with the specified source location and parameters.
func NewMkdisk(line, column int, params map[string]string) *Mkdisk {
	return &Mkdisk{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.MKDISK, // Command type for creating a disk.
			Line:   line,         // Source code line where the command is defined.
			Column: column,       // Source code column where the command is defined.
		},
		Params: params, // Command-specific parameters.
	}
}

// Exec executes the MKDISK command.
// It validates the parameters, creates a disk file with the specified size,
// writes empty data in batches to the file, creates a new MBR structure, and writes
// the MBR to the disk. Finally, it logs the outcome.
func (m *Mkdisk) Exec() {
	// Validate command parameters.
	if err := m.validateParams(); err != nil {
		m.AppendError(err.Error())
		return
	}

	// Retrieve the file path from the parameters.
	filePath := m.Params["path"]

	// Create the disk file (if it doesn't already exist).
	err := Utils.CreateFile(filePath)
	if err != nil {
		m.AppendError(err.Error())
		return
	}

	// Open the disk file for writing.
	file, err := Utils.OpenFile(filePath)
	if err != nil {
		m.AppendError(err.Error())
		return
	}

	// Recalculate units (K or M) based on the provided parameter.
	units := m.recalculateUnits()
	// Convert the "size" parameter to an integer.
	size, _ := strconv.Atoi(m.Params["size"])
	totalSize := units * size

	// Write empty data to the file in batches (to optimize resource usage).
	batchSize := 1024 * 1024 // 1 MB batch size.
	emptyData := make([]byte, batchSize)
	bytesWritten := 0
	for bytesWritten < totalSize {
		writeSize := batchSize
		if totalSize-bytesWritten < batchSize {
			// Write the remaining bytes if less than a full batch.
			writeSize = totalSize - bytesWritten
		}
		file.Write(emptyData[:writeSize])
		bytesWritten += writeSize
	}

	// Create a new Master Boot Record (MBR) using the total disk size and the chosen fit.
	newMBR := Structs.NewMBR(int32(totalSize), m.getFit())
	// Write the MBR to the beginning of the disk file.
	if err := Utils.WriteObject(file, newMBR, 0); err != nil {
		m.AppendError(err.Error())
		return
	}
	defer file.Close()

	// Log the details of the new MBR.
	m.LogConsole("=================MKDISK=================\n" + newMBR.ToString() + "\n=================END MKDISK=================\n")
}

// validateParams checks that all required parameters for the MKDISK command are provided,
// sets default values for optional parameters, and ensures that the "size" parameter is valid.
// It also ensures that the disk file has the correct extension.
func (m *Mkdisk) validateParams() error {
	// Check for required "size" and "path" parameters.
	if _, ok := m.Params["size"]; !ok {
		return errors.New("missing parameter -size")
	} else if _, ok := m.Params["path"]; !ok {
		return errors.New("missing parameter -path")
	}

	// Set default values if not provided.
	if _, ok := m.Params["fit"]; !ok {
		m.Params["fit"] = "FF"
	}
	if _, ok := m.Params["unit"]; !ok {
		m.Params["unit"] = "M"
	}

	// Ensure that the size is a positive integer.
	if val, _ := strconv.Atoi(m.Params["size"]); val <= 0 {
		return errors.New("size must be greater than 0")
	}
	// Ensure the file path has the ".mia" extension (case-insensitive).
	if !strings.EqualFold(filepath.Ext(m.Params["path"]), ".mia") {
		m.Params["path"] = m.Params["path"] + ".mia"
	}

	return nil
}

// recalculateUnits returns the multiplier for the disk size based on the provided unit.
// If the unit is "K", it returns 1024; otherwise, it returns 1024*1024 for megabytes.
func (m *Mkdisk) recalculateUnits() int {
	m.Params["unit"] = strings.ToUpper(m.Params["unit"])
	if m.Params["unit"] == "K" {
		return 1024
	}
	return 1024 * 1024
}

// getFit returns the fit type as a 1-byte array based on the "fit" parameter.
// "FF" maps to 'F', "BF" maps to 'B', and any other value defaults to 'W'.
func (m *Mkdisk) getFit() [1]byte {
	m.Params["fit"] = strings.ToUpper(m.Params["fit"])
	if m.Params["fit"] == "FF" {
		return [1]byte{'F'}
	} else if m.Params["fit"] == "BF" {
		return [1]byte{'B'}
	}
	return [1]byte{'W'}
}
