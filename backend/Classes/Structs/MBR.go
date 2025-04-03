// Package Structs provides data structures for representing disk metadata,
// including the Master Boot Record (MBR) and its related components.
package Structs

import (
	"fmt"
	"golang.org/x/exp/rand"
	"time"
)

// MBR represents the Master Boot Record of a disk.
// It stores information about the disk size, creation date, a unique signature,
// the partition fit type, and an array of partitions.
type MBR struct {
	MbrSize      int32        // MbrSize indicates the total size of the MBR in bytes.
	CreationDate [10]byte     // CreationDate holds the creation date in "YYYY-MM-DD" format.
	Signature    int32        // Signature is a randomly generated identifier for the MBR.
	Fit          [1]byte      // Fit represents the partition fit type (e.g., first, best, or worst fit).
	Partitions   [4]Partition // Partitions is an array containing up to 4 partitions.
}

// NewMBR creates a new MBR instance with the provided size and fit parameters.
// It initializes the creation date with the current date, generates a random signature,
// and sets up an empty partitions array.
func NewMBR(size int32, fit [1]byte) *MBR {
	// Format the current time as "YYYY-MM-DD"
	timeString := time.Now().Format("2006-01-02")
	// Create a byte array to store the formatted creation date
	var timeByteArr [10]byte
	// Copy the formatted time string into the byte array
	copy(timeByteArr[:], timeString)

	// Return a pointer to a new MBR instance with initialized values.
	return &MBR{
		MbrSize:      size,
		CreationDate: timeByteArr,
		Signature:    rand.Int31(),
		Fit:          fit,
		Partitions:   [4]Partition{},
	}
}

// ToString returns a formatted string representation of the MBR.
// The string includes details such as the MBR size, creation date, fit type,
// and a list of each partition's string representation.
func (m *MBR) ToString() string {
	// Build the string representation for the partitions.
	partitions := "\nPartitions:"
	for _, p := range m.Partitions {
		partitions += "\n\t" + p.ToString()
	}
	// Return the combined MBR details and partition information.
	return fmt.Sprintf("Size: %d, CreationDate: %s, Fit: %s", m.MbrSize, m.CreationDate, m.Fit) + partitions
}
