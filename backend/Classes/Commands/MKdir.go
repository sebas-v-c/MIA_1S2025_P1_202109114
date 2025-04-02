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

type Mkdir struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewMkdir(line, column int, params map[string]string) *Mkdir {
	return &Mkdir{
		Interfaces.CommandStruct{
			Type:   Utils.MKDIR,
			Line:   line,
			Column: column,
		},
		params,
	}
}

func (m *Mkdir) Exec() {
	var consoleString strings.Builder
	consoleString.WriteString("=================MKDIR=================\n")

	if err := m.validateParams(); err != nil {
		m.AppendError(err.Error())
		return
	}

	if env.CurrentUser == nil {
		m.AppendError("No user is logged in")
		return
	}

	// Verify disc integrity
	var mbrPartition *Structs.Partition
	var file *os.File
	var err error
	_, mbrPartition, file, err = env.VerifyDiscStatus(env.CurrentUser.MountedPartition.Id)
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	defer file.Close()

	var superBlock Structs.SuperBlock
	if err = Utils.ReadObject(file, &superBlock, int64(mbrPartition.Start)); err != nil {
		m.AppendError(err.Error())
		return
	}

	// Get dir tree
	var dirTree = Structs.NewDirTree(superBlock, file)

	var createDirs = false
	if _, ok := m.Params["p"]; ok {
		createDirs = true
	}
	splitedPath := strings.Split(m.Params["path"], "/")
	if splitedPath[0] == "" || splitedPath[0] == "/" {
		splitedPath = splitedPath[1:] // Remove leading empty component from path
	}

	if err, consoleString = CreateDirs(dirTree, splitedPath, createDirs, &consoleString, file, superBlock); err != nil {
		m.AppendError(err.Error())
		return
	}

	consoleString.WriteString("\n=================END MKDIR=================\n")
	m.LogConsole(consoleString.String())
}

func CreateDirs(dirTree *Structs.DirTree, splitedPath []string, createDirs bool, consoleString *strings.Builder, file *os.File, superBlock Structs.SuperBlock) (error, strings.Builder) {
	var parentDirInode *Structs.Inode
	var parentDirInodeIndex int32
	var err error
	parentDirInodeIndex, parentDirInode, err = dirTree.GetInodeByPath("/")
	if err != nil {
		return err, *consoleString
	}

	var joinedPath = "/"
	var folderCreated = false

	for dirI, dir := range splitedPath {
		// check permissions on parent folder
		_, write, _, err := env.CheckFilePermissions(*env.CurrentUser, parentDirInode)
		if err != nil {
			return err, *consoleString
		}
		if !write {
			return errors.New("you dont have write permissions in this folder"), *consoleString
		}

		dirInodeIndex, dirInode, err := dirTree.GetInodeByPath(joinedPath + dir)
		// if dirInodeIndex is -1 that means the path to the dir does not exist
		if dirInodeIndex == -1 {
			// if createDirs flag is true or if this is the last dir of the path then try and create directory
			if createDirs || dirI == len(splitedPath)-1 {
				// WARNING: this function also updates the parentDirNode
				UID := env.CurrentUser.User.Id
				GID, err := env.GetUserGID(env.CurrentUser.User.Group, env.CurrentUser.MountedPartition)
				if err != nil {
					return err, *consoleString
				}

				dirInodeIndex, dirInode, err = dirTree.CreateNewDir(parentDirInodeIndex, parentDirInode, dir, UID, GID)
				if err != nil {
					return err, *consoleString
				}
				consoleString.WriteString(fmt.Sprintf("\nDirectory %s, created at %s\n", dir, joinedPath))

				var writtenInode Structs.Inode
				if err := Utils.ReadObject(file, &writtenInode, int64(superBlock.InodeStart+dirInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
					return err, *consoleString
				}
				folderCreated = true
				consoleString.WriteString("Written Inode:\n")
				consoleString.WriteString(writtenInode.ToString())
			} else if err != nil { // if createDirs is false and is not the last dir then return an error
				return errors.New("path to directory does not exist: " + joinedPath + dir), *consoleString
			}
		}

		joinedPath += dir + "/"
		// update parents to new created folder
		parentDirInodeIndex = dirInodeIndex
		parentDirInode = dirInode
	}

	if !folderCreated {
		return errors.New("target already exist!!"), *consoleString
	}
	return nil, *consoleString
}

func (m *Mkdir) validateParams() error {
	if _, ok := m.Params["path"]; !ok {
		return errors.New("missing parameter -user")
	}
	return nil
}
