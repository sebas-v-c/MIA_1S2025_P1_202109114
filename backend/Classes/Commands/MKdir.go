package Commands

import (
	"backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Mkdir represents the MKDIR command, which is used to create a new directory (or directories)
// in the filesystem. It embeds the common CommandStruct to inherit shared metadata and behavior,
// and holds a map of parameters for the command.
type Mkdir struct {
	Interfaces.CommandStruct                   // Embeds common command metadata (Type, Line, Column, etc.)
	Params                   map[string]string // Params holds command parameters (e.g., "path").
}

// NewMkdir creates a new instance of the Mkdir command with the given source location and parameters.
func NewMkdir(line, column int, params map[string]string) *Mkdir {
	return &Mkdir{
		Interfaces.CommandStruct{
			Type:   Utils.MKDIR, // Command type constant for MKDIR.
			Line:   line,        // Line number where the command is defined.
			Column: column,      // Column number where the command is defined.
		},
		params, // Set command-specific parameters.
	}
}

// Exec executes the MKDIR command to create a new directory.
// It validates parameters, verifies disk integrity, and then attempts to create the directory
// (or directories) along the provided path. The final updated superblock is written to disk,
// and a log is generated.
func (m *Mkdir) Exec() {
	var consoleString strings.Builder
	consoleString.WriteString("=================MKDIR=================\n")

	// Validate that the required parameters are provided.
	if err := m.validateParams(); err != nil {
		m.AppendError(err.Error())
		return
	}

	// Ensure that a user is logged in.
	if Env.CurrentUser == nil {
		m.AppendError("No user is logged in")
		return
	}

	// Verify disk integrity by checking the mounted partition status.
	var mbrPartition *Structs.Partition
	var file *os.File
	var err error
	_, mbrPartition, file, err = Env.VerifyDiscStatus(Env.CurrentUser.MountedPartition.Id)
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

	// Check if the "-p" flag is provided, indicating that missing directories should be created.
	var createDirs = false
	if _, ok := m.Params["p"]; ok {
		createDirs = true
	}

	// Split the provided path into components.
	splitedPath := strings.Split(m.Params["path"], "/")
	// Remove a leading empty component if the path starts with a slash.
	if splitedPath[0] == "" || splitedPath[0] == "/" {
		splitedPath = splitedPath[1:]
	}

	// Attempt to create directories along the path.
	// The CreateDirs function will try to create each directory if it doesn't exist.
	// It returns an error and the current console log.
	if err, consoleString = CreateDirs(dirTree, splitedPath, createDirs, &consoleString, file, superBlock); err != nil {
		m.AppendError(err.Error())
		return
	}

	// Update the superblock on disk to finalize the changes made by the directory creation process.
	if err := Utils.WriteObject(file, dirTree.SuperBlock, int64(mbrPartition.Start)); err != nil {
		m.AppendError(err.Error())
		return
	}

	consoleString.WriteString("\n=================END MKDIR=================\n")
	// Log the complete console output.
	m.LogConsole(consoleString.String())
}

// CreateDirs is a helper function that creates directories along a specified path.
// It starts from the root ("/") and iterates through each directory in the provided path.
// If a directory does not exist and the createDirs flag is set (or if it's the last directory),
// it creates the directory and updates the parent inode accordingly.
// It returns an error (if any) and the updated console log.
func CreateDirs(dirTree *Structs.DirTree, splitedPath []string, createDirs bool, consoleString *strings.Builder, file *os.File, superBlock Structs.SuperBlock) (error, strings.Builder) {
	var parentDirInode *Structs.Inode
	var parentDirInodeIndex int32
	var err error

	// Start with the root directory inode.
	parentDirInodeIndex, parentDirInode, err = dirTree.GetInodeByPath("/")
	if err != nil {
		return err, *consoleString
	}

	var joinedPath = "/" // This variable holds the accumulated path.
	var folderCreated = false

	// Iterate through each directory component in the provided path.
	for dirI, dir := range splitedPath {
		// Check write permissions on the current parent directory.
		_, write, _, err := Env.CheckFilePermissions(*Env.CurrentUser, parentDirInode)
		if err != nil {
			return err, *consoleString
		}
		if !write {
			return errors.New("you dont have write permissions in this folder"), *consoleString
		}

		// Try to get the inode for the current directory (joinedPath + current dir).
		dirInodeIndex, dirInode, err := dirTree.GetInodeByPath(joinedPath + dir)
		// If the directory does not exist (dirInodeIndex == -1)...
		if dirInodeIndex == -1 {
			// If the createDirs flag is true, or if this is the last directory in the path,
			// attempt to create the directory.
			if createDirs || dirI == len(splitedPath)-1 {
				// Retrieve UID and GID from the current user.
				UID := Env.CurrentUser.User.Id
				GID, err := Env.GetUserGID(Env.CurrentUser.User.Group, Env.CurrentUser.MountedPartition)
				if err != nil {
					return err, *consoleString
				}

				// Create the new directory in the directory tree.
				dirInodeIndex, dirInode, err = dirTree.CreateNewDir(parentDirInodeIndex, parentDirInode, dir, UID, GID)
				if err != nil {
					return err, *consoleString
				}
				// Log that the directory was created.
				consoleString.WriteString(fmt.Sprintf("\nDirectory %s, created at %s\n", dir, joinedPath))

				// Read back the new directory inode from disk for verification.
				var writtenInode Structs.Inode
				if err := Utils.ReadObject(file, &writtenInode, int64(superBlock.InodeStart+dirInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
					return err, *consoleString
				}
				folderCreated = true
				consoleString.WriteString("Written Inode:\n")
				consoleString.WriteString(writtenInode.ToString())
			} else if err != nil {
				// If createDirs is false and the directory doesn't exist, return an error.
				return errors.New("path to directory does not exist: " + joinedPath + dir), *consoleString
			}
		}

		// Append the current directory to the joined path.
		joinedPath += dir + "/"
		// Update the parent directory inode and its index for the next iteration.
		parentDirInodeIndex = dirInodeIndex
		parentDirInode = dirInode
	}

	// If no new folder was created during the process, return an error indicating the target already exists.
	if !folderCreated {
		return errors.New("target already exist!!"), *consoleString
	}
	return nil, *consoleString
}

// validateParams checks that the required parameters for the MKDIR command are provided.
// It returns an error if the "path" parameter is missing.
func (m *Mkdir) validateParams() error {
	if _, ok := m.Params["path"]; !ok {
		return errors.New("missing parameter -user")
	}
	return nil
}
