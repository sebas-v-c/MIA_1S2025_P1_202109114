package Structs

import (
	"fmt"
	"time"
)

type Inode struct {
	UID    int32
	GID    int32
	Size   int32
	ATime  [17]byte
	CTime  [17]byte
	MTime  [17]byte
	IBlock [15]int32
	Type   [1]byte
	Perm   [3]byte
}

// NewInode creates a new inode with some default parameters, inodeType must be '0' = for folder and '1' = for file
func NewInode(inodeType [1]byte) *Inode {
	var date [17]byte
	var perm [3]byte
	copy(date[:], time.Now().Format("2006-01-02 15:04"))
	copy(perm[:], "644")
	var blocks [15]int32
	for i := 0; i < 15; i++ {
		blocks[i] = -1
	}

	return &Inode{
		UID:    1,
		GID:    1,
		Size:   0,
		ATime:  date,
		CTime:  date,
		MTime:  date,
		IBlock: blocks,
		Type:   inodeType,
		Perm:   perm,
	}
}

func (i *Inode) ToString() string {
	return fmt.Sprintf(`        i_UID %v
        i_GID %v
        i_Size %v
        i_Atime %s
        i_CTime %s
        i_MTime %s
        i_Block %v
        i_Type %v
        i_Perm %v`,
		i.UID,
		i.GID,
		i.Size,
		string(i.ATime[:]),
		string(i.CTime[:]),
		string(i.MTime[:]),
		i.IBlock,
		string(i.Type[:]),
		string(i.Perm[:]))
}
