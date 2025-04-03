// Package Commands provides implementations for various filesystem commands,
// including changing the group of a user.
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
	"time"
)

// Chgrp represents the CHGRP command, which is used to change the group associated
// with a user in the "/users.txt" file. It embeds the common CommandStruct for shared
// command metadata and behavior, and holds a map of command-specific parameters.
type Chgrp struct {
	Interfaces.CommandStruct                   // Embedded command structure for metadata (Type, Line, Column, etc.)
	Params                   map[string]string // Params holds the command parameters (e.g., "user" and "grp").
}

// NewChgrp creates a new instance of the Chgrp command with the specified source location and parameters.
func NewChgrp(line, column int, params map[string]string) *Chgrp {
	return &Chgrp{
		Interfaces.CommandStruct{
			Type:   Utils.CHRGP, // Command type for CHGRP.
			Line:   line,        // Source code line where the command is defined.
			Column: column,      // Source code column where the command is defined.
		},
		params, // Set command-specific parameters.
	}
}

// Exec executes the CHGRP command.
// It validates parameters, ensures the logged user is root, verifies the disk integrity,
// retrieves the "/users.txt" file, and then updates the group for the specified user.
func (c *Chgrp) Exec() {
	var consoleString strings.Builder
	consoleString.WriteString("=================CHGRP=================\n")

	// Validate the required parameters.
	if err := c.validateParams(); err != nil {
		c.AppendError(err.Error())
		return
	}

	// Ensure that a user is logged in.
	if env.CurrentUser == nil {
		c.AppendError("No user is logged in")
		return
	}

	// Ensure that only the root user can execute this command.
	if env.CurrentUser.User.Name != "root" || env.CurrentUser.User.Id != 1 || env.CurrentUser.User.Group != "root" {
		c.AppendError("You need to be root user to execute this command")
		return
	}

	// Verify disk integrity by checking the mounted partition.
	var mbrPartition *Structs.Partition
	var file *os.File
	var err error
	_, mbrPartition, file, err = env.VerifyDiscStatus(env.CurrentUser.MountedPartition.Id)
	if err != nil {
		c.AppendError(err.Error())
		return
	}
	defer file.Close()

	// Read the superblock from disk.
	var superBlock Structs.SuperBlock
	if err = Utils.ReadObject(file, &superBlock, int64(mbrPartition.Start)); err != nil {
		c.AppendError(err.Error())
		return
	}

	// Build the directory tree from the superblock.
	var dirTree = Structs.NewDirTree(superBlock, file)
	// Retrieve the inode for the "/users.txt" file.
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

	// Read the current content of the "/users.txt" file.
	var fileContent string
	fileContent, err = dirTree.GetFileContentByInode(fileInode)
	if err != nil {
		c.AppendError(err.Error())
		return
	}

	consoleString.WriteString("Inode content:\n")
	consoleString.WriteString(fileContent)
	consoleString.WriteByte('\n')

	// Split file content into lines.
	fileContentLines := strings.Split(fileContent, "\n")
	// Remove the last empty line.
	fileContentLines = fileContentLines[:len(fileContentLines)-1]

	// Prepare temporary structs to store target user and group information.
	targetUser := struct {
		Index int    // Line index of the user entry.
		User  string // Username.
		Pass  string // Password.
		Num   string // User identifier (number) as stored in the file.
	}{Index: -1, User: "", Pass: "", Num: ""}
	targetGroup := struct {
		Index   int    // Line index of the group entry.
		Content string // Group name.
	}{Index: -1, Content: ""}

	// Iterate over each line in the file content.
	for i, line := range fileContentLines {
		// Split each line into fields by commas.
		words := strings.Split(line, ",")
		// Check if the line represents a group entry (3 fields and type "G").
		if len(words) == 3 && words[1] == "G" {
			// If the target group is marked as removed (ID == "0"), report error.
			if words[2] == c.Params["grp"] && words[0] == "0" {
				c.AppendError("Target group has been removed")
				return
			} else if words[2] == c.Params["grp"] { // Otherwise, record the target group.
				targetGroup.Index = i
				targetGroup.Content = words[2]
			}
		} else if len(words) == 5 && words[1] == "U" { // Check if the line represents a user entry.
			// If the target user is marked as removed (ID == "0"), report error.
			if words[3] == c.Params["user"] && words[0] == "0" {
				c.AppendError("Target user has been removed")
				return
			} else if words[3] == c.Params["user"] { // Otherwise, record the target user information.
				targetUser.Index = i
				targetUser.User = words[3]
				targetUser.Pass = words[4]
				targetUser.Num = words[0]
			}
		}
	}

	// Ensure that both the target user and group were found.
	if targetUser.Index == -1 {
		c.AppendError("Target group not found")
		return
	}
	if targetGroup.Index == -1 {
		c.AppendError("Target group not found")
		return
	}

	// Update the user entry with the group information.
	// The user entry is reconstructed using its original identifier and password,
	// while assigning the group from the target group.
	fileContentLines[targetUser.Index] = fmt.Sprintf("%s,U,%s,%s,%s", targetUser.Num, targetGroup.Content, targetUser.User, targetUser.Pass)
	// Reassemble the file content.
	fileContent = strings.Join(fileContentLines, "\n") + "\n"

	// Free all blocks associated with the current inode content.
	err = dirTree.FreeInodeBlockContent(fileInode)
	if err != nil {
		c.AppendError(err.Error())
		return
	}

	consoleString.WriteString("Freed Inode:\n")
	consoleString.WriteString(fileInode.ToString())
	consoleString.WriteByte('\n')

	// Append the modified file content to the inode.
	err = dirTree.AppendToFileInode(fileContent, fileInode)
	if err != nil {
		c.AppendError(err.Error())
		return
	}
	// Update the inode's modification time.
	copy(fileInode.MTime[:], time.Now().Format("2006-01-02 15:04"))
	// Write the updated inode back to disk.
	if err := Utils.WriteObject(file, *fileInode, int64(superBlock.InodeStart+fileInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
		c.AppendError(err.Error())
		return
	}

	// Read back the written inode for verification.
	var writtenInode Structs.Inode
	if err := Utils.ReadObject(file, &writtenInode, int64(superBlock.InodeStart+fileInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
		c.AppendError(err.Error())
		return
	}
	consoleString.WriteString("Written Inode:\n")
	consoleString.WriteString(writtenInode.ToString())

	// Retrieve the new inode content after update.
	fileContent, err = dirTree.GetFileContentByInode(&writtenInode)
	if err != nil {
		c.AppendError(err.Error())
		return
	}
	consoleString.WriteString("\nNew Inode content:\n")
	consoleString.WriteString(fileContent)
	consoleString.WriteString("\n=================END RMUSR=================\n")
	// Log the final output.
	c.LogConsole(consoleString.String())
}

// validateParams checks that all required parameters for the CHGRP command are provided.
// It returns an error if the "user" or "grp" parameter is missing.
func (c *Chgrp) validateParams() error {
	if _, ok := c.Params["user"]; !ok {
		return errors.New("missing parameter -user")
	}
	if _, ok := c.Params["grp"]; !ok {
		return errors.New("missing parameter -grp")
	}
	return nil
}
