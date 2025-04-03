// Package Env provides utilities for handling disk and partition status,
// including verifying disk health and retrieving mounted partitions.
package Env

import (
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"errors"
	"os"
)

// VerifyDiscStatus checks the status of a disk partition given its ID.
// It attempts to retrieve the mounted partition from RAM using the provided partition ID,
// then validates the health of the partition by reading the disk's MBR.
// It returns pointers to the mounted partition, its corresponding MBR partition entry,
// the disk file, and an error if any step fails.
func VerifyDiscStatus(partition [4]byte) (*Structs.MountedPartition, *Structs.Partition, *os.File, error) {
	// Retrieve the mounted partition from memory using the provided partition ID.
	var mountedPartition *Structs.MountedPartition
	// If the partition is not found in RAM, return an error.
	if mountedPartition = getPartitionInRam(partition); mountedPartition == nil {
		return nil, nil, nil, errors.New("No partition was founded given id")
	}

	var mbrPartition *Structs.Partition
	var discFile *os.File
	var err error

	// Check the health of the mounted partition by verifying its MBR entry.
	mbrPartition, discFile, err = checkMountedPartitionHealth(mountedPartition)
	if err != nil {
		return nil, nil, nil, err
	}
	return mountedPartition, mbrPartition, discFile, nil
}

// checkMountedPartitionHealth validates the health of a mounted partition.
// It opens the disk file associated with the mounted partition, reads the MBR,
// and searches for the partition's entry in the MBR. If found, it returns the partition entry and file.
// Otherwise, it returns an error.
func checkMountedPartitionHealth(mountedPartition *Structs.MountedPartition) (*Structs.Partition, *os.File, error) {
	// Open the disk file for the mounted partition.
	file, err := Utils.OpenFile(mountedPartition.Path)
	if err != nil {
		return nil, nil, err
	}

	// Read the MBR from the beginning of the disk.
	var discMBR Structs.MBR
	if err := Utils.ReadObject(file, &discMBR, 0); err != nil {
		return nil, nil, err
	}

	// Search for the mounted partition within the MBR's partition entries.
	for _, partition := range discMBR.Partitions {
		// If the partition ID matches the mounted partition's ID, return it.
		if partition.Id == mountedPartition.Id {
			return &partition, file, nil
		}
	}

	// Return an error if the mounted partition is not found in the MBR.
	return nil, nil, errors.New("Mounted partition does not exist at disc MBR")
}

// getPartitionInRam searches for a mounted partition in memory (RAM) by comparing partition IDs.
// It iterates over the global list of mounted partitions and returns the matching partition if found,
// or nil if no match is found.
func getPartitionInRam(partitionId [4]byte) *Structs.MountedPartition {
	// Iterate over the list of mounted partitions.
	for _, part := range GetPartitions() {
		// Compare partition IDs by converting the byte arrays to strings.
		if string(part.Id[:]) == string(partitionId[:]) {
			return part
		}
	}
	// Return nil if no matching partition is found.
	return nil
}
