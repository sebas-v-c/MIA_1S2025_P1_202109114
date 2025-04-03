// Package Commands provides implementations for various commands that operate on the filesystem,
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

// Rmusr represents the command to remove a user.
// It embeds the common CommandStruct and includes a map of parameters.
type Rmusr struct {
	Interfaces.CommandStruct                   // Embedded command structure providing basic command metadata.
	Params                   map[string]string // Params holds the command parameters, e.g., the username to remove.
}

// NewRmusr creates a new instance of the Rmusr command with the given source position and parameters.
func NewRmusr(line, column int, params map[string]string) *Rmusr {
	return &Rmusr{
		Interfaces.CommandStruct{
			Type:   Utils.RMUSR, // Command type for removing a user.
			Line:   line,        // Line number where the command is defined.
			Column: column,      // Column number where the command is defined.
		},
		params,
	}
}

// Exec executes the Rmusr command logic.
// It validates parameters, checks if the current user is root,
// verifies disk integrity, locates the "/users.txt" inode,
// and then marks the specified user as removed in the file content.
func (r *Rmusr) Exec() {
	var consoleString strings.Builder
	consoleString.WriteString("=================RMGRP=================\n")

	// Validate required parameters.
	if err := r.validateParams(); err != nil {
		r.AppendError(err.Error())
		return
	}
	consoleString.WriteString("User to remove: " + r.Params["user"])
	consoleString.WriteByte('\n')

	// Ensure a user is logged in.
	if env.CurrentUser == nil {
		r.AppendError("No user is logged in")
		return
	}

	// Check if the logged user is root.
	if env.CurrentUser.User.Name != "root" || env.CurrentUser.User.Id != 1 || env.CurrentUser.User.Group != "root" {
		r.AppendError("You need to be root user to execute this command")
		return
	}

	// Verify disk integrity by checking the mounted partition status.
	var mbrPartition *Structs.Partition
	var file *os.File
	var err error
	_, mbrPartition, file, err = env.VerifyDiscStatus(env.CurrentUser.MountedPartition.Id)
	if err != nil {
		r.AppendError(err.Error())
		return
	}
	defer file.Close()

	// Read the superblock from the disk.
	var superBlock Structs.SuperBlock
	if err = Utils.ReadObject(file, &superBlock, int64(mbrPartition.Start)); err != nil {
		r.AppendError(err.Error())
		return
	}

	// Build the directory tree from the superblock.
	var dirTree = Structs.NewDirTree(superBlock, file)
	// Retrieve the inode corresponding to the "/users.txt" file.
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

	// Split the file content into lines.
	fileContentLines := strings.Split(fileContent, "\n")
	fileContentLines = fileContentLines[:len(fileContentLines)-1]

	// Define a structure to store the index and new content for the user to be removed.
	userToRemove := struct {
		Index   int
		Content string
	}{Index: -1, Content: ""}
	// Iterate through the file lines to locate the user.
	for i, line := range fileContentLines {
		// Split the line by commas.
		words := strings.Split(line, ",")
		// Verify the line is valid and of type "U" (User).
		if len(words) != 5 || words[1] != "U" {
			continue
		}

		// If the user has already been removed (ID field is "0"), report an error.
		if words[3] == r.Params["user"] && words[0] == "0" {
			r.AppendError("user has already been removed")
			return
		} else if words[3] == r.Params["user"] {
			// Mark the user as removed by setting the ID to "0".
			words[0] = "0"
			userToRemove.Index = i
			userToRemove.Content = strings.Join(words, ",")
			break
		}
	}
	// Return an error if the user was not found.
	if userToRemove.Index == -1 {
		r.AppendError("user to remove does not exist")
		return
	}
	// Update the specific line in the file content with the modified user entry.
	fileContentLines[userToRemove.Index] = userToRemove.Content
	fileContent = strings.Join(fileContentLines, "\n") + "\n"

	// Free the inode's current block content.
	err = dirTree.FreeInodeBlockContent(fileInode)
	if err != nil {
		r.AppendError(err.Error())
		return
	}

	consoleString.WriteString("Freed Inode:\n")
	consoleString.WriteString(fileInode.ToString())
	consoleString.WriteByte('\n')

	// Append the modified content back into the file inode.
	err = dirTree.AppendToFileInode(fileContent, fileInode)
	if err != nil {
		r.AppendError(err.Error())
		return
	}
	// Update the inode's modification time.
	copy(fileInode.MTime[:], time.Now().Format("2006-01-02 15:04"))
	// Write the updated inode to disk.
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

	// Retrieve and print the stored content from the updated inode.
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
	consoleString.WriteString("\n=================END RMUSR=================\n")
	// Log the console output.
	r.LogConsole(consoleString.String())
}

// validateParams checks that the required parameters for the Rmusr command are present.
// Returns an error if the "user" parameter is missing.
func (r *Rmusr) validateParams() error {
	if _, ok := r.Params["user"]; !ok {
		return errors.New("missing parameter -user")
	}
	return nil
}
