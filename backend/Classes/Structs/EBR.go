package Structs

import "fmt"

type EBR struct {
	Mount byte
	Fit   byte
	Start int32
	Size  int32
	Next  int32
	Name  [16]byte
}

func NewEBR(fit byte, start, size int32, name [16]byte) *EBR {
	return &EBR{
		Mount: '0',
		Fit:   fit,
		Start: start,
		Size:  size,
		Next:  -1,
		Name:  name,
	}
}

func (e *EBR) ToString() string {
	return fmt.Sprintf("Mount: %-5v, Start: %d, Size: %d, Name: %s, Next: %d", e.Mount, e.Start, e.Size, e.Name, e.Next)
}
