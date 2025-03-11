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
