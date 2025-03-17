package Structs

type FolderBlock struct {
	Content [4]Content
}

func (fb *FolderBlock) ToString() string {
	str := "\t\tFolderBlock Content:"
	for _, content := range fb.Content {
		str += "\n\t\t\t" + content.ToString()
	}
	return str
}
