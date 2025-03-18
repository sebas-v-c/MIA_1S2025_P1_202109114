package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Utils"
	"errors"
	"fmt"
)

type Mkgrp struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewMkgrp(line, column int, params map[string]string) *Mkgrp {
	return &Mkgrp{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.MKGROUP,
			Line:   line,
			Column: column,
		},
		Params: params,
	}
}

func (m *Mkgrp) Exec() {
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
	mountedPartition, mbrPartition, file, err := env.VerifyDiscStatus(*env.CurrentUser)
	if err != nil {
		m.AppendError(err.Error())
		return
	}
	defer file.Close()
	fmt.Println(mountedPartition.ToString(), mbrPartition.ToString())

}

func (m *Mkgrp) validateParams() error {
	if _, ok := m.Params["name"]; !ok {
		return errors.New("missing parameter -name")
	}
	return nil
}
