// Package Commands provides implementations for various filesystem commands.
package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"encoding/binary"
	"os"
	"strings"
	"time"
)

// Cat represents the CAT command, which is used to display the contents of one or more files.
// It embeds the common CommandStruct for shared metadata and behavior,
// and holds a slice of file paths to be read.
type Cat struct {
	Interfaces.CommandStruct          // Embeds common command metadata (Type, Line, Column, etc.)
	Files                    []string // Files is a slice of file paths to be displayed.
}

// NewCat creates a new instance of the Cat command with the specified source location and list of file paths.
func NewCat(line, column int, files []string) *Cat {
	return &Cat{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.CAT, // Command type for CAT.
			Line:   line,      // Source code line where the command is defined.
			Column: column,    // Source code column where the command is defined.
		},
		Files: files, // List of file paths to be read.
	}
}

// Exec executes the CAT command.
// It verifies that a user is logged in, checks the disk integrity, and then iterates over the provided file paths.
// For each file, it verifies read permissions, retrieves the file content, updates the access time,
// and appends the content to the console log.
func (c *Cat) Exec() {
	var consoleString strings.Builder
	consoleString.WriteString("=================CAT=================\n")

	// Ensure that a user is logged in.
	if env.CurrentUser == nil {
		c.AppendError("No user is logged in")
		return
	}

	// Verify disk integrity using the mounted partition.
	var mbrPartition *Structs.Partition
	var file *os.File
	var err error
	_, mbrPartition, file, err = env.VerifyDiscStatus(env.CurrentUser.MountedPartition.Id)
	if err != nil {
		c.AppendError(err.Error())
		return
	}
	defer file.Close()

	// Read the superblock from disk using the partition start offset.
	var superBlock Structs.SuperBlock
	if err = Utils.ReadObject(file, &superBlock, int64(mbrPartition.Start)); err != nil {
		c.AppendError(err.Error())
		return
	}

	// Build the directory tree from the superblock.
	var dirTree = Structs.NewDirTree(superBlock, file)
	// Iterate over each file path provided in the command.
	for _, fileName := range c.Files {
		var fileInode *Structs.Inode
		var fileInodeIndex int32
		// Retrieve the inode for the file by its path.
		fileInodeIndex, fileInode, err = dirTree.GetInodeByPath(fileName)
		if err != nil {
			c.AppendError(err.Error())
			return
		}

		// Check that the current user has read permissions for the file.
		var read bool
		read, _, _, err = env.CheckFilePermissions(*env.CurrentUser, fileInode)
		if !read {
			c.AppendError("You do not have permission to read file " + fileName)
			continue
		}

		// Retrieve the file content from the inode.
		var fileContent string
		fileContent, err = dirTree.GetFileContentByInode(fileInode)
		if err != nil {
			c.AppendError(err.Error())
			return
		}

		// Update the access time of the inode.
		copy(fileInode.ATime[:], time.Now().Format("2006-01-02 15:04"))
		if err = Utils.WriteObject(file, *fileInode, int64(superBlock.InodeStart+fileInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
			c.AppendError(err.Error())
			return
		}

		// Append the file content to the console log.
		consoleString.WriteString(fileContent)
		consoleString.WriteByte('\n')
	}

	consoleString.WriteString("=================END CAT=================\n")
	c.LogConsole(consoleString.String())
}
