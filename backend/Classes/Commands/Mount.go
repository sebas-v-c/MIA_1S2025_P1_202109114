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

	nameBytes := [16]byte{}
	copy(nameBytes[:], m.Params["name"])

	partitionIndex := -1
	for i, part := range discMBR.Partitions {
		if part.Type[0] == 'P' && bytes.Equal(part.Name[:], nameBytes[:]) {
			partitionIndex = i
			break
		}
	}

	m.LogConsole("=================MOUNT=================")
	m.LogConsole(fmt.Sprintf("\tMounting partition '%s' at '%s'...", m.Params["name"], m.Params["path"]))
	if partitionIndex == -1 {
		m.AppendError("Partition not found or partition is not a primary partition")
		return
	}

	// Here we are checking in the MBR of the disc if the partition is already mounted
	// also checking if the partition is mounted in RAM
	mbrPartition := &discMBR.Partitions[partitionIndex]
	if mbrPartition.Status[0] == '1' || m.isPartitionRAMMounted(nameBytes, discMBR.Signature) {
		m.AppendError("Partition is already mounted")
		return
	}

	partitionId := m.generatePartitionId(discMBR.Signature)
	// change status of the partition from unmounted (0) to mounted (1) and change the value of the id
	mbrPartition.Status = [1]byte{'1'}
	mbrPartition.Id = partitionId
	// Add the partition to the loaded in ram partitions
	env.AddPartition(&Structs.MountedPartition{
		Partition:     *mbrPartition,
		DiscSignature: discMBR.Signature,
		DiscTag:       rune(partitionId[3]),
	})

	if err := Utils.WriteObject(file, &mbrPartition, 0); err != nil {
		m.AppendError(err.Error())
		return
	}

	m.LogConsole(fmt.Sprintf("\tPartition '%s' has been mounted with ID: '%s'", mbrPartition.Name, mbrPartition.Id))
	m.LogConsole("\tUpdated MBR:\n" + mbrPartition.ToString())
	m.LogConsole(m.printMountedPartitions())
	m.LogConsole("=================END MOUNT=================")
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

func (m *Mount) isPartitionRAMMounted(name [16]byte, mbrSignature int32) bool {
	for _, part := range env.GetPartitions() {
		if part.DiscSignature == mbrSignature && bytes.Equal(part.Name[:], name[:]) {
			return true
		}
	}
	return false
}

// generateDiscId generates a new disc ID based on the last two digits of my student ID -->202109114<--
func (m *Mount) generatePartitionId(signature int32) [4]byte {
	var id [4]byte
	copy(id[:], "14")

	discPartitionCount := 1
	lastDiscTag := 'A'
	partitionDiscTag := 'A'
	for _, part := range env.GetPartitions() {
		if lastDiscTag < part.DiscTag {
			lastDiscTag = part.DiscTag
		}
		if part.DiscSignature == signature {
			discPartitionCount++
			partitionDiscTag = part.DiscTag
		}
	}
	id[2] = byte(discPartitionCount + 1)
	id[3] = byte(partitionDiscTag)
	if discPartitionCount == 1 {
		id[3] = byte(lastDiscTag + 1)
	}
	return id
}

func (m *Mount) printMountedPartitions() string {
	var mountedParts = "Mounted Partitions:\n"

	for _, part := range env.GetPartitions() {
		mountedParts += part.ToString() + "\n"
	}
	return mountedParts
}
