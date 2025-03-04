package Commands

import "backend/Classes/Utils"

type Mkdisk struct {
	Result string
	Type   Utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewMkdisk(line, column int, params map[string]string) *Mkdisk {
	return &Mkdisk{
		Type:   Utils.MKDISK,
		Params: params,
		Line:   line,
		Column: column,
	}
}

func (c *Mkdisk) GetLine() int {
	return c.Line
}

func (c *Mkdisk) GetColumn() int {
	return c.Column
}

func (c *Mkdisk) GetType() Utils.Type {
	return c.Type
}

func (c *Mkdisk) Exec() {
	// TODO
	panic("implement me")
}

func (c *Mkdisk) GetResult() string {
	//TODO implement me
	panic("implement me")
}
