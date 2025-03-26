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

type Chgrp struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewChgrp(line, column int, params map[string]string) *Chgrp {
	return &Chgrp{
		Interfaces.CommandStruct{
			Type:   Utils.CHRGP,
			Line:   line,
			Column: column,
		},
		params,
	}
}

func (c *Chgrp) Exec() {
	var consoleString strings.Builder
	consoleString.WriteString("=================CHGRP=================\n")

	if err := c.validateParams(); err != nil {
		c.AppendError(err.Error())
		return
	}
	consoleString.WriteString("Group to remove: " + c.Params["name"])
	consoleString.WriteByte('\n')

	if env.CurrentUser == nil {
		c.AppendError("No user is logged in")
		return
	}

	if env.CurrentUser.User.Name != "root" || env.CurrentUser.User.Id != 1 || env.CurrentUser.User.Group != "root" {
		c.AppendError("You need to be root user to execute this command")
		return
	}

	// Verify disc integrity
	var mbrPartition *Structs.Partition
	var file *os.File
	var err error
	_, mbrPartition, file, err = env.VerifyDiscStatus(*env.CurrentUser)
	if err != nil {
		c.AppendError(err.Error())
		return
	}
	defer file.Close()

	var superBlock Structs.SuperBlock
	if err = Utils.ReadObject(file, &superBlock, int64(mbrPartition.Start)); err != nil {
		c.AppendError(err.Error())
		return
	}

	var dirTree = Structs.NewDirTree(superBlock, file)
	var fileInode *Structs.Inode
	var fileInodeIndex int32
	fileInodeIndex, fileInode, err = dirTree.GetInodeByPath("/users.txt")
	if err != nil {
		c.AppendError(err.Error())
		return
	}
	consoleString.WriteString("Founded Inode:\n")
	consoleString.WriteString(fileInode.ToString())
	consoleString.WriteByte('\n')

	var fileContent string
	fileContent, err = dirTree.GetFileContentByInode(fileInode)
	if err != nil {
		c.AppendError(err.Error())
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

		if words[3] == c.Params["user"] && words[0] == "0" {
			c.AppendError("user has already been removed")
			return
		} else if words[3] == c.Params["user"] {
			words[0] = "0"
			userToRemove.Index = i
			userToRemove.Content = strings.Join(words, ",")
			break
		}
	}
	// if the user is not found the index is still -1
	if userToRemove.Index == -1 {
		c.AppendError("user to remove does not exist")
		return
	}
	// store the updated content here
	fileContentLines[userToRemove.Index] = userToRemove.Content
	fileContent = strings.Join(fileContentLines, "\n") + "\n"
	// Liberate all Inode Content
	err = dirTree.FreeInodeBlockContent(fileInode)
	if err != nil {
		c.AppendError(err.Error())
		return
	}

	consoleString.WriteString("Freed Inode:\n")
	consoleString.WriteString(fileInode.ToString())
	consoleString.WriteByte('\n')

	// Append modifiedcontent here
	err = dirTree.AppendToFileInode(fileContent, fileInode)
	if err != nil {
		c.AppendError(err.Error())
		return
	}
	// Update Modified inode time
	copy(fileInode.MTime[:], time.Now().Format("2006-01-02 15:04"))
	if err := Utils.WriteObject(file, *fileInode, int64(superBlock.InodeStart+fileInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
		c.AppendError(err.Error())
		return
	}

	var writtenInode Structs.Inode
	if err := Utils.ReadObject(file, &writtenInode, int64(superBlock.InodeStart+fileInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
		c.AppendError(err.Error())
		return
	}
	consoleString.WriteString("Written Inode:\n")
	consoleString.WriteString(writtenInode.ToString())

	// print stored content
	fileContent, err = dirTree.GetFileContentByInode(&writtenInode)
	if err != nil {
		c.AppendError(err.Error())
		return
	}
	consoleString.WriteString("\nNew Inode content:\n")
	consoleString.WriteString(fileContent)
	consoleString.WriteString("\n=================END RMUSR=================\n")
	c.LogConsole(consoleString.String())
}

func (c *Chgrp) validateParams() error {
	if _, ok := c.Params["user"]; !ok {
		return errors.New("missing parameter -user")
	}
	if _, ok := c.Params["grp"]; !ok {
		return errors.New("missing parameter -grp")
	}
	return nil
}
