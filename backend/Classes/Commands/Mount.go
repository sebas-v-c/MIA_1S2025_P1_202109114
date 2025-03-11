package Commands

import (
	"backend/Classes/Interfaces"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"bytes"
	"errors"
	"fmt"
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

	file, err := Utils.OpenFile(m.Params["path"])
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	defer file.Close()

	// read the MBR of the disc file
	var discMBR Structs.MBR
	if err := Utils.ReadObject(file, &discMBR, 0); err != nil {
		m.AppendError(err.Error())
		return
	}

	m.LogConsole("=================MOUNT=================")
	m.LogConsole(fmt.Sprintf("\tSearching for partition '%s'", m.Params["name"]))

	nameBytes := [16]byte{}
	copy(nameBytes[:], m.Params["name"])

	partitionIndex := -1
	for i, part := range discMBR.Partitions {
		if part.Type[0] == 'P' && bytes.Equal(part.Name[:], nameBytes[:]) {
			partitionIndex = i
			break
		}
	}

	if partitionIndex == -1 {
		m.AppendError("Partition not found or partition is not a primary partition")
		return
	}

	// Here we are checking in the MBR of the disc if the partition is already mounted
	partition := &discMBR.Partitions[partitionIndex]
	if partition.Status[0] == '1' {
		m.AppendError("Partition is already mounted")
	}

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
