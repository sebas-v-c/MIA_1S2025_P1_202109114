package Handlers

import (
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"
)

type RequestPath struct {
	Path string `json:"path"`
}

type FileInfo struct {
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	Type         string    `json:"type"`
	CreatedTime  time.Time `json:"createdTime"`
	ModifiedTime time.Time `json:"modifiedTime"`
	Content      string    `json:"content"`
	GroupId      int       `json:"groupId"`
	UserId       int       `json:"userId"`
}

type DiscInfo struct {
	Name         string          `json:"name"`
	MbrSize      int32           `json:"mbrSize"`
	CreationDate string          `json:"creationDate"`
	Signature    int32           `json:"signature"`
	Fit          string          `json:"fit"`
	Partitions   []PartitionInfo `json:"partitions"`
}

type PartitionInfo struct {
	Status      string `json:"status"`
	Type        string `json:"type"`
	Fit         string `json:"fit"`
	Start       int32  `json:"start"` // Changed from int to int32 to match Partition
	Size        int32  `json:"size"`  // Changed from int to int32 to match Partition
	Name        string `json:"name"`
	Correlative int32  `json:"correlative"` // Changed from int to int32 to match Partition
	Id          string `json:"id"`          // Changed to string since Id in Partition is [4]byte
}

// getDiscsDirectory returns the root Discs directory relative to this handler
func getDiscsDirectory() string {
	_, currentFile, _, _ := runtime.Caller(0)
	handlersDir := filepath.Dir(currentFile)
	// Go up one level to project root, then into Discs folder
	return filepath.Join(filepath.Dir(handlersDir), "Discs")
}

// Convert MBR to DiscInfo
func MBRToDiscInfo(mbr *Structs.MBR, name string) DiscInfo {
	partitions := make([]PartitionInfo, 0)
	for _, p := range mbr.Partitions {
		if p.Size > 0 { // Only include non-empty partitions
			partitions = append(partitions, PartitionToPartitionInfo(&p))
		}
	}

	return DiscInfo{
		Name:         name,
		MbrSize:      mbr.MbrSize,
		CreationDate: string(mbr.CreationDate[:]),
		Signature:    mbr.Signature,
		Fit:          string(mbr.Fit[:]),
		Partitions:   partitions,
	}
}

// Convert Partition to PartitionInfo
func PartitionToPartitionInfo(p *Structs.Partition) PartitionInfo {
	return PartitionInfo{
		Status:      string(p.Status[:]),
		Type:        string(p.Type[:]),
		Fit:         string(p.Fit[:]),
		Start:       p.Start,
		Size:        p.Size,
		Name:        string(p.Name[:]),
		Correlative: p.Correlative,
		Id:          string(p.Id[:]),
	}
}

// listDirectoryContents reads the given path and returns FileInfo for each entry
func listDirectoryContents(path string) ([]DiscInfo, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var discs []DiscInfo
	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())
		_, err := entry.Info()
		if err != nil {
			continue
		}

		file, err := Utils.OpenFile(fullPath)
		if err != nil {
			return nil, err
		}

		// Read the disk's Master Boot Record (MBR) from the file.
		var diskMBR Structs.MBR
		if err := Utils.ReadObject(file, &diskMBR, 0); err != nil {
			return nil, err
		}

		discs = append(discs, MBRToDiscInfo(&diskMBR, entry.Name()))
	}
	return discs, nil
}

func DiscsHandler(c *gin.Context) {
	var req RequestPath
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	if req.Path == "" {
		c.JSON(400, gin.H{"error": "Path cannot be empty"})
		return
	}

	if !strings.HasPrefix(req.Path, "%") {
		c.JSON(400, gin.H{"error": "Path must start with %"})
		return
	}

	discsDir := getDiscsDirectory()
	path := strings.TrimPrefix(req.Path, "%")

	// Case 1: list all discs
	if path == "" {
		list, err := listDirectoryContents(discsDir)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to list discs directory"})
			return
		}
		c.JSON(200, gin.H{"type": "list_discs", "contents": list})
		return
	}

	// Split disc and subpath
	parts := strings.SplitN(path, "%", 2)
	discName := parts[0]
	discPath := filepath.Join(discsDir, discName)

	// Case 2: disc info only
	if len(parts) == 1 {
		// Check existence
		if _, err := os.Stat(discPath); os.IsNotExist(err) {
			c.JSON(404, gin.H{"error": "Disc not found"})
			return
		}

		fi, err := os.Stat(discPath)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to get disc information"})
			return
		}

		statT, _ := fi.Sys().(*syscall.Stat_t)
		c.JSON(200, gin.H{
			"type": "disc_operation",
			"disc": FileInfo{
				Name:         fi.Name(),
				Size:         fi.Size(),
				Type:         "Disc",
				ModifiedTime: fi.ModTime(),
				CreatedTime:  fi.ModTime(),
				Content:      "",
				GroupId:      int(statT.Gid),
				UserId:       int(statT.Uid),
			},
		})
		return
	}

	// Case 3: full path (list contents of subpath)
	subPath := filepath.Join(discPath, parts[1])
	contents, err := listDirectoryContents(subPath)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to list path contents"})
		return
	}
	c.JSON(200, gin.H{"type": "path_operation", "disc": discName, "contents": contents})
}
