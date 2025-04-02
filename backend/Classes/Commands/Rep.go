package Commands

import (
	"backend/Classes/Interfaces"
	"backend/Classes/Utils"
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

}
