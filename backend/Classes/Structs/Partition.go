package Structs

import "fmt"

type Partition struct {
	Status      [1]byte
	Type        [1]byte
	Fit         [1]byte
	Start       int32
	Size        int32
	Name        [16]byte
	Correlative int32
	Id          [4]byte
}

func (p *Partition) ToString() string {
	return fmt.Sprintf("Start: %d, Status: %c, Size: %d, Name: %s", p.Start, p.Status, p.Size, p.Name)
}
