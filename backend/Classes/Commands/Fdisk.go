package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"encoding/binary"
	"errors"
	"path/filepath"
	"strconv"
	"strings"
)

type Fdisk struct {
	Result string
	Type   Utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewFdisk(line, column int, params map[string]string) *Fdisk {
	return &Fdisk{
		Type:   Utils.FDISK,
		Params: params,
		Line:   line,
		Column: column,
	}
}

func (f *Fdisk) GetLine() int {
	return f.Line
}

func (f *Fdisk) GetColumn() int {
	return f.Column
}

func (f *Fdisk) GetType() Utils.Type {
	return f.Type
}

func (f *Fdisk) Exec() {
	if err, ok := f.validParams(); !ok {
		env.Errors = append(env.Errors, env.RuntimeError{
			Line:    f.Line,
			Column:  f.Column,
			Command: Utils.FDISK,
			Msg:     err.Error(),
		})
		return
	}

	// open file
	file, err := Utils.OpenFile(f.Params["path"])
	if err != nil {
		env.Errors = append(env.Errors, env.RuntimeError{
			Line:    f.Line,
			Column:  f.Column,
			Command: Utils.FDISK,
			Msg:     err.Error(),
		})
		return
	}

	// Read MBR from binary disk
	var diskMBR Structs.MBR
	if err := Utils.ReadObject(file, &diskMBR, 0); err != nil {
		env.Errors = append(env.Errors, env.RuntimeError{
			Line:    f.Line,
			Column:  f.Column,
			Command: Utils.FDISK,
			Msg:     err.Error(),
		})
		return
	}

	// Validate Partitions in current MBR
	// recalculate size of the partition
	//units := f.recalculateUnits()
	size, _ := strconv.Atoi(f.Params["size"])
	totalSize := f.recalculateUnits() * size
	err, totalPartitions := f.validatePartitions(diskMBR, totalSize)
	if err != nil {
		env.Errors = append(env.Errors, env.RuntimeError{
			Line:    f.Line,
			Column:  f.Column,
			Command: Utils.FDISK,
			Msg:     err.Error(),
		})
		return
	}

	// calculating the next empty space
	var gap = int32(binary.Size(diskMBR)) // binary size of the MBR struct
	// if there are at least 1 partition loaded calculate the next available space
	if totalPartitions > 0 {
		// adding the last partition start + the last partition size
		gap = diskMBR.Partitions[totalPartitions-1].Start + diskMBR.Partitions[totalPartitions-1].Size
	}

	// Send to console
	var consoleString string
	consoleString = "=================FDISK=================\n"

	// Create a primary or extended partition and store it in the MBR
	for i := range diskMBR.Partitions {
		// If partition size == 0 means the partition is empty, the check if the partition to be created is Primary or Extended
		if diskMBR.Partitions[i].Size == 0 && (f.getPartType() == 'P' || f.getPartType() == 'E') {
			// create partition in MBR
			diskMBR.Partitions[i] = Structs.Partition{
				Start:       gap, // the start is placed at the next available space
				Size:        int32(totalSize),
				Correlative: int32(totalPartitions + 1),
				Status:      [1]byte{'0'},
				Fit:         f.getFit(),
			}
			copy(diskMBR.Partitions[i].Name[:], f.Params["name"])
			copy(diskMBR.Partitions[i].Type[:], string(f.getPartType()))
			// If partition is an extended partition created his first EBR
			if f.getPartType() == 'E' {
				partitionEBR := Structs.EBR{
					Fit:   f.getFit()[0], // get fit from parameters
					Start: gap,           // first EBR is placed at the start of the extended partition
					Size:  0,
					Next:  -1,
				}
				copy(partitionEBR.Name[:], "")
				// Write object to disk
				if err := Utils.WriteObject(file, partitionEBR, int64(gap)); err != nil {
					env.Errors = append(env.Errors, env.RuntimeError{
						Line:    f.Line,
						Column:  f.Column,
						Command: Utils.FDISK,
						Msg:     err.Error(),
					})
					return
				}
				consoleString += "EBR Created:\n" + "\t" + partitionEBR.ToString() + "\n"
			}
			break
		}
	}

	// if the partition is of logical type then
	if f.getPartType() == 'L' {
		// iterate over the array of partitions in the MBR
		for i := range diskMBR.Partitions {
			// Since there could be only 1 extended partition then the logical partition should be created here
			if diskMBR.Partitions[i].Type == [1]byte{'E'} {
				// The Ebr starting position is equal to this extended partition start
				// Since the EBR is set at the beginning of an extended partition
				ebrPosition := diskMBR.Partitions[i].Start
				var ebr Structs.EBR

				// Walk the ebr list until finding the last one
				for {
					Utils.ReadObject(file, &ebr, int64(ebrPosition))
					if ebr.Next == -1 { // We know this is the last EBR since the next parameter is set to -1
						break
					}
					ebrPosition = ebr.Next // if is not the last partition get the position of the next EBR
				}

				// The position of the next EBR is set to the start of the last EBR + it's size
				// Placing this new EBR at the end of the las logical partition
				newEBRPosition := ebr.Start + ebr.Size
				// The start of the new logical partition is set to the newEBR position + the size of the EBR struct
				logicalPartitionStart := newEBRPosition + int32(binary.Size(ebr))
				// Update the last EBR to set it's Next atribute to the new calculated EBR position
				ebr.Next = newEBRPosition
				// Writing the old EBR into the file at the EBRPosition
				if err := Utils.WriteObject(file, ebr, int64(ebrPosition)); err != nil {
					env.Errors = append(env.Errors, env.RuntimeError{
						Line:    f.Line,
						Column:  f.Column,
						Command: Utils.FDISK,
						Msg:     err.Error(),
					})
					return
				}

				// Create and write the new EBR with the Next parameter to -1
				// indicating that this is the last EBR at the partition
				newEBR := Structs.EBR{
					Fit:   f.getFit()[0],
					Start: logicalPartitionStart,
					Size:  int32(totalSize),
					Next:  -1,
				}
				copy(newEBR.Name[:], f.Params["name"])
				// Writing the new EBR at the end of the last Partition
				if err := Utils.WriteObject(file, newEBR, int64(newEBRPosition)); err != nil {
					env.Errors = append(env.Errors, env.RuntimeError{
						Line:    f.Line,
						Column:  f.Column,
						Command: Utils.FDISK,
						Msg:     err.Error(),
					})
					return
				}

				// Print Written EBR
				consoleString += ":New EBR Created\n"
				consoleString += "\t" + newEBR.ToString() + "\n"
				// Print all EBRs from the current extended partition
				consoleString += "Printing All EBR from partition: \n"
				ebrPos := diskMBR.Partitions[i].Start
				for {
					if err := Utils.ReadObject(file, &ebr, int64(ebrPos)); err != nil {
						env.Errors = append(env.Errors, env.RuntimeError{
							Line:    f.Line,
							Column:  f.Column,
							Command: Utils.FDISK,
							Msg:     err.Error(),
						})
						return
					}
					consoleString += "\t" + ebr.ToString() + "\n"
					if ebr.Next == -1 {
						break
					}
					ebrPos = ebr.Next
				}
				break
			}
		}
	}

	if err := Utils.WriteObject(file, diskMBR, 0); err != nil {
		env.Errors = append(env.Errors, env.RuntimeError{
			Line:    f.Line,
			Column:  f.Column,
			Command: Utils.FDISK,
			Msg:     err.Error(),
		})
		return
	}

	var updatedMBR Structs.MBR
	if err := Utils.ReadObject(file, &updatedMBR, 0); err != nil {
		env.Errors = append(env.Errors, env.RuntimeError{
			Line:    f.Line,
			Column:  f.Column,
			Command: Utils.FDISK,
			Msg:     err.Error(),
		})
		return
	}

	// Send consoles
	consoleString += "Updated MBR:\n" + updatedMBR.ToString() + "\n=================END FDISK================="
	env.CommandLog = append(env.CommandLog, consoleString)
	defer file.Close()
}

func (f *Fdisk) validParams() (error, bool) {
	// obligatory parameters
	if _, ok := f.Params["size"]; !ok {
		return errors.New("missing parameter -size"), false
	} else if _, ok := f.Params["path"]; !ok {
		return errors.New("missing parameter -path"), false
	} else if _, ok := f.Params["name"]; !ok {
		return errors.New("missing parameter -name"), false
	}

	// optional parameters
	if _, ok := f.Params["type"]; !ok {
		f.Params["type"] = "P"
	}
	if _, ok := f.Params["fit"]; !ok {
		f.Params["fit"] = "WF"
	}
	if _, ok := f.Params["unit"]; !ok {
		f.Params["unit"] = "K"
	}

	// check if size is greater than 0
	if val, _ := strconv.Atoi(f.Params["size"]); val <= 0 {
		return errors.New("size must be greater than 0"), false
	}

	// check if file is a disk
	if strings.EqualFold(filepath.Ext(f.Params["path"]), ".mia") {
		return nil, true
	}
	f.Params["path"] = f.Params["path"] + ".mia"
	return nil, true
}

func (f *Fdisk) GetResult() string {
	//TODO implement me
	panic("implement me")
}

func (f *Fdisk) recalculateUnits() int {
	// if size is b multiply value by 1
	// if size is k multiply value by 1024
	// if the unit is m multiply by 1024*1024
	f.Params["unit"] = strings.ToUpper(f.Params["unit"])
	if f.Params["unit"] == "B" {
		return 1
	} else if f.Params["unit"] == "K" {
		return 1024
	}
	return 1024 * 1024
}

func (f *Fdisk) validatePartitions(mbr Structs.MBR, size int) (error, int) {
	var primaryCount, extendedCount, totalPartitions int
	var usedSpace int32
	for i := 0; i < 4; i++ {
		if mbr.Partitions[i].Size != 0 {
			totalPartitions++
			usedSpace += mbr.Partitions[i].Size

			if mbr.Partitions[i].Type == [1]byte{'P'} {
				primaryCount++
			} else if mbr.Partitions[i].Type == [1]byte{'E'} {
				extendedCount++
			}
		}
	}

	if totalPartitions >= 4 {
		return errors.New("cannot create more than 4 primary or extended partitions"), totalPartitions
	}
	if f.getPartType() == 'E' && extendedCount >= 1 {
		return errors.New("only one extended partition is supported"), totalPartitions
	}
	if f.getPartType() == 'L' && extendedCount == 0 {
		return errors.New("cannot create a logical partition without an extended partition"), totalPartitions
	}
	if usedSpace+int32(size) > mbr.MbrSize {
		return errors.New("not enough disk space to create this partition"), totalPartitions
	}
	return nil, totalPartitions
}

func (f *Fdisk) getPartType() rune {
	return []rune(strings.ToUpper(f.Params["type"]))[0]
}

func (f *Fdisk) getFit() [1]byte {
	f.Params["fit"] = strings.ToUpper(f.Params["fit"])
	if f.Params["fit"] == "FF" {
		return [1]byte{'F'}
	} else if f.Params["fit"] == "BF" {
		return [1]byte{'B'}
	}
	return [1]byte{'W'}
}
