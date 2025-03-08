package Commands

import (
	"backend/Classes/Env"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"encoding/binary"
	"errors"
	"fmt"
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

func (f *Fdisk) Exec(env *Env.Env) {
	if err, ok := f.validParams(); !ok {
		env.Errors = append(env.Errors, Env.RuntimeError{
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
		env.Errors = append(env.Errors, Env.RuntimeError{
			Line:    f.Line,
			Column:  f.Column,
			Command: Utils.FDISK,
			Msg:     err.Error(),
		})
		return
	}

	// Read MBR from binary disk
	var tempMBR Structs.MBR
	if err := Utils.ReadObject(file, &tempMBR, 0); err != nil {
		env.Errors = append(env.Errors, Env.RuntimeError{
			Line:    f.Line,
			Column:  f.Column,
			Command: Utils.FDISK,
			Msg:     err.Error(),
		})
		return
	}
	env.CommandLog = append(env.CommandLog, "----------------FDISK-------------------\n"+tempMBR.ToString()+"\n")
	fmt.Println("----------------FDISK-------------------\n" + tempMBR.ToString() + "\n")

	// Validate Partitions in current MBR
	// recalculate size of the partition
	//units := f.recalculateUnits()
	size, _ := strconv.Atoi(f.Params["size"])
	totalSize := f.recalculateUnits() * size
	//err, primaryCount, extendedCount, totalPartitions := f.validatePartitions(tempMBR, totalSize)
	err, _, _, totalPartitions := f.validatePartitions(tempMBR, totalSize)
	if err != nil {
		env.Errors = append(env.Errors, Env.RuntimeError{
			Line:    f.Line,
			Column:  f.Column,
			Command: Utils.FDISK,
			Msg:     err.Error(),
		})
		return
	}

	//var gap int32 = int32(binary.Size(tempMBR))
	var _ int32 = int32(binary.Size(tempMBR))
	if totalPartitions > 0 {
		_ = tempMBR.Partitions[totalPartitions-1].Start + tempMBR.Partitions[totalPartitions-1].Size
	}
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
	return errors.New("the specified file is not a disk"), false
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

func (f *Fdisk) validatePartitions(mbr Structs.MBR, size int) (error, int, int, int) {
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
		return errors.New("cannot create more than 4 primary or extended partitions"), primaryCount, extendedCount, totalPartitions
	}
	if f.Type == 'E' && extendedCount >= 1 {
		return errors.New("only one extended partition is supported"), primaryCount, extendedCount, totalPartitions
	}
	if f.Type == 'L' && extendedCount == 0 {
		return errors.New("cannot create a logical partition without an extended partition"), primaryCount, extendedCount, totalPartitions
	}
	if usedSpace+int32(size) > mbr.MbrSize {
		return errors.New("not enough disk space to create this partition"), primaryCount, extendedCount, totalPartitions
	}
	return nil, primaryCount, extendedCount, totalPartitions
}
