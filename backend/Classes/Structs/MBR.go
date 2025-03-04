package Structs

type MBR struct {
	MbrSize      int32
	CreationDate [10]byte
	Signature    int32
	Fit          [1]byte
}
