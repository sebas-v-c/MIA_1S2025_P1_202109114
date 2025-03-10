package Commands

import (
	"backend/Classes/Env"
	"backend/Classes/Utils"
)

type Mount struct {
	Result string
	Type   Utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewMount(line, column int, params map[string]string) *Mount {
	return &Mount{Type: Utils.MOUNT, Line: line, Column: column, Params: params}
}

func (m Mount) GetLine() int {
	return m.Line
}

func (m Mount) GetColumn() int {
	return m.Column
}

func (m Mount) GetType() Utils.Type {
	return m.Type
}

// Using my last two digits of my ID ==> *******14 <==

func (m Mount) Exec(env *Env.Env) {
	if
}

func (m Mount) GetResult() string {
	//TODO implement me
	panic("implement me")
}
