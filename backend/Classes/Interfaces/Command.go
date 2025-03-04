package Interfaces

import "backend/Classes/Utils"

type Command interface {
	GetLine() int
	GetColumn() int
	GetType() Utils.Type
	Exec()
	GetResult() string
}
