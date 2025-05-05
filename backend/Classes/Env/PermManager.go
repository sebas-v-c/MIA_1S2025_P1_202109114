// Package Env provides utilities for handling environment operations such as
// file permission checking, user identification, and related functionalities.
package Env

import (
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"errors"
	"strconv"
	"strings"
)

// CheckFilePermissions checks if the given logged user has read, write, and execute permissions
// for the file represented by the provided inode. It returns three booleans corresponding to
// read, write, and execute permissions respectively, along with an error if one occurs.
// The function determines permissions based on the file's owner (UID), group (GID),
// and a lookup of the user's group ID from the mounted partition.
func CheckFilePermissions(loggedUser LoggedUser, fileInode *Structs.Inode) (bool, bool, bool, error) {
	// Extract permission bytes and ownership information from the file inode.
	perms := fileInode.Perm
	fileUID := fileInode.UID
	fileGID := fileInode.GID

	// Retrieve the group ID for the logged user from the mounted partition.
	userGID, err := GetUserGID(loggedUser.User.Group, loggedUser.MountedPartition)
	if err != nil {
		return false, false, false, err
	}

	// Root user (with GID 1 and matching credentials) is granted full permissions.
	if userGID == 1 && loggedUser.User.Name == "root" && loggedUser.User.Group == "root" && loggedUser.User.Id == 1 {
		return true, true, true, nil
	}

	// Determine permissions based on file ownership.
	if fileUID == loggedUser.User.Id {
		// File owner: use the first permission digit.
		read, write, execute := CalculatePermissions(perms[0])
		return read, write, execute, nil
	} else if fileGID == userGID {
		// Group owner: use the second permission digit.
		read, write, execute := CalculatePermissions(perms[1])
		return read, write, execute, nil
	} else {
		// Others: use the third permission digit.
		read, write, execute := CalculatePermissions(perms[2])
		return read, write, execute, nil
	}
}

// CalculatePermissions converts a permission byte (expressed as a character '0'-'7')
// into boolean values representing read, write, and execute permissions.
// The permission digit corresponds to a specific combination:
// '0': no permissions, '1': execute, '2': write, '3': write and execute,
// '4': read, '5': read and execute, '6': read and write, '7': read, write and execute.
func CalculatePermissions(perm byte) (bool, bool, bool) {
	switch perm {
	case '0':
		return false, false, false
	case '1':
		return false, false, true
	case '2':
		return false, true, false
	case '3':
		return false, true, true
	case '4':
		return true, false, false
	case '5':
		return true, false, true
	case '6':
		return true, true, false
	case '7':
		return true, true, true
	default:
		return false, false, false
	}
}

// GetUserGID retrieves the group ID (GID) for a user based on the provided group name and
// mounted partition. It opens the file system corresponding to the mounted partition,
// reads the superblock and directory tree, and then searches for the "/users.txt" file,
// which contains user and group information.
// If found, the function returns the group ID as an int32; otherwise, it returns an error.
func GetUserGID(groupName string, mountedPartition Structs.MountedPartition) (int32, error) {
	// Open the disk file using the mounted partition's path.
	file, err := Utils.OpenFile(mountedPartition.Path)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	// Read the superblock from the disk at the mounted partition's start position.
	var superBlock Structs.SuperBlock
	if err := Utils.ReadObject(file, &superBlock, int64(mountedPartition.Start)); err != nil {
		return -1, err
	}

	// Create a directory tree from the superblock.
	dirTree := Structs.NewDirTree(superBlock, file)
	var fileInode *Structs.Inode

	// Retrieve the inode corresponding to the "/users.txt" file.
	_, fileInode, err = dirTree.GetInodeByPath("/users.txt")
	if err != nil {
		return -1, err
	}

	// Get the content of the "/users.txt" file.
	var fileContent string
	fileContent, err = dirTree.GetFileContentByInode(fileInode)
	if err != nil {
		return -1, err
	}
	// Split the content into lines.
	fileContentLines := strings.Split(fileContent, "\n")

	// Iterate over each line to search for the group entry.
	for _, line := range fileContentLines {
		// Each line is expected to be comma-separated.
		words := strings.Split(line, ",")
		id, _ := strconv.Atoi(words[0])

		if len(words) <= 1 {
			continue
		}

		// Validate the line:
		// - If id == 0, the user was deleted and is invalid.
		// - The second field must be "G" to denote a group.
		// - The line should have exactly 3 fields.
		if words[1] != "G" || id == 0 || len(words) != 3 {
			continue
		}

		// If the group name matches, return the group ID.
		if words[2] == groupName {
			gid, _ := strconv.Atoi(words[0])
			return int32(gid), nil
		}
	}
	// Return an error if the group ID was not found.
	return -1, errors.New("user GID was not found")
}

// GetGroupNameByID retrieves the group name for a given group ID from the mounted partition.
// It returns the group name as a string and an error if the group is not found.
func GetGroupNameByID(groupID int32, mountedPartition Structs.MountedPartition) (string, error) {
	file, err := Utils.OpenFile(mountedPartition.Path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var superBlock Structs.SuperBlock
	if err := Utils.ReadObject(file, &superBlock, int64(mountedPartition.Start)); err != nil {
		return "", err
	}

	dirTree := Structs.NewDirTree(superBlock, file)
	var fileInode *Structs.Inode

	_, fileInode, err = dirTree.GetInodeByPath("/users.txt")
	if err != nil {
		return "", err
	}

	fileContent, err := dirTree.GetFileContentByInode(fileInode)
	if err != nil {
		return "", err
	}

	fileContentLines := strings.Split(fileContent, "\n")

	for _, line := range fileContentLines {
		words := strings.Split(line, ",")
		id, err := strconv.Atoi(words[0])
		if err != nil {
			continue
		}

		if len(words) <= 1 {
			continue
		}

		// Look for group entries (type "G") with matching ID
		if words[1] == "G" && id != 0 && len(words) == 3 && int32(id) == groupID {
			return words[2], nil
		}
	}

	return "", errors.New("group name was not found")
}

// GetUserNameByID retrieves the username for a given user ID from the mounted partition.
// It returns the username as a string and an error if the user is not found.
func GetUserNameByID(userID int32, mountedPartition Structs.MountedPartition) (string, error) {
	file, err := Utils.OpenFile(mountedPartition.Path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var superBlock Structs.SuperBlock
	if err := Utils.ReadObject(file, &superBlock, int64(mountedPartition.Start)); err != nil {
		return "", err
	}

	dirTree := Structs.NewDirTree(superBlock, file)
	var fileInode *Structs.Inode

	_, fileInode, err = dirTree.GetInodeByPath("/users.txt")
	if err != nil {
		return "", err
	}

	fileContent, err := dirTree.GetFileContentByInode(fileInode)
	if err != nil {
		return "", err
	}

	fileContentLines := strings.Split(fileContent, "\n")

	for _, line := range fileContentLines {
		words := strings.Split(line, ",")
		id, err := strconv.Atoi(words[0])
		if err != nil {
			continue
		}

		if len(words) <= 1 {
			continue
		}

		// Look for user entries (type "U") with matching ID
		if words[1] == "U" && id != 0 && len(words) == 5 && int32(id) == userID {
			return words[3], nil
		}
	}

	return "", errors.New("username was not found")
}
