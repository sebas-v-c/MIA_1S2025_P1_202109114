package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Utils"
	"fmt"
)

type Mounted struct {
	Interfaces.CommandStruct
}

func NewMounted(line, column int) *Mounted {
	return &Mounted{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.MOUNTED,
			Line:   line,
			Column: column,
		},
	}
}

func (m *Mounted) Exec() {
	consoleString := "=================MOUNTED=================\n"
	if len(env.GetPartitions()) == 0 {
		consoleString += "No partition has been mounted...\n" + "=================END MOUNTED=================\n"
		return
	}

	consoleString += "Loaded partition(s):\n"
	for _, partition := range env.GetPartitions() {
		consoleString += fmt.Sprintf("%s\n", partition.ToString())
	}

	consoleString += "=================END MOUNTED=================\n"
	m.LogConsole(consoleString)
}
