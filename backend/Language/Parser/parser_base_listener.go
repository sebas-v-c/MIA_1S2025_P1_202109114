// Code generated from Parser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Parser

import "github.com/antlr4-go/antlr/v4"

// BaseParserListener is a complete listener for a parse tree produced by ParserParser.
type BaseParserListener struct{}

var _ ParserListener = &BaseParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInit is called when production init is entered.
func (s *BaseParserListener) EnterInit(ctx *InitContext) {}

// ExitInit is called when production init is exited.
func (s *BaseParserListener) ExitInit(ctx *InitContext) {}

// EnterCommands is called when production commands is entered.
func (s *BaseParserListener) EnterCommands(ctx *CommandsContext) {}

// ExitCommands is called when production commands is exited.
func (s *BaseParserListener) ExitCommands(ctx *CommandsContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseParserListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseParserListener) ExitCommand(ctx *CommandContext) {}

// EnterMkdisk is called when production mkdisk is entered.
func (s *BaseParserListener) EnterMkdisk(ctx *MkdiskContext) {}

// ExitMkdisk is called when production mkdisk is exited.
func (s *BaseParserListener) ExitMkdisk(ctx *MkdiskContext) {}

// EnterMkdiskparams is called when production mkdiskparams is entered.
func (s *BaseParserListener) EnterMkdiskparams(ctx *MkdiskparamsContext) {}

// ExitMkdiskparams is called when production mkdiskparams is exited.
func (s *BaseParserListener) ExitMkdiskparams(ctx *MkdiskparamsContext) {}

// EnterMkdiskparam is called when production mkdiskparam is entered.
func (s *BaseParserListener) EnterMkdiskparam(ctx *MkdiskparamContext) {}

// ExitMkdiskparam is called when production mkdiskparam is exited.
func (s *BaseParserListener) ExitMkdiskparam(ctx *MkdiskparamContext) {}

// EnterRmdisk is called when production rmdisk is entered.
func (s *BaseParserListener) EnterRmdisk(ctx *RmdiskContext) {}

// ExitRmdisk is called when production rmdisk is exited.
func (s *BaseParserListener) ExitRmdisk(ctx *RmdiskContext) {}
