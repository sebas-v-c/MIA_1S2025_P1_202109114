// Package Structs provides data structures for representing disk partitions
// and their mounted state within a filesystem.
package Structs

import "fmt"

// Partition represents a disk partition.
// It holds metadata including its status, type, fit, starting position, size,
// name, a correlative number, and a unique identifier.
type Partition struct {
	Status      [1]byte  // Status indicates the partition's status (e.g., active/inactive).
	Type        [1]byte  // Type specifies the partition type (primary, extended, etc.).
	Fit         [1]byte  // Fit defines the partition's fit type (best, first, or worst fit).
	Start       int32    // Start is the starting byte position of the partition on the disk.
	Size        int32    // Size indicates the total size of the partition in bytes.
	Name        [16]byte // Name holds the partition's name.
	Correlative int32    // Correlative is a sequential number used for identification.
	Id          [4]byte  // Id is a unique identifier for the partition.
}

// NewPartition creates a new Partition with the provided parameters.
// It returns a pointer to the newly created Partition.
func NewPartition(status, type_, fit [1]byte, start, size int32, name [16]byte, correlative int32) *Partition {
	return &Partition{
		Status:      status,
		Type:        type_,
		Fit:         fit,
		Start:       start,
		Size:        size,
		Name:        name,
		Correlative: correlative,
	}
}

// ToString returns a formatted string representation of the Partition.
// The output includes details such as start position, status, size, name,
// type, fit, ID, and correlative number.
func (p *Partition) ToString() string {
	return fmt.Sprintf("Start: %d, Status: %c, Size: %d, Name: %s, Type: %s, Fit: %s, ID: %s, Correlative: %d",
		p.Start, p.Status[:], p.Size, p.Name, p.Type, p.Fit, p.Id[:], p.Correlative)
}

// MountedPartition represents a partition that has been mounted on the system.
// It embeds a Partition and adds disk-specific information such as disk signature,
// disk tag, and the mount path.
type MountedPartition struct {
	Partition            // Embedded Partition struct.
	DiscSignature int32  // DiscSignature is a numeric signature for the disk.
	DiscTag       rune   // DiscTag is a character identifier for the disk.
	Path          string // Path is the filesystem path where the partition is mounted.
}

// ToString returns a formatted string representation of the MountedPartition.
// It includes both the partition details and the disk-specific metadata.
func (m *MountedPartition) ToString() string {
	return fmt.Sprintf("\tPartition:\n\t\t%s\n\t\tDisc Signature: %d, Disc Tag: %c",
		m.Partition.ToString(), m.DiscSignature, m.DiscTag)
}
