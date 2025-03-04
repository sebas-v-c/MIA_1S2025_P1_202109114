package Language

import (
	interfaces "backend/Classes/Interfaces"
	parser "backend/Language/Parser"
)

type EXT2Listener struct {
	*parser.BaseParserListener
	Execute []interfaces.Command
}

func NewEXT2Listener() *EXT2Listener {
	return new(EXT2Listener)
}

func (l *EXT2Listener) ExitInit(c *parser.InitContext) {
	l.Execute = c.GetResult()
}
