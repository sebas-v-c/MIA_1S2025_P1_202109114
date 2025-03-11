package Commands

import (
	"backend/Classes/Env"
	"backend/Classes/Utils"
	"errors"
	"path/filepath"
	"strings"
)

type Mount struct {
	Result string
	Type   Utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewMount(line, column int, params map[string]string) *Mount {
	return &Mount{Type: Utils.MOUNT, Line: line, Column: column, Params: params}
}

func (m *Mount) GetLine() int {
	return m.Line
}

func (m *Mount) GetColumn() int {
	return m.Column
}

func (m *Mount) GetType() Utils.Type {
	return m.Type
}

// Using my last two digits of my ID ==> *******14 <==

func (m *Mount) Exec(env *Env.Env) {
	if err := m.validParams(); err != nil {
		env.Errors = append(env.Errors, m.makeError(err.Error()))
		return
	}
}

func (m *Mount) GetResult() string {
	//TODO implement me
	panic("implement me")
}

func (m *Mount) makeError(msg string) Env.RuntimeError {
	return Env.RuntimeError{m.Line, m.Column, Utils.MOUNT, msg}
}

func (m *Mount) validParams() error {
	if _, ok := m.Params["path"]; !ok {
		return errors.New("missing parameter -path")
	} else if _, ok := m.Params["name"]; !ok {
		return errors.New("missing parameter -name")
	}

	// check if file is a disk
	if strings.EqualFold(filepath.Ext(m.Params["path"]), ".mia") {
		return nil
	}
	m.Params["path"] = m.Params["path"] + ".mia"
	return nil
}
