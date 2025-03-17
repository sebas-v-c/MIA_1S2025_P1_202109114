package Structs

import (
	"bytes"
	"fmt"
)

type FileBlock struct {
	Content [64]byte
}

func (fb *FileBlock) ToString() string {
	contentArr := [64]byte{}
	copy(contentArr[:], fb.Content[:])
	trimmed := fb.Content[:bytes.IndexByte(fb.Content[:], 0)]

	fileBlockString := fmt.Sprintf("%q", trimmed)
	return fileBlockString[1 : len(fileBlockString)-1]
}
