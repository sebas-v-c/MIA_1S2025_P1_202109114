package Env

import "backend/Classes/Utils"

type RuntimeError struct {
	Line    int
	Column  int
	Command Utils.Type
	Msg     string
}
type Env struct {
	Errors     []RuntimeError
	CommandLog []string
}

func NewEnv() *Env {
	return &Env{}
}
