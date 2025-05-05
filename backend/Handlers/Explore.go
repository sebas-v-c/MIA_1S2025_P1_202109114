package Handlers

import (
	"backend/Classes/Env"
	env "backend/Classes/Env"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"github.com/gin-gonic/gin"
	"os"
)

type ExploreRequest struct {
	Path string `json:"path"`
}

type FileResponse struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Size        int32  `json:"size"`
	ATime       string `json:"atime"`
	MTime       string `json:"mtime"`
	CTime       string `json:"ctime"`
	Owner       string `json:"owner"`
	Group       string `json:"group"`
	Permissions string `json:"permissions"`
	Content     string `json:"content"`
}
type ExploreResponse struct {
	Files []FileResponse `json:"files"`
}

func ExploreHandler(c *gin.Context) {
	var request ExploreRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	if Env.CurrentUser == nil {
		c.JSON(401, gin.H{"error": "No user is currently logged in"})
		return
	}

	// Verify disk integrity using the mounted partition.
	var mbrPartition *Structs.Partition
	var file *os.File
	var err error
	_, mbrPartition, file, err = env.VerifyDiscStatus(env.CurrentUser.MountedPartition.Id)
	if err != nil {
		c.JSON(401, gin.H{"error": err})
		return
	}
	defer file.Close()

	// Read the superblock from disk using the partition start offset.
	var superBlock Structs.SuperBlock
	if err = Utils.ReadObject(file, &superBlock, int64(mbrPartition.Start)); err != nil {
		c.JSON(401, gin.H{"error": err})
		return
	}

	// Build the directory tree from the superblock.
	var dirTree = Structs.NewDirTree(superBlock, file)
	// Iterate over each file path provided in the command.
	fileName := request.Path
	var fileInode *Structs.Inode
	// Retrieve the inode for the file by its path.
	_, fileInode, err = dirTree.GetInodeByPath(fileName)
	if err != nil {
		c.JSON(401, gin.H{"error": err})
		return
	}

	// Check that the current user has read permissions for the file.
	var read bool
	read, _, _, err = env.CheckFilePermissions(*env.CurrentUser, fileInode)
	if !read {
		c.JSON(401, gin.H{"error": "You do not have permission to read file " + fileName})
		return
	}

	if fileInode.Type == [1]byte{1} {
		c.JSON(401, gin.H{"error": "File " + fileName + " is not a directory"})
	}
	fileList, err := dirTree.GetFolderContent(fileInode)
	if err != nil {
		c.JSON(401, gin.H{"error": err})
		return
	}

	var response ExploreResponse
	for _, fileItem := range fileList {
		var fileContent = ""
		var fileType = "0"
		if fileItem.Inode.Type == [1]byte{1} {
			fileType = "1"
			fileContent, err = dirTree.GetFileContentByInode(&fileItem.Inode)
			if err != nil {
				c.JSON(401, gin.H{"error": err})
				return
			}
		}

		owner, err := env.GetUserNameByID(fileItem.Inode.UID, env.CurrentUser.MountedPartition)
		if err != nil {
			c.JSON(401, gin.H{"error": err})
			return
		}
		group, err := env.GetGroupNameByID(fileItem.Inode.GID, env.CurrentUser.MountedPartition)
		if err != nil {
			c.JSON(401, gin.H{"error": err})
		}

		fileResponse := FileResponse{
			Name:        fileItem.Name,
			Type:        fileType,
			Size:        fileItem.Inode.Size,
			ATime:       string(fileItem.Inode.ATime[:]),
			MTime:       string(fileItem.Inode.MTime[:]),
			CTime:       string(fileItem.Inode.CTime[:]),
			Owner:       owner,
			Group:       group,
			Permissions: string(fileItem.Inode.Perm[:]),
			Content:     fileContent,
		}

		response.Files = append(response.Files, fileResponse)

	}

	/*
		// Retrieve the file content from the inode.
		var fileContent string
		fileContent, err = dirTree.GetFileContentByInode(fileInode)
		if err != nil {
			c.JSON(401, gin.H{"error": err})
			return
		}
	*/

	c.JSON(200, response)

}
