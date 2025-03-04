package Language

import "github.com/antlr4-go/antlr/v4"

type CustomSyntaxError struct {
	Line   int
	Column int
	Msg    string
}

type EXT2ErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []*CustomSyntaxError
}

func NewEXT2ErrorListener() *EXT2ErrorListener {
	return new(EXT2ErrorListener)
}

func (c *EXT2ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, &CustomSyntaxError{
		Line:   line,
		Column: column + 1,
		Msg:    msg,
	})
}
