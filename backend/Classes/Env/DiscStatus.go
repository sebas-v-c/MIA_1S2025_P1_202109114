package Env

import (
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"errors"
	"os"
)

func VerifyDiscStatus(partition [4]byte) (*Structs.MountedPartition, *Structs.Partition, *os.File, error) {
	// Check state of the disk
	var mountedPartition *Structs.MountedPartition
	//if mountedPartition = getPartitionInRam(loggedUser.MountedPartition.Id); mountedPartition == nil {
	if mountedPartition = getPartitionInRam(partition); mountedPartition == nil {
		return nil, nil, nil, errors.New("No partition was founded given id")
	}

	var mbrPartition *Structs.Partition
	var discFile *os.File
	var err error
	mbrPartition, discFile, err = checkMountedPartitionHealth(mountedPartition)
	if err != nil {
		return nil, nil, nil, err
	}
	return mountedPartition, mbrPartition, discFile, nil
}

func checkMountedPartitionHealth(mountedPartition *Structs.MountedPartition) (*Structs.Partition, *os.File, error) {
	file, err := Utils.OpenFile(mountedPartition.Path)
	if err != nil {
		return nil, nil, err
	}
	var discMBR Structs.MBR
	if err := Utils.ReadObject(file, &discMBR, 0); err != nil {
		return nil, nil, err
	}
	for _, partition := range discMBR.Partitions {
		//fmt.Println(partition.ToString())
		if partition.Id == mountedPartition.Id {
			return &partition, file, nil
		}
	}

	return nil, nil, errors.New("Mounted partition does not exist at disc MBR")
}

func getPartitionInRam(partitionId [4]byte) *Structs.MountedPartition {
	for _, part := range GetPartitions() {
		if string(part.Id[:]) == string(partitionId[:]) {
			return part
		}
	}
	return nil
}
