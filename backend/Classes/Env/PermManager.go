package Env

import (
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"errors"
	"strconv"
	"strings"
)

// CheckFilePermissions given a user checks if the user can Read, Write or Execute the given file
func CheckFilePermissions(loggedUser LoggedUser, fileInode *Structs.Inode) (bool, bool, bool, error) {
	perms := fileInode.Perm
	fileUID := fileInode.UID
	fileGID := fileInode.GID
	userGID, err := GetUserGID(loggedUser.User.Group, loggedUser.MountedPartition)
	if err != nil {
		return false, false, false, err
	}

	// Root user (GID 1) has permissions over any file
	if userGID == 1 && loggedUser.User.Name == "root" && loggedUser.User.Group == "root" && loggedUser.User.Id == 1 {
		return true, true, true, nil
	}

	if fileUID == loggedUser.User.Id {
		read, write, execute := CalculatePermissions(perms[0])
		return read, write, execute, nil
	} else if fileGID == userGID {
		read, write, execute := CalculatePermissions(perms[1])
		return read, write, execute, nil
	} else {
		read, write, execute := CalculatePermissions(perms[2])
		return read, write, execute, nil
	}
}

func CalculatePermissions(perm byte) (bool, bool, bool) {
	switch perm {
	case '0':
		// case 0: no permissions
		return false, false, false
	case '1':
		// Execute permission
		return false, false, true
	case '2':
		// Write permission
		return false, true, false
	case '3':
		// Execute and Write permission
		return false, true, true
	case '4':
		// Read permission
		return true, false, false
	case '5':
		// Read and Execute permission
		return true, false, true
	case '6':
		// Read and Write permission
		return true, true, false
	case '7':
		// Read, Write and Execute permission
		return true, true, true
	default:
		return false, false, false
	}
}

func GetUserGID(groupName string, mountedPartition Structs.MountedPartition) (int32, error) {
	file, err := Utils.OpenFile(mountedPartition.Path)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	var superBlock Structs.SuperBlock
	if err := Utils.ReadObject(file, &superBlock, int64(mountedPartition.Start)); err != nil {
		return -1, err
	}

	dirTree := Structs.NewDirTree(superBlock, file)
	var fileInode *Structs.Inode
	_, fileInode, err = dirTree.GetInodeByPath("/users.txt")
	if err != nil {
		return -1, err
	}
	// Get file content given the inode
	var fileContent string
	fileContent, err = dirTree.GetFileContentByInode(fileInode)
	if err != nil {
		return -1, err
	}
	fileContentLines := strings.Split(fileContent, "\n")

	for _, line := range fileContentLines {
		// get the coma separated values from the line and get the id
		words := strings.Split(line, ",")
		id, _ := strconv.Atoi(words[0])
		// if id == 0 then user was deleted and is invalid
		// if is not of type "G" is also invalid
		// if the length of the line is not 3 then is not valid
		if words[1] != "G" || id == 0 || len(words) != 3 {
			continue
		}

		if words[2] == groupName {
			gid, _ := strconv.Atoi(words[0])
			return int32(gid), nil
		}
	}
	return -1, errors.New("user GID was not found")
}
