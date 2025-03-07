package Structs

import (
	"golang.org/x/exp/rand"
	"time"
)

type MBR struct {
	MbrSize      int32
	CreationDate [10]byte
	Signature    int32
	Fit          [1]byte
	Partitions   [4]Partition
}

func NewMBR(size int32, fit [1]byte) *MBR {
	timeString := time.Now().Format("2006-01-02")
	var timeByteArr [10]byte
	copy(timeByteArr[:], timeString)

	return &MBR{
		MbrSize:      size,
		CreationDate: timeByteArr,
		Signature:    rand.Int31(),
		Fit:          fit,
		Partitions:   [4]Partition{},
	}
}

/* WAIT
// maybe use this, for the moment use lab code
func (mbr *MBR) Encode() []byte {
	var buffer bytes.Buffer
	return buffer.Bytes()
}

func Decode(data []byte) *MBR {

}
*/
