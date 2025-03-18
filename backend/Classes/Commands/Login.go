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

type Login struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewLogin(line, column int, params map[string]string) *Login {
	return &Login{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.LOGIN,
			Line:   line,
			Column: column,
		},
		Params: params,
	}
}

func (l *Login) Exec() {
	if err := l.validateParams(); err != nil {
		l.AppendError(err.Error())
		return
	}

	var consoleString string
	consoleString = "=================LOGIN=================\n"
	consoleString += fmt.Sprintf("Logging in: user: %s , pass: %s , at '%s'\n", l.Params["user"], l.Params["pass"], l.Params["id"])

	if env.CurrentUser != nil {
		l.AppendError("Already logged in")
		return
	}

	var mountedPartition *Structs.MountedPartition
	if mountedPartition = l.getPartitionInRam(); mountedPartition == nil {
		l.AppendError("No partition was founded given id")
		return
	}

	var mbrPartition *Structs.Partition
	var discFile *os.File
	mbrPartition, discFile = l.checkMountedPartitionHealth(mountedPartition)
	if mbrPartition == nil {
		l.AppendError("Mounted partition does not exist at disc MBR")
		return
	}
	defer discFile.Close()

	var loggedUser *Structs.User
	loggedUser, err := l.loginUser(l.Params["user"], l.Params["pass"], mbrPartition, discFile)
	if err != nil {
		l.AppendError(err.Error())
		return
	}
	consoleString += "User '" + loggedUser.Name + "' logged in successfully\n"
	env.CurrentUser = &env.LoggedUser{
		User:             *loggedUser,
		MountedPartition: *mountedPartition,
	}
	consoleString += "=================END LOGIN=================\n"
	l.LogConsole(consoleString)
}

func (l *Login) loginUser(user string, password string, partition *Structs.Partition, file *os.File) (*Structs.User, error) {
	var fsSuperBlock Structs.SuperBlock
	if err := Utils.ReadObject(file, &fsSuperBlock, int64(partition.Start)); err != nil {
		return nil, err
	}

	dirTree := Structs.NewDirTree(fsSuperBlock, file)
	//var inodeIndex int32
	var fileInode *Structs.Inode
	var err error
	_, fileInode, err = dirTree.GetInodeByPath("/users.txt")
	if err != nil {
		return nil, err
	}
	// Get file content given the inode
	var fileContent string
	fileContent, err = dirTree.GetFileContentByInode(fileInode)
	if err != nil {
		return nil, err
	}
	// split in lines the file content and iterate over it
	fileContentLines := strings.Split(fileContent, "\n")
	for _, line := range fileContentLines {
		// get the coma separated values from the line and get the id
		words := strings.Split(line, ",")
		id, _ := strconv.Atoi(words[0])
		// if id == 0 then user was deleted and is invalid
		// if is not of type "U" is also invalid
		// if the length of the line is not 5 then is not valid
		if words[1] != "U" || id == 0 || len(words) != 5 {
			continue
		}

		if words[3] == user && words[4] == password {
			return &Structs.User{
				Id:       id,
				Group:    words[2],
				Name:     words[3],
				Password: words[4],
			}, nil
		}
	}
	return nil, errors.New("User '" + user + "' does not exist")
}

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
	for _, partition := range discMBR.Partitions {
		//fmt.Println(partition.ToString())
		if partition.Id == mountedPartition.Id {
			return &partition, file
		}
	}

	return nil, nil
}

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

func (l *Login) getPartitionInRam() *Structs.MountedPartition {
	for _, part := range env.GetPartitions() {
		if string(part.Id[:]) == l.Params["id"] {
			return part
		}
	}
	return nil
}
