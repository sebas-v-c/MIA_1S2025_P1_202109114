// Package Structs provides data structures for representing filesystem components,
// including blocks for storing file data.
package Structs

import (
	"bytes"
	"fmt"
)

// FileBlock represents a block of data for file storage.
// It contains a fixed-size array of 64 bytes that holds file content.
type FileBlock struct {
	Content [64]byte // Content holds the raw file data.
}

// ToString returns a formatted string representation of the FileBlock's content.
// It trims the byte array at the first null byte (0) and returns the resulting string.
// This is useful for converting the fixed-size byte array into a human-readable string.
func (fb *FileBlock) ToString() string {
	// Create a copy of the Content array (this step can be useful for preserving data)
	contentArr := [64]byte{}
	copy(contentArr[:], fb.Content[:])

	// Find the index of the first null byte and slice the content up to that index.
	trimmed := fb.Content[:bytes.IndexByte(fb.Content[:], 0)]

	// Format the trimmed content as a quoted string.
	fileBlockString := fmt.Sprintf("%q", trimmed)
	// Remove the surrounding quotes and return the clean string.
	return fileBlockString[1 : len(fileBlockString)-1]
}
