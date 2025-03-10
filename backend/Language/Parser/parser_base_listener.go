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

// EnterFdisk is called when production fdisk is entered.
func (s *BaseParserListener) EnterFdisk(ctx *FdiskContext) {}

// ExitFdisk is called when production fdisk is exited.
func (s *BaseParserListener) ExitFdisk(ctx *FdiskContext) {}

// EnterFdiskparams is called when production fdiskparams is entered.
func (s *BaseParserListener) EnterFdiskparams(ctx *FdiskparamsContext) {}

// ExitFdiskparams is called when production fdiskparams is exited.
func (s *BaseParserListener) ExitFdiskparams(ctx *FdiskparamsContext) {}

// EnterFdiskparam is called when production fdiskparam is entered.
func (s *BaseParserListener) EnterFdiskparam(ctx *FdiskparamContext) {}

// ExitFdiskparam is called when production fdiskparam is exited.
func (s *BaseParserListener) ExitFdiskparam(ctx *FdiskparamContext) {}

// EnterMount is called when production mount is entered.
func (s *BaseParserListener) EnterMount(ctx *MountContext) {}

// ExitMount is called when production mount is exited.
func (s *BaseParserListener) ExitMount(ctx *MountContext) {}

// EnterMountparams is called when production mountparams is entered.
func (s *BaseParserListener) EnterMountparams(ctx *MountparamsContext) {}

// ExitMountparams is called when production mountparams is exited.
func (s *BaseParserListener) ExitMountparams(ctx *MountparamsContext) {}

// EnterMountparam is called when production mountparam is entered.
func (s *BaseParserListener) EnterMountparam(ctx *MountparamContext) {}

// ExitMountparam is called when production mountparam is exited.
func (s *BaseParserListener) ExitMountparam(ctx *MountparamContext) {}
