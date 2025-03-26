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

	// EnterMount is called when entering the mount production.
	EnterMount(c *MountContext)

	// EnterMountparams is called when entering the mountparams production.
	EnterMountparams(c *MountparamsContext)

	// EnterMountparam is called when entering the mountparam production.
	EnterMountparam(c *MountparamContext)

	// EnterMounted is called when entering the mounted production.
	EnterMounted(c *MountedContext)

	// EnterMkfs is called when entering the mkfs production.
	EnterMkfs(c *MkfsContext)

	// EnterMkfsparams is called when entering the mkfsparams production.
	EnterMkfsparams(c *MkfsparamsContext)

	// EnterMkfsparam is called when entering the mkfsparam production.
	EnterMkfsparam(c *MkfsparamContext)

	// EnterCat is called when entering the cat production.
	EnterCat(c *CatContext)

	// EnterCatparams is called when entering the catparams production.
	EnterCatparams(c *CatparamsContext)

	// EnterCatparam is called when entering the catparam production.
	EnterCatparam(c *CatparamContext)

	// EnterLogin is called when entering the login production.
	EnterLogin(c *LoginContext)

	// EnterLoginparams is called when entering the loginparams production.
	EnterLoginparams(c *LoginparamsContext)

	// EnterLoginparam is called when entering the loginparam production.
	EnterLoginparam(c *LoginparamContext)

	// EnterLogout is called when entering the logout production.
	EnterLogout(c *LogoutContext)

	// EnterMkgrp is called when entering the mkgrp production.
	EnterMkgrp(c *MkgrpContext)

	// EnterMkgrpparams is called when entering the mkgrpparams production.
	EnterMkgrpparams(c *MkgrpparamsContext)

	// EnterMkgrpparam is called when entering the mkgrpparam production.
	EnterMkgrpparam(c *MkgrpparamContext)

	// EnterRmgrp is called when entering the rmgrp production.
	EnterRmgrp(c *RmgrpContext)

	// EnterRmgrpparams is called when entering the rmgrpparams production.
	EnterRmgrpparams(c *RmgrpparamsContext)

	// EnterRmgrpparam is called when entering the rmgrpparam production.
	EnterRmgrpparam(c *RmgrpparamContext)

	// EnterMkusr is called when entering the mkusr production.
	EnterMkusr(c *MkusrContext)

	// EnterMkusrparams is called when entering the mkusrparams production.
	EnterMkusrparams(c *MkusrparamsContext)

	// EnterMkusrparam is called when entering the mkusrparam production.
	EnterMkusrparam(c *MkusrparamContext)

	// EnterRmusr is called when entering the rmusr production.
	EnterRmusr(c *RmusrContext)

	// EnterRmusrparams is called when entering the rmusrparams production.
	EnterRmusrparams(c *RmusrparamsContext)

	// EnterRmusrparam is called when entering the rmusrparam production.
	EnterRmusrparam(c *RmusrparamContext)

	// EnterChgrp is called when entering the chgrp production.
	EnterChgrp(c *ChgrpContext)

	// EnterChgrpparams is called when entering the chgrpparams production.
	EnterChgrpparams(c *ChgrpparamsContext)

	// EnterChgrpparam is called when entering the chgrpparam production.
	EnterChgrpparam(c *ChgrpparamContext)

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

	// ExitMount is called when exiting the mount production.
	ExitMount(c *MountContext)

	// ExitMountparams is called when exiting the mountparams production.
	ExitMountparams(c *MountparamsContext)

	// ExitMountparam is called when exiting the mountparam production.
	ExitMountparam(c *MountparamContext)

	// ExitMounted is called when exiting the mounted production.
	ExitMounted(c *MountedContext)

	// ExitMkfs is called when exiting the mkfs production.
	ExitMkfs(c *MkfsContext)

	// ExitMkfsparams is called when exiting the mkfsparams production.
	ExitMkfsparams(c *MkfsparamsContext)

	// ExitMkfsparam is called when exiting the mkfsparam production.
	ExitMkfsparam(c *MkfsparamContext)

	// ExitCat is called when exiting the cat production.
	ExitCat(c *CatContext)

	// ExitCatparams is called when exiting the catparams production.
	ExitCatparams(c *CatparamsContext)

	// ExitCatparam is called when exiting the catparam production.
	ExitCatparam(c *CatparamContext)

	// ExitLogin is called when exiting the login production.
	ExitLogin(c *LoginContext)

	// ExitLoginparams is called when exiting the loginparams production.
	ExitLoginparams(c *LoginparamsContext)

	// ExitLoginparam is called when exiting the loginparam production.
	ExitLoginparam(c *LoginparamContext)

	// ExitLogout is called when exiting the logout production.
	ExitLogout(c *LogoutContext)

	// ExitMkgrp is called when exiting the mkgrp production.
	ExitMkgrp(c *MkgrpContext)

	// ExitMkgrpparams is called when exiting the mkgrpparams production.
	ExitMkgrpparams(c *MkgrpparamsContext)

	// ExitMkgrpparam is called when exiting the mkgrpparam production.
	ExitMkgrpparam(c *MkgrpparamContext)

	// ExitRmgrp is called when exiting the rmgrp production.
	ExitRmgrp(c *RmgrpContext)

	// ExitRmgrpparams is called when exiting the rmgrpparams production.
	ExitRmgrpparams(c *RmgrpparamsContext)

	// ExitRmgrpparam is called when exiting the rmgrpparam production.
	ExitRmgrpparam(c *RmgrpparamContext)

	// ExitMkusr is called when exiting the mkusr production.
	ExitMkusr(c *MkusrContext)

	// ExitMkusrparams is called when exiting the mkusrparams production.
	ExitMkusrparams(c *MkusrparamsContext)

	// ExitMkusrparam is called when exiting the mkusrparam production.
	ExitMkusrparam(c *MkusrparamContext)

	// ExitRmusr is called when exiting the rmusr production.
	ExitRmusr(c *RmusrContext)

	// ExitRmusrparams is called when exiting the rmusrparams production.
	ExitRmusrparams(c *RmusrparamsContext)

	// ExitRmusrparam is called when exiting the rmusrparam production.
	ExitRmusrparam(c *RmusrparamContext)

	// ExitChgrp is called when exiting the chgrp production.
	ExitChgrp(c *ChgrpContext)

	// ExitChgrpparams is called when exiting the chgrpparams production.
	ExitChgrpparams(c *ChgrpparamsContext)

	// ExitChgrpparam is called when exiting the chgrpparam production.
	ExitChgrpparam(c *ChgrpparamContext)
}
