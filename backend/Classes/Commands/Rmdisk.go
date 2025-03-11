package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Utils"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type Rmdisk struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewRmdisk(line, column int, params map[string]string) *Rmdisk {
	return &Rmdisk{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.RMDISK,
			Line:   line,
			Column: column,
		},
		Params: params,
	}
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
	if err := r.validateParams(); err != nil {
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

func (r *Rmdisk) validateParams() error {
	if _, ok := r.Params["path"]; ok {
		if !strings.EqualFold(filepath.Ext(r.Params["path"]), ".mia") {
			r.Params["path"] = r.Params["path"] + ".mia"
		}
		return nil
	}
	return errors.New("obligatory parameter not exist")
}

func (r *Rmdisk) GetResult() string {
	//TODO implement me
	panic("implement me")
}
