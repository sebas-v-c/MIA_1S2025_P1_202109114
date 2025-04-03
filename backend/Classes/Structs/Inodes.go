// Package Structs provides data structures for representing filesystem metadata,
// including inodes that store information about files and directories.
package Structs

import (
	"fmt"
	"time"
)

// Inode represents a file or directory inode in the filesystem.
// It contains metadata such as the user and group IDs, file size,
// timestamps for access, creation, and modification, data block pointers,
// the type of the inode, and its permissions.
type Inode struct {
	UID    int32     // UID is the user identifier that owns the inode.
	GID    int32     // GID is the group identifier that owns the inode.
	Size   int32     // Size represents the size of the file or directory in bytes.
	ATime  [17]byte  // ATime is the last access time in "YYYY-MM-DD HH:MM" format.
	CTime  [17]byte  // CTime is the creation time in "YYYY-MM-DD HH:MM" format.
	MTime  [17]byte  // MTime is the last modification time in "YYYY-MM-DD HH:MM" format.
	IBlock [15]int32 // IBlock is an array of pointers to data blocks.
	Type   [1]byte   // Type denotes the inode type: '0' for folder and '1' for file.
	Perm   [3]byte   // Perm represents the permission settings (e.g., "644").
}

// NewInode creates a new inode with default parameters.
// The inodeType parameter must be '0' for a folder and '1' for a file.
// It initializes the timestamps to the current time and sets default permissions.
func NewInode(inodeType [1]byte) *Inode {
	var date [17]byte
	var perm [3]byte

	// Set the current time formatted as "YYYY-MM-DD HH:MM".
	copy(date[:], time.Now().Format("2006-01-02 15:04"))
	// Set default permission to "644".
	copy(perm[:], "644")

	// Initialize the data block pointers to -1 indicating unused blocks.
	var blocks [15]int32
	for i := 0; i < 15; i++ {
		blocks[i] = -1
	}

	return &Inode{
		UID:    1,         // Default UID is set to 1.
		GID:    1,         // Default GID is set to 1.
		Size:   0,         // Default size is 0.
		ATime:  date,      // Set the last access time.
		CTime:  date,      // Set the creation time.
		MTime:  date,      // Set the last modification time.
		IBlock: blocks,    // Initialize data block pointers.
		Type:   inodeType, // Set the inode type.
		Perm:   perm,      // Set the default permissions.
	}
}

// ToString returns a formatted string representation of the inode.
// It includes details such as UID, GID, size, timestamps, block pointers,
// inode type, and permissions.
func (i *Inode) ToString() string {
	var inodeType string
	// Determine the inode type: '0' for folder, '1' for file.
	if i.Type == [1]byte{0} {
		inodeType = "0"
	} else {
		inodeType = "1"
	}
	return fmt.Sprintf(`        i_UID %v
        i_GID %v
        i_Size %v
        i_Atime %s
        i_CTime %s
        i_MTime %s
        i_Block %v
        i_Type %s
        i_Perm %v`,
		i.UID,
		i.GID,
		i.Size,
		string(i.ATime[:]),
		string(i.CTime[:]),
		string(i.MTime[:]),
		i.IBlock,
		inodeType,
		string(i.Perm[:]))
}
