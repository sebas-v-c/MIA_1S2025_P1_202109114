package Structs

import "fmt"

type SuperBlock struct {
	FilesystemType  int32
	InodesCount     int32
	BlocksCount     int32
	FreeBlocksCount int32
	FreeInodesCount int32
	MTime           [17]byte
	UMTime          [17]byte
	MntCount        int32
	Magic           int32
	InodeSize       int32
	BlockSize       int32
	FirstInode      int32
	FirstBlock      int32
	BmInodeStart    int32
	BmBlockStart    int32
	InodeStart      int32
	BlockStart      int32
}

func (sb *SuperBlock) ToString() string {
	return fmt.Sprintf(`filesystem_type %v
inodes_count %v
blocks_count %v
free_inodes_count %v
free_blocks_count %v
mtime %v
umtime %v
mnt_count %v
magic %v
inode_size %v
block_size %v
first_ino %v
first_blo %v
bm_inode_start %v
bm_block_start %v
inode_start %v
block_start %v`,
		sb.FilesystemType,
		sb.InodesCount,
		sb.BlocksCount,
		sb.FreeInodesCount,
		sb.FreeBlocksCount,
		sb.MTime,
		sb.UMTime,
		sb.MntCount,
		sb.Magic,
		sb.InodeSize,
		sb.BlockSize,
		sb.FirstInode,
		sb.FirstBlock,
		sb.BmInodeStart,
		sb.BmBlockStart,
		sb.InodeStart,
		sb.BlockStart,
	)
}
