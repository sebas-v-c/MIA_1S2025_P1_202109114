package Commands

import (
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

func (r *Rmdisk) Exec() {
	if err := r.validateParams(); err != nil {
		r.AppendError(err.Error())
		return
	}

	file := os.Remove(r.Params["path"])
	if file != nil {
		r.AppendError("Error deleting file: " + r.Params["path"])
		return
	}

	r.LogConsole("=================RMDISK=================\n" + "File succesfully removed" + "\n=================END RMDISK=================\n")
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
