package Commands

import (
	"backend/Classes/Env"
	"backend/Classes/Utils"
)

type Fdisk struct {
	Result string
	Type   Utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewFdisk(line, column int, params map[string]string) *Fdisk {
	return &Fdisk{
		Type:   Utils.FDISK,
		Params: params,
		Line:   line,
		Column: column,
	}
}

func (f *Fdisk) GetLine() int {
	return f.Line
}

func (f *Fdisk) GetColumn() int {
	return f.Column
}

func (f *Fdisk) GetType() Utils.Type {
	return f.Type
}

func (f *Fdisk) Exec(env *Env.Env) {
	if !f.validParams() {
		env.Errors = append(env.Errors, Env.RuntimeError{
			Line:    f.Line,
			Column:  f.Column,
			Command: Utils.FDISK,
			Msg:     "Missing parameters",
		})
		return
	}
}

func (f *Fdisk) validParams() bool {
	if _, ok := f.Params["size"]; !ok {
		return false
	} else if _, ok := f.Params["path"]; !ok {
		return false
	} else if _, ok := f.Params["name"]; !ok {
		return false
	}

	if _, ok := f.Params["type"]; !ok {
		f.Params["type"] = "P"
	}
	if _, ok := f.Params["fit"]; !ok {
		f.Params["fit"] = "WF"
	}
	if _, ok := f.Params["unit"]; !ok {
		f.Params["unit"] = "K"
	}

	return true
}

func (f *Fdisk) GetResult() string {
	//TODO implement me
	panic("implement me")
}
