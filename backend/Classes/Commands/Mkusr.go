package Commands

import (
	"backend/Classes/Interfaces"
	"backend/Classes/Utils"
)

type Mkusr struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewMkusr(line, column int, params map[string]string) *Mkusr {
	return &Mkusr{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.MKUSR,
			Line:   line,
			Column: column,
		},
		Params: params,
	}
}

func (m *Mkusr) Exec() {

}
