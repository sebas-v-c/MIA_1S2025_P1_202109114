package Structs

import "fmt"

type FileBlock struct {
	Content [64]byte
}

func (fb *FileBlock) ToString() string {
	return fmt.Sprintf("%s", fb.Content[:])
}
