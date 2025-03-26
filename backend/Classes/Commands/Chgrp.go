package Commands

import (
	"backend/Classes/Interfaces"
	"backend/Classes/Utils"
)

type Chgrp struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewChgrp(line, column int, params map[string]string) *Chgrp {
	return &Chgrp{
		Interfaces.CommandStruct{
			Type:   Utils.CHRGP,
			Line:   line,
			Column: column,
		},
		params,
	}
}

func (c *Chgrp) Exec() {

}
