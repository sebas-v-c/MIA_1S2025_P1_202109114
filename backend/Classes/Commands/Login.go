package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"errors"
	"fmt"
	"os"
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
	var discMBR *Structs.MBR
	var discFile *os.File
	mbrPartition, discMBR, discFile = l.checkMountedPartitionHealth(mountedPartition)
	if mbrPartition == nil {
		l.AppendError("Mounted partition does not exist at disc MBR")
		return
	}
	defer discFile.Close()

	var loggedUser *Structs.User
	loggedUser, err := l.loginUser(l.Params["user"], l.Params["pass"], mbrPartition, discMBR, discFile)
	if err != nil {
		l.AppendError(err.Error())
		return
	}
	consoleString += "User '" + loggedUser.ToString() + "' logged in successfully\n"
	env.CurrentUser = &env.LoggedUser{
		User:             *loggedUser,
		MountedPartition: *mountedPartition,
	}
	consoleString = "=================END LOGIN=================\n"
	l.LogConsole(consoleString)
}

func (l *Login) loginUser(user string, password string, partition *Structs.Partition, mbr *Structs.MBR, file *os.File) (*Structs.User, error) {
	panic("implement me")
}

func (l *Login) checkMountedPartitionHealth(mountedPartition *Structs.MountedPartition) (*Structs.Partition, *Structs.MBR, *os.File) {
	file, err := Utils.OpenFile(mountedPartition.Path)
	if err != nil {
		l.AppendError(err.Error())
		return nil, nil, nil
	}
	defer file.Close()
	var discMBR Structs.MBR
	if err := Utils.ReadObject(file, &discMBR, 0); err != nil {
		l.AppendError(err.Error())
		return nil, nil, nil
	}
	for _, partition := range discMBR.Partitions {
		//fmt.Println(partition.ToString())
		if partition.Id == mountedPartition.Id {
			return &partition, &discMBR, file
		}
	}

	return nil, nil, nil
}

func (l *Login) validateParams() error {
	if _, ok := l.Params["path"]; !ok {
		return errors.New("missing parameter -path")
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
