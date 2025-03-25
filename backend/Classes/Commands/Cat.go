package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"os"
	"strings"
)

type Cat struct {
	Interfaces.CommandStruct
	Files []string
}

func NewCat(line, column int, files []string) *Cat {
	return &Cat{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.CAT,
			Line:   line,
			Column: column,
		},
		Files: files,
	}
}

func (c *Cat) Exec() {
	var consoleString strings.Builder
	consoleString.WriteString("=================CAT=================\n")

	if env.CurrentUser == nil {
		c.AppendError("No user is logged in")
		return
	}

	// Verify disc integrity
	var mbrPartition *Structs.Partition
	var file *os.File
	var err error
	_, mbrPartition, file, err = env.VerifyDiscStatus(*env.CurrentUser)
	if err != nil {
		c.AppendError(err.Error())
		return
	}
	defer file.Close()

	var superBlock Structs.SuperBlock
	if err = Utils.ReadObject(file, &superBlock, int64(mbrPartition.Start)); err != nil {
		c.AppendError(err.Error())
		return
	}

	// creating variables
	var dirTree = Structs.NewDirTree(superBlock, file)
	for _, fileName := range c.Files {
		var fileInode *Structs.Inode
		_, fileInode, err = dirTree.GetInodeByPath(fileName)
		if err != nil {
			c.AppendError(err.Error())
			return
		}

		var read bool
		read, _, _, err = env.CheckFilePermissions(*env.CurrentUser, fileInode)
		if !read {
			c.AppendError("You do not have permission to read file " + fileName)
			continue
		}

		var fileContent string
		fileContent, err = dirTree.GetFileContentByInode(fileInode)
		if err != nil {
			c.AppendError(err.Error())
			return
		}

		consoleString.WriteString(fileContent)
		consoleString.WriteByte('\n')
	}

	consoleString.WriteString("=================END CAT=================\n")
	c.LogConsole(consoleString.String())
}
