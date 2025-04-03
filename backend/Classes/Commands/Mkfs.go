// Package Commands provides implementations for various filesystem commands,
// including commands for formatting partitions and creating EXT2 filesystems.
package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Mkfs represents the MKFS command, which is used to format a mounted partition
// with an EXT2 filesystem. It embeds the common CommandStruct for shared metadata
// and includes a map for command parameters.
type Mkfs struct {
	Interfaces.CommandStruct                   // Embedded command structure (provides Type, Line, Column, etc.)
	Params                   map[string]string // Params contains command parameters such as "id" and "type".
}

// NewMkfs creates a new Mkfs command instance with the specified source location and parameters.
func NewMkfs(line, column int, params map[string]string) *Mkfs {
	return &Mkfs{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.MKFS, // Command type constant for MKFS.
			Line:   line,       // Source code line where the command is defined.
			Column: column,     // Source code column where the command is defined.
		},
		Params: params,
	}
}

// Exec executes the MKFS command to format the partition with an EXT2 filesystem.
// It validates parameters, verifies the mounted partition, calculates filesystem parameters,
// creates a new superblock, initializes inodes and blocks, creates root and users files,
// updates bitmaps, and writes the final superblock back to disk.
func (m *Mkfs) Exec() {
	// Validate command parameters.
	if err := m.validateParams(); err != nil {
		m.AppendError(err.Error())
		return
	}

	// Start building a console log for the operation.
	var consoleString string
	consoleString = "=================MKFS=================\n"
	consoleString += fmt.Sprintf("Id: %s, Type: %s, Fs: EXT2\n", m.Params["id"], m.Params["type"])

	// Retrieve the mounted partition using the provided partition id.
	var mountedPartition *Structs.MountedPartition
	if mountedPartition = m.getPartition(); mountedPartition == nil {
		m.AppendError("Partition ID not found")
		return
	}

	// Ensure that the partition is mounted.
	if mountedPartition.Status != [1]byte{'1'} {
		m.AppendError("Partition is not mounted")
		return
	}

	// Verify that the mounted partition exists in the disk's MBR.
	mbrPartition := m.checkMountedPartitionHealth(mountedPartition, consoleString)
	if mbrPartition == nil {
		m.AppendError("Mounted partition does not exist at disc MBR")
		return
	}

	// Calculate the number of inodes (n) based on the partition size minus the superblock size.
	// The denominator is calculated based on overhead (4 bytes) plus sizes of an inode and three file blocks.
	numerator := mbrPartition.Size - int32(binary.Size(Structs.SuperBlock{}))
	baseDenominator := 4 + int32(binary.Size(Structs.Inode{})) + 3*int32(binary.Size(Structs.FileBlock{}))
	var temp int32 = 0
	denominator := baseDenominator + temp
	n := int32(numerator / denominator)
	consoleString += fmt.Sprintf("I-NODES qnt: %d\n", n)

	// Create the EXT2 filesystem structure.
	resultString := m.makeEXT2(n, mbrPartition, mountedPartition, consoleString)
	if resultString == nil {
		return
	}

	consoleString = *resultString
	// Append footer to the console output.
	consoleString += "=================END MKFS=================\n"
	m.LogConsole(consoleString)
}

// makeEXT2 creates an EXT2 filesystem on the given partition.
// It builds a new superblock, calculates starting offsets for the bitmaps,
// inode area, and block area, initializes inodes and file blocks, and creates
// the root folder and users file. It returns a pointer to the updated console log string.
func (m *Mkfs) makeEXT2(n int32, mbrPartition *Structs.Partition, mountedPartition *Structs.MountedPartition, consoleString string) *string {
	// Initialize a new superblock with EXT2 settings.
	var superBlock Structs.SuperBlock
	superBlock.FilesystemType = 2
	superBlock.InodesCount = n
	superBlock.BlocksCount = 3 * n
	superBlock.FreeBlocksCount = 3 * n
	superBlock.FreeInodesCount = n
	copy(superBlock.MTime[:], time.Now().Format("2006-01-02 15:04"))
	copy(superBlock.UMTime[:], time.Now().Format("2006-01-02 15:04"))
	superBlock.MntCount = 1
	superBlock.Magic = 0xEF53
	superBlock.InodeSize = int32(binary.Size(Structs.Inode{}))
	superBlock.BlockSize = int32(binary.Size(Structs.FileBlock{}))

	// Calculate starting offsets for bitmaps, inodes, and file blocks.
	superBlock.BmInodeStart = mbrPartition.Start + int32(binary.Size(Structs.SuperBlock{}))
	superBlock.BmBlockStart = superBlock.BmInodeStart + n
	superBlock.InodeStart = superBlock.BmBlockStart + 3*n
	superBlock.BlockStart = superBlock.InodeStart + n*superBlock.InodeSize

	consoleString += "_________________Creating EXT2_________________\n"
	consoleString += superBlock.ToString() + "\n\n"

	// Open the disk file for writing.
	file, err := Utils.OpenFile(mountedPartition.Path)
	if err != nil {
		m.AppendError(err.Error())
		return nil
	}
	defer file.Close()

	// Initialize the block bitmap: write 0 for each block (total of 3*n blocks).
	for i := int32(0); i < 3*n; i++ {
		if err := Utils.WriteObject(file, byte(0), int64(superBlock.BmInodeStart)); err != nil {
			m.AppendError(err.Error())
			return nil
		}
	}

	// Initialize all inodes and file blocks to default values.
	if err := initInodesAndBlocks(n, superBlock, file); err != nil {
		m.AppendError(err.Error())
		return nil
	}

	// Create the root folder and the users file in the filesystem.
	if err := createRootAndUsersFile(superBlock, file); err != nil {
		m.AppendError(err.Error())
		return nil
	}
	// Update free counts (subtract reserved entries for root and users).
	superBlock.FreeBlocksCount = 3*n - 2
	superBlock.FreeInodesCount = n - 2

	// Write the updated superblock at the start of the partition.
	if err := Utils.WriteObject(file, superBlock, int64(mbrPartition.Start)); err != nil {
		m.AppendError(err.Error())
		return nil
	}

	// Mark reserved inodes and blocks as used.
	if err := markUsedInodesAndNodes(file, superBlock); err != nil {
		m.AppendError(err.Error())
		return nil
	}

	// Print the first few inodes for debugging purposes.
	consoleString += "\tPrinting Inodes:\n"
	for i := int32(0); i < 3; i++ {
		var inode Structs.Inode
		offset := int64(superBlock.InodeStart + i*int32(binary.Size(Structs.Inode{})))
		if err := Utils.ReadObject(file, &inode, offset); err != nil {
			m.AppendError(err.Error())
			return nil
		}
		consoleString += fmt.Sprintf("\t\tInode %d:\n", i)
		consoleString += inode.ToString() + "\n\n"
	}
	consoleString += "\t\t......\n\n"

	// Print the first folder block.
	consoleString += "\tPrinting Folder Block 0:\n"
	var folderBlock Structs.FolderBlock
	if err := Utils.ReadObject(file, &folderBlock, int64(superBlock.BlockStart)); err != nil {
		m.AppendError(err.Error())
		return nil
	}
	consoleString += folderBlock.ToString() + "\n"

	// Print the first file block (located after the folder block).
	consoleString += "\n\tPrinting File Block 0:\n"
	var fileBlock Structs.FileBlock
	offset := int64(superBlock.BlockStart + int32(binary.Size(Structs.FolderBlock{})))
	if err := Utils.ReadObject(file, &fileBlock, offset); err != nil {
		m.AppendError(err.Error())
		return nil
	}
	consoleString += "\t\t" + fileBlock.ToString() + "\n\n"
	consoleString += "Final Super Block:\n"
	consoleString += superBlock.ToString() + "\n"
	consoleString += "_________________End Creating EXT2_________________\n"

	return &consoleString
}

// markUsedInodesAndNodes marks reserved inodes and blocks as used in the filesystem's bitmaps.
// It writes a byte value of 1 for the first two inodes and the first two blocks.
func markUsedInodesAndNodes(file *os.File, superBlock Structs.SuperBlock) error {
	// Mark the first inode (root inode) as used.
	if err := Utils.WriteObject(file, byte(1), int64(superBlock.BmInodeStart)); err != nil {
		return err
	}
	// Mark the second inode (users.txt inode) as used.
	if err := Utils.WriteObject(file, byte(1), int64(superBlock.BmInodeStart+1)); err != nil {
		return err
	}
	// Mark the first block (root folder block) as used.
	if err := Utils.WriteObject(file, byte(1), int64(superBlock.BmBlockStart)); err != nil {
		return err
	}
	// Mark the second block (users.txt file block) as used.
	if err := Utils.WriteObject(file, byte(1), int64(superBlock.BmBlockStart+1)); err != nil {
		return err
	}
	return nil
}

// createRootAndUsersFile creates the root directory and the users file ("/users.txt").
// It creates corresponding inodes, a folder block for the root directory, and a file block for users data.
func createRootAndUsersFile(superBlock Structs.SuperBlock, file *os.File) error {
	// Create the root inode (type folder) and the users inode (type file).
	rootInode := Structs.NewInode([1]byte{0})
	usersInode := Structs.NewInode([1]byte{1})

	// Set the first block pointers:
	// Root inode points to folder block 0, and users inode points to file block 1.
	rootInode.IBlock[0] = 0
	usersInode.IBlock[0] = 1

	// Initialize the users file content.
	data := "1,G,root\n1,U,root,root,123\n"
	actualSize := int32(len(data))
	usersInode.Size = actualSize

	// Create a file block for the users file and copy the data.
	var usersFileBlock Structs.FileBlock
	copy(usersFileBlock.Content[:], data[:])

	// Create a folder block for the root directory ("/").
	var rootFolderBlock Structs.FolderBlock
	// Entry for current directory "."
	rootFolderBlock.Content[0].Inode = 0
	copy(rootFolderBlock.Content[0].Name[:], ".")
	// Entry for parent directory ".." (same as root in this case)
	rootFolderBlock.Content[1].Inode = 0
	copy(rootFolderBlock.Content[1].Name[:], "..")
	// Entry for the users file "users.txt"
	rootFolderBlock.Content[2].Inode = 1
	copy(rootFolderBlock.Content[2].Name[:], "users.txt")
	// Mark remaining entry as unused.
	rootFolderBlock.Content[3].Inode = -1

	// Write the root inode at the beginning of the inode area.
	if err := Utils.WriteObject(file, rootInode, int64(superBlock.InodeStart)); err != nil {
		return err
	}
	// Write the users inode right after the root inode.
	if err := Utils.WriteObject(file, usersInode, int64(superBlock.InodeStart+int32(binary.Size(Structs.Inode{})))); err != nil {
		return err
	}
	// Write the root folder block at the beginning of the block area.
	if err := Utils.WriteObject(file, rootFolderBlock, int64(superBlock.BlockStart)); err != nil {
		return err
	}
	// Write the users file block next to the folder block.
	if err := Utils.WriteObject(file, usersFileBlock, int64(superBlock.BlockStart+int32(binary.Size(Structs.FolderBlock{})))); err != nil {
		return err
	}

	return nil
}

// initInodesAndBlocks initializes all inodes and file blocks on the filesystem.
// It writes a default inode (with all IBlock pointers set to -1) for each inode slot,
// and writes a default (empty) file block for each of the 3*n file blocks.
func initInodesAndBlocks(n int32, superBlock Structs.SuperBlock, file *os.File) error {
	// Prepare a default inode with IBlock pointers initialized to -1.
	var newInode Structs.Inode
	for i := 0; i < 15; i++ {
		newInode.IBlock[i] = -1
	}

	// Write n inodes to the inode area.
	for i := int32(0); i < n; i++ {
		if err := Utils.WriteObject(file, newInode, int64(superBlock.InodeStart+i*int32(binary.Size(Structs.Inode{})))); err != nil {
			return err
		}
	}

	// Prepare a default file block.
	var newFileBlock Structs.FileBlock
	// Write 3*n file blocks to the block area.
	for i := int32(0); i < 3*n; i++ {
		if err := Utils.WriteObject(file, newFileBlock, int64(superBlock.BlockStart+i*int32(binary.Size(Structs.FileBlock{})))); err != nil {
			return err
		}
	}
	return nil
}

// checkMountedPartitionHealth verifies that the mounted partition exists on disk by reading the MBR.
// It opens the disk file for the mounted partition, reads the MBR, logs the MBR info to the console string,
// and searches for the partition that matches the mounted partition's ID.
func (m *Mkfs) checkMountedPartitionHealth(mountedPartition *Structs.MountedPartition, consoleString string) *Structs.Partition {
	file, err := Utils.OpenFile(mountedPartition.Path)
	if err != nil {
		m.AppendError(err.Error())
		return nil
	}
	defer file.Close()
	var discMBR Structs.MBR
	if err := Utils.ReadObject(file, &discMBR, 0); err != nil {
		m.AppendError(err.Error())
		return nil
	}
	consoleString += "Founded MBR:\n" + discMBR.ToString() + "\n"
	// Search through the MBR partitions for the one that matches the mounted partition's ID.
	for _, partition := range discMBR.Partitions {
		if partition.Id == mountedPartition.Id {
			return &partition
		}
	}

	return nil
}

// validateParams ensures that the required parameters for the MKFS command are provided.
// It checks that the "id" parameter exists and sets a default "type" if not provided.
// It also converts the "id" parameter to uppercase.
func (m *Mkfs) validateParams() error {
	if _, ok := m.Params["id"]; !ok {
		return errors.New("missing parameter -id")
	}
	if _, ok := m.Params["type"]; !ok {
		m.Params["type"] = "full"
	}
	m.Params["id"] = strings.ToUpper(m.Params["id"])
	return nil
}

// getPartition searches for a mounted partition that matches the provided "id" parameter.
// It iterates over the global list of mounted partitions and returns the matching partition.
func (m *Mkfs) getPartition() *Structs.MountedPartition {
	for _, part := range env.GetPartitions() {
		if string(part.Id[:]) == m.Params["id"] {
			return part
		}
	}
	return nil
}
