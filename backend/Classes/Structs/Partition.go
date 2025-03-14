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

func NewPartition(status, type_, fit [1]byte, start, size int32, name [16]byte, correlative int32) *Partition {
	return &Partition{
		Status:      status,
		Type:        type_,
		Fit:         fit,
		Start:       start,
		Size:        size,
		Name:        name,
		Correlative: correlative,
	}
}

func (p *Partition) ToString() string {
	return fmt.Sprintf("Start: %d, Status: %c, Size: %d, Name: %s, Type: %s, Fit: %s, ID: %s, Correlative: %d", p.Start, p.Status[:], p.Size, p.Name, p.Type, p.Fit, p.Id[:], p.Correlative)
}

type MountedPartition struct {
	Partition
	DiscSignature int32
	DiscTag       rune
	Path          string
}

func (m *MountedPartition) ToString() string {
	return fmt.Sprintf("\tPartition:\n\t\t%s\n\t\tDisc Signature: %d, Disc Tag: %c", m.Partition.ToString(), m.DiscSignature, m.DiscTag)
}
