package Structs

import "fmt"

type PointerBlock struct {
	Pointers [16]int32
}

func (p *PointerBlock) ToString() string {
	return fmt.Sprintf("Pointers: %v", p.Pointers)
}
