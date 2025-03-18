package Structs

import (
	"backend/Classes/Utils"
	"encoding/binary"
	"errors"
	"os"
	"strings"
)

// DirTree represents a directory tree structure containing SuperBlock and the file reference
type DirTree struct {
	SuperBlock *SuperBlock // Pointer to the superblock of the file system
	File       *os.File    // Reference to the file representing the disk
}

// NewDirTree creates a new instance of DirTree with the provided superblock and file
func NewDirTree(superBlock SuperBlock, file *os.File) *DirTree {
	return &DirTree{SuperBlock: &superBlock, File: file}
}

// GetInodeByPath retrieves an inode based on the provided path
// It splits the path into components and searches through the directory tree to find the correct inode.
func (dt *DirTree) GetInodeByPath(path string) (int32, *Inode, error) {
	// Split the path into components
	splitedPath := strings.Split(path, "/")
	if splitedPath[0] == "" {
		splitedPath = splitedPath[1:] // Remove leading empty component from path
	}

	// Read the first inode (root directory)
	var firstInode Inode
	if err := Utils.ReadObject(dt.File, &firstInode, int64(dt.SuperBlock.InodeStart)); err != nil {
		return -1, nil, err
	}

	// Call the recursive function to search for the inode based on the path
	return dt.searchInodeByPath(splitedPath, &firstInode)
}

// searchInodeByPath searches for an inode recursively by following the path components
func (dt *DirTree) searchInodeByPath(path []string, inode *Inode) (int32, *Inode, error) {
	// Iterate through the blocks of the current inode to find the matching path component
	for i, block := range inode.IBlock {
		if block == -1 {
			continue // Skip empty blocks
		}

		// If index is >= 12, it means we are dealing with an indirect access (pointer block)
		if i >= 12 {
			// Load the pointer block and pass it to loadDirectoryPointerBlock
			var pointerBlock PointerBlock
			if err := Utils.ReadObject(dt.File, &pointerBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(PointerBlock{})))); err != nil {
				return -1, nil, err
			}
			// Load the next inode from the pointer block
			index, newInode, err := dt.loadDirectoryPointerBlock(&pointerBlock, path, i-12)
			if err == nil {
				return index, newInode, nil
			}
			continue
		}

		// Iterate over the folder block to find the folder or file
		if index, newInode, err := dt.iterateFolderBlock(block, path); err == nil {
			return index, newInode, nil
		}
	}

	// Return error if inode is not found
	return -1, nil, errors.New("invalid Inode loaded")
}

// loadDirectoryPointerBlock handles loading of pointer blocks and iterates recursively for indirect blocks
func (dt *DirTree) loadDirectoryPointerBlock(pointerBlock *PointerBlock, path []string, level int) (int32, *Inode, error) {
	// Iterate over the pointer block's pointers
	for _, block := range pointerBlock.Pointers {
		if block == -1 {
			continue
		}

		// Recursively load the next pointer block if the level is greater than 0
		if level > 0 {
			var ptBlock PointerBlock
			if err := Utils.ReadObject(dt.File, &ptBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(PointerBlock{})))); err != nil {
				return -1, nil, err
			}
			// Recursive call to load the next pointer block
			index, inode, err := dt.loadDirectoryPointerBlock(&ptBlock, path, level-1)
			if err == nil {
				return index, inode, nil
			}
			continue
		}

		// Once at the correct level, load the folder block and search for the inode
		if index, inode, err := dt.iterateFolderBlock(block, path); err == nil {
			return index, inode, nil
		}
	}
	// Return error if the file is not found
	return -1, nil, errors.New("file not found")
}

// iterateFolderBlock loads a folder block, searches through its contents, and finds the inode
func (dt *DirTree) iterateFolderBlock(block int32, path []string) (int32, *Inode, error) {
	// Load the folder block from the file
	var folderBlock FolderBlock
	if err := Utils.ReadObject(dt.File, &folderBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(FolderBlock{})))); err != nil {
		return -1, nil, err
	}
	// Search through the folder content for the matching name
	for _, content := range folderBlock.Content {
		// Compare the folder name with the path component
		if strings.TrimRight(string(content.Name[:]), "\x00") != path[0] {
			continue
		}

		// Load the inode for the matched folder/file
		var nextInode Inode
		if err := Utils.ReadObject(dt.File, &nextInode, int64(dt.SuperBlock.InodeStart+content.Inode*int32(binary.Size(Inode{})))); err != nil {
			return -1, nil, err
		}

		// If the path is fully traversed, return the inode
		if len(path) == 1 {
			return content.Inode, &nextInode, nil
		} else {
			// Recursively search for the next path component
			return dt.searchInodeByPath(path[1:], &nextInode)
		}
	}
	// Return error if inode is not found
	return -1, nil, errors.New("inode not found")
}

// GetFileContentByInode retrieves the content of a file by its inode
func (dt *DirTree) GetFileContentByInode(inode *Inode) (string, error) {
	// Check if the inode is of type file (type 1)
	if inode.Type != [1]byte{1} {
		return "", errors.New("inode is not of type file")
	}

	var resultString strings.Builder // Use a StringBuilder for efficient string concatenation

	// Iterate over the iblocks to retrieve the file content
	for i, block := range inode.IBlock {
		if block == -1 {
			continue // Skip empty blocks
		}

		// If the index is greater than or equal to 12, it's an indirect block
		if i >= 12 {
			var pointerBlock PointerBlock
			if err := Utils.ReadObject(dt.File, &pointerBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(PointerBlock{})))); err != nil {
				return resultString.String(), err
			}
			result, err := dt.loadFilePointerBlock(&pointerBlock, resultString.String(), i-12)
			if err != nil {
				return result, err
			}
			resultString.WriteString(result)
			continue
		}

		// Load the file block and append content
		var fileBlock FileBlock
		if err := Utils.ReadObject(dt.File, &fileBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(FileBlock{})))); err != nil {
			return "", err
		}
		resultString.WriteString(strings.TrimRight(string(fileBlock.Content[:]), "\x00"))
	}

	return resultString.String(), nil
}

// loadFilePointerBlock recursively handles the indirect blocks for large files
func (dt *DirTree) loadFilePointerBlock(pointerBlock *PointerBlock, resultString string, level int) (string, error) {
	for _, block := range pointerBlock.Pointers {
		if block == -1 {
			continue
		}

		// If level > 0, recursively load the pointer block
		if level > 0 {
			var ptBlock PointerBlock
			if err := Utils.ReadObject(dt.File, &ptBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(PointerBlock{})))); err != nil {
				return resultString, err
			}
			result, err := dt.loadFilePointerBlock(&ptBlock, resultString, level-1)
			if err != nil {
				return result, err
			}
			resultString += result
			continue
		}

		// Load file block content
		var fileBlock FileBlock
		if err := Utils.ReadObject(dt.File, &fileBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(FileBlock{})))); err != nil {
			return resultString, err
		}
		resultString += strings.TrimRight(string(fileBlock.Content[:]), "\x00")
	}
	return resultString, nil
}
