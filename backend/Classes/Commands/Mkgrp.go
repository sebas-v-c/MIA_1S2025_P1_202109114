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

// Mkgrp represents the MKGRP command, which is used to create a new group.
// It embeds the common CommandStruct to inherit shared metadata and behavior,
// and holds a map of parameters required for creating the group.
type Mkgrp struct {
	Interfaces.CommandStruct                   // Embedded command structure (provides Type, Line, Column, etc.)
	Params                   map[string]string // Params contains command parameters, e.g., "name" of the group.
}

// NewMkgrp creates a new Mkgrp command instance with the specified source location and parameters.
func NewMkgrp(line, column int, params map[string]string) *Mkgrp {
	return &Mkgrp{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.MKGROUP, // Command type for creating a new group.
			Line:   line,          // Source code line where the command is defined.
			Column: column,        // Source code column where the command is defined.
		},
		Params: params, // Command-specific parameters.
	}
}

// Exec executes the MKGRP command to create a new group.
// It validates parameters, ensures that the current user is root,
// verifies disk integrity, locates the "/users.txt" inode that stores user and group entries,
// and then appends a new group entry if it does not already exist.
func (m *Mkgrp) Exec() {
	// Create a builder for accumulating console log output.
	var consoleString strings.Builder
	consoleString.WriteString("=================MKGRP=================\n")

	// Validate that the required parameter "name" is provided.
	if err := m.validateParams(); err != nil {
		m.AppendError(err.Error())
		return
	}

	// Ensure a user is currently logged in.
	if env.CurrentUser == nil {
		m.AppendError("No user is logged in")
		return
	}

	// Ensure that only the root user can execute this command.
	if env.CurrentUser.User.Name != "root" || env.CurrentUser.User.Id != 1 || env.CurrentUser.User.Group != "root" {
		m.AppendError("You need to be root user to execute this command")
		return
	}

	// Verify disk integrity by retrieving the mounted partition.
	var mbrPartition *Structs.Partition
	var file *os.File
	var err error
	_, mbrPartition, file, err = env.VerifyDiscStatus(env.CurrentUser.MountedPartition.Id)
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	defer file.Close()

	// Read the superblock from disk using the starting offset of the partition.
	var superBlock Structs.SuperBlock
	if err = Utils.ReadObject(file, &superBlock, int64(mbrPartition.Start)); err != nil {
		m.AppendError(err.Error())
		return
	}

	// Build the directory tree from the superblock.
	var dirTree = Structs.NewDirTree(superBlock, file)
	// Retrieve the inode that contains the user and group entries ("/users.txt").
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

	// Read the current content of the "/users.txt" file.
	var fileContent string
	fileContent, err = dirTree.GetFileContentByInode(fileInode)
	if err != nil {
		m.AppendError(err.Error())
		return
	}

	consoleString.WriteString("Inode content:\n")
	consoleString.WriteString(fileContent)
	consoleString.WriteByte('\n')

	// Split the file content into lines for processing.
	fileContentLines := strings.Split(fileContent, "\n")
	// Remove the trailing empty line.
	fileContentLines = fileContentLines[:len(fileContentLines)-1]

	// Initialize a counter for existing group entries.
	var groupCounter = 0
	// Iterate over each line to check for duplicate group names.
	for _, line := range fileContentLines {
		// Split the line into its comma-separated fields.
		words := strings.Split(line, ",")
		// Check if the line represents a group entry (3 fields with type "G").
		if len(words) != 3 || words[1] != "G" {
			continue
		}
		// If a group with the provided name already exists, log an error.
		if words[2] == m.Params["name"] {
			m.AppendError("group already exists")
			return
		}
		// Count the existing group entries.
		groupCounter++
	}

	// Prepare a new group entry.
	// The new group's ID is one greater than the current group count.
	newGroup := fmt.Sprintf("%d,G,%s\n", groupCounter+1, strings.TrimSpace(m.Params["name"]))
	// Append the new group entry to the file by updating the inode's content.
	if err := dirTree.AppendToFileInode(newGroup, fileInode); err != nil {
		m.AppendError(err.Error())
		return
	}

	// Write the updated inode back to disk.
	if err := Utils.WriteObject(file, *fileInode, int64(superBlock.InodeStart+fileInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
		m.AppendError(err.Error())
		return
	}

	// Read back the written inode for verification.
	var writtenInode Structs.Inode
	if err := Utils.ReadObject(file, &writtenInode, int64(superBlock.InodeStart+fileInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
		m.AppendError(err.Error())
		return
	}
	consoleString.WriteString("Written Inode:\n")
	consoleString.WriteString(writtenInode.ToString())

	// Retrieve and log the updated content from the inode.
	fileContent, err = dirTree.GetFileContentByInode(&writtenInode)
	if err != nil {
		m.AppendError(err.Error())
		return
	}

	// Update the superblock on disk to finalize changes to the directory tree.
	if err := Utils.WriteObject(file, dirTree.SuperBlock, int64(mbrPartition.Start)); err != nil {
		m.AppendError(err.Error())
		return
	}

	consoleString.WriteString("\nNew Inode content:\n")
	consoleString.WriteString(fileContent)
	consoleString.WriteString("\n=================END MKGRP=================\n")
	// Log the final output to the console.
	m.LogConsole(consoleString.String())
}

// validateParams ensures that the required parameter "name" for the MKGRP command is provided.
// Returns an error if the parameter is missing.
func (m *Mkgrp) validateParams() error {
	if _, ok := m.Params["name"]; !ok {
		return errors.New("missing parameter -name")
	}
	return nil
}
