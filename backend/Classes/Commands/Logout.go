package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Utils"
)

type Logout struct {
	Interfaces.CommandStruct
}

func NewLogout(line, column int) *Logout {
	return &Logout{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.LOGOUT,
			Line:   line,
			Column: column,
		},
	}
}

func (l *Logout) Exec() {
	if env.CurrentUser == nil {
		l.AppendError("You are not logged in")
		return
	}
	var consoleString string
	consoleString += "=================LOGOUT=================\n"
	env.CurrentUser = nil
	consoleString += "Logout success!\n"
	consoleString += "=================END LOGOUT=================\n"
	l.LogConsole(consoleString)
}
