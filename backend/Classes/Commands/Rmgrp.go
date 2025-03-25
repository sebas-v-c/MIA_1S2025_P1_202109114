package Commands

import (
	"backend/Classes/Interfaces"
	"backend/Classes/Utils"
)

type Rmgrp struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewRmgrp(line, column int, params map[string]string) *Rmgrp {
	return &Rmgrp{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.RMGROUP,
			Line:   line,
			Column: column,
		},
		Params: params,
	}
}

func (r *Rmgrp) Exec() {

}
