// Package Structs provides data structures for representing various filesystem components.
package Structs

import "fmt"

// Content represents an entry within a folder block.
// It contains a fixed-size name and an inode number that points to the associated file or directory.
type Content struct {
	Name  [12]byte // Name holds the content's name as a fixed-size byte array.
	Inode int32    // Inode is the index of the inode associated with this content entry.
}

// ToString returns a formatted string representation of the Content.
// It converts the Name field from a byte array to a string and displays the inode number.
func (c *Content) ToString() string {
	return fmt.Sprintf(`Name: %s, Inode: %d`, string(c.Name[:]), c.Inode)
}
