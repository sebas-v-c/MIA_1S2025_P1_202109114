// Package Structs provides data structures for representing various filesystem
// components, including the Extended Boot Record (EBR) used for managing logical partitions.
package Structs

import "fmt"

// EBR represents an Extended Boot Record.
// It stores metadata for a logical partition including its mount status, fit type,
// starting position, size, pointer to the next EBR, and its name.
type EBR struct {
	Mount byte     // Mount indicates the mount status of the partition ('0' for unmounted).
	Fit   byte     // Fit specifies the partition's fit type (e.g., first, best, or worst fit).
	Start int32    // Start is the starting byte position of the partition.
	Size  int32    // Size indicates the total size of the partition in bytes.
	Next  int32    // Next points to the starting byte of the next EBR (-1 if none).
	Name  [16]byte // Name holds the partition's name.
}

// NewEBR creates a new Extended Boot Record with the provided fit type, starting position,
// size, and name. The Mount field is initialized to '0' (unmounted) and Next is set to -1,
// indicating that there is no subsequent EBR.
func NewEBR(fit byte, start, size int32, name [16]byte) *EBR {
	return &EBR{
		Mount: '0',   // Set initial mount status to '0' (unmounted).
		Fit:   fit,   // Set the fit type.
		Start: start, // Set the starting position.
		Size:  size,  // Set the partition size.
		Next:  -1,    // Initialize the next pointer to -1 (no next EBR).
		Name:  name,  // Set the partition name.
	}
}

// ToString returns a formatted string representation of the EBR.
// It includes details such as mount status, start position, size, name, and the pointer to the next EBR.
func (e *EBR) ToString() string {
	return fmt.Sprintf("Mount: %-5v, Start: %d, Size: %d, Name: %s, Next: %d",
		e.Mount, e.Start, e.Size, e.Name, e.Next)
}
