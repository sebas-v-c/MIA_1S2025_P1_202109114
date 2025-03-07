package Utils

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(name string) error {
	// Make sure file exist
	dir := filepath.Dir(name)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Println("Err CreateFile dir==", err)
		return err
	}

	// Create file
	if _, err := os.Stat(name); os.IsNotExist(err) {
		file, err := os.Create(name)
		if err != nil {
			fmt.Println("Err CreateFile create==", err)
			return err
		}
		defer file.Close()
	}
	return nil
}

func OpenFile(name string) (*os.File, error) {
	file, err := os.OpenFile(name, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Err OpenFile==", err)
		return nil, err
	}
	return file, nil
}

func WriteObject(file *os.File, data interface{}, position int64) error {
	// Move to correct position
	_, err := file.Seek(position, 0)
	if err != nil {
		fmt.Println("Err WriteObject seek==", err)
		return err
	}

	// Use a buffered writer to improve performance
	bufferedWriter := bufio.NewWriter(file)
	err = binary.Write(bufferedWriter, binary.LittleEndian, data)
	if err != nil {
		fmt.Println("Err WriteObject==", err)
		return err
	}

	// Flush the buffer (write to disk in large chunks)
	err = bufferedWriter.Flush()
	if err != nil {
		fmt.Println("Err WriteObject flush==", err)
		return err
	}
	return nil
}

func ReadObject(file *os.File, data interface{}, position int64) error {
	// Move to the correct position
	_, err := file.Seek(position, 0)
	if err != nil {
		fmt.Println("Err ReadObject seek==", err)
		return err
	}

	// Use a buffered reader to speed up reading
	bufferedReader := bufio.NewReader(file)
	err = binary.Read(bufferedReader, binary.LittleEndian, data)
	if err != nil {
		fmt.Println("Err ReadObject==", err)
		return err
	}
	return nil
}
