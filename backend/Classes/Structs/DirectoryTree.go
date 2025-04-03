// Package Structs provides data structures and methods for managing a filesystem's
// directory tree structure, including inodes, file blocks, folder blocks, and related metadata.
package Structs

import (
	"backend/Classes/Utils"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// DirTree represents a directory tree structure for the filesystem.
// It contains a pointer to the superblock, a reference to the disk file, and bitmaps for blocks and inodes.
type DirTree struct {
	SuperBlock  *SuperBlock // Pointer to the superblock of the filesystem.
	File        *os.File    // Reference to the file representing the disk.
	BlockBitMap []byte      // BlockBitMap indicates the allocation status of each block.
	InodeBitMap []byte      // InodeBitMap indicates the allocation status of each inode.
}

// NewDirTree creates a new DirTree instance using the provided superblock and file.
// It reads the block and inode bitmaps from the disk file based on the superblock information.
func NewDirTree(superBlock SuperBlock, file *os.File) *DirTree {
	// Allocate slices for the block and inode bitmaps.
	var blockBitMap = make([]byte, superBlock.BlocksCount)
	var inodeBitMap = make([]byte, superBlock.InodesCount)

	// Read the block bitmap from the disk.
	if err := Utils.ReadObject(file, &blockBitMap, int64(superBlock.BmBlockStart)); err != nil {
		fmt.Println(err)
		return nil
	}
	// Read the inode bitmap from the disk.
	if err := Utils.ReadObject(file, &inodeBitMap, int64(superBlock.BmInodeStart)); err != nil {
		fmt.Println(err)
		return nil
	}

	// Return a new DirTree instance with the initialized values.
	return &DirTree{SuperBlock: &superBlock, File: file, BlockBitMap: blockBitMap, InodeBitMap: inodeBitMap}
}

// GetInodeByPath retrieves an inode by traversing the directory tree using the provided path.
// It returns the inode index, a pointer to the Inode, and an error if the inode cannot be found.
func (dt *DirTree) GetInodeByPath(path string) (int32, *Inode, error) {
	// Split the path into components.
	splitedPath := strings.Split(path, "/")
	if splitedPath[0] == "" {
		splitedPath = splitedPath[1:] // Remove leading empty component.
	}
	// TODO: Check if this handling is correct.
	if splitedPath[0] == "" || splitedPath[0] == "/" {
		splitedPath[0] = "."
	}

	// Read the root inode (first inode) from the disk.
	var firstInode Inode
	if err := Utils.ReadObject(dt.File, &firstInode, int64(dt.SuperBlock.InodeStart)); err != nil {
		return -1, nil, err
	}

	// Recursively search for the inode matching the given path.
	return dt.searchInodeByPath(splitedPath, &firstInode)
}

// searchInodeByPath recursively searches for an inode by following the path components.
// It returns the inode index, a pointer to the Inode, and an error if the inode cannot be found.
func (dt *DirTree) searchInodeByPath(path []string, inode *Inode) (int32, *Inode, error) {
	// Iterate through the inode's data block pointers.
	for i, block := range inode.IBlock {
		if block == -1 {
			continue // Skip unused blocks.
		}

		// For indices >= 12, handle indirect addressing using a pointer block.
		if i >= 12 {
			var pointerBlock PointerBlock
			// Read the pointer block from the disk.
			if err := Utils.ReadObject(dt.File, &pointerBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(PointerBlock{})))); err != nil {
				return -1, nil, err
			}
			// Attempt to retrieve the inode from the pointer block.
			index, newInode, err := dt.loadDirectoryPointerBlockContent(&pointerBlock, path, i-12)
			if err == nil {
				return index, newInode, nil
			}
			continue
		}

		// For direct blocks, iterate through the folder block to find the matching entry.
		if index, newInode, err := dt.iterateFolderBlock(block, path); err == nil {
			return index, newInode, nil
		}
	}
	// If no matching inode is found, return an error.
	return -1, nil, errors.New("file inode not found")
}

// loadDirectoryPointerBlockContent processes pointer blocks for indirect addressing.
// It recursively traverses pointer blocks until reaching the correct level to load a folder block.
// Returns the inode index, pointer to the Inode, and error if not found.
func (dt *DirTree) loadDirectoryPointerBlockContent(pointerBlock *PointerBlock, path []string, level int) (int32, *Inode, error) {
	// Iterate over each pointer in the pointer block.
	for _, block := range pointerBlock.Pointers {
		if block == -1 {
			continue // Skip unused pointers.
		}

		// If further levels remain, recursively load the next pointer block.
		if level > 0 {
			var ptBlock PointerBlock
			if err := Utils.ReadObject(dt.File, &ptBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(PointerBlock{})))); err != nil {
				return -1, nil, err
			}
			index, inode, err := dt.loadDirectoryPointerBlockContent(&ptBlock, path, level-1)
			if err == nil {
				return index, inode, nil
			}
			continue
		}

		// At the correct level, load the folder block and search for the inode.
		if index, inode, err := dt.iterateFolderBlock(block, path); err == nil {
			return index, inode, nil
		}
	}
	// Return an error if the file is not found.
	return -1, nil, errors.New("file not found")
}

// iterateFolderBlock loads a folder block from the disk and searches its contents for the matching path component.
// Returns the inode index, pointer to the Inode, and an error if the matching entry is not found.
func (dt *DirTree) iterateFolderBlock(block int32, path []string) (int32, *Inode, error) {
	// Load the folder block from the disk.
	var folderBlock FolderBlock
	if err := Utils.ReadObject(dt.File, &folderBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(FolderBlock{})))); err != nil {
		return -1, nil, err
	}

	// Iterate over the entries in the folder block.
	for _, content := range folderBlock.Content {
		// Compare the entry name (trimmed of null bytes) with the current path component.
		if strings.TrimRight(string(content.Name[:]), "\x00") != path[0] || content.Inode == -1 {
			continue
		}

		// Load the inode corresponding to the matched entry.
		var nextInode Inode
		if err := Utils.ReadObject(dt.File, &nextInode, int64(dt.SuperBlock.InodeStart+content.Inode*int32(binary.Size(Inode{})))); err != nil {
			return -1, nil, err
		}

		// If this is the last component in the path, return the inode.
		if len(path) == 1 {
			return content.Inode, &nextInode, nil
		} else {
			// Otherwise, recursively search using the remaining path components.
			return dt.searchInodeByPath(path[1:], &nextInode)
		}
	}
	// Return an error if no matching inode is found.
	return -1, nil, errors.New("inode not found")
}

// GetFileContentByInode retrieves the content of a file represented by the given inode.
// It verifies the inode type, then reads and concatenates file blocks (including indirect blocks)
// into a single string. The inode's access time is updated upon retrieval.
func (dt *DirTree) GetFileContentByInode(inode *Inode) (string, error) {
	// Check if the inode represents a file (type 1).
	if inode.Type != [1]byte{1} {
		return "", errors.New("inode is not of type file")
	}

	var resultString strings.Builder // Efficiently build the file content string.

	// Iterate over the inode's block pointers.
	for i, block := range inode.IBlock {
		if block == -1 {
			continue // Skip unused blocks.
		}

		// Handle indirect blocks if index is >= 12.
		if i >= 12 {
			var pointerBlock PointerBlock
			if err := Utils.ReadObject(dt.File, &pointerBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(PointerBlock{})))); err != nil {
				return resultString.String(), err
			}
			result, err := dt.loadFilePointerBlockContent(&pointerBlock, resultString.String(), i-12)
			if err != nil {
				return result, err
			}
			resultString.WriteString(result)
			continue
		}

		// Load the direct file block and append its content.
		var fileBlock FileBlock
		if err := Utils.ReadObject(dt.File, &fileBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(FileBlock{})))); err != nil {
			return "", err
		}
		resultString.WriteString(strings.TrimRight(string(fileBlock.Content[:]), "\x00"))
	}

	// Update the inode's access time.
	copy(inode.ATime[:], time.Now().Format("2006-01-02 15:04"))
	return resultString.String(), nil
}

// loadFilePointerBlockContent recursively processes indirect file blocks.
// It traverses pointer blocks until reaching file blocks, concatenating their content.
// Returns the updated result string and an error if encountered.
func (dt *DirTree) loadFilePointerBlockContent(pointerBlock *PointerBlock, resultString string, level int) (string, error) {
	// Iterate through the pointers in the pointer block.
	for _, block := range pointerBlock.Pointers {
		if block == -1 {
			continue // Skip unused pointers.
		}

		// If there are additional levels, recursively process the next pointer block.
		if level > 0 {
			var ptBlock PointerBlock
			if err := Utils.ReadObject(dt.File, &ptBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(PointerBlock{})))); err != nil {
				return resultString, err
			}
			result, err := dt.loadFilePointerBlockContent(&ptBlock, resultString, level-1)
			if err != nil {
				return result, err
			}
			resultString += result
			continue
		}

		// Load the file block content at the correct level.
		var fileBlock FileBlock
		if err := Utils.ReadObject(dt.File, &fileBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(FileBlock{})))); err != nil {
			return resultString, err
		}
		resultString += strings.TrimRight(string(fileBlock.Content[:]), "\x00")
	}
	return resultString, nil
}

// AppendToFileInode appends the given string content to a file represented by the inode.
// It determines the last file block, appends the new content, splits the result into chunks,
// and writes updated file blocks to disk. Note: Indirect addressing is not fully supported.
func (dt *DirTree) AppendToFileInode(s string, inode *Inode) error {
	// WARNING: Modifying inode data here is not ideal; ideally, this method should only return updated data.
	var lastBlock int32 = -1
	var lastIndex = 0

	// Find the last used block in the inode's block pointers.
	for i, block := range inode.IBlock {
		if block == -1 {
			break
		}
		// Indirect addresses are not implemented.
		if i >= 12 {
			fmt.Println("Cannot load indirect addresses, functionality not implemented")
			return errors.New("implement indirect address")
		}
		lastBlock = block
		lastIndex = i
	}

	// Load the last file block if it exists.
	var lastFileBlock FileBlock
	if lastBlock != -1 {
		if err := Utils.ReadObject(dt.File, &lastFileBlock, int64(dt.SuperBlock.BlockStart+lastBlock*int32(binary.Size(FileBlock{})))); err != nil {
			return err
		}
	}

	// Append new content to the current content of the last file block.
	content := strings.TrimRight(string(lastFileBlock.Content[:]), "\x00")
	content += s

	// Split the updated content into 64-byte chunks.
	splittedBytesContent := dt.SplitRawBytes(content, binary.Size(FileBlock{}))
	for _, block := range splittedBytesContent {
		var newFileBlock FileBlock
		copy(newFileBlock.Content[:], block)

		// TODO: Check logic if lastIndex is in the last 3 IBlock positions (indirect addressing).
		if lastIndex >= 12 {
			fmt.Println("Cannot handle indirect addresses, functionality not implemented")
			return errors.New("implement indirect addresses creation")
		}

		// If no block exists, allocate a new file block.
		if lastBlock == -1 {
			var err error
			lastBlock, err = dt.GetAvailableBlockAddress()
			if err != nil {
				return err
			}
		}

		// Write the new file block to disk.
		if err := Utils.WriteObject(dt.File, newFileBlock, int64(dt.SuperBlock.BlockStart+lastBlock*int32(binary.Size(FileBlock{})))); err != nil {
			return err
		}

		// Update the superblock bitmap for the used block.
		if err := Utils.WriteObject(dt.File, byte(1), int64(dt.SuperBlock.BmBlockStart+lastBlock)); err != nil {
			return err
		}
		dt.BlockBitMap[lastBlock] = 1
		dt.SuperBlock.FreeBlocksCount = dt.freeBlocks()
		dt.SuperBlock.FirstInode, _ = dt.GetAvailableInodeAdress()

		// Update the inode to point to the current file block.
		inode.IBlock[lastIndex] = lastBlock

		// Prepare for the next block.
		lastBlock = -1
		lastIndex += 1
	}

	// Update inode metadata.
	inode.Size = int32(len(content))
	copy(inode.MTime[:], time.Now().Format("2006-01-02 15:04"))
	return nil
}

// GetAvailableBlockAddress searches the block bitmap for a free block and returns its address.
func (dt *DirTree) GetAvailableBlockAddress() (int32, error) {
	for i, block := range dt.BlockBitMap {
		if block == 0 {
			return int32(i), nil
		}
	}
	return -1, errors.New("no more free blocks")
}

// GetAvailableInodeAdress searches the inode bitmap for a free inode and returns its address.
func (dt *DirTree) GetAvailableInodeAdress() (int32, error) {
	for i, block := range dt.InodeBitMap {
		if block == 0 {
			return int32(i), nil
		}
	}
	return -1, errors.New("no more free inodes")
}

// SplitRawBytes splits a string into chunks of the given size and returns a slice of string chunks.
func (dt *DirTree) SplitRawBytes(s string, chunkSize int) []string {
	b := []byte(s)
	var chunks []string
	for i := 0; i < len(b); i += chunkSize {
		end := i + chunkSize
		if end > len(b) {
			end = len(b)
		}
		chunks = append(chunks, string(b[i:end]))
	}
	return chunks
}

// FreeInodeBlockContent frees all data blocks associated with an inode by updating the block bitmap
// and resetting the inode's block pointers. It writes the updated bitmap to disk.
func (dt *DirTree) FreeInodeBlockContent(inode *Inode) error {
	// Iterate over each block pointer in the inode.
	for i, block := range inode.IBlock {
		// Skip indirect addressing.
		if i >= 12 {
			continue
		}
		// If the block is in use, free it.
		if block != -1 {
			dt.BlockBitMap[block] = 0
			inode.IBlock[i] = -1
		}
	}
	// Write the updated block bitmap back to disk.
	if err := Utils.WriteObject(dt.File, dt.BlockBitMap, int64(dt.SuperBlock.BmBlockStart)); err != nil {
		return err
	}
	return nil
}

// CreateNewInode creates a new inode entry within a directory represented by the parent inode.
// It adds the new inode entry to a folder block, updates bitmaps, writes the new inode to disk,
// and updates the parent inode accordingly.
// Returns the new inode index, a pointer to the new Inode, and an error if encountered.
func (dt *DirTree) CreateNewInode(index int32, inode *Inode, name string, UID, GID int32, inodeType [1]byte) (int32, *Inode, error) {
	// WARNING: Modifying the parent inode here is not ideal; this implementation should ideally only return new data.
	var lastBlock int32 = -1
	var lastIndex = 0

	// Find the last used block in the parent inode.
	for i, block := range inode.IBlock {
		if block == -1 {
			break
		}
		// Indirect addresses are not implemented.
		if i >= 12 {
			fmt.Println("Cannot load indirect addresses, functionality not implemented")
			return -1, nil, errors.New("implement indirect addresses inode creation")
		}
		lastBlock = block
		lastIndex = i
	}

	var lastFolderBlock FolderBlock
	var emptyContentIndex = -1
	// Load the current folder block from the parent inode.
	if err := Utils.ReadObject(dt.File, &lastFolderBlock, int64(dt.SuperBlock.BlockStart+lastBlock*int32(binary.Size(FolderBlock{})))); err != nil {
		return -1, nil, err
	}
	// Search for an empty slot in the folder block.
	for i, content := range lastFolderBlock.Content {
		if content.Inode == -1 {
			emptyContentIndex = i
			break
		}
	}

	// If no empty slot exists, create a new folder block.
	if emptyContentIndex == -1 {
		lastFolderBlock = *NewFolderBlock()
		emptyContentIndex = 0
		lastIndex += 1
		var err error
		lastBlock, err = dt.GetAvailableBlockAddress()
		if err != nil {
			return -1, nil, err
		}
		// TODO: Implement indirect addresses creation for inode if needed.
		if lastIndex >= 12 {
			fmt.Println("Cannot handle indirect addresses, functionality not implemented")
			return -1, nil, errors.New("implement indirect addresses creation inode")
		}
		// Update the parent inode with the new folder block address.
		inode.IBlock[lastIndex] = lastBlock
	}
	// Update the parent inode's modification time.
	copy(inode.MTime[:], time.Now().Format("2006-01-02 15:04"))

	// Create and initialize the new inode.
	newInode := NewInode(inodeType)
	newInode.UID = UID
	newInode.GID = GID
	newInodeIndex, err := dt.GetAvailableInodeAdress()
	if err != nil {
		return -1, nil, err
	}

	// Update the folder block entry with the new inode's information.
	copy(lastFolderBlock.Content[emptyContentIndex].Name[:], name)
	lastFolderBlock.Content[emptyContentIndex].Inode = newInodeIndex

	// Write the updated folder block to disk.
	if err := Utils.WriteObject(dt.File, &lastFolderBlock, int64(dt.SuperBlock.BlockStart+lastBlock*int32(binary.Size(FolderBlock{})))); err != nil {
		return -1, nil, err
	}
	// Update the block bitmap for the new folder block.
	if err := Utils.WriteObject(dt.File, byte(1), int64(dt.SuperBlock.BmBlockStart+lastBlock)); err != nil {
		return -1, nil, err
	}
	dt.BlockBitMap[lastBlock] = 1
	dt.SuperBlock.FreeBlocksCount = dt.freeBlocks()
	dt.SuperBlock.FirstBlock, _ = dt.GetAvailableBlockAddress()

	// Write the new inode to disk.
	if err := Utils.WriteObject(dt.File, newInode, int64(dt.SuperBlock.InodeStart+newInodeIndex*int32(binary.Size(Inode{})))); err != nil {
		return -1, nil, err
	}
	if err := Utils.WriteObject(dt.File, byte(1), int64(dt.SuperBlock.BmInodeStart+newInodeIndex)); err != nil {
		return -1, nil, err
	}
	dt.InodeBitMap[newInodeIndex] = 1
	dt.SuperBlock.FreeInodesCount = dt.freeInodes()
	dt.SuperBlock.FirstInode, _ = dt.GetAvailableInodeAdress()

	// Write the updated parent inode to disk.
	if err := Utils.WriteObject(dt.File, inode, int64(dt.SuperBlock.InodeStart+index*int32(binary.Size(Inode{})))); err != nil {
		return -1, nil, err
	}
	return newInodeIndex, newInode, nil
}

// CreateNewDir creates a new directory under a parent directory represented by the given inode.
// It creates a new inode for the directory, initializes its folder block with default entries ("." and ".."),
// writes the new folder block and inode to disk, and updates the parent's inode.
// Returns the new directory inode index, a pointer to the new Inode, and an error if encountered.
func (dt *DirTree) CreateNewDir(index int32, inode *Inode, dirName string, uid int32, gid int32) (int32, *Inode, error) {
	newInodeIndex, newInode, err := dt.CreateNewInode(index, inode, dirName, uid, gid, [1]byte{0})
	if err != nil {
		return -1, nil, err
	}

	// Create a new folder block for the new directory.
	newFolderBlock := NewFolderBlock()
	newFolderBlockIndex, err := dt.GetAvailableBlockAddress()
	if err != nil {
		return -1, nil, err
	}

	// Set default folder entries: current directory "." and parent directory "..".
	newFolderBlock.Content[0].Inode = newInodeIndex
	copy(newFolderBlock.Content[0].Name[:], ".")
	newFolderBlock.Content[1].Inode = index
	copy(newFolderBlock.Content[1].Name[:], "..")

	// Assign the new folder block address to the new inode.
	newInode.IBlock[0] = newFolderBlockIndex

	// Write the new folder block to disk.
	if err := Utils.WriteObject(dt.File, newFolderBlock, int64(dt.SuperBlock.BlockStart+newFolderBlockIndex*int32(binary.Size(FolderBlock{})))); err != nil {
		return -1, nil, err
	}
	if err := Utils.WriteObject(dt.File, byte(1), int64(dt.SuperBlock.BmBlockStart+newFolderBlockIndex)); err != nil {
		return -1, nil, err
	}
	dt.BlockBitMap[newFolderBlockIndex] = 1
	dt.SuperBlock.FreeBlocksCount = dt.freeBlocks()
	dt.SuperBlock.FirstBlock, _ = dt.GetAvailableBlockAddress()

	// Write the updated new inode to disk.
	if err := Utils.WriteObject(dt.File, newInode, int64(dt.SuperBlock.InodeStart+newInodeIndex*int32(binary.Size(Inode{})))); err != nil {
		return -1, nil, err
	}
	return newInodeIndex, newInode, err
}

// CreateNewFile creates a new file under a directory represented by the given inode.
// It creates a new inode for the file, appends the initial content to the file's inode,
// writes the new inode to disk, and returns the new file inode index, pointer, and error if any.
func (dt *DirTree) CreateNewFile(index int32, inode *Inode, fileName string, content string, uid int32, gid int32) (int32, *Inode, error) {
	newInodeIndex, newInode, err := dt.CreateNewInode(index, inode, fileName, uid, gid, [1]byte{1})
	if err != nil {
		return -1, nil, err
	}

	// Append the initial content to the new file inode.
	if err := dt.AppendToFileInode(content, newInode); err != nil {
		return -1, nil, err
	}

	// Write the updated new inode to disk.
	if err := Utils.WriteObject(dt.File, newInode, int64(dt.SuperBlock.InodeStart+newInodeIndex*int32(binary.Size(Inode{})))); err != nil {
		return -1, nil, err
	}
	return newInodeIndex, newInode, err
}

// GetInodeByIndex retrieves an inode from the disk given its index.
// It reads the inode from the disk using the superblock's inode start position.
func (dt *DirTree) GetInodeByIndex(index int32) (*Inode, error) {
	var inode Inode
	if err := Utils.ReadObject(dt.File, &inode, int64(dt.SuperBlock.InodeStart+index*int32(binary.Size(Inode{})))); err != nil {
		return nil, err
	}
	return &inode, nil
}

// freeBlocks calculates and returns the number of free blocks based on the block bitmap.
func (dt *DirTree) freeBlocks() int32 {
	var counter int32 = 0
	for _, bit := range dt.BlockBitMap {
		if bit == 0 {
			counter++
		}
	}
	return counter
}

// freeInodes calculates and returns the number of free inodes based on the inode bitmap.
func (dt *DirTree) freeInodes() int32 {
	var counter int32 = 0
	for _, bit := range dt.InodeBitMap {
		if bit == 0 {
			counter++
		}
	}
	return counter
}
