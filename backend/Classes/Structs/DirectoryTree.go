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

// DirTree represents a directory tree structure containing SuperBlock and the file reference
type DirTree struct {
	SuperBlock  *SuperBlock // Pointer to the superblock of the file system
	File        *os.File    // Reference to the file representing the disk
	BlockBitMap []byte
	InodeBitMap []byte
}

// NewDirTree creates a new instance of DirTree with the provided superblock and file
func NewDirTree(superBlock SuperBlock, file *os.File) *DirTree {
	var blockBitMap = make([]byte, superBlock.BlocksCount)
	var inodeBitMap = make([]byte, superBlock.InodesCount)
	if err := Utils.ReadObject(file, &blockBitMap, int64(superBlock.BmBlockStart)); err != nil {
		fmt.Println(err)
		return nil
	}
	if err := Utils.ReadObject(file, &inodeBitMap, int64(superBlock.BmInodeStart)); err != nil {
		fmt.Println(err)
		return nil
	}

	return &DirTree{SuperBlock: &superBlock, File: file, BlockBitMap: blockBitMap, InodeBitMap: inodeBitMap}
}

// GetInodeByPath retrieves an inode based on the provided path
// It splits the path into components and searches through the directory tree to find the correct inode.
func (dt *DirTree) GetInodeByPath(path string) (int32, *Inode, error) {
	// Split the path into components
	splitedPath := strings.Split(path, "/")
	if splitedPath[0] == "" {
		splitedPath = splitedPath[1:] // Remove leading empty component from path
	}
	// TODO check if this is correct
	if splitedPath[0] == "" || splitedPath[0] == "/" {
		splitedPath[0] = "."
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
			// Load the pointer block and pass it to loadDirectoryPointerBlockContent
			var pointerBlock PointerBlock
			if err := Utils.ReadObject(dt.File, &pointerBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(PointerBlock{})))); err != nil {
				return -1, nil, err
			}
			// Load the next inode from the pointer block
			index, newInode, err := dt.loadDirectoryPointerBlockContent(&pointerBlock, path, i-12)
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
	return -1, nil, errors.New("file inode not found")
}

// loadDirectoryPointerBlockContent handles loading of pointer blocks and iterates recursively for indirect blocks
func (dt *DirTree) loadDirectoryPointerBlockContent(pointerBlock *PointerBlock, path []string, level int) (int32, *Inode, error) {
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
			index, inode, err := dt.loadDirectoryPointerBlockContent(&ptBlock, path, level-1)
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
		if strings.TrimRight(string(content.Name[:]), "\x00") != path[0] || content.Inode == -1 {
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
			result, err := dt.loadFilePointerBlockContent(&pointerBlock, resultString.String(), i-12)
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

	copy(inode.ATime[:], time.Now().Format("2006-01-02 15:04"))
	return resultString.String(), nil
}

// loadFilePointerBlockContent recursively handles the indirect blocks for large files
func (dt *DirTree) loadFilePointerBlockContent(pointerBlock *PointerBlock, resultString string, level int) (string, error) {
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
			result, err := dt.loadFilePointerBlockContent(&ptBlock, resultString, level-1)
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

func (dt *DirTree) AppendToFileInode(s string, inode *Inode) error {
	// ALL OF THIS IS BAD DESIGN SINCE INODE SHOULD NOT BE MODIFIED HERE, JUST RETURN
	var lastBlock int32 = -1
	var lastIndex = 0
	for i, block := range inode.IBlock {
		if block == -1 {
			break
		}
		// indirect addresses
		if i >= 12 {
			fmt.Println("Cannot load indirect addresses, functionality not implemented")
			return errors.New("implement indirect address")
		}

		lastBlock = block
		lastIndex = i
	}

	// load the last file block
	var lastFileBlock FileBlock
	// if last block is not an empty block, meaning that the inode is empty
	if lastBlock != -1 {
		if err := Utils.ReadObject(dt.File, &lastFileBlock, int64(dt.SuperBlock.BlockStart+lastBlock*int32(binary.Size(FileBlock{})))); err != nil {
			return err
		}
	}
	// adding the content to append to the end of the file block
	content := strings.TrimRight(string(lastFileBlock.Content[:]), "\x00")
	content += s

	// getting the content splited in chunks of 64 bytes
	splittedBytesContent := dt.SplitRawBytes(content, binary.Size(FileBlock{}))
	for _, block := range splittedBytesContent {
		var newFileBlock FileBlock
		copy(newFileBlock.Content[:], block)
		// writing to disc logic
		// TODO Check logic if last index is in the last 3 Iblock position
		if lastIndex >= 12 {
			fmt.Println("Cannot handle indirect addresses, functionality not implemented")
			return errors.New("implement indirect addresses creation")
		}
		// if last block was -1 it means that we need a new direction to store the File Block at the last index
		if lastBlock == -1 {
			// load a new File Block Direction
			var err error
			lastBlock, err = dt.GetAvailableBlockAddress()
			if err != nil {
				return err
			}
		}
		// Store the new FileBlock
		if err := Utils.WriteObject(dt.File, newFileBlock, int64(dt.SuperBlock.BlockStart+lastBlock*int32(binary.Size(FileBlock{})))); err != nil {
			return err
		}

		// Update superblock bitmap
		if err := Utils.WriteObject(dt.File, byte(1), int64(dt.SuperBlock.BmBlockStart+lastBlock)); err != nil {
			return err
		}
		dt.BlockBitMap[lastBlock] = 1

		// update inode index block address
		inode.IBlock[lastIndex] = lastBlock

		// next block is going to be a new block
		lastBlock = -1
		// add +1 to the index count
		lastIndex += 1

	}

	// Updating inode info
	inode.Size = int32(len(content))
	copy(inode.MTime[:], time.Now().Format("2006-01-02 15:04"))
	//copy(inode.ATime[:], time.Now().Format("2006-01-02 15:04"))
	//fmt.Println(inode.ToString())
	// updating bitmaps

	return nil
}

/*
func (dt *DirTree) loadFilePointerBlock(pointerBlock *PointerBlock, level int) (int32, error) {
	var lastBlock int32 = -1
	for _, block := range pointerBlock.Pointers {
		if block == -1 {
			break
		}
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
		lastBlock = block
	}

}
*/

func (dt *DirTree) GetAvailableBlockAddress() (int32, error) {
	for i, block := range dt.BlockBitMap {
		if block == 0 {
			return int32(i), nil
		}
	}
	return -1, errors.New("no more free blocks")
}

func (dt *DirTree) GetAvailableInodeAdress() (int32, error) {
	for i, block := range dt.BlockBitMap {
		if block == 0 {
			return int32(i), nil
		}
	}
	return -1, errors.New("no more free inodes")
}

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

func (dt *DirTree) FreeInodeBlockContent(inode *Inode) error {
	for i, block := range inode.IBlock {
		if i >= 12 {
			continue
			// TODO free pointer blocks
			//fmt.Println("Cannot handle indirect addresses, functionality not implemented")
			//panic("Implement indirect addresses creation")
		}

		if block != -1 {
			dt.BlockBitMap[block] = 0
			inode.IBlock[i] = -1
		}
	}
	if err := Utils.WriteObject(dt.File, dt.BlockBitMap, int64(dt.SuperBlock.BmBlockStart)); err != nil {
		return err
	}
	return nil
}

func (dt *DirTree) CreateNewInode(index int32, inode *Inode, name string, UID, GID int32, inodeType [1]byte) (int32, *Inode, error) {
	// ALL OF THIS IS BAD DESIGN SINCE INODE SHOULD NOT BE MODIFIED HERE, JUST RETURN
	var lastBlock int32 = -1
	var lastIndex = 0
	for i, block := range inode.IBlock {
		if block == -1 {
			break
		}
		// indirect addresses
		if i >= 12 {
			fmt.Println("Cannot load indirect addresses, functionality not implemented")
			return -1, nil, errors.New("implement indirect addresses inode creation")
		}

		lastBlock = block
		lastIndex = i
	}

	// load the last file block
	var lastFolderBlock FolderBlock
	var emptyContentIndex = -1
	// last block could never be a -1 since every folder inode has at least 2 folder created
	if err := Utils.ReadObject(dt.File, &lastFolderBlock, int64(dt.SuperBlock.BlockStart+lastBlock*int32(binary.Size(FolderBlock{})))); err != nil {
		return -1, nil, err
	}
	// look for any empty content inode in current folder block
	for i, content := range lastFolderBlock.Content {
		if content.Inode == -1 {
			emptyContentIndex = i
			break
		}
	}

	// if empty content block is -1 then there is a need to create a new folder block for parent inode
	if emptyContentIndex == -1 {
		// new folder block
		lastFolderBlock = *NewFolderBlock()
		// set the index content to 0 since is a new folder block
		emptyContentIndex = 0
		// set the last index of the i node blocks to be +1
		lastIndex += 1
		// get a new block address for new Folder block
		var err error
		lastBlock, err = dt.GetAvailableBlockAddress()
		if err != nil {
			return -1, nil, err
		}
		// TODO implement indirect addresses creation inode
		if lastIndex >= 12 {
			fmt.Println("Cannot handle indirect addresses, functionality not implemented")
			return -1, nil, errors.New("implement indirect addresses creation inode")
		}

		// Update parent inode
		inode.IBlock[lastIndex] = lastBlock
	}
	// update index block
	copy(inode.MTime[:], time.Now().Format("2006-01-02 15:04"))

	// Create new Inode
	newInode := NewInode(inodeType)
	newInode.UID = UID
	newInode.GID = GID
	newInodeIndex, err := dt.GetAvailableInodeAdress()
	if err != nil {
		return -1, nil, err
	}

	// Update information of new entry at the last folder block
	copy(lastFolderBlock.Content[emptyContentIndex].Name[:], name)
	lastFolderBlock.Content[emptyContentIndex].Inode = newInodeIndex

	// Write updated  folder block
	if err := Utils.WriteObject(dt.File, &lastFolderBlock, int64(dt.SuperBlock.BlockStart+lastBlock*int32(binary.Size(FolderBlock{})))); err != nil {
		return -1, nil, err
	}
	if err := Utils.WriteObject(dt.File, byte(1), int64(dt.SuperBlock.BmBlockStart+lastBlock)); err != nil {
		return -1, nil, err
	}
	dt.BlockBitMap[lastBlock] = 1
	// Write new Inode
	if err := Utils.WriteObject(dt.File, newInode, int64(dt.SuperBlock.InodeStart+newInodeIndex*int32(binary.Size(Inode{})))); err != nil {
		return -1, nil, err
	}
	if err := Utils.WriteObject(dt.File, byte(1), int64(dt.SuperBlock.BmInodeStart+newInodeIndex)); err != nil {
		return -1, nil, err
	}
	dt.InodeBitMap[newInodeIndex] = 1
	// Write updated root inode
	if err := Utils.WriteObject(dt.File, inode, int64(dt.SuperBlock.InodeStart+index*int32(binary.Size(Inode{})))); err != nil {
		return -1, nil, err
	}
	return newInodeIndex, newInode, nil
}

func (dt *DirTree) CreateNewDir(index int32, inode *Inode, dirName string, uid int32, gid int32) (int32, *Inode, error) {
	newInodeIndex, newInode, err := dt.CreateNewInode(index, inode, dirName, uid, gid, [1]byte{0})
	if err != nil {
		return -1, nil, err
	}

	newFolderBlock := NewFolderBlock()
	newFolderBlockIndex, err := dt.GetAvailableBlockAddress()
	if err != nil {
		return -1, nil, err
	}

	// default folder block directions
	newFolderBlock.Content[0].Inode = newInodeIndex
	copy(newFolderBlock.Content[0].Name[:], ".")
	newFolderBlock.Content[1].Inode = index
	copy(newFolderBlock.Content[1].Name[:], "..")

	// assign new inode the new folder block
	newInode.IBlock[0] = newFolderBlockIndex

	// write new folder block
	if err := Utils.WriteObject(dt.File, newFolderBlock, int64(dt.SuperBlock.BlockStart+newFolderBlockIndex*int32(binary.Size(FolderBlock{})))); err != nil {
		return -1, nil, err
	}
	if err := Utils.WriteObject(dt.File, byte(1), int64(dt.SuperBlock.BmBlockStart+newFolderBlockIndex)); err != nil {
		return -1, nil, err
	}
	dt.BlockBitMap[newFolderBlockIndex] = 1
	// Write updated new inode
	if err := Utils.WriteObject(dt.File, newInode, int64(dt.SuperBlock.InodeStart+newInodeIndex*int32(binary.Size(Inode{})))); err != nil {
		return -1, nil, err
	}

	return newInodeIndex, newInode, err
}

func (dt *DirTree) CreateNewFile(index int32, inode *Inode, fileName string, content string, uid int32, gid int32) (int32, *Inode, error) {
	newInodeIndex, newInode, err := dt.CreateNewInode(index, inode, fileName, uid, gid, [1]byte{0})
	if err != nil {
		return -1, nil, err
	}

	// load content into file
	if err := dt.AppendToFileInode(content, newInode); err != nil {
		return -1, nil, err
	}

	// Write updated new inode
	if err := Utils.WriteObject(dt.File, newInode, int64(dt.SuperBlock.InodeStart+newInodeIndex*int32(binary.Size(Inode{})))); err != nil {
		return -1, nil, err
	}
	// update old inode

	return newInodeIndex, newInode, err

}
