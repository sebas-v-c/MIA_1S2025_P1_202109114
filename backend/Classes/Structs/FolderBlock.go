package Structs

type FolderBlock struct {
	Content [4]Content
}

func (fb *FolderBlock) ToString() string {
	str := "FolderBlock Content:"
	for _, content := range fb.Content {
		str += "\n" + content.ToString()
	}
	return str
}
