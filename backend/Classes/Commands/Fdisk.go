// Package Commands provides implementations for various filesystem commands,
// including the creation and management of disk partitions.
package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"encoding/binary"
	"errors"
	"path/filepath"
	"strconv"
	"strings"
)

// Fdisk represents the FDISK command, which is used to create a new partition
// on a disk file. It embeds the common CommandStruct for shared command metadata
// and behavior, and holds a map of command parameters.
type Fdisk struct {
	Interfaces.CommandStruct                   // Embedded command structure (provides Type, Line, Column, etc.)
	Params                   map[string]string // Params holds command-specific parameters like "size", "path", "name", "type", "fit", and "unit".
}

// NewFdisk creates a new Fdisk command instance with the specified source location and parameters.
func NewFdisk(line, column int, params map[string]string) *Fdisk {
	return &Fdisk{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.FDISK, // Command type constant for FDISK.
			Line:   line,        // Source code line where the command is defined.
			Column: column,      // Source code column where the command is defined.
		},
		Params: params, // Command-specific parameters.
	}
}

// Exec executes the FDISK command to create a new partition.
// It validates parameters, opens the disk file, reads the disk MBR,
// and creates a new primary, extended, or logical partition based on the provided parameters.
// The command calculates the appropriate start position and size for the partition,
// updates the MBR, and (if needed) creates an EBR for logical partitions.
func (f *Fdisk) Exec() {
	// Validate the command parameters.
	if err := f.validateParams(); err != nil {
		f.AppendError(err.Error())
		return
	}

	// Open the disk file specified by the "path" parameter.
	file, err := Utils.OpenFile(f.Params["path"])
	if err != nil {
		f.AppendError(err.Error())
		return
	}

	// Read the disk's Master Boot Record (MBR) from the file.
	var diskMBR Structs.MBR
	if err := Utils.ReadObject(file, &diskMBR, 0); err != nil {
		f.AppendError(err.Error())
		return
	}

	// Recalculate total partition size from provided "size" parameter.
	size, _ := strconv.Atoi(f.Params["size"])
	totalSize := f.recalculateUnits() * size

	// Validate that there is enough space and partition limits are not exceeded.
	err, totalPartitions := f.validatePartitions(diskMBR, totalSize)
	if err != nil {
		f.AppendError(err.Error())
		return
	}

	// Calculate the starting position for the new partition.
	// Initially, the gap is set to the binary size of the MBR.
	var gap = int32(binary.Size(diskMBR))
	// If there are already partitions, gap is set to the end of the last partition.
	if totalPartitions > 0 {
		lastPartition := diskMBR.Partitions[totalPartitions-1]
		gap = lastPartition.Start + lastPartition.Size
	}

	// Build a console output log.
	var consoleString string
	consoleString = "=================FDISK=================\n"

	// Create a primary or extended partition.
	for i := range diskMBR.Partitions {
		// Check if the partition slot is empty (Size == 0) and if the requested type is Primary or Extended.
		if diskMBR.Partitions[i].Size == 0 && (f.getPartType() == 'P' || f.getPartType() == 'E') {
			// Create the partition entry in the MBR.
			diskMBR.Partitions[i] = Structs.Partition{
				Start:       gap,                        // Start at the next available space.
				Size:        int32(totalSize),           // Partition size.
				Correlative: int32(totalPartitions + 1), // Sequential number.
				Status:      [1]byte{'0'},               // Initially unmounted.
				Fit:         f.getFit(),                 // Partition fit type.
			}
			// Set the partition name and type.
			copy(diskMBR.Partitions[i].Name[:], f.Params["name"])
			copy(diskMBR.Partitions[i].Type[:], string(f.getPartType()))
			// If the partition is Extended, create the first EBR.
			if f.getPartType() == 'E' {
				partitionEBR := Structs.EBR{
					Fit:   f.getFit()[0], // Use the fit type for the EBR.
					Start: gap,           // Place the first EBR at the start of the extended partition.
					Size:  0,             // Initial EBR size is 0.
					Next:  -1,            // Indicates no subsequent EBR.
				}
				// Set an empty name for the EBR.
				copy(partitionEBR.Name[:], "")
				// Write the EBR to disk.
				if err := Utils.WriteObject(file, partitionEBR, int64(gap)); err != nil {
					f.AppendError(err.Error())
					return
				}
				consoleString += "EBR Created:\n" + "\t" + partitionEBR.ToString() + "\n"
			}
			break
		}
	}

	// Handle logical partition creation if requested.
	if f.getPartType() == 'L' {
		// Iterate over MBR partitions to find the extended partition.
		for i := range diskMBR.Partitions {
			if diskMBR.Partitions[i].Type == [1]byte{'E'} {
				// The EBR for logical partitions is inside the extended partition.
				ebrPosition := diskMBR.Partitions[i].Start
				var ebr Structs.EBR

				// Walk the linked list of EBRs until the last one is found.
				for {
					Utils.ReadObject(file, &ebr, int64(ebrPosition))
					if ebr.Next == -1 { // Last EBR found.
						break
					}
					ebrPosition = ebr.Next
				}

				// Calculate new EBR position:
				newEBRPosition := ebr.Start + ebr.Size
				// The new logical partition will start right after the new EBR.
				logicalPartitionStart := newEBRPosition + int32(binary.Size(ebr))
				// Update the last EBR's Next field to point to the new EBR.
				ebr.Next = newEBRPosition
				// Write the updated last EBR back to disk.
				if err := Utils.WriteObject(file, ebr, int64(ebrPosition)); err != nil {
					f.AppendError(err.Error())
					return
				}

				// Create a new EBR for the new logical partition.
				newEBR := Structs.EBR{
					Fit:   f.getFit()[0],
					Start: logicalPartitionStart,
					Size:  int32(totalSize),
					Next:  -1, // Indicates this is the last EBR.
				}
				copy(newEBR.Name[:], f.Params["name"])
				// Write the new EBR at its calculated position.
				if err := Utils.WriteObject(file, newEBR, int64(newEBRPosition)); err != nil {
					f.AppendError(err.Error())
					return
				}

				consoleString += ":New EBR Created\n"
				consoleString += "\t" + newEBR.ToString() + "\n"
				// Optionally print all EBRs for the extended partition.
				consoleString += "Printing All EBR from partition: \n"
				ebrPos := diskMBR.Partitions[i].Start
				for {
					if err := Utils.ReadObject(file, &ebr, int64(ebrPos)); err != nil {
						f.AppendError(err.Error())
						return
					}
					consoleString += "\t" + ebr.ToString() + "\n"
					if ebr.Next == -1 {
						break
					}
					ebrPos = ebr.Next
				}
				break
			}
		}
	}

	// Write the updated MBR back to disk.
	if err := Utils.WriteObject(file, diskMBR, 0); err != nil {
		f.AppendError(err.Error())
		return
	}

	// Read back the updated MBR (optional, for logging).
	var updatedMBR Structs.MBR
	if err := Utils.ReadObject(file, &updatedMBR, 0); err != nil {
		f.AppendError(err.Error())
		return
	}

	// Append final log output.
	consoleString += "Updated MBR:\n" + updatedMBR.ToString() + "\n=================END FDISK================="
	env.AddCommandLog(consoleString)
	defer file.Close()
}

// validateParams checks that all obligatory parameters for the FDISK command are provided,
// sets default values for optional parameters, and verifies that the size is greater than zero.
// It also ensures that the disk file has the correct ".mia" extension.
func (f *Fdisk) validateParams() error {
	// Check required parameters.
	if _, ok := f.Params["size"]; !ok {
		return errors.New("missing parameter -size")
	} else if _, ok := f.Params["path"]; !ok {
		return errors.New("missing parameter -path")
	} else if _, ok := f.Params["name"]; !ok {
		return errors.New("missing parameter -name")
	}

	// Set default values for optional parameters if not provided.
	if _, ok := f.Params["type"]; !ok {
		f.Params["type"] = "P"
	}
	if _, ok := f.Params["fit"]; !ok {
		f.Params["fit"] = "WF"
	}
	if _, ok := f.Params["unit"]; !ok {
		f.Params["unit"] = "K"
	}

	// Ensure size is greater than 0.
	if val, _ := strconv.Atoi(f.Params["size"]); val <= 0 {
		return errors.New("size must be greater than 0")
	}

	// Ensure the disk file has the ".mia" extension.
	if !strings.EqualFold(filepath.Ext(f.Params["path"]), ".mia") {
		f.Params["path"] = f.Params["path"] + ".mia"
	}
	return nil
}

// recalculateUnits returns the multiplier for the partition size based on the specified unit.
// If the unit is "B", it returns 1; if "K", it returns 1024; otherwise (e.g., "M"), it returns 1024*1024.
func (f *Fdisk) recalculateUnits() int {
	f.Params["unit"] = strings.ToUpper(f.Params["unit"])
	if f.Params["unit"] == "B" {
		return 1
	} else if f.Params["unit"] == "K" {
		return 1024
	}
	return 1024 * 1024
}

// validatePartitions checks the current state of partitions in the MBR and validates if a new partition
// can be created with the requested size. It returns an error if limits are exceeded or disk space is insufficient,
// along with the total number of partitions currently in use.
func (f *Fdisk) validatePartitions(mbr Structs.MBR, size int) (error, int) {
	var primaryCount, extendedCount, totalPartitions int
	var usedSpace int32
	for i := 0; i < 4; i++ {
		if mbr.Partitions[i].Size != 0 {
			totalPartitions++
			usedSpace += mbr.Partitions[i].Size

			if mbr.Partitions[i].Type == [1]byte{'P'} {
				primaryCount++
			} else if mbr.Partitions[i].Type == [1]byte{'E'} {
				extendedCount++
			}
		}
	}

	if totalPartitions >= 4 {
		return errors.New("cannot create more than 4 primary or extended partitions"), totalPartitions
	}
	if f.getPartType() == 'E' && extendedCount >= 1 {
		return errors.New("only one extended partition is supported"), totalPartitions
	}
	if f.getPartType() == 'L' && extendedCount == 0 {
		return errors.New("cannot create a logical partition without an extended partition"), totalPartitions
	}
	if usedSpace+int32(size) > mbr.MbrSize {
		return errors.New("not enough disk space to create this partition"), totalPartitions
	}
	return nil, totalPartitions
}

// getPartType returns the partition type as a rune based on the "type" parameter.
// It converts the parameter to uppercase and returns its first rune.
func (f *Fdisk) getPartType() rune {
	return []rune(strings.ToUpper(f.Params["type"]))[0]
}

// getFit returns the fit type as a 1-byte array based on the "fit" parameter.
// "FF" maps to 'F', "BF" maps to 'B', and any other value defaults to 'W'.
func (f *Fdisk) getFit() [1]byte {
	f.Params["fit"] = strings.ToUpper(f.Params["fit"])
	if f.Params["fit"] == "FF" {
		return [1]byte{'F'}
	} else if f.Params["fit"] == "BF" {
		return [1]byte{'B'}
	}
	return [1]byte{'W'}
}
