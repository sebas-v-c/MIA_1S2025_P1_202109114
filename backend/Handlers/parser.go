package Handlers

import (
	"backend/Classes/Env"
	listener "backend/Language"
	parser "backend/Language/Parser"
	"github.com/antlr4-go/antlr/v4"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CodeRequest struct {
	Code string `json:"code" binding:"required"`
}

type JSONResponse struct {
	SynErrors   []listener.CustomSyntaxError `json:"synErrors" binding:"required"`
	RunErrors   []Env.RuntimeError           `json:"runErrors" binding:"required"`
	CommandLogs []string                     `json:"commandLogs" binding:"required"`
}

func ParseCodeHandler(c *gin.Context) {
	// TODO: parse code and return the result
	var req CodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	inputStream := antlr.NewInputStream(req.Code)
	scanner := parser.NewScanner(inputStream)
	tokens := antlr.NewCommonTokenStream(scanner, antlr.TokenDefaultChannel)
	parserParser := parser.NewParserParser(tokens)
	parserParser.RemoveErrorListeners()
	parserErrors := listener.NewEXT2ErrorListener()
	parserParser.AddErrorListener(parserErrors)

	parserParser.BuildParseTrees = true
	tree := parserParser.Init()
	var ext2Listener = listener.NewEXT2Listener()
	antlr.ParseTreeWalkerDefault.Walk(ext2Listener, tree)

	var synErrors []listener.CustomSyntaxError
	for _, fail := range parserErrors.Errors {
		synErrors = append(synErrors, listener.CustomSyntaxError{Line: fail.Line, Column: fail.Column, Msg: fail.Msg})
	}

	if synErrors != nil {
		c.IndentedJSON(http.StatusAccepted, JSONResponse{
			SynErrors:   synErrors,
			RunErrors:   nil,
			CommandLogs: nil,
		})
		return
	}

	env := Env.NewEnv()
	for _, cmd := range ext2Listener.Execute {
		cmd.Exec(env)
	}

	if len(env.Errors) > 0 {
		c.IndentedJSON(http.StatusAccepted, JSONResponse{SynErrors: synErrors, RunErrors: env.Errors, CommandLogs: env.CommandLog})
		return
	}

	c.IndentedJSON(http.StatusOK, JSONResponse{SynErrors: synErrors, RunErrors: env.Errors, CommandLogs: env.CommandLog})
}
