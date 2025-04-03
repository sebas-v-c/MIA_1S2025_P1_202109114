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

type Mkusr struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewMkusr(line, column int, params map[string]string) *Mkusr {
	return &Mkusr{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.MKUSR,
			Line:   line,
			Column: column,
		},
		Params: params,
	}
}

func (m *Mkusr) Exec() {
	var consoleString strings.Builder
	consoleString.WriteString("=================MKUSR=================\n")

	if err := m.validateParams(); err != nil {
		m.AppendError(err.Error())
		return
	}

	if env.CurrentUser == nil {
		m.AppendError("No user is logged in")
		return
	}

	if env.CurrentUser.User.Name != "root" || env.CurrentUser.User.Id != 1 || env.CurrentUser.User.Group != "root" {
		m.AppendError("You need to be root user to execute this command")
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

	var dirTree = Structs.NewDirTree(superBlock, file)
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

	var fileContent string
	fileContent, err = dirTree.GetFileContentByInode(fileInode)
	if err != nil {
		m.AppendError(err.Error())
		return
	}

	consoleString.WriteString("Inode content:\n")
	consoleString.WriteString(fileContent)
	consoleString.WriteByte('\n')

	fileContentLines := strings.Split(fileContent, "\n")
	fileContentLines = fileContentLines[:len(fileContentLines)-1]
	var userCounter = 0
	var groupExist = false
	for _, line := range fileContentLines {
		// get the coma separated values from the line and get the id
		words := strings.Split(line, ",")
		// if line is of type group and of length then is not a user
		if len(words) == 3 && words[1] == "G" {
			if words[2] == m.Params["grp"] {
				groupExist = true
			}
			continue
		}
		// if current line is a user
		if len(words) == 5 && words[1] == "U" {
			userCounter++
		}
		// If the user already exist
		if words[3] == m.Params["user"] {
			m.AppendError("user already exists")
			return
		}
	}

	if !groupExist {
		m.AppendError("group '" + m.Params["grp"] + "' not exists")
		return
	}

	// Add one more so if the counter is 1 the next group will be 2
	newGroup := fmt.Sprintf("%d,U,%s,%s,%s\n", userCounter+1, m.Params["grp"], m.Params["user"], m.Params["pass"])
	if err := dirTree.AppendToFileInode(newGroup, fileInode); err != nil {
		m.AppendError(err.Error())
		return
	}

	// update fileInode to disc
	if err := Utils.WriteObject(file, *fileInode, int64(superBlock.InodeStart+fileInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
		m.AppendError(err.Error())
		return
	}

	var writtenInode Structs.Inode
	if err := Utils.ReadObject(file, &writtenInode, int64(superBlock.InodeStart+fileInodeIndex*int32(binary.Size(Structs.Inode{})))); err != nil {
		m.AppendError(err.Error())
		return
	}
	consoleString.WriteString("Written Inode:\n")
	consoleString.WriteString(writtenInode.ToString())

	// print stored content
	fileContent, err = dirTree.GetFileContentByInode(&writtenInode)
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	// Close dirtree by updating the superblock
	if err := Utils.WriteObject(file, dirTree.SuperBlock, int64(mbrPartition.Start)); err != nil {
		m.AppendError(err.Error())
		return
	}

	consoleString.WriteString("\nNew Inode content:\n")
	consoleString.WriteString(fileContent)
	consoleString.WriteString("\n=================END MKUSR=================\n")
	m.LogConsole(consoleString.String())
}

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
