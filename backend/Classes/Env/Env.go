package Env

import "backend/Classes/Utils"

type RuntimeError struct {
	Line    int
	Column  int
	Command Utils.Type
	Msg     string
}

var Errors []RuntimeError
var CommandLog []string

func CleanConsole() {
	Errors = []RuntimeError{}
	CommandLog = []string{}
}

func AddCommandLog(msg string) {
	CommandLog = append(CommandLog, msg)
}

func AddError(err RuntimeError) {
	Errors = append(Errors, err)
}
