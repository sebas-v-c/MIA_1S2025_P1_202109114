package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"encoding/binary"
	"errors"
	"os"
	"strings"
	"time"
)

type Rmgrp struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewRmgrp(line, column int, params map[string]string) *Rmgrp {
	return &Rmgrp{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.RMGROUP,
			Line:   line,
			Column: column,
		},
		Params: params,
	}
}

func (r *Rmgrp) Exec() {
	var consoleString strings.Builder
	consoleString.WriteString("=================RMGRP=================\n")

	if err := r.validateParams(); err != nil {
		r.AppendError(err.Error())
		return
	}
	consoleString.WriteString("Group to remove: " + r.Params["name"])
	consoleString.WriteByte('\n')

	if env.CurrentUser == nil {
		r.AppendError("No user is logged in")
		return
	}

	if env.CurrentUser.User.Name != "root" || env.CurrentUser.User.Id != 1 || env.CurrentUser.User.Group != "root" {
		r.AppendError("You need to be root user to execute this command")
		return
	}

	// Verify disc integrity
	var mbrPartition *Structs.Partition
	var file *os.File
	var err error
	_, mbrPartition, file, err = env.VerifyDiscStatus(*env.CurrentUser)
	if err != nil {
		r.AppendError(err.Error())
		return
	}
	defer file.Close()

	var superBlock Structs.SuperBlock
	if err = Utils.ReadObject(file, &superBlock, int64(mbrPartition.Start)); err != nil {
		r.AppendError(err.Error())
		return
	}

	var dirTree = Structs.NewDirTree(superBlock, file)
	var fileInode *Structs.Inode
	var fileInodeIndex int32
	fileInodeIndex, fileInode, err = dirTree.GetInodeByPath("/users.txt")
	if err != nil {
		r.AppendError(err.Error())
		return
	}
	consoleString.WriteString("Founded Inode:\n")
	consoleString.WriteString(fileInode.ToString())
	consoleString.WriteByte('\n')

	/*
		ROOT USER HAS PERMISSIONS OVER ANY FILE
		var write bool
		_, write, _, err = env.CheckFilePermissions(*env.CurrentUser, fileInode)
		if !write {
			r.AppendError("You do not have permission to write this file")
		}
	*/

	var fileContent string
	fileContent, err = dirTree.GetFileContentByInode(fileInode)
	if err != nil {
		r.AppendError(err.Error())
		return
	}

	consoleString.WriteString("Inode content:\n")
	consoleString.WriteString(fileContent)
	consoleString.WriteByte('\n')

	fileContentLines := strings.Split(fileContent, "\n")
	fileContentLines = fileContentLines[:len(fileContentLines)-1]
	groupToRemove := struct {
		Index   int
		Content string
	}{Index: -1, Content: ""}
	for i, line := range fileContentLines {
		// get the coma separated values from the line and get the id
		words := strings.Split(line, ",")
		// if is not of type "G" is also invalid
		// if the length of the line is not 3 then is not valid
		if len(words) != 3 || words[1] != "G" {
			continue
		}

		if words[2] == r.Params["name"] && words[0] == "0" {
			r.AppendError("group has already been removed")
			return
		} else if words[2] == r.Params["name"] {
			words[0] = "0"
			groupToRemove.Index = i
			groupToRemove.Content = strings.Join(words, ",")
		}
	}
	// if the group is not found the index is still -1
	if groupToRemove.Index == -1 {
		r.AppendError("group does not exist")
		return
	}
	// store the updated content here
	fileContentLines[groupToRemove.Index] = groupToRemove.Content
	fileContent = strings.Join(fileContentLines, "\n") + "\n"
	// Liberate all Inode Content
	//consoleString.WriteString(fmt.Sprintf("Actual Bitmap %v\n", dirTree.BlockBitMap[:10]))
	err = dirTree.FreeInodeBlockContent(fileInode)
	if err != nil {
		r.AppendError(err.Error())
		return
	}

	//consoleString.WriteString(fmt.Sprintf("Freed Bitmap %v\n", dirTree.BlockBitMap[:10]))
	consoleString.WriteString("Freed Inode:\n")
	consoleString.WriteString(fileInode.ToString())
	consoleString.WriteByte('\n')

	// Append modifiedcontent here
	err = dirTree.AppendToFileInode(fileContent, fileInode)
	if err != nil {
		r.AppendError(err.Error())
		return
	}
	//consoleString.WriteString(fmt.Sprintf("Appended Bitmap %v\n", dirTree.BlockBitMap[:10]))

	/*
		// Get the index of the change here
		changedCharContentIndex := strings.Index(fileContent, groupToRemove.Content)
		// Split file content into an array of 64 characters
		splittedContentBytes := dirTree.SplitRawBytes(fileContent, binary.Size(Structs.FileBlock{}))
		// Get estimated index of the byte block where the changed byte is located
		changedCharBlockIndex := changedCharContentIndex / binary.Size(Structs.FileBlock{})
		// Copy content of the byte block at the changed byte index into a new FileBlock
		var changedFileBlock Structs.FileBlock
		copy(changedFileBlock.Content[:], splittedContentBytes[changedCharBlockIndex])
		// TODO take into account the pointer blocks
		if changedCharBlockIndex >= 12 {
			fmt.Println("Update pointer block not implemented yet")
			panic("Update pointer blocks not implemented yet")
		}
		// Get the address of the block where the original content in the inode is stored
		inodeBlockAddress := fileInode.IBlock[changedCharBlockIndex]

		// Rewrite the new content into the inode Block address
		if err := Utils.WriteObject(file, changedFileBlock, int64(superBlock.BlockStart+inodeBlockAddress*int32(binary.Size(Structs.FileBlock{})))); err != nil {
			r.AppendError(err.Error())
			return
		}
	*/

	// Update Modified inode time
	copy(fileInode.MTime[:], time.Now().Format("2006-01-02 15:04"))
	if err := Utils.WriteObject(file, *fileInode, int64(superBlock.InodeStart+fileInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
		r.AppendError(err.Error())
		return
	}

	var writtenInode Structs.Inode
	if err := Utils.ReadObject(file, &writtenInode, int64(superBlock.InodeStart+fileInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
		r.AppendError(err.Error())
		return
	}
	consoleString.WriteString("Written Inode:\n")
	consoleString.WriteString(writtenInode.ToString())

	// print stored content
	fileContent, err = dirTree.GetFileContentByInode(&writtenInode)
	if err != nil {
		r.AppendError(err.Error())
		return
	}
	consoleString.WriteString("\nNew Inode content:\n")
	consoleString.WriteString(fileContent)
	consoleString.WriteString("\n=================END RMGRP=================\n")
	r.LogConsole(consoleString.String())
}

func (r *Rmgrp) validateParams() error {
	if _, ok := r.Params["name"]; !ok {
		return errors.New("missing parameter -name")
	}
	return nil
}
