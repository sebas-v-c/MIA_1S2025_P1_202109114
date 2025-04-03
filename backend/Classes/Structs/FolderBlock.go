// Package Structs provides data structures for representing filesystem components,
// including folder blocks which store directory entries.
package Structs

// FolderBlock represents a block in the filesystem that holds directory content.
// It contains an array of Content structures, each representing an entry in a folder.
type FolderBlock struct {
	Content [4]Content // Content is an array of 4 Content items representing directory entries.
}

// NewFolderBlock creates and returns a new FolderBlock instance.
// It initializes the Content array with default values where each Content's Inode is set to -1,
// indicating that the entry is unused.
func NewFolderBlock() *FolderBlock {
	return &FolderBlock{
		Content: [4]Content{
			{Inode: -1}, // Default value for unused content entry.
			{Inode: -1}, // Default value for unused content entry.
			{Inode: -1}, // Default value for unused content entry.
			{Inode: -1}, // Default value for unused content entry.
		},
	}
}

// ToString returns a formatted string representation of the FolderBlock.
// It iterates over the Content array and concatenates each entry's string representation.
func (fb *FolderBlock) ToString() string {
	str := "\t\tFolderBlock Content:"
	for _, content := range fb.Content {
		str += "\n\t\t\t" + content.ToString() // Append each Content's string representation.
	}
	return str
}
