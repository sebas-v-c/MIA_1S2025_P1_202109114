package Structs

import (
	"backend/Classes/Utils"
	"encoding/binary"
	"errors"
	"os"
	"strings"
)

type DirTree struct {
	SuperBlock *SuperBlock
	File       *os.File
}

func NewDirTree(superBlock SuperBlock, file *os.File) *DirTree {
	return &DirTree{SuperBlock: &superBlock, File: file}
}

func (dt *DirTree) GetInodeByPath(path string) (int32, *Inode, error) {
	// TODO add a path verification if using "." or ".."
	splitedPath := strings.Split(path, "/")
	if splitedPath[0] == "" {
		splitedPath = splitedPath[1:]
	}

	var firstInode Inode
	if err := Utils.ReadObject(dt.File, &firstInode, int64(dt.SuperBlock.InodeStart)); err != nil {
		return -1, nil, err
	}

	return dt.searchInodeByPath(splitedPath, firstInode)
}

// searchInodeByPath receives the path at the moment and the current inode
func (dt *DirTree) searchInodeByPath(path []string, inode Inode) (int32, *Inode, error) {
	// given current inode look at the i blocks and search for the current item at path
	for i, block := range inode.IBlock {
		// If block is equal to -1 then is empty
		if block == -1 {
			continue
		}

		// TODO if the block is in the last 13 spaces then is an indirect access
		if i > 12 {
			// dt.loadPointerBlock()
			panic("implement indirect block access")
		}

		// Now check if the current inode is a file or a folder
		//if inode.Type == [1]byte{'0'} { // is a folder
		// load the folder block
		var folderBlock FolderBlock
		if err := Utils.ReadObject(dt.File, &folderBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(FolderBlock{})))); err != nil {
			return -1, nil, err
		}
		// loop through the content of the folder block
		for _, content := range folderBlock.Content {
			// if the name of the file is not equal to the current name then ignore
			if strings.TrimRight(string(content.Name[:]), "\x00") != path[0] {
				continue
			}

			// load the next inode
			var nextInode Inode
			if err := Utils.ReadObject(dt.File, &nextInode, int64(dt.SuperBlock.InodeStart+content.Inode*int32(binary.Size(Inode{})))); err != nil {
				return -1, nil, err
			}
			// if there is no more path to follow then return the inode of the current content
			if len(path) == 1 {
				return content.Inode, &nextInode, nil
			} else { // continue to the next iteration removing the first element of the path
				return dt.searchInodeByPath(path[1:], nextInode)
			}
		}
		//} else { // is a file

		//}
	}

	return -1, nil, errors.New("Invalid Inode loaded")
}

func (dt *DirTree) GetFileContentByInode(inode *Inode) (string, error) {
	// Check if inode is of type 1 (file)
	if inode.Type != [1]byte{1} {
		return "", errors.New("inode is not of type file")
	}
	var resultString = ""
	// iterate over the iblocks
	for i, block := range inode.IBlock {
		// if iblock is -1 then is empty, continue
		if block == -1 {
			continue
		}
		// TODO if the index is greater than 12 then its an indirect block so must be implemented for larger files
		if i > 12 {
			panic("implement indirect block access")
		}

		// load fileblock from memory
		var fileBlock FileBlock
		if err := Utils.ReadObject(dt.File, &fileBlock, int64(dt.SuperBlock.BlockStart+block*int32(binary.Size(FileBlock{})))); err != nil {
			return "", err
		}
		// add file content to the result string
		resultString += strings.TrimRight(string(fileBlock.Content[:]), "\x00")
	}
	return resultString, nil
}
