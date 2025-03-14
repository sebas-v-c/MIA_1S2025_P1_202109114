package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"errors"
	"fmt"
)

type Mkfs struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewMkfs(line, column int, params map[string]string) *Mkfs {
	return &Mkfs{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.MKFS,
			Line:   line,
			Column: column,
		},
		Params: params,
	}
}

func (m *Mkfs) Exec() {
	if err := m.validateParams(); err != nil {
		m.AppendError(err.Error())
		return
	}

	var consoleString string
	consoleString = "=================MKFS=================\n"
	consoleString += fmt.Sprintf("Id: %s, Type: %s, Fs: EXT2\n", m.Params["id"], m.Params["type"])

	var mountedPartition *Structs.MountedPartition
	if mountedPartition = m.getPartition(); mountedPartition == nil {
		m.AppendError("Partition ID not found")
		return
	}

	if mountedPartition.Status != [1]byte{'1'} {
		m.AppendError("Partition is not mounted")
		return
	}

}

func (m *Mkfs) validateParams() error {
	if _, ok := m.Params["id"]; !ok {
		return errors.New("missing parameter -id")
	}
	if _, ok := m.Params["type"]; !ok {
		m.Params["type"] = "full"
	}
	return nil
}

func (m *Mkfs) getPartition() *Structs.MountedPartition {
	for _, part := range env.GetPartitions() {
		if string(part.Id[:]) == m.Params["id"] {
			return part
		}
	}
	return nil
}
