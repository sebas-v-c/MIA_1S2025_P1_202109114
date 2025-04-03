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

type Rmusr struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewRmusr(line, column int, params map[string]string) *Rmusr {
	return &Rmusr{
		Interfaces.CommandStruct{
			Type:   Utils.RMUSR,
			Line:   line,
			Column: column,
		},
		params,
	}
}

func (r *Rmusr) Exec() {
	var consoleString strings.Builder
	consoleString.WriteString("=================RMGRP=================\n")

	if err := r.validateParams(); err != nil {
		r.AppendError(err.Error())
		return
	}
	consoleString.WriteString("User to remove: " + r.Params["user"])
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
	_, mbrPartition, file, err = env.VerifyDiscStatus(env.CurrentUser.MountedPartition.Id)
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
	userToRemove := struct {
		Index   int
		Content string
	}{Index: -1, Content: ""}
	for i, line := range fileContentLines {
		// get the coma separated values from the line and get the id
		words := strings.Split(line, ",")
		// if is not of type "G" is also invalid
		// if the length of the line is not 3 then is not valid
		if len(words) != 5 || words[1] != "U" {
			continue
		}

		if words[3] == r.Params["user"] && words[0] == "0" {
			r.AppendError("user has already been removed")
			return
		} else if words[3] == r.Params["user"] {
			words[0] = "0"
			userToRemove.Index = i
			userToRemove.Content = strings.Join(words, ",")
			break
		}
	}
	// if the user is not found the index is still -1
	if userToRemove.Index == -1 {
		r.AppendError("user to remove does not exist")
		return
	}
	// store the updated content here
	fileContentLines[userToRemove.Index] = userToRemove.Content
	fileContent = strings.Join(fileContentLines, "\n") + "\n"
	// Liberate all Inode Content
	err = dirTree.FreeInodeBlockContent(fileInode)
	if err != nil {
		r.AppendError(err.Error())
		return
	}

	consoleString.WriteString("Freed Inode:\n")
	consoleString.WriteString(fileInode.ToString())
	consoleString.WriteByte('\n')

	// Append modifiedcontent here
	err = dirTree.AppendToFileInode(fileContent, fileInode)
	if err != nil {
		r.AppendError(err.Error())
		return
	}
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

	// Close dirtree by updating the superblock
	if err := Utils.WriteObject(file, dirTree.SuperBlock, int64(mbrPartition.Start)); err != nil {
		r.AppendError(err.Error())
		return
	}

	consoleString.WriteString("\nNew Inode content:\n")
	consoleString.WriteString(fileContent)
	consoleString.WriteString("\n=================END RMUSR=================\n")
	r.LogConsole(consoleString.String())
}

func (r *Rmusr) validateParams() error {
	if _, ok := r.Params["user"]; !ok {
		return errors.New("missing parameter -user")
	}
	return nil
}
