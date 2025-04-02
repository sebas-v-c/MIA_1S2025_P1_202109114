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

type Mkfile struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewMkfile(line, column int, params map[string]string) *Mkfile {
	return &Mkfile{
		Interfaces.CommandStruct{
			Type:   Utils.MKFILE,
			Line:   line,
			Column: column,
		},
		params,
	}
}

func (m *Mkfile) Exec() {
	var consoleString strings.Builder
	consoleString.WriteString("=================MKFILE=================\n")

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
	_, mbrPartition, file, err = env.VerifyDiscStatus(*env.CurrentUser)
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
	if _, ok := m.Params["r"]; ok {
		createDirs = true
	}
	splitedPath := strings.Split(m.Params["path"], "/")
	if splitedPath[0] == "" || splitedPath[0] == "/" {
		splitedPath = splitedPath[1:] // Remove leading empty component from path
	}

	// If -r flag is set then create the files
	if _, ok := m.Params["r"]; ok {
		// here we send the splitedPath but not the last item since is not a folder, is a file
		if err, consoleString = CreateDirs(dirTree, splitedPath[:len(splitedPath)-1], createDirs, &consoleString, file, superBlock); err != nil {
			m.AppendError(err.Error())
			return
		}
	}

	var fileContent = ""
	// if size flag is set, then fill the file content with number from 0 to 9
	if _, ok := m.Params["size"]; ok {
		if num, err := strconv.Atoi(m.Params["size"]); err == nil {
			if num < 0 {
				m.AppendError("size must be greater than zero")
				return
			}
			for i := 0; i < num; i++ {
				fileContent += fmt.Sprintf("%d", i%10)
			}
		} else {
			m.AppendError(err.Error())
			return
		}
	}

	// if cont flag is set, then read the content of the other file
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

	// File
	fileDirPath := filepath.Dir(m.Params["path"])
	fileName := filepath.Base(m.Params["path"])
	parentInodeIndex, parentInode, err := dirTree.GetInodeByPath(fileDirPath)
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	UID := env.CurrentUser.User.Id
	GID, err := env.GetUserGID(env.CurrentUser.User.Group, env.CurrentUser.MountedPartition)
	if err != nil {
		m.AppendError(err.Error())
		return
	}

	newInodeIndex, _, err := dirTree.CreateNewFile(parentInodeIndex, parentInode, fileName, fileContent, UID, GID)
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	// just because store again the new inode
	/*
		if err := Utils.WriteObject(file, newInode, int64(superBlock.InodeStart+newInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
			m.AppendError(err.Error())
			return
		}
	*/

	var writtenInode Structs.Inode
	if err := Utils.ReadObject(file, &writtenInode, int64(superBlock.InodeStart+newInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
		m.AppendError(err.Error())
		return
	}
	consoleString.WriteString("\n\nWritten Inode:\n")
	consoleString.WriteString(writtenInode.ToString())

	// print stored content
	fileContent, err = dirTree.GetFileContentByInode(&writtenInode)
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	consoleString.WriteString("\nNew Inode content:\n")
	consoleString.WriteString(fileContent)

	consoleString.WriteString("\n=================END MKFILE=================\n")
	m.LogConsole(consoleString.String())
}

func (m *Mkfile) validateParams() error {
	if _, ok := m.Params["path"]; !ok {
		return errors.New("missing parameter -user")
	}
	if _, ok := m.Params["cont"]; ok {
		if _, ok := m.Params["size"]; ok {
			delete(m.Params, "size")
		}
	}
	return nil
}
