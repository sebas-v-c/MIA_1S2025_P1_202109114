package Structs

import "fmt"

type Inode struct {
	UID   int32
	GID   int32
	Size  int32
	ATime [17]byte
	CTime [17]byte
	MTime [17]byte
	Block [15]int32
	Type  [1]byte
	Perm  [3]byte
}

func (i *Inode) ToString() string {
	return fmt.Sprintf(`i_UID %v
i_GID %v
i_Size %v
i_Atime %v
i_CTime %v
i_MTime %v
i_Block %v
i_Type %v
i_Perm %v`,
		i.UID,
		i.GID,
		i.Size,
		i.ATime,
		i.CTime,
		i.MTime,
		i.Block,
		i.Type,
		i.Perm)
}
