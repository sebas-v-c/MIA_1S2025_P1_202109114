// Package Structs provides data structures for use within the application.
package Structs

import "fmt"

// PointerBlock represents a block of pointers in the filesystem.
// It contains an array of 16 pointers represented as int32 values.
type PointerBlock struct {
	Pointers [16]int32 // Pointers holds an array of pointer values.
}

// ToString returns a string representation of the PointerBlock.
// It formats the Pointers array into a readable string.
func (p *PointerBlock) ToString() string {
	return fmt.Sprintf("Pointers: %v", p.Pointers)
}
