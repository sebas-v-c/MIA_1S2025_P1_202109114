package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Utils"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type Rmdisk struct {
	Result string
	Type   Utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewRmdisk(line, column int, params map[string]string) *Rmdisk {
	return &Rmdisk{Type: Utils.RMDISK, Line: line, Column: column, Params: params}
}

func (r *Rmdisk) GetLine() int {
	return r.Line
}

func (r *Rmdisk) GetColumn() int {
	return r.Column
}

func (r *Rmdisk) GetType() Utils.Type {
	return r.Type
}

func (r *Rmdisk) Exec() {
	if err, ok := r.validParams(); !ok {
		env.Errors = append(env.Errors, env.RuntimeError{
			Line:    r.Line,
			Column:  r.Column,
			Command: Utils.RMDISK,
			Msg:     err.Error(),
		})
		return
	}

	file := os.Remove(r.Params["path"])
	if file != nil {
		env.Errors = append(env.Errors, env.RuntimeError{
			Line:    r.Line,
			Column:  r.Column,
			Command: Utils.RMDISK,
			Msg:     "Error removing file",
		})
		return
	}

	env.CommandLog = append(env.CommandLog, "=================RMDISK=================\n"+"File succesfully removed"+"\n=================END RMDISK=================\n")
}

func (r *Rmdisk) validParams() (error, bool) {
	if _, ok := r.Params["path"]; ok {
		if strings.EqualFold(filepath.Ext(r.Params["path"]), ".mia") {
			return nil, true
		}
		r.Params["path"] = r.Params["path"] + ".mia"
		return nil, true
	}
	return errors.New("obligatory parameter not exist"), false
}

func (r *Rmdisk) GetResult() string {
	//TODO implement me
	panic("implement me")
}
