package Utils

import (
	"bufio"
	"encoding/binary"
	"os"
	"path/filepath"
)

func CreateFile(name string) error {
	// Make sure file exist
	dir := filepath.Dir(name)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// Create file
	if _, err := os.Stat(name); os.IsNotExist(err) {
		file, err := os.Create(name)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

func OpenFile(name string) (*os.File, error) {
	file, err := os.OpenFile(name, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func WriteObject(file *os.File, data interface{}, position int64) error {
	// Move to correct position
	_, err := file.Seek(position, 0)
	if err != nil {
		return err
	}

	// Use a buffered writer to improve performance
	bufferedWriter := bufio.NewWriter(file)
	err = binary.Write(bufferedWriter, binary.LittleEndian, data)
	if err != nil {
		return err
	}

	// Flush the buffer (write to disk in large chunks)
	err = bufferedWriter.Flush()
	if err != nil {
		return err
	}
	return nil
}

func ReadObject(file *os.File, data interface{}, position int64) error {
	// Move to the correct position
	_, err := file.Seek(position, 0)
	if err != nil {
		return err
	}

	// Use a buffered reader to speed up reading
	bufferedReader := bufio.NewReader(file)
	err = binary.Read(bufferedReader, binary.LittleEndian, data)
	if err != nil {
		return err
	}
	return nil
}

func StringPointer(s string) *string { return &s }
