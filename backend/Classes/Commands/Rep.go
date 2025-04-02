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

	switch r.Params["name"] {
	case "mbr":
		r.createMBRRep()
	case "disk":
		r.createDiskRep()
	case "inode":
		r.createInodeRep()
	case "block":
		r.createBlockRep()
	case "bm_inode":
		r.createBMInodeRep()
	case "bm_block":
		r.createBMBlockRep()
	case "tree":
		r.createTreeRep()
	case "sb":
		r.createSbRep()
	case "file":
		r.createFileRep()
	case "ls":
		r.createLsRep()
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

	name, ok := r.Params["name"]
	if !ok {
		return errors.New("missing parameter -name")
	}

	switch name {
	case "ls", "file":
		if _, ok := r.Params["pfl"]; !ok {
			return errors.New("parameter -path_file_ls is required when generating 'ls' and 'file' reports")
		}
	case "mbr", "disk", "inode", "block", "bm_inode", "bm_block", "tree", "sb":

	default:
		return errors.New("invalid value for -name. Allowed: mbr, disk, inode, block, bm_inode, bm_block, tree, sb, file, ls")
	}

	return nil
}
