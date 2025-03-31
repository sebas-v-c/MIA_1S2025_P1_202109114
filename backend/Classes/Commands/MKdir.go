package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"errors"
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

	//var path = m.Params["path"]
	//var createDirs = false
	if _, ok := m.Params["path"]; ok {
		//createDirs = true
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

	// Getting root inode
	var rootFileInode *Structs.Inode
	_, rootFileInode, err = dirTree.GetInodeByPath("/")
	if err != nil {
		m.AppendError(err.Error())
		return
	}

	consoleString.WriteString("Founded Inode:\n")
	consoleString.WriteString(rootFileInode.ToString())
	consoleString.WriteByte('\n')

	_, write, _, err := env.CheckFilePermissions(*env.CurrentUser, rootFileInode)
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	if !write {
		m.AppendError("You dont have write permissions in this folder")
	}

	m.LogConsole(consoleString.String())

}

func (m *Mkdir) validateParams() error {
	if _, ok := m.Params["path"]; !ok {
		return errors.New("missing parameter -user")
	}
	return nil
}
