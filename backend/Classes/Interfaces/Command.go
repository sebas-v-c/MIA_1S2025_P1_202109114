package Interfaces

import (
	"backend/Classes/Env"
	"backend/Classes/Utils"
)

type Command interface {
	GetLine() int
	GetColumn() int
	GetType() Utils.Type
	Exec(env *Env.Env)
	GetResult() string
}
