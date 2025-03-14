package Commands

import (
	"backend/Classes/Interfaces"
	"backend/Classes/Utils"
	"errors"
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
