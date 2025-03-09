package Commands

import (
	Env "backend/Classes/Env"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"errors"
	"path/filepath"
	"strconv"
	"strings"
)

type Mkdisk struct {
	Result string
	Type   Utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewMkdisk(line, column int, params map[string]string) *Mkdisk {
	return &Mkdisk{
		Type:   Utils.MKDISK,
		Params: params,
		Line:   line,
		Column: column,
	}
}

func (m *Mkdisk) GetLine() int {
	return m.Line
}

func (m *Mkdisk) GetColumn() int {
	return m.Column
}

func (m *Mkdisk) GetType() Utils.Type {
	return m.Type
}

func (m *Mkdisk) Exec(env *Env.Env) {
	// if the parameters are not valid
	if err, ok := m.validParams(); !ok {
		env.Errors = append(env.Errors, Env.RuntimeError{
			Line:    m.Line,
			Column:  m.Column,
			Command: Utils.MKDISK,
			Msg:     err.Error(),
		})
		return
	}

	filePath := m.Params["path"]

	err := Utils.CreateFile(filePath)
	if err != nil {
		env.Errors = append(env.Errors, Env.RuntimeError{
			Line:    m.Line,
			Column:  m.Column,
			Command: Utils.MKDISK,
			Msg:     err.Error(),
		})
		return
	}

	file, err := Utils.OpenFile(filePath)
	if err != nil {
		env.Errors = append(env.Errors, Env.RuntimeError{
			Line:    m.Line,
			Column:  m.Column,
			Command: Utils.MKDISK,
			Msg:     err.Error(),
		})
		return
	}

	units := m.recalculateUnits()
	size, _ := strconv.Atoi(m.Params["size"])
	totalSize := units * size

	batchSize := 1024 * 1024
	emptyData := make([]byte, batchSize)
	bytesWritten := 0
	// Writing bytes in batches to save resources
	for bytesWritten < totalSize {
		writeSize := batchSize
		if totalSize-bytesWritten < batchSize {
			writeSize = totalSize - bytesWritten // write the last batch
		}
		file.Write(emptyData[:writeSize])
		bytesWritten += writeSize
	}
	// Write MBR to the file
	newMBR := Structs.NewMBR(int32(totalSize), m.getFit())
	if err := Utils.WriteObject(file, newMBR, 0); err != nil {
		env.Errors = append(env.Errors, Env.RuntimeError{
			Line:    m.Line,
			Column:  m.Column,
			Command: Utils.MKDISK,
			Msg:     err.Error(),
		})
		return
	}
	defer file.Close()
	env.CommandLog = append(env.CommandLog, "=================MKDISK=================\n"+newMBR.ToString()+"\n=================END MKDISK=================\n")
}

func (m *Mkdisk) validParams() (error, bool) {
	if _, ok := m.Params["size"]; !ok {
		return errors.New("missing parameter -size"), false
	} else if _, ok := m.Params["path"]; !ok {
		return errors.New("missing parameter -path"), false
	}

	if _, ok := m.Params["fit"]; !ok {
		m.Params["fit"] = "FF"
	}
	if _, ok := m.Params["unit"]; !ok {
		m.Params["unit"] = "M"
	}

	if val, _ := strconv.Atoi(m.Params["size"]); val <= 0 {
		return errors.New("size must be greater than 0"), false
	}
	// check if file is a disk
	if !strings.EqualFold(filepath.Ext(m.Params["path"]), ".mia") {
		m.Params["path"] = m.Params["path"] + ".mia"
	}

	return nil, true
}

func (m *Mkdisk) GetResult() string {
	//TODO implement me
	panic("implement me")
}

func (m *Mkdisk) recalculateUnits() int {
	// if size is k multiply value by 1024
	// if the unit is m multiply by 1024*1024
	m.Params["unit"] = strings.ToUpper(m.Params["unit"])
	if m.Params["unit"] == "K" {
		return 1024
	}
	return 1024 * 1024
}

func (m *Mkdisk) getFit() [1]byte {
	m.Params["fit"] = strings.ToUpper(m.Params["fit"])
	if m.Params["fit"] == "FF" {
		return [1]byte{'F'}
	} else if m.Params["fit"] == "BF" {
		return [1]byte{'B'}
	}
	return [1]byte{'W'}
}
