package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Mkgrp struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewMkgrp(line, column int, params map[string]string) *Mkgrp {
	return &Mkgrp{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.MKGROUP,
			Line:   line,
			Column: column,
		},
		Params: params,
	}
}

func (m *Mkgrp) Exec() {
	var consoleString strings.Builder
	consoleString.WriteString("=================MKGRP=================\n")

	if err := m.validateParams(); err != nil {
		m.AppendError(err.Error())
		return
	}

	if env.CurrentUser == nil {
		m.AppendError("No user is logged in")
		return
	}

	if env.CurrentUser.User.Name != "root" || env.CurrentUser.User.Id != 1 || env.CurrentUser.User.Group != "root" {
		m.AppendError("You need to be root user to execute this command")
		return
	}

	// Verify disc integrity
	var mbrPartition *Structs.Partition
	var file *os.File
	var err error
	_, mbrPartition, file, err = env.VerifyDiscStatus(*env.CurrentUser)
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	defer file.Close()

	var superBlock Structs.SuperBlock
	if err = Utils.ReadObject(file, &superBlock, int64(mbrPartition.Start)); err != nil {
		m.AppendError(err.Error())
		return
	}

	var dirTree = Structs.NewDirTree(superBlock, file)
	var fileInode *Structs.Inode
	var fileInodeIndex int32
	fileInodeIndex, fileInode, err = dirTree.GetInodeByPath("/users.txt")
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	consoleString.WriteString("Founded Inode:\n")
	consoleString.WriteString(fileInode.ToString())
	consoleString.WriteByte('\n')

	var read, write bool
	read, write, _, err = env.CheckFilePermissions(*env.CurrentUser, fileInode)
	if !read || !write {
		m.AppendError("You do not have permission to read or write this file")
		return
	}

	var fileContent string
	fileContent, err = dirTree.GetFileContentByInode(fileInode)
	if err != nil {
		m.AppendError(err.Error())
		return
	}

	consoleString.WriteString("Inode content:\n")
	consoleString.WriteString(fileContent)
	consoleString.WriteByte('\n')

	fileContentLines := strings.Split(fileContent, "\n")
	fileContentLines = fileContentLines[:len(fileContentLines)-1]
	var groupCounter = 0
	for _, line := range fileContentLines {
		// get the coma separated values from the line and get the id
		words := strings.Split(line, ",")
		// if is not of type "G" is also invalid
		// if the length of the line is not 3 then is not valid
		if len(words) != 3 || words[1] != "G" {
			continue
		}
		if words[2] == m.Params["name"] {
			m.AppendError("group already exists")
			return
		}

		groupCounter++
	}

	// Add one more so if the counter is 1 the next group will be 2
	newGroup := fmt.Sprintf("%d,G,%s\n", groupCounter+1, strings.TrimSpace(m.Params["name"]))
	if err := dirTree.AppendToFileInode(newGroup, fileInode); err != nil {
		m.AppendError(err.Error())
		return
	}

	// update fileInode to disc
	if err := Utils.WriteObject(file, *fileInode, int64(superBlock.InodeStart+fileInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
		m.AppendError(err.Error())
		return
	}

	var writtenInode Structs.Inode
	if err := Utils.ReadObject(file, &writtenInode, int64(superBlock.InodeStart+fileInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
		m.AppendError(err.Error())
		return
	}
	consoleString.WriteString("Written Inode:\n")
	consoleString.WriteString(writtenInode.ToString())

	// print stored content
	fileContent, err = dirTree.GetFileContentByInode(&writtenInode)
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	consoleString.WriteString("\nNew Inode content:\n")
	consoleString.WriteString(fileContent)
	consoleString.WriteString("\n=================END MKGRP=================\n")
	m.LogConsole(consoleString.String())
}

func (m *Mkgrp) validateParams() error {
	if _, ok := m.Params["name"]; !ok {
		return errors.New("missing parameter -name")
	}
	return nil
}
