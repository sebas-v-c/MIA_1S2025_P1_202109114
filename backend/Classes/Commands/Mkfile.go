// Package Commands provides implementations for various filesystem commands,
// including the creation of new files.
package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Mkfile represents the MKFILE command, which is used to create a new file
// in the filesystem. It embeds the common CommandStruct for shared behavior
// and includes a map for command parameters.
type Mkfile struct {
	Interfaces.CommandStruct                   // Embedded command structure providing metadata (Type, Line, Column, etc.)
	Params                   map[string]string // Params holds command-specific parameters, such as "path", "size", and "cont".
}

// NewMkfile creates a new Mkfile command instance with the given source location and parameters.
func NewMkfile(line, column int, params map[string]string) *Mkfile {
	return &Mkfile{
		Interfaces.CommandStruct{
			Type:   Utils.MKFILE, // Command type for creating a file.
			Line:   line,         // Source code line where the command is defined.
			Column: column,       // Source code column where the command is defined.
		},
		params,
	}
}

// Exec executes the MKFILE command to create a new file in the filesystem.
// It validates parameters, verifies disk integrity, optionally creates missing directories,
// prepares file content (from a given size or an external file), locates the parent directory inode,
// and creates a new file inode with the specified content. The updated inode and file content
// are then written to disk, and a log message is generated.
func (m *Mkfile) Exec() {
	var consoleString strings.Builder
	consoleString.WriteString("=================MKFILE=================\n")

	// Validate the command parameters.
	if err := m.validateParams(); err != nil {
		m.AppendError(err.Error())
		return
	}

	// Ensure that a user is logged in.
	if env.CurrentUser == nil {
		m.AppendError("No user is logged in")
		return
	}

	// Verify the disk integrity by checking the status of the mounted partition.
	var mbrPartition *Structs.Partition
	var file *os.File
	var err error
	_, mbrPartition, file, err = env.VerifyDiscStatus(env.CurrentUser.MountedPartition.Id)
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	defer file.Close()

	// Read the superblock from the disk.
	var superBlock Structs.SuperBlock
	if err = Utils.ReadObject(file, &superBlock, int64(mbrPartition.Start)); err != nil {
		m.AppendError(err.Error())
		return
	}

	// Build the directory tree using the superblock and disk file.
	var dirTree = Structs.NewDirTree(superBlock, file)

	// Determine if the "-r" flag is set (which indicates that missing directories should be created).
	var createDirs = false
	if _, ok := m.Params["r"]; ok {
		createDirs = true
	}

	// Split the file path into components.
	splitedPath := strings.Split(m.Params["path"], "/")
	// Remove any leading empty components (i.e. if the path starts with a slash).
	if splitedPath[0] == "" || splitedPath[0] == "/" {
		splitedPath = splitedPath[1:]
	}

	// If the "-r" flag is set, create missing directories.
	// The directories to create are the components of the path excluding the last element (which is the file name).
	if _, ok := m.Params["r"]; ok {
		// Create directories along the given path. If an error occurs, log and return.
		if err, consoleString = CreateDirs(dirTree, splitedPath[:len(splitedPath)-1], createDirs, &consoleString, file, superBlock); err != nil {
			m.AppendError(err.Error())
			return
		}
	}

	// Prepare the content to be written into the file.
	var fileContent = ""
	// If the "size" parameter is provided, generate content by repeating numbers.
	if _, ok := m.Params["size"]; ok {
		if num, err := strconv.Atoi(m.Params["size"]); err == nil {
			if num < 0 {
				m.AppendError("size must be greater than zero")
				return
			}
			// Append digits 0-9 repeatedly to form the content.
			for i := 0; i < num; i++ {
				fileContent += fmt.Sprintf("%d", i%10)
			}
		} else {
			m.AppendError(err.Error())
			return
		}
	}

	// If the "cont" parameter is set, read the content from the specified external file.
	if _, ok := m.Params["cont"]; ok {
		contentFile, err := Utils.OpenFile(m.Params["cont"])
		if err != nil {
			m.AppendError(err.Error())
			return
		}
		defer contentFile.Close()
		content, err := io.ReadAll(contentFile)
		if err != nil {
			m.AppendError(err.Error())
			return
		}
		fileContent = string(content)
		// Alternative approach (commented out) to retrieve file content from an existing inode:
		/*
			_, tempInode, err := dirTree.GetInodeByPath(m.Params["cont"])
			if err != nil {
				m.AppendError(err.Error())
				return
			}
			fileContent, err = dirTree.GetFileContentByInode(tempInode)
			if err != nil {
				m.AppendError(err.Error())
				return
			}
		*/
	}

	// Determine the parent directory for the new file.
	fileDirPath := filepath.Dir(m.Params["path"])
	// Extract the file name from the provided path.
	fileName := filepath.Base(m.Params["path"])
	// Retrieve the parent directory's inode by its path.
	parentInodeIndex, parentInode, err := dirTree.GetInodeByPath(fileDirPath)
	if err != nil {
		m.AppendError(err.Error())
		return
	}

	// Retrieve user and group IDs for the new file.
	UID := env.CurrentUser.User.Id
	GID, err := env.GetUserGID(env.CurrentUser.User.Group, env.CurrentUser.MountedPartition)
	if err != nil {
		m.AppendError(err.Error())
		return
	}

	// Create the new file by updating the parent directory.
	newInodeIndex, _, err := dirTree.CreateNewFile(parentInodeIndex, parentInode, fileName, fileContent, UID, GID)
	if err != nil {
		m.AppendError(err.Error())
		return
	}

	// Read the newly created inode from disk for verification.
	var writtenInode Structs.Inode
	if err := Utils.ReadObject(file, &writtenInode, int64(superBlock.InodeStart+newInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
		m.AppendError(err.Error())
		return
	}
	consoleString.WriteString("\n\nWritten Inode:\n")
	consoleString.WriteString(writtenInode.ToString())

	// Retrieve the updated content from the newly written inode.
	fileContent, err = dirTree.GetFileContentByInode(&writtenInode)
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	consoleString.WriteString("\nNew Inode content:\n")
	consoleString.WriteString(fileContent)

	// Finalize the directory tree by writing the updated superblock to disk.
	if err := Utils.WriteObject(file, dirTree.SuperBlock, int64(mbrPartition.Start)); err != nil {
		m.AppendError(err.Error())
		return
	}
	consoleString.WriteString("\n=================END MKFILE=================\n")
	// Log the complete console output.
	m.LogConsole(consoleString.String())
}

// validateParams ensures that the required parameters for the MKFILE command are provided.
// If both "cont" (content file) and "size" parameters are present, the "size" parameter is removed.
func (m *Mkfile) validateParams() error {
	// Check that the "path" parameter exists.
	if _, ok := m.Params["path"]; !ok {
		return errors.New("missing parameter -user")
	}
	// If "cont" is provided along with "size", remove "size" to avoid conflict.
	if _, ok := m.Params["cont"]; ok {
		if _, ok := m.Params["size"]; ok {
			delete(m.Params, "size")
		}
	}
	return nil
}
