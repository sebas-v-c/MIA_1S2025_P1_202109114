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

type Mkfs struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewMkfs(line, column int, params map[string]string) *Mkfs {
	return &Mkfs{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.MKFS,
			Line:   line,
			Column: column,
		},
		Params: params,
	}
}

func (m *Mkfs) Exec() {
	if err := m.validateParams(); err != nil {
		m.AppendError(err.Error())
		return
	}

	var consoleString string
	consoleString = "=================MKFS=================\n"
	consoleString += fmt.Sprintf("Id: %s, Type: %s, Fs: EXT2\n", m.Params["id"], m.Params["type"])

	// Check if the given id belong to a partition that is mounted
	var mountedPartition *Structs.MountedPartition
	if mountedPartition = m.getPartition(); mountedPartition == nil {
		m.AppendError("Partition ID not found")
		return
	}

	// Check if mounted partition has the status of mounted
	if mountedPartition.Status != [1]byte{'1'} {
		m.AppendError("Partition is not mounted")
		return
	}

	// Checking if the mounted partition is the same as the partition in disc
	mbrPartition := m.checkMountedPartitionHealth(mountedPartition, consoleString)
	if mbrPartition == nil {
		m.AppendError("Partition does not exist in MBR")
		return
	}

	// Calculating the size of the inodes
	numerator := mbrPartition.Size - int32(binary.Size(Structs.SuperBlock{}))
	baseDenominator := 4 + int32(binary.Size(Structs.Inode{})) + 3*int32(binary.Size(Structs.FileBlock{}))
	var temp int32 = 0
	denominator := baseDenominator + temp
	n := int32(numerator / denominator)
	consoleString += fmt.Sprintf("I-NODES qnt: %d\n", n)

	resultString := m.makeEXT2(n, mbrPartition, mountedPartition, consoleString)
	if resultString == nil {
		return
	}

	consoleString = *resultString
	// Only FS2 filesystem right now
	consoleString += "=================END MKFS=================\n"
	m.LogConsole(consoleString)
}

func (m *Mkfs) makeEXT2(n int32, mbrPartition *Structs.Partition, mountedPartition *Structs.MountedPartition, consoleString string) *string {
	// Creating super block
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

	// Calculate start position
	superBlock.BmInodeStart = mbrPartition.Start + int32(binary.Size(Structs.SuperBlock{}))
	superBlock.BmBlockStart = superBlock.BmInodeStart + n
	superBlock.InodeStart = superBlock.BmBlockStart + 3*n
	superBlock.BlockStart = superBlock.InodeStart + n*superBlock.InodeSize

	consoleString += "_________________Creating EXT2_________________\n"
	consoleString += superBlock.ToString() + "\n\n"

	file, err := Utils.OpenFile(mountedPartition.Path)
	if err != nil {
		m.AppendError(err.Error())
		return nil
	}
	defer file.Close()

	for i := int32(0); i < 3*n; i++ {
		if err := Utils.WriteObject(file, byte(0), int64(superBlock.BmInodeStart)); err != nil {
			m.AppendError(err.Error())
			return nil
		}
	}

	// initialize inodes and blocks to file
	if err := initInodesAndBlocks(n, superBlock, file); err != nil {
		m.AppendError(err.Error())
		return nil
	}

	// Create the root folder and the users file
	if err := createRootAndUsersFile(superBlock, file); err != nil {
		m.AppendError(err.Error())
		return nil
	}
	superBlock.FreeBlocksCount = 3*n - 2
	superBlock.FreeInodesCount = n - 2

	// Write super block at the start of the partition
	if err := Utils.WriteObject(file, superBlock, int64(mbrPartition.Start)); err != nil {
		m.AppendError(err.Error())
		return nil
	}

	// Mark the first inodes and blocks as used
	if err := markUsedInodesAndNodes(file, superBlock); err != nil {
		m.AppendError(err.Error())
		return nil
	}

	// printing the list of inodes in partition
	// Just printing the first 3 since the first 2 are the users and root folder and the third one is the same for the rest
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

	// Printing the only folder block available at the moment
	consoleString += "\tPrinting Folder Block 0:\n"
	var folderBlock Structs.FolderBlock
	if err := Utils.ReadObject(file, &folderBlock, int64(superBlock.BlockStart)); err != nil {
		m.AppendError(err.Error())
		return nil
	}
	consoleString += folderBlock.ToString() + "\n"

	// printing the only file block at the moment, next to the first folder block
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

func markUsedInodesAndNodes(file *os.File, superBlock Structs.SuperBlock) error {
	// Marking the first inode (root inode) as used in the first item of the bitmap
	if err := Utils.WriteObject(file, byte(1), int64(superBlock.BmInodeStart)); err != nil {
		return err
	}
	// Marking the second inode (users.txt inode) as used in the second item of the bitmap
	if err := Utils.WriteObject(file, byte(1), int64(superBlock.BmInodeStart+1)); err != nil {
		return err
	}
	// Marking the first block (root folder block) as used in the first item of the bitmap
	if err := Utils.WriteObject(file, byte(1), int64(superBlock.BmBlockStart)); err != nil {
		return err
	}
	// Marking the second block (users.txt file block) as used in the second item of the bitmap
	if err := Utils.WriteObject(file, byte(1), int64(superBlock.BmBlockStart+1)); err != nil {
		return err
	}
	return nil
}

func createRootAndUsersFile(superBlock Structs.SuperBlock, file *os.File) error {
	rootInode := Structs.NewInode([1]byte{'0'})
	usersInode := Structs.NewInode([1]byte{'1'})

	// this is set to 0 because we are using the first folder Block available
	rootInode.IBlock[0] = '0'
	// This is set to 1 because here we are pointing to the next file block
	usersInode.IBlock[0] = '1'

	// Assign the size of the content of the file into the users Inode
	data := "1,G,root\n1,U,root,root,123\n"
	actualSize := int32(len(data))
	usersInode.Size = actualSize

	// File block for storing the data of the users
	var usersFileBlock Structs.FileBlock
	copy(usersFileBlock.Content[:], data[:])

	// Creating the root folder block for "/"
	var rootFolderBlock Structs.FolderBlock
	// The current folder is "." referring to "/"
	rootFolderBlock.Content[0].Inode = 0
	copy(rootFolderBlock.Content[0].Name[:], ".")
	// The parent folder is "." too, so still referring to "/"
	rootFolderBlock.Content[1].Inode = 0
	copy(rootFolderBlock.Content[1].Name[:], "..")
	// The next available space is for the file users.txt
	rootFolderBlock.Content[2].Inode = 1
	copy(rootFolderBlock.Content[2].Name[:], "users.txt")

	// Writing the root inode at the start of the Inode part
	if err := Utils.WriteObject(file, rootInode, int64(superBlock.InodeStart)); err != nil {
		return err
	}
	// Writing the users inode next to the first inode
	if err := Utils.WriteObject(file, usersInode, int64(superBlock.InodeStart+int32(binary.Size(Structs.Inode{})))); err != nil {
		return err
	}
	// Writing the root folder block
	if err := Utils.WriteObject(file, rootFolderBlock, int64(superBlock.BlockStart)); err != nil {
		return err
	}
	// Writing the users file block next to the root folder block
	if err := Utils.WriteObject(file, usersFileBlock, int64(superBlock.BlockStart+int32(binary.Size(Structs.FolderBlock{})))); err != nil {
		return err
	}

	return nil
}

func initInodesAndBlocks(n int32, superBlock Structs.SuperBlock, file *os.File) error {
	var newInode Structs.Inode
	for i := int32(0); i < 15; i++ {
		newInode.IBlock[i] = -1
	}

	for i := int32(0); i < n; i++ {
		if err := Utils.WriteObject(file, newInode, int64(superBlock.InodeStart+i*int32(binary.Size(Structs.Inode{})))); err != nil {
			return err
		}
	}

	var newFileBlock Structs.FileBlock
	for i := int32(0); i < 3*n; i++ {
		if err := Utils.WriteObject(file, newFileBlock, int64(superBlock.BlockStart+i*int32(binary.Size(Structs.FileBlock{})))); err != nil {
			return err
		}
	}
	return nil
}

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
	for _, partition := range discMBR.Partitions {
		//fmt.Println(partition.ToString())
		if partition.Id == mountedPartition.Id {
			return &partition
		}
	}

	return nil
}

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

func (m *Mkfs) getPartition() *Structs.MountedPartition {
	for _, part := range env.GetPartitions() {
		if string(part.Id[:]) == m.Params["id"] {
			return part
		}
	}
	return nil
}
