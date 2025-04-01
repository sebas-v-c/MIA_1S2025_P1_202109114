package Structs

type FolderBlock struct {
	Content [4]Content
}

func NewFolderBlock() *FolderBlock {
	return &FolderBlock{
		Content: [4]Content{
			{Inode: -1},
			{Inode: -1},
			{Inode: -1},
			{Inode: -1},
		},
	}
}

func (fb *FolderBlock) ToString() string {
	str := "\t\tFolderBlock Content:"
	for _, content := range fb.Content {
		str += "\n\t\t\t" + content.ToString()
	}
	return str
}
