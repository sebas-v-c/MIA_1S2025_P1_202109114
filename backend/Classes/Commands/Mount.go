// Package Commands provides implementations for various filesystem commands,
// such as mounting disks, removing disks, users, groups, etc.
package Commands

import (
	env "backend/Classes/Env"
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
// It embeds the common CommandStruct for shared behavior and includes a map for command parameters.
type Mount struct {
	Interfaces.CommandStruct                   // Embeds common command metadata (type, line, column, etc.)
	Params                   map[string]string // Params holds the command parameters, such as "path" and "name".
}

// NewMount creates a new Mount command instance with the provided source position and parameters.
func NewMount(line, column int, params map[string]string) *Mount {
	return &Mount{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.MOUNT, // Command type for mounting a disk.
			Line:   line,        // Line number where the command is defined.
			Column: column,      // Column number where the command is defined.
		},
		Params: params, // Set the command-specific parameters.
	}
}

// Exec executes the MOUNT command.
// It validates parameters, opens the disk file, reads the MBR, and mounts a primary partition.
// It then updates the partition's status and ID, adds the partition to the global RAM partitions,
// writes the updated MBR back to disk, and logs the operation.
func (m *Mount) Exec() {
	// Validate the required parameters ("path" and "name").
	if err := m.validateParams(); err != nil {
		m.AppendError(err.Error()) // Log error if validation fails.
		return
	}

	// Open the disk file specified by the "path" parameter.
	file, err := Utils.OpenFile(m.Params["path"])
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	defer file.Close()

	// Read the Master Boot Record (MBR) from the disk.
	var discMBR Structs.MBR
	if err := Utils.ReadObject(file, &discMBR, 0); err != nil {
		m.AppendError(err.Error())
		return
	}

	// Prepare the partition name by copying the provided name into a fixed-size array.
	nameBytes := [16]byte{}
	copy(nameBytes[:], m.Params["name"])

	// Search for the primary partition matching the given name in the MBR.
	partitionIndex := -1
	for i, part := range discMBR.Partitions {
		// Check if the partition is primary ('P') and its name matches.
		if part.Type[0] == 'P' && bytes.Equal(part.Name[:], nameBytes[:]) {
			partitionIndex = i
			break
		}
	}

	// Build the console log output.
	consoleString := "=================MOUNT=================\n"
	consoleString += fmt.Sprintf("Mounting partition '%s' at '%s'...\n", m.Params["name"], m.Params["path"])
	// If the partition was not found, log an error.
	if partitionIndex == -1 {
		m.AppendError("Partition not found or partition is not a primary partition")
		return
	}

	// Retrieve the partition from the MBR.
	mbrPartition := &discMBR.Partitions[partitionIndex]
	// Check if the partition is already marked as mounted or is mounted in RAM.
	if mbrPartition.Status[0] == '1' || m.isPartitionRAMMounted(nameBytes, discMBR.Signature) {
		m.AppendError("Partition is already mounted")
		return
	}

	// Generate a new partition ID based on the MBR signature and existing mounted partitions.
	partitionId := m.generatePartitionId(discMBR.Signature)
	// Update the partition's status to mounted and set the new ID.
	mbrPartition.Status = [1]byte{'1'}
	mbrPartition.Id = partitionId

	// Add the partition to the global list of mounted partitions in RAM.
	env.AddPartition(&Structs.MountedPartition{
		Partition:     *mbrPartition,
		DiscSignature: discMBR.Signature,
		DiscTag:       rune(partitionId[3]),
		Path:          m.Params["path"],
	})

	// Write the updated MBR back to the disk.
	if err := Utils.WriteObject(file, discMBR, 0); err != nil {
		m.AppendError(err.Error())
		return
	}

	// Log details about the mounted partition.
	consoleString += fmt.Sprintf("Partition '%s' has been mounted with ID: '%s'\n", mbrPartition.Name, mbrPartition.Id)
	consoleString += "Updated MBR:\n\t" + mbrPartition.ToString() + "\n"
	consoleString += "=================END MOUNT================="
	// Output the final console log.
	m.LogConsole(consoleString)
}

// validateParams checks if the required parameters for the MOUNT command are provided.
// It ensures that both "path" and "name" are present, and appends ".mia" to the path if needed.
func (m *Mount) validateParams() error {
	// Check for the existence of the "path" parameter.
	if _, ok := m.Params["path"]; !ok {
		return errors.New("missing parameter -path")
	}
	// Check for the existence of the "name" parameter.
	if _, ok := m.Params["name"]; !ok {
		return errors.New("missing parameter -name")
	}

	// Ensure that the file path ends with the ".mia" extension.
	if !strings.EqualFold(filepath.Ext(m.Params["path"]), ".mia") {
		m.Params["path"] = m.Params["path"] + ".mia"
	}

	return nil
}

// isPartitionRAMMounted checks if a partition with the given name and MBR signature is already mounted in RAM.
// It iterates over the global list of mounted partitions and returns true if a match is found.
func (m *Mount) isPartitionRAMMounted(name [16]byte, mbrSignature int32) bool {
	for _, part := range env.GetPartitions() {
		if part.DiscSignature == mbrSignature && bytes.Equal(part.Name[:], name[:]) {
			return true
		}
	}
	return false
}

// generatePartitionId generates a new partition ID based on the provided MBR signature and currently mounted partitions.
// It uses fixed prefix "14" (derived from the last two digits of a student ID) and adjusts the remaining bytes
// based on the number of partitions already mounted and their tags.
func (m *Mount) generatePartitionId(signature int32) [4]byte {
	var id [4]byte
	// Set a fixed prefix (example: "14").
	copy(id[:], "14")

	discPartitionCount := '0'
	lastDiscTag := 'A'
	partitionDiscTag := 'A'
	// Iterate over the currently mounted partitions to determine the partition count and last used tag.
	for _, part := range env.GetPartitions() {
		if lastDiscTag < part.DiscTag {
			lastDiscTag = part.DiscTag
		}
		if part.DiscSignature == signature {
			discPartitionCount++
			partitionDiscTag = part.DiscTag
		}
	}
	// Set the third byte as the updated partition count.
	id[2] = byte(discPartitionCount + 1)
	// Set the fourth byte as the partition tag.
	id[3] = byte(partitionDiscTag)
	// If no partition has been counted and there are mounted partitions, increment the last tag.
	if discPartitionCount == '0' && len(env.GetPartitions()) != 0 {
		id[3] = byte(lastDiscTag + 1)
	}
	return id
}

// printMountedPartitions returns a string representation of all mounted partitions.
// It iterates over the global list of mounted partitions and concatenates their string representations.
func (m *Mount) printMountedPartitions() string {
	var mountedParts = "Mounted Partitions:\n"
	for _, part := range env.GetPartitions() {
		mountedParts += part.ToString() + "\n"
	}
	return mountedParts
}
