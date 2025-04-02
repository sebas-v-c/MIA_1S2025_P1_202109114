package Commands

import (
	"backend/Classes/Interfaces"
	"backend/Classes/Utils"
	"errors"
	"strings"
)

type Rep struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewRep(line, column int, params map[string]string) *Rep {
	return &Rep{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.REP,
			Line:   line,
			Column: column,
		},
		Params: params,
	}
}

func (r *Rep) Exec() {
	var consoleString strings.Builder
	consoleString.WriteString("=================REP=================\n")

	if err := r.validateParams(); err != nil {
		r.AppendError(err.Error())
		return
	}

	consoleString.WriteString("\n=================REP=================\n")
	r.LogConsole(consoleString.String())
}

func (r *Rep) validateParams() error {
	if _, ok := r.Params["path"]; !ok {
		return errors.New("missing parameter -path")
	}

	if _, ok := r.Params["id"]; !ok {
		return errors.New("missing parameter -id")
	}

	if _, ok := r.Params["pfl"]; ok {
		if name, ok := r.Params["name"]; ok {
			if name != "ls" && name != "file" {
				return errors.New("parameter -name must be 'ls' or 'file' if parameter -path_file_ls is present")
			}
			return nil
		} else {
			return errors.New("missing parameter -name")
		}
	}

	if name, ok := r.Params["name"]; ok {
		if name != "mbr" && name != "disk" && name != "inode" && name != "block" && name != "bm_inode" && name != "bm_block" && name != "tree" && name != "sb" {
			return errors.New("report name could only be mbr, disk, inode, block, bm_inode, bm_block, tree, sb, file or ls")
		}
		if name == "ls" || name == "file" {
			return errors.New("ls and file only could be generated if parameter 'pdf' is present")
		}
	} else {
		return errors.New("missing parameter -name")
	}

	return nil
}
