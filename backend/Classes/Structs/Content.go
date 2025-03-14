package Structs

import "fmt"

type Content struct {
	Name  [12]byte
	Inode int32
}

func (c *Content) ToString() string {
	return fmt.Sprintf(`Name: %s, Inode: %d`, string(c.Name[:]), c.Inode)
}
