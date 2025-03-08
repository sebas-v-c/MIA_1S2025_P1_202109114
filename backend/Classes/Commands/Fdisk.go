package Commands

import (
	"backend/Classes/Env"
	"backend/Classes/Utils"
	"errors"
	"path/filepath"
	"strconv"
	"strings"
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
	if err, ok := f.validParams(); !ok {
		env.Errors = append(env.Errors, Env.RuntimeError{
			Line:    f.Line,
			Column:  f.Column,
			Command: Utils.FDISK,
			Msg:     err.Error(),
		})
		return
	}
}

func (f *Fdisk) validParams() (error, bool) {
	// obligatory parameters
	if _, ok := f.Params["size"]; !ok {
		return errors.New("missing parameter -size"), false
	} else if _, ok := f.Params["path"]; !ok {
		return errors.New("missing parameter -path"), false
	} else if _, ok := f.Params["name"]; !ok {
		return errors.New("missing parameter -name"), false
	}

	// optional parameters
	if _, ok := f.Params["type"]; !ok {
		f.Params["type"] = "P"
	}
	if _, ok := f.Params["fit"]; !ok {
		f.Params["fit"] = "WF"
	}
	if _, ok := f.Params["unit"]; !ok {
		f.Params["unit"] = "K"
	}

	// check if size is greater than 0
	if val, _ := strconv.Atoi(f.Params["size"]); val <= 0 {
		return errors.New("size must be greater than 0"), false
	}

	// check if file is a disk
	if strings.EqualFold(filepath.Ext(f.Params["path"]), ".mia") {
		return nil, true
	}
	return errors.New("the specified file is not a disk"), false
}

func (f *Fdisk) GetResult() string {
	//TODO implement me
	panic("implement me")
}
