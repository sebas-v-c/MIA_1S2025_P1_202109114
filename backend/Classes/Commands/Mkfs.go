package Commands

import "backend/Classes/Interfaces"

type Mkfs struct {
	Interfaces.CommandStruct
	Params map[string]string
}
