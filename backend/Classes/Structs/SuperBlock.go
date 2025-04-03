// Package Structs provides data structures for filesystem metadata.
package Structs

import "fmt"

// SuperBlock holds the metadata for a filesystem's superblock.
// It contains various fields describing the filesystem's properties such as
// inode and block counts, timestamps, and layout information.
type SuperBlock struct {
	FilesystemType  int32    // FilesystemType indicates the type of the filesystem.
	InodesCount     int32    // InodesCount is the total number of inodes in the filesystem.
	BlocksCount     int32    // BlocksCount is the total number of blocks in the filesystem.
	FreeBlocksCount int32    // FreeBlocksCount is the number of free blocks available.
	FreeInodesCount int32    // FreeInodesCount is the number of free inodes available.
	MTime           [17]byte // MTime represents the last modification time (e.g., "2025-09-20 12:00:00").
	UMTime          [17]byte // UMTime represents the last unmount time.
	MntCount        int32    // MntCount is the count of how many times the filesystem has been mounted.
	Magic           int32    // Magic holds a magic number for identifying or validating the filesystem.
	InodeSize       int32    // InodeSize specifies the size of each inode in bytes.
	BlockSize       int32    // BlockSize specifies the size of each block in bytes.
	FirstInode      int32    // FirstInode is the index of the first inode.
	FirstBlock      int32    // FirstBlock is the index of the first block.
	BmInodeStart    int32    // BmInodeStart is the starting position of the inode bitmap.
	BmBlockStart    int32    // BmBlockStart is the starting position of the block bitmap.
	InodeStart      int32    // InodeStart is the starting position of the inodes.
	BlockStart      int32    // BlockStart is the starting position of the blocks.
}

// ToString returns a formatted multi-line string representation of the SuperBlock.
// It formats all the fields of the superblock into a readable string.
func (sb *SuperBlock) ToString() string {
	return fmt.Sprintf(`Super Block:
filesystem_type %v
inodes_count %v
blocks_count %v
free_inodes_count %v
free_blocks_count %v
mtime %s
umtime %s
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
		string(sb.MTime[:]),
		string(sb.UMTime[:]),
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
