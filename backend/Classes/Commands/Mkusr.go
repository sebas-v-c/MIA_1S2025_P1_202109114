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

// Mkusr represents the MKUSR command, which is used to create a new user.
// It embeds the common CommandStruct to inherit shared metadata and behavior,
// and it holds a map of parameters required for user creation.
type Mkusr struct {
	Interfaces.CommandStruct                   // Embedded command structure (provides Type, Line, Column, etc.)
	Params                   map[string]string // Params contains command parameters such as "user", "pass", and "grp".
}

// NewMkusr creates a new Mkusr command instance with the specified source location and parameters.
func NewMkusr(line, column int, params map[string]string) *Mkusr {
	return &Mkusr{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.MKUSR, // Command type constant for MKUSR.
			Line:   line,        // Line number where the command is defined.
			Column: column,      // Column number where the command is defined.
		},
		Params: params, // Command-specific parameters.
	}
}

// Exec executes the MKUSR command to create a new user.
// It validates parameters, checks that the current user is root, verifies disk integrity,
// locates the "/users.txt" inode, and appends a new user entry to that file.
func (m *Mkusr) Exec() {
	// Build a console log to accumulate messages during execution.
	var consoleString strings.Builder
	consoleString.WriteString("=================MKUSR=================\n")

	// Validate that required parameters ("user", "pass", "grp") are provided.
	if err := m.validateParams(); err != nil {
		m.AppendError(err.Error())
		return
	}

	// Check if a user is currently logged in.
	if env.CurrentUser == nil {
		m.AppendError("No user is logged in")
		return
	}

	// Ensure that the logged-in user is the root user.
	if env.CurrentUser.User.Name != "root" || env.CurrentUser.User.Id != 1 || env.CurrentUser.User.Group != "root" {
		m.AppendError("You need to be root user to execute this command")
		return
	}

	// Verify disk integrity by retrieving the mounted partition information.
	var mbrPartition *Structs.Partition
	var file *os.File
	var err error
	_, mbrPartition, file, err = env.VerifyDiscStatus(env.CurrentUser.MountedPartition.Id)
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	defer file.Close()

	// Read the superblock from the disk using the starting position of the partition.
	var superBlock Structs.SuperBlock
	if err = Utils.ReadObject(file, &superBlock, int64(mbrPartition.Start)); err != nil {
		m.AppendError(err.Error())
		return
	}

	// Build the directory tree from the superblock.
	var dirTree = Structs.NewDirTree(superBlock, file)
	// Retrieve the inode for the "/users.txt" file, which stores user and group information.
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

	// Read the content of the "/users.txt" file.
	var fileContent string
	fileContent, err = dirTree.GetFileContentByInode(fileInode)
	if err != nil {
		m.AppendError(err.Error())
		return
	}

	consoleString.WriteString("Inode content:\n")
	consoleString.WriteString(fileContent)
	consoleString.WriteByte('\n')

	// Split file content into lines for processing.
	fileContentLines := strings.Split(fileContent, "\n")
	// Remove the last empty line.
	fileContentLines = fileContentLines[:len(fileContentLines)-1]

	// Initialize variables for counting existing users and validating group existence.
	var userCounter = 0
	var groupExist = false

	// Iterate over each line to process user and group entries.
	for _, line := range fileContentLines {
		// Split the line by commas to extract individual fields.
		words := strings.Split(line, ",")
		// If the line represents a group entry (3 fields and type "G"), check if the specified group exists.
		if len(words) == 3 && words[1] == "G" {
			if words[2] == m.Params["grp"] {
				groupExist = true
			}
			continue
		}
		// If the line represents a user entry (5 fields and type "U"), count it.
		if len(words) == 5 && words[1] == "U" {
			userCounter++
		}
		// Check if a user with the provided username already exists.
		if words[3] == m.Params["user"] {
			m.AppendError("user already exists")
			return
		}
	}

	// If the specified group does not exist, log an error.
	if !groupExist {
		m.AppendError("group '" + m.Params["grp"] + "' not exists")
		return
	}

	// Prepare the new user entry.
	// The new user's ID is one greater than the current number of user entries.
	newGroup := fmt.Sprintf("%d,U,%s,%s,%s\n", userCounter+1, m.Params["grp"], m.Params["user"], m.Params["pass"])
	// Append the new user entry to the "/users.txt" file by updating the inode's content.
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

	// Retrieve the updated content from the written inode.
	fileContent, err = dirTree.GetFileContentByInode(&writtenInode)
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	// Update the superblock in the directory tree to finalize changes.
	if err := Utils.WriteObject(file, dirTree.SuperBlock, int64(mbrPartition.Start)); err != nil {
		m.AppendError(err.Error())
		return
	}

	consoleString.WriteString("\nNew Inode content:\n")
	consoleString.WriteString(fileContent)
	consoleString.WriteString("\n=================END MKUSR=================\n")
	// Log the complete console output.
	m.LogConsole(consoleString.String())
}

// validateParams ensures that all required parameters for the MKUSR command are provided
// and that their lengths do not exceed allowed limits.
// Returns an error if any parameter is missing or too large.
func (m *Mkusr) validateParams() error {
	if _, ok := m.Params["user"]; !ok {
		return errors.New("missing parameter -name")
	}
	if _, ok := m.Params["pass"]; !ok {
		return errors.New("missing parameter -pass")
	}
	if _, ok := m.Params["grp"]; !ok {
		return errors.New("missing parameter -grp")
	}

	if len(m.Params["user"]) > 10 {
		return errors.New("parameter -user too large")
	}
	if len(m.Params["pass"]) > 10 {
		return errors.New("parameter -pass too large")
	}
	if len(m.Params["grp"]) > 10 {
		return errors.New("parameter -grp too large")
	}

	return nil
}
