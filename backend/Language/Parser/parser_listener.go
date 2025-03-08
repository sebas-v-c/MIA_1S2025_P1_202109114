// Code generated from Parser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Parser

import "github.com/antlr4-go/antlr/v4"

// ParserListener is a complete listener for a parse tree produced by ParserParser.
type ParserListener interface {
	antlr.ParseTreeListener

	// EnterInit is called when entering the init production.
	EnterInit(c *InitContext)

	// EnterCommands is called when entering the commands production.
	EnterCommands(c *CommandsContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterMkdisk is called when entering the mkdisk production.
	EnterMkdisk(c *MkdiskContext)

	// EnterMkdiskparams is called when entering the mkdiskparams production.
	EnterMkdiskparams(c *MkdiskparamsContext)

	// EnterMkdiskparam is called when entering the mkdiskparam production.
	EnterMkdiskparam(c *MkdiskparamContext)

	// EnterRmdisk is called when entering the rmdisk production.
	EnterRmdisk(c *RmdiskContext)

	// EnterFdisk is called when entering the fdisk production.
	EnterFdisk(c *FdiskContext)

	// EnterFdiskparams is called when entering the fdiskparams production.
	EnterFdiskparams(c *FdiskparamsContext)

	// EnterFdiskparam is called when entering the fdiskparam production.
	EnterFdiskparam(c *FdiskparamContext)

	// ExitInit is called when exiting the init production.
	ExitInit(c *InitContext)

	// ExitCommands is called when exiting the commands production.
	ExitCommands(c *CommandsContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitMkdisk is called when exiting the mkdisk production.
	ExitMkdisk(c *MkdiskContext)

	// ExitMkdiskparams is called when exiting the mkdiskparams production.
	ExitMkdiskparams(c *MkdiskparamsContext)

	// ExitMkdiskparam is called when exiting the mkdiskparam production.
	ExitMkdiskparam(c *MkdiskparamContext)

	// ExitRmdisk is called when exiting the rmdisk production.
	ExitRmdisk(c *RmdiskContext)

	// ExitFdisk is called when exiting the fdisk production.
	ExitFdisk(c *FdiskContext)

	// ExitFdiskparams is called when exiting the fdiskparams production.
	ExitFdiskparams(c *FdiskparamsContext)

	// ExitFdiskparam is called when exiting the fdiskparam production.
	ExitFdiskparam(c *FdiskparamContext)
}
