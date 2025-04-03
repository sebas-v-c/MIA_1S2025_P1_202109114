// Package Commands provides implementations for various filesystem commands,
// such as adding or removing users and groups.
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

// Rmgrp represents the command to remove a group.
// It embeds the common CommandStruct and includes a map of parameters.
type Rmgrp struct {
	Interfaces.CommandStruct                   // Embedded command structure providing basic metadata.
	Params                   map[string]string // Params holds the command parameters (e.g., the group name to remove).
}

// NewRmgrp creates a new instance of the Rmgrp command with the specified source position and parameters.
func NewRmgrp(line, column int, params map[string]string) *Rmgrp {
	return &Rmgrp{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.RMGROUP, // Command type for removing a group.
			Line:   line,          // Line number where the command is defined.
			Column: column,        // Column number where the command is defined.
		},
		Params: params,
	}
}

// Exec executes the Rmgrp command logic.
// It validates parameters, ensures that the current user is root,
// verifies disk integrity, locates the "/users.txt" inode (which also contains group entries),
// and then marks the specified group as removed in the file content.
func (r *Rmgrp) Exec() {
	var consoleString strings.Builder
	consoleString.WriteString("=================RMGRP=================\n")

	// Validate required parameters.
	if err := r.validateParams(); err != nil {
		r.AppendError(err.Error())
		return
	}
	consoleString.WriteString("Group to remove: " + r.Params["name"])
	consoleString.WriteByte('\n')

	// Ensure a user is logged in.
	if env.CurrentUser == nil {
		r.AppendError("No user is logged in")
		return
	}

	// Check if the current user is root.
	if env.CurrentUser.User.Name != "root" || env.CurrentUser.User.Id != 1 || env.CurrentUser.User.Group != "root" {
		r.AppendError("You need to be root user to execute this command")
		return
	}

	// Verify disk integrity by checking the status of the mounted partition.
	var mbrPartition *Structs.Partition
	var file *os.File
	var err error
	_, mbrPartition, file, err = env.VerifyDiscStatus(env.CurrentUser.MountedPartition.Id)
	if err != nil {
		r.AppendError(err.Error())
		return
	}
	// Ensure the file is closed when the function exits.
	defer file.Close()

	// Read the superblock from disk using the partition's start offset.
	var superBlock Structs.SuperBlock
	if err = Utils.ReadObject(file, &superBlock, int64(mbrPartition.Start)); err != nil {
		r.AppendError(err.Error())
		return
	}

	// Build the directory tree from the superblock.
	var dirTree = Structs.NewDirTree(superBlock, file)
	// Retrieve the inode corresponding to the "/users.txt" file,
	// which contains the user and group entries.
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

	// Retrieve the content of the "/users.txt" file.
	var fileContent string
	fileContent, err = dirTree.GetFileContentByInode(fileInode)
	if err != nil {
		r.AppendError(err.Error())
		return
	}

	consoleString.WriteString("Inode content:\n")
	consoleString.WriteString(fileContent)
	consoleString.WriteByte('\n')

	// Split the file content into individual lines.
	fileContentLines := strings.Split(fileContent, "\n")
	// Remove the last empty line.
	fileContentLines = fileContentLines[:len(fileContentLines)-1]

	// Define a structure to store the index and modified content of the group to be removed.
	groupToRemove := struct {
		Index   int
		Content string
	}{Index: -1, Content: ""}
	// Iterate through the lines to locate the group entry.
	for i, line := range fileContentLines {
		// Split each line by commas.
		words := strings.Split(line, ",")
		// Validate that the line represents a group entry (must have exactly 3 fields and type "G").
		if len(words) != 3 || words[1] != "G" {
			continue
		}

		// Check if the group is already removed (ID field is "0").
		if words[2] == r.Params["name"] && words[0] == "0" {
			r.AppendError("group has already been removed")
			return
		} else if words[2] == r.Params["name"] {
			// Mark the group as removed by setting its ID to "0".
			words[0] = "0"
			groupToRemove.Index = i
			groupToRemove.Content = strings.Join(words, ",")
			break
		}
	}
	// If the group was not found, report an error.
	if groupToRemove.Index == -1 {
		r.AppendError("group does not exist")
		return
	}
	// Update the specified line in the file content with the modified group entry.
	fileContentLines[groupToRemove.Index] = groupToRemove.Content
	// Reassemble the file content.
	fileContent = strings.Join(fileContentLines, "\n") + "\n"

	// Free the current content of the inode (clear its block content).
	err = dirTree.FreeInodeBlockContent(fileInode)
	if err != nil {
		r.AppendError(err.Error())
		return
	}

	consoleString.WriteString("Freed Inode:\n")
	consoleString.WriteString(fileInode.ToString())
	consoleString.WriteByte('\n')

	// Append the modified file content back into the inode.
	err = dirTree.AppendToFileInode(fileContent, fileInode)
	if err != nil {
		r.AppendError(err.Error())
		return
	}

	// Update the inode's modification time.
	copy(fileInode.MTime[:], time.Now().Format("2006-01-02 15:04"))
	// Write the updated inode back to disk.
	if err := Utils.WriteObject(file, *fileInode, int64(superBlock.InodeStart+fileInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
		r.AppendError(err.Error())
		return
	}

	// Read back the written inode for verification.
	var writtenInode Structs.Inode
	if err := Utils.ReadObject(file, &writtenInode, int64(superBlock.InodeStart+fileInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
		r.AppendError(err.Error())
		return
	}
	consoleString.WriteString("Written Inode:\n")
	consoleString.WriteString(writtenInode.ToString())

	// Retrieve and log the stored content from the updated inode.
	fileContent, err = dirTree.GetFileContentByInode(&writtenInode)
	if err != nil {
		r.AppendError(err.Error())
		return
	}

	// Update the superblock in the directory tree to close the dirtree.
	if err := Utils.WriteObject(file, dirTree.SuperBlock, int64(mbrPartition.Start)); err != nil {
		r.AppendError(err.Error())
		return
	}
	consoleString.WriteString("\nNew Inode content:\n")
	consoleString.WriteString(fileContent)
	consoleString.WriteString("\n=================END RMGRP=================\n")
	// Log the complete output to the console.
	r.LogConsole(consoleString.String())
}

// validateParams ensures that the required parameters for the Rmgrp command are present.
// Returns an error if the "name" parameter is missing.
func (r *Rmgrp) validateParams() error {
	if _, ok := r.Params["name"]; !ok {
		return errors.New("missing parameter -name")
	}
	return nil
}
