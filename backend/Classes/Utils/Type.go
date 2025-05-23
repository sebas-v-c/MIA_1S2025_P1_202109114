package Utils

type Type int

const (
	MKDISK Type = iota
	RMDISK
	FDISK
	MOUNT
	MOUNTED
	MKFS
	CAT
	LOGIN
	LOGOUT
	MKGROUP
	RMGROUP
	MKUSR
	RMUSR
	CHRGP
	MKFILE
	MKDIR
	REP
)
