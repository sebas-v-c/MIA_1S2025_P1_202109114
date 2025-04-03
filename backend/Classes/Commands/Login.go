// Package Commands provides implementations for various filesystem commands,
// including user login, logout, file/directory creation, and disk management.
package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Login represents the LOGIN command, which is used to authenticate a user
// against the filesystem's user records stored in "/users.txt".
// It embeds the common CommandStruct for shared behavior and holds a map of parameters.
type Login struct {
	Interfaces.CommandStruct                   // Embeds common command metadata (Type, Line, Column, etc.)
	Params                   map[string]string // Params contains command-specific parameters (e.g., "user", "pass", "id").
}

// NewLogin creates a new instance of the Login command with the specified source location and parameters.
func NewLogin(line, column int, params map[string]string) *Login {
	return &Login{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.LOGIN, // Command type constant for LOGIN.
			Line:   line,        // Source code line where the command is defined.
			Column: column,      // Source code column where the command is defined.
		},
		Params: params, // Command-specific parameters.
	}
}

// Exec executes the LOGIN command.
// It validates the provided parameters, retrieves the mounted partition,
// checks the partition's integrity, and attempts to locate and authenticate the user
// using the information from the "/users.txt" file. If successful, it sets the global
// current user; otherwise, it logs an appropriate error.
func (l *Login) Exec() {
	// Validate that required parameters are provided.
	if err := l.validateParams(); err != nil {
		l.AppendError(err.Error())
		return
	}

	// Begin constructing the console output log.
	var consoleString string
	consoleString = "=================LOGIN=================\n"
	consoleString += fmt.Sprintf("Logging in: user: %s , pass: %s , at '%s'\n", l.Params["user"], l.Params["pass"], l.Params["id"])

	// Check if a user is already logged in.
	if env.CurrentUser != nil {
		l.AppendError("Already logged in")
		return
	}

	// Retrieve the mounted partition from RAM using the provided partition id.
	var mountedPartition *Structs.MountedPartition
	if mountedPartition = l.getPartitionInRam(); mountedPartition == nil {
		l.AppendError("No partition was founded given id")
		return
	}

	// Check the health of the mounted partition by comparing it with the MBR on disk.
	var mbrPartition *Structs.Partition
	var discFile *os.File
	mbrPartition, discFile = l.checkMountedPartitionHealth(mountedPartition)
	if mbrPartition == nil {
		l.AppendError("Mounted partition does not exist at disc MBR")
		return
	}
	defer discFile.Close()

	// Attempt to log in the user by reading the "/users.txt" file from the filesystem.
	var loggedUser *Structs.User
	loggedUser, err := l.loginUser(l.Params["user"], l.Params["pass"], mbrPartition, discFile)
	if err != nil {
		l.AppendError(err.Error())
		return
	}
	consoleString += "User '" + loggedUser.Name + "' logged in successfully\n"

	// Set the global current user with the authenticated user and the mounted partition.
	env.CurrentUser = &env.LoggedUser{
		User:             *loggedUser,
		MountedPartition: *mountedPartition,
	}
	consoleString += "=================END LOGIN=================\n"
	l.LogConsole(consoleString)
}

// loginUser attempts to locate and authenticate a user by searching the "/users.txt" file.
// It reads the filesystem superblock and directory tree, then retrieves the user inode and its content.
// The function returns the matching user if found (and if the provided password matches), or an error otherwise.
func (l *Login) loginUser(user string, password string, partition *Structs.Partition, file *os.File) (*Structs.User, error) {
	var fsSuperBlock Structs.SuperBlock
	if err := Utils.ReadObject(file, &fsSuperBlock, int64(partition.Start)); err != nil {
		return nil, err
	}

	// Build the directory tree from the superblock.
	dirTree := Structs.NewDirTree(fsSuperBlock, file)
	var fileInode *Structs.Inode
	var err error
	_, fileInode, err = dirTree.GetInodeByPath("/users.txt")
	if err != nil {
		return nil, err
	}

	// Retrieve the content of the "/users.txt" file.
	var fileContent string
	fileContent, err = dirTree.GetFileContentByInode(fileInode)
	if err != nil {
		return nil, err
	}

	// Split the file content into lines and iterate over each line.
	fileContentLines := strings.Split(fileContent, "\n")
	for _, line := range fileContentLines {
		// Split each line by commas to extract user details.
		words := strings.Split(line, ",")
		id, _ := strconv.Atoi(words[0])
		// Skip lines that are invalid: user deleted (id == 0), not a user entry, or malformed.
		if words[1] != "U" || id == 0 || len(words) != 5 {
			continue
		}

		// If the username and password match, return the corresponding user.
		if words[3] == user && words[4] == password {
			return &Structs.User{
				Id:       int32(id),
				Group:    words[2],
				Name:     words[3],
				Password: words[4],
			}, nil
		}
	}
	return nil, errors.New("User '" + user + "' does not exist")
}

// checkMountedPartitionHealth verifies that the mounted partition exists on disk by reading the MBR.
// It opens the disk file for the mounted partition, reads the MBR, and searches for the partition
// matching the mounted partition's id. It returns the partition (if found) along with the file pointer.
func (l *Login) checkMountedPartitionHealth(mountedPartition *Structs.MountedPartition) (*Structs.Partition, *os.File) {
	file, err := Utils.OpenFile(mountedPartition.Path)
	if err != nil {
		l.AppendError(err.Error())
		return nil, nil
	}
	var discMBR Structs.MBR
	if err := Utils.ReadObject(file, &discMBR, 0); err != nil {
		l.AppendError(err.Error())
		return nil, nil
	}
	// Search the partitions in the MBR for a matching partition id.
	for _, partition := range discMBR.Partitions {
		if partition.Id == mountedPartition.Id {
			return &partition, file
		}
	}

	return nil, nil
}

// validateParams checks that all required parameters for the LOGIN command are provided.
// It returns an error if any required parameter ("user", "pass", or "id") is missing.
func (l *Login) validateParams() error {
	if _, ok := l.Params["user"]; !ok {
		return errors.New("missing parameter -user")
	} else if _, ok := l.Params["pass"]; !ok {
		return errors.New("missing parameter -pass")
	} else if _, ok := l.Params["id"]; !ok {
		return errors.New("missing parameter -id")
	}
	return nil
}

// getPartitionInRam searches for a mounted partition in memory (RAM) that matches the provided "id" parameter.
// It iterates over the global list of mounted partitions and returns the matching partition, or nil if not found.
func (l *Login) getPartitionInRam() *Structs.MountedPartition {
	for _, part := range env.GetPartitions() {
		if string(part.Id[:]) == l.Params["id"] {
			return part
		}
	}
	return nil
}
