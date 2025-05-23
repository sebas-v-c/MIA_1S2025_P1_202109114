package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type Rep struct {
	Interfaces.CommandStruct
	Params map[string]string
}

func NewRep(line, column int, params map[string]string) *Rep {
	return &Rep{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.REP,
			Line:   line,
			Column: column,
		},
		Params: params,
	}
}

func (r *Rep) Exec() {
	var consoleString strings.Builder
	consoleString.WriteString("=================REP=================\n")

	if err := r.validateParams(); err != nil {
		r.AppendError(err.Error())
		return
	}

	// Verify disc integrity
	var mbrPartition *Structs.Partition
	var file *os.File
	var err error
	var id [4]byte
	copy(id[:], r.Params["id"])
	_, mbrPartition, file, err = env.VerifyDiscStatus(id)
	if err != nil {
		r.AppendError(err.Error())
		return
	}
	defer file.Close()

	var superBlock Structs.SuperBlock
	if err = Utils.ReadObject(file, &superBlock, int64(mbrPartition.Start)); err != nil {
		r.AppendError(err.Error())
		return
	}

	// Get dir tree
	var dirTree = Structs.NewDirTree(superBlock, file)
	var repContent string

	switch r.Params["name"] {
	case "mbr":
		repContent, err = r.createMBRRep(file)
		if err != nil {
			r.AppendError(err.Error())
			return
		}
	case "disk":
		repContent, err = r.createDiskRep(file)
		if err != nil {
			r.AppendError(err.Error())
			return
		}
	case "inode":
		repContent, err = r.createInodeRep(dirTree)
		if err != nil {
			r.AppendError(err.Error())
			return
		}
	case "block":
		//r.createBlockRep()
		if err != nil {
			r.AppendError(err.Error())
			return
		}
	case "bm_inode":
		repContent, err = r.createBMRep(dirTree.InodeBitMap, "BM_INODE")
		if err != nil {
			r.AppendError(err.Error())
			return
		}
	case "bm_block":
		repContent, err = r.createBMRep(dirTree.BlockBitMap, "BM_BLOCK")
		if err != nil {
			r.AppendError(err.Error())
			return
		}
	case "tree":
		repContent, err = r.createTreeRep(dirTree)
		if err != nil {
			r.AppendError(err.Error())
			return
		}
	case "sb":
		repContent, err = r.createSbRep(dirTree)
		if err != nil {
			r.AppendError(err.Error())
			return
		}
	case "file":
		repContent, err = r.createFileRep(dirTree)
		if err != nil {
			r.AppendError(err.Error())
			return
		}
	case "ls":
		repContent, err = r.createLsRep(dirTree)
		if err != nil {
			r.AppendError(err.Error())
			return
		}
	default:
		r.AppendError("Unknown command: " + r.Params["name"])
		return
	}

	// Create report file and write it
	if err := Utils.CreateFile(r.Params["path"] + ".dot"); err != nil {
		r.AppendError(err.Error())
		return
	}
	repFile, err := Utils.OpenFile(r.Params["path"] + ".dot")
	if err != nil {
		r.AppendError(err.Error())
		return
	}

	if err := Utils.WriteObject(repFile, []byte(repContent), int64(0)); err != nil {
		r.AppendError(err.Error())
		return
	}

	// Convert .dot to image
	// Generate image using dot command (requires Graphviz installed)

	cmd := exec.Command("dot", "-Tjpg", r.Params["path"]+".dot", "-o", r.Params["path"]+".jpg")
	err = cmd.Run()
	if err != nil {
		r.AppendError(err.Error())
		return
	}

	consoleString.WriteString("\n=================REP=================\n")
	r.LogConsole(consoleString.String())
}

func (r *Rep) validateParams() error {
	if _, ok := r.Params["path"]; !ok {
		return errors.New("missing parameter -path")
	}

	if _, ok := r.Params["id"]; !ok {
		return errors.New("missing parameter -id")
	}
	r.Params["id"] = strings.ToUpper(r.Params["id"])

	name, ok := r.Params["name"]
	if !ok {
		return errors.New("missing parameter -name")
	}

	switch name {
	case "ls", "file":
		if _, ok := r.Params["pfl"]; !ok {
			return errors.New("parameter -path_file_ls is required when generating 'ls' and 'file' reports")
		}
	case "mbr", "disk", "inode", "block", "bm_inode", "bm_block", "tree", "sb":

	default:
		return errors.New("invalid value for -name. Allowed: mbr, disk, inode, block, bm_inode, bm_block, tree, sb, file, ls")
	}

	return nil
}

func (r *Rep) createMBRRep(file *os.File) (string, error) {
	var sb strings.Builder

	var mbr Structs.MBR
	if err := Utils.ReadObject(file, &mbr, 0); err != nil {
		return "", err
	}

	sb.WriteString("digraph MBR {\n")
	sb.WriteString("    node [shape=plaintext];\n\n")
	sb.WriteString("    mbr_report [\n")
	sb.WriteString("        label=<\n")
	sb.WriteString("            <table border=\"0\" cellborder=\"1\" cellspacing=\"0\" cellpadding=\"4\">\n")
	sb.WriteString("                <tr><td colspan=\"2\" bgcolor=\"#4B0082\"><font color=\"white\"><b>REPORTE DE MBR</b></font></td></tr>\n")
	sb.WriteString(fmt.Sprintf("                <tr><td><b>mbr_tamano</b></td><td>%d</td></tr>\n", mbr.MbrSize))
	sb.WriteString(fmt.Sprintf("                <tr><td><b>mbr_fecha_creacion</b></td><td>%s</td></tr>\n", string(mbr.CreationDate[:])))
	sb.WriteString(fmt.Sprintf("                <tr><td><b>mbr_disk_signature</b></td><td>%d</td></tr>\n", mbr.Signature))

	for _, part := range mbr.Partitions {
		partName := strings.TrimRight(string(part.Name[:]), "\x00")
		partType := string(part.Type[:])
		partFit := string(part.Fit[:])
		partStatus := string(part.Status[:])

		if partName == "" || part.Size == 0 {
			continue // Skip empty partitions
		}

		color := "#5D3FD3"
		label := "Particion"

		sb.WriteString(fmt.Sprintf("                <tr><td colspan=\"2\" bgcolor=\"%s\"><b>%s</b></td></tr>\n", color, label))
		sb.WriteString(fmt.Sprintf("                <tr><td><b>part_status</b></td><td>%s</td></tr>\n", partStatus))
		sb.WriteString(fmt.Sprintf("                <tr><td><b>part_type</b></td><td>%s</td></tr>\n", partType))
		sb.WriteString(fmt.Sprintf("                <tr><td><b>part_fit</b></td><td>%s</td></tr>\n", partFit))
		sb.WriteString(fmt.Sprintf("				   <tr><td><b>part_start</b></td><td>%d</td></tr>\n", part.Start))
		sb.WriteString(fmt.Sprintf("                <tr><td><b>part_size</b></td><td>%d</td></tr>\n", part.Size))
		sb.WriteString(fmt.Sprintf("                <tr><td><b>part_name</b></td><td>%s</td></tr>\n", partName))

		if partType == "E" {
			// Leer EBRs encadenados
			currentEBRPos := part.Start
			for currentEBRPos != -1 {
				var ebr Structs.EBR
				if err := Utils.ReadObject(file, &ebr, int64(currentEBRPos)); err != nil {
					break // error leyendo EBR, salir del bucle
				}

				var ebrMount string
				if ebr.Mount == 0 {
					ebrMount = "0"
				} else {
					ebrMount = "1"
				}

				ebrName := strings.TrimRight(string(ebr.Name[:]), "\x00")
				sb.WriteString("                <tr><td colspan=\"2\" bgcolor=\"#FF9999\"><b>Particion Logica</b></td></tr>\n")
				sb.WriteString(fmt.Sprintf("                <tr><td><b>part_mount</b></td><td>%s</td></tr>\n", ebrMount))
				sb.WriteString(fmt.Sprintf("                <tr><td><b>part_next</b></td><td>%d</td></tr>\n", ebr.Next))
				sb.WriteString(fmt.Sprintf("                <tr><td><b>part_fit</b></td><td>%s</td></tr>\n", string(ebr.Fit)))
				sb.WriteString(fmt.Sprintf("                <tr><td><b>part_start</b></td><td>%d</td></tr>\n", ebr.Start))
				sb.WriteString(fmt.Sprintf("                <tr><td><b>part_size</b></td><td>%d</td></tr>\n", ebr.Size))
				sb.WriteString(fmt.Sprintf("                <tr><td><b>part_name</b></td><td>%s</td></tr>\n", ebrName))

				// Actualizar posición del siguiente EBR
				currentEBRPos = ebr.Next
			}
			continue // ya se procesó la extendida
		}
	}
	sb.WriteString("                <tr><td colspan=\"2\" align=\"center\"><b>Reporte de MBR</b></td></tr>\n")
	sb.WriteString("            </table>\n")
	sb.WriteString("        >\n")
	sb.WriteString("    ];\n")
	sb.WriteString("}\n")

	return sb.String(), nil
}

func (r *Rep) createDiskRep(file *os.File) (string, error) {
	var sb strings.Builder

	var mbr Structs.MBR
	if err := Utils.ReadObject(file, &mbr, 0); err != nil {
		return "", err
	}

	totalSize := float64(mbr.MbrSize)

	sb.WriteString("digraph DiskLayout {\n")
	sb.WriteString("    node [shape=plaintext];\n")
	sb.WriteString("    disk [label=<\n")
	sb.WriteString("    <table border=\"0\" cellborder=\"1\" cellspacing=\"0\">\n")
	sb.WriteString("        <tr>\n")

	// MBR fixed block
	sb.WriteString("            <td><b>MBR</b></td>\n")
	used := (float64(binary.Size(Structs.MBR{})) / totalSize) * 100

	for _, part := range mbr.Partitions {
		if part.Size == 0 {
			continue
		}

		partName := strings.TrimRight(string(part.Name[:]), "\x00")
		partType := strings.TrimSpace(string(part.Type[:]))
		partSize := float64(part.Size)
		partPercentage := (partSize / totalSize) * 100

		if partType == "E" {
			sb.WriteString("            <td>\n")
			sb.WriteString("                <table border=\"0\" cellborder=\"1\">\n")
			sb.WriteString("                    <tr><td colspan=\"2\"><b>Extendida</b></td></tr>\n")

			// Add the EBR metadata size (if needed)
			used += (float64(binary.Size(Structs.EBR{})) / totalSize) * 100

			currentEBRPos := part.Start
			endOfExtended := part.Start + part.Size
			lastLogicalEnd := part.Start

			for currentEBRPos != -1 {
				var ebr Structs.EBR
				if err := Utils.ReadObject(file, &ebr, int64(currentEBRPos)); err != nil {
					break
				}

				ebrName := strings.TrimRight(string(ebr.Name[:]), "\x00")
				if ebrName == "" {
					ebrName = "Logica"
				}

				logicalStart := ebr.Start
				logicalEnd := ebr.Start + ebr.Size
				logicalSize := float64(ebr.Size)
				logicalPercentage := (logicalSize / totalSize) * 100

				// Internal free space between logicals
				if logicalStart > lastLogicalEnd {
					freeSize := float64(logicalStart - lastLogicalEnd)
					freePercentage := (freeSize / totalSize) * 100
					//sb.WriteString(fmt.Sprintf("                    <tr><td colspan=\"2\">Libre<br/>%.2f%%</td></tr>\n", freePercentage))
					used += freePercentage
				}

				// Add EBR + logical partition
				sb.WriteString("                    <tr>\n")
				sb.WriteString(fmt.Sprintf("                        <td>EBR</td><td>%s<br/>%.2f%%</td>\n", ebrName, logicalPercentage))
				sb.WriteString("                    </tr>\n")
				used += logicalPercentage
				lastLogicalEnd = logicalEnd
				currentEBRPos = ebr.Next
			}

			// Final internal free space in extended partition
			if endOfExtended > lastLogicalEnd {
				freeSize := float64(endOfExtended - lastLogicalEnd)
				freePercentage := (freeSize / totalSize) * 100
				sb.WriteString(fmt.Sprintf("                    <tr><td colspan=\"2\">Libre<br/>%.2f%%</td></tr>\n", freePercentage))
				used += freePercentage
			}

			sb.WriteString("                </table>\n")
			sb.WriteString("            </td>\n")
			//used += partPercentage // count full extended partition
		} else {
			sb.WriteString(fmt.Sprintf("            <td><b>%s</b><br/>%.2f%%</td>\n", partName, partPercentage))
			used += partPercentage
		}
	}

	// Remaining free space in the disk
	free := 100.0 - used
	if free < 0 {
		free = 0
	}
	sb.WriteString(fmt.Sprintf("            <td>Libre<br/>%.2f%%</td>\n", free))

	sb.WriteString("        </tr>\n</table>>];\n}")

	return sb.String(), nil
}

func (r *Rep) createTreeRep(dirTree *Structs.DirTree) (string, error) {
	var sb strings.Builder

	sb.WriteString("digraph FileSystem {\n")
	sb.WriteString("    rankdir=LR;\n")
	sb.WriteString("    splines=false;\n")
	sb.WriteString("    node [shape=plaintext, style=\"filled,rounded\", fillcolor=\"#f0f8ff\"];\n")

	// here we store block and file addresses so we can recover this addresses and graph them
	var folderBlockAddresses []int32
	var fileBlockAddresses []int32

	// Structure type
	type ConnectionType int
	const (
		ItoF ConnectionType = iota
		ItoD
		DtoI
	)
	type Connection struct {
		Type  ConnectionType
		Inode int32
		Block int32
		From  int32
	}
	var connections []Connection

	// Loop through the BitMap of inodes
	for i, bit := range dirTree.InodeBitMap {
		index := int32(i)
		if bit != 1 {
			continue
		}

		var bitMapInode Structs.Inode
		if err := Utils.ReadObject(dirTree.File, &bitMapInode, int64(dirTree.SuperBlock.InodeStart+index*int32(binary.Size(Structs.Inode{})))); err != nil {
			return "", err
		}

		sb.WriteString(fmt.Sprintf("\n    inode%d [\n", index))
		sb.WriteString(fmt.Sprintf("        label=<\n"))
		sb.WriteString(fmt.Sprintf("        <TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\">\n"))
		sb.WriteString(fmt.Sprintf("            <TR><TD COLSPAN=\"2\" BGCOLOR=\"#cce5ff\"><B>Inode%d</B></TD></TR>\n", index))
		sb.WriteString(fmt.Sprintf("            <TR><TD>UID</TD><TD>%d</TD></TR>\n", bitMapInode.UID))
		sb.WriteString(fmt.Sprintf("            <TR><TD>GID</TD><TD>%d</TD></TR>\n", bitMapInode.GID))
		sb.WriteString(fmt.Sprintf("            <TR><TD>Size</TD><TD>%d</TD></TR>\n", bitMapInode.Size))
		time := strings.TrimRight(string(bitMapInode.ATime[:]), "\x00")
		sb.WriteString(fmt.Sprintf("            <TR><TD>ATime</TD><TD>%s</TD></TR>\n", time))
		time = strings.TrimRight(string(bitMapInode.CTime[:]), "\x00")
		sb.WriteString(fmt.Sprintf("            <TR><TD>CTime</TD><TD>%s</TD></TR>\n", time))
		time = strings.TrimRight(string(bitMapInode.MTime[:]), "\x00")
		sb.WriteString(fmt.Sprintf("            <TR><TD>MTime</TD><TD>%s</TD></TR>\n", time))
		sb.WriteString(fmt.Sprintf("            <TR>\n"))
		sb.WriteString(fmt.Sprintf("                <TD>IBlock</TD>\n"))
		sb.WriteString(fmt.Sprintf("                <TD>\n"))
		sb.WriteString(fmt.Sprintf("                    <TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\">\n"))
		for j, addr := range bitMapInode.IBlock {
			sb.WriteString(fmt.Sprintf("                        <TR><TD PORT=\"ib%d\">block_%d: %d</TD></TR>\n", j, j, addr))
			if addr == -1 {
				continue
			}
			// Set connection type
			var connType ConnectionType
			if bitMapInode.Type == [1]byte{0} {
				connType = ItoD
				folderBlockAddresses = append(folderBlockAddresses, addr)
			} else {
				connType = ItoF
				fileBlockAddresses = append(fileBlockAddresses, addr)
			}
			connections = append(connections, Connection{
				Type:  connType,
				Inode: index,
				Block: addr,
				From:  int32(j),
			})
		}
		sb.WriteString(fmt.Sprintf("                    </TABLE>\n"))
		sb.WriteString(fmt.Sprintf("                </TD>\n"))
		sb.WriteString(fmt.Sprintf("            </TR>\n"))
		if bitMapInode.Type == [1]byte{0} {
			sb.WriteString(fmt.Sprintf("            <TR><TD>Type</TD><TD>%d</TD></TR>\n", 0))
		} else {
			sb.WriteString(fmt.Sprintf("            <TR><TD>Type</TD><TD>%d</TD></TR>\n", 1))
		}
		sb.WriteString(fmt.Sprintf("            <TR><TD>Perm</TD><TD>%s</TD></TR>\n", string(bitMapInode.Perm[:])))
		sb.WriteString(fmt.Sprintf("        </TABLE>\n"))

		sb.WriteString(fmt.Sprintf("        >\n"))
		sb.WriteString(fmt.Sprintf("    ];\n"))

	}

	for _, blockIndex := range fileBlockAddresses {
		var fileBlock Structs.FileBlock
		if err := Utils.ReadObject(dirTree.File, &fileBlock, int64(dirTree.SuperBlock.BlockStart+blockIndex*int32(binary.Size(Structs.FileBlock{})))); err != nil {
			return "", err
		}
		blockContent := strings.TrimRight(string(fileBlock.Content[:]), "\x00")

		sb.WriteString(fmt.Sprintf("\n    fileBlock%d [\n", blockIndex))
		sb.WriteString(fmt.Sprintf("        label=<\n"))
		sb.WriteString(fmt.Sprintf("            <TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\">\n"))
		sb.WriteString(fmt.Sprintf("                <TR><TD BGCOLOR=\"#ffe0b3\"><B>File Block %d</B></TD></TR>\n", blockIndex))
		sb.WriteString(fmt.Sprintf("                <TR><TD>%s</TD></TR>\n", blockContent))
		sb.WriteString(fmt.Sprintf("            </TABLE>\n"))
		sb.WriteString(fmt.Sprintf("        >\n"))
		sb.WriteString(fmt.Sprintf("    ];\n"))
	}

	for _, blockIndex := range folderBlockAddresses {
		var folderBlock Structs.FolderBlock
		if err := Utils.ReadObject(dirTree.File, &folderBlock, int64(dirTree.SuperBlock.BlockStart+blockIndex*int32(binary.Size(Structs.FolderBlock{})))); err != nil {
			return "", err
		}

		sb.WriteString(fmt.Sprintf("\n    folderBlock%d [\n", blockIndex))
		sb.WriteString(fmt.Sprintf("        label=<\n"))
		sb.WriteString(fmt.Sprintf("            <TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\">\n"))
		sb.WriteString(fmt.Sprintf("                <TR><TD COLSPAN=\"2\" BGCOLOR=\"#b3ffcc\"><B>Folder Block %d</B></TD></TR>\n", blockIndex))
		for j, content := range folderBlock.Content {
			blockName := strings.TrimRight(string(content.Name[:]), "\x00")
			sb.WriteString(fmt.Sprintf("                <TR><TD >%s</TD><TD PORT=\"ib%d\">%d</TD></TR>\n", blockName, j, content.Inode))
			if content.Inode == -1 {
				continue
			}
			if blockName == "." || blockName == ".." {
				continue
			}
			connections = append(connections, Connection{
				Type:  DtoI,
				Inode: content.Inode,
				Block: blockIndex,
				From:  int32(j),
			})
		}
		sb.WriteString(fmt.Sprintf("            </TABLE>\n"))
		sb.WriteString(fmt.Sprintf("        >\n"))
		sb.WriteString(fmt.Sprintf("    ];\n"))

	}

	for _, connection := range connections {
		switch connection.Type {
		case ItoF:
			sb.WriteString(fmt.Sprintf("    inode%d:ib%d -> fileBlock%d\n", connection.Inode, connection.From, connection.Block))
		case ItoD:
			sb.WriteString(fmt.Sprintf("    inode%d:ib%d -> folderBlock%d\n", connection.Inode, connection.From, connection.Block))
		case DtoI:
			sb.WriteString(fmt.Sprintf("    folderBlock%d:ib%d -> inode%d\n", connection.Block, connection.From, connection.Inode))
		}
	}

	sb.WriteString("}\n")
	return sb.String(), nil
}

func (r *Rep) createSbRep(dirTree *Structs.DirTree) (string, error) {
	var sb strings.Builder

	sb.WriteString("digraph SuperBlockReport {\n")
	sb.WriteString("    rankdir=TB;\n")
	sb.WriteString("    node [shape=plaintext];\n\n")
	sb.WriteString("    superblock [\n")
	sb.WriteString("        label=<\n")
	sb.WriteString("            <TABLE BORDER=\"1\" CELLBORDER=\"1\" CELLSPACING=\"0\" CELLPADDING=\"5\">\n")
	sb.WriteString("                <TR>\n")
	sb.WriteString("                    <TD COLSPAN=\"2\" BGCOLOR=\"#6bbf6b\">\n")
	sb.WriteString("                        <FONT COLOR=\"white\"><B>Reporte de SUPERBLOQUE</B></FONT>\n")
	sb.WriteString("                    </TD>\n")
	sb.WriteString("                </TR>\n")

	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#c8ecc8\"><TD><B>FilesystemType</B></TD><TD>%d</TD></TR>\n", dirTree.SuperBlock.FilesystemType))
	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#d2f0d2\"><TD><B>InodesCount</B></TD><TD>%d</TD></TR>\n", dirTree.SuperBlock.InodesCount))
	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#c8ecc8\"><TD><B>BlocksCount</B></TD><TD>%d</TD></TR>\n", dirTree.SuperBlock.BlocksCount))
	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#d2f0d2\"><TD><B>FreeBlocksCount</B></TD><TD>%d</TD></TR>\n", dirTree.SuperBlock.FreeBlocksCount))
	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#c8ecc8\"><TD><B>FreeInodesCount</B></TD><TD>%d</TD></TR>\n", dirTree.SuperBlock.FreeInodesCount))
	// Convert the byte arrays to string
	time := strings.TrimRight(string(dirTree.SuperBlock.MTime[:]), "\x00")
	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#d2f0d2\"><TD><B>MTime</B></TD><TD>%s</TD></TR>\n", time))
	time = strings.TrimRight(string(dirTree.SuperBlock.UMTime[:]), "\x00")
	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#c8ecc8\"><TD><B>UMTime</B></TD><TD>%s</TD></TR>\n", time))
	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#d2f0d2\"><TD><B>MntCount</B></TD><TD>%d</TD></TR>\n", dirTree.SuperBlock.MntCount))
	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#c8ecc8\"><TD><B>Magic</B></TD><TD>%d</TD></TR>\n", dirTree.SuperBlock.Magic))
	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#d2f0d2\"><TD><B>InodeSize</B></TD><TD>%d</TD></TR>\n", dirTree.SuperBlock.InodeSize))
	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#c8ecc8\"><TD><B>BlockSize</B></TD><TD>%d</TD></TR>\n", dirTree.SuperBlock.BlockSize))
	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#d2f0d2\"><TD><B>FirstInode</B></TD><TD>%d</TD></TR>\n", dirTree.SuperBlock.FirstInode))
	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#c8ecc8\"><TD><B>FirstBlock</B></TD><TD>%d</TD></TR>\n", dirTree.SuperBlock.FirstBlock))
	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#d2f0d2\"><TD><B>BmInodeStart</B></TD><TD>%d</TD></TR>\n", dirTree.SuperBlock.BmInodeStart))
	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#c8ecc8\"><TD><B>BmBlockStart</B></TD><TD>%d</TD></TR>\n", dirTree.SuperBlock.BmBlockStart))
	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#d2f0d2\"><TD><B>InodeStart</B></TD><TD>%d</TD></TR>\n", dirTree.SuperBlock.InodeStart))
	sb.WriteString(fmt.Sprintf("                <TR BGCOLOR=\"#c8ecc8\"><TD><B>BlockStart</B></TD><TD>%d</TD></TR>\n", dirTree.SuperBlock.BlockStart))

	sb.WriteString("                <TR>\n")
	sb.WriteString("                    <TD COLSPAN=\"2\" BGCOLOR=\"#6bbf6b\">\n")
	sb.WriteString("                        <FONT COLOR=\"white\"><I>Fin del Reporte</I></FONT>\n")
	sb.WriteString("                    </TD>\n")
	sb.WriteString("                </TR>\n")
	sb.WriteString("            </TABLE>\n")
	sb.WriteString("        >\n")
	sb.WriteString("    ];\n")
	sb.WriteString("}\n")

	return sb.String(), nil
}

func (r *Rep) createFileRep(dirTree *Structs.DirTree) (string, error) {
	_, inode, err := dirTree.GetInodeByPath(r.Params["pfl"])
	if err != nil {
		return "", err
	}
	var content string
	content, err = dirTree.GetFileContentByInode(inode)
	if err != nil {
		return "", err
	}
	fileName := filepath.Base(r.Params["pfl"])

	var sb strings.Builder

	// Start the DOT graph.
	sb.WriteString("digraph FileReport {\n")
	sb.WriteString("    rankdir=TB;\n")
	sb.WriteString("    node [shape=plaintext];\n")
	sb.WriteString("\n")
	sb.WriteString("    file_report [\n")
	sb.WriteString("        label=<\n")
	sb.WriteString("            <TABLE BORDER=\"1\" CELLBORDER=\"1\" CELLSPACING=\"0\" CELLPADDING=\"5\">\n")
	// Header row with the file name.
	sb.WriteString(fmt.Sprintf("                <TR><TD COLSPAN=\"1\" BGCOLOR=\"#6bbf6b\"><FONT COLOR=\"white\"><B>%s</B></FONT></TD></TR>\n", fileName))

	// Split the content into 64-character lines.
	for i := 0; i < len(content); i += 64 {
		end := i + 64
		if end > len(content) {
			end = len(content)
		}
		line := content[i:end]
		sb.WriteString(fmt.Sprintf("                <TR><TD>%s</TD></TR>\n", line))
	}

	sb.WriteString("            </TABLE>\n")
	sb.WriteString("        >\n")
	sb.WriteString("    ];\n")
	sb.WriteString("}\n")

	return sb.String(), nil
}

func (r *Rep) createLsRep(dirTree *Structs.DirTree) (string, error) {
	_, inode, err := dirTree.GetInodeByPath(r.Params["pfl"])
	if err != nil {
		return "", err
	}

	if inode.Type != [1]byte{0} {
		return "", errors.New("Inode is not a directory")
	}

	var folderBlocks []Structs.FolderBlock
	for _, block := range inode.IBlock {
		if block == -1 {
			continue
		}
		var folderBlock Structs.FolderBlock
		if err := Utils.ReadObject(dirTree.File, &folderBlock, int64(dirTree.SuperBlock.BlockStart+block*int32(binary.Size(Structs.FolderBlock{})))); err != nil {
			return "", err
		}
		folderBlocks = append(folderBlocks, folderBlock)
	}

	var sb strings.Builder

	// Start the DOT graph.
	sb.WriteString("digraph FileTable {\n")
	sb.WriteString("    rankdir=TB;\n")
	sb.WriteString("    node [shape=plaintext];\n")
	sb.WriteString("\n")
	sb.WriteString("    ls_report [\n")
	sb.WriteString("        label=<\n")
	sb.WriteString("            <TABLE BORDER=\"1\" CELLBORDER=\"1\" CELLSPACING=\"0\" CELLPADDING=\"5\">\n")
	// Header row with the file name.
	sb.WriteString(fmt.Sprintf("                <TR>\n"))
	sb.WriteString(fmt.Sprintf("                    <TD><B>Permisos</B></TD>\n"))
	sb.WriteString(fmt.Sprintf("                    <TD><B>Owner</B></TD>\n"))
	sb.WriteString(fmt.Sprintf("                    <TD><B>Grupo</B></TD>\n"))
	sb.WriteString(fmt.Sprintf("                    <TD><B>Size (en Bytes)</B></TD>\n"))
	sb.WriteString(fmt.Sprintf("                    <TD><B>Fecha</B></TD>\n"))
	sb.WriteString(fmt.Sprintf("                    <TD><B>Hora</B></TD>\n"))
	sb.WriteString(fmt.Sprintf("                    <TD><B>Tipo</B></TD>\n"))
	sb.WriteString(fmt.Sprintf("                    <TD><B>Name</B></TD>\n"))
	sb.WriteString(fmt.Sprintf("                </TR>\n"))

	for _, block := range folderBlocks {
		for _, content := range block.Content {
			if content.Inode == -1 {
				continue
			}
			fileName := strings.TrimRight(string(content.Name[:]), "\x00")
			if fileName == "." || fileName == ".." {
				continue
			}

			var contentInode Structs.Inode
			if err := Utils.ReadObject(dirTree.File, &contentInode, int64(dirTree.SuperBlock.InodeStart+content.Inode*int32(binary.Size(Structs.Inode{})))); err != nil {
				return "", err
			}

			sb.WriteString(fmt.Sprintf("                <TR>\n"))
			sb.WriteString(fmt.Sprintf("                    <TD>-%s</TD>\n", r.getPermString(contentInode.Perm)))

			sb.WriteString(fmt.Sprintf("                    <TD>%s</TD>\n", r.getOwnerNameByGID(contentInode.GID, dirTree)))
			sb.WriteString(fmt.Sprintf("                    <TD>%s</TD>\n", r.getOwnerNameByUID(contentInode.UID, dirTree)))
			sb.WriteString(fmt.Sprintf("                    <TD>%d</TD>\n", contentInode.Size))

			time := strings.TrimRight(string(contentInode.MTime[:]), "\x00")
			date := strings.Split(time, " ")[0]
			hour := strings.Split(time, " ")[1]
			sb.WriteString(fmt.Sprintf("                    <TD>%s</TD>\n", date))
			sb.WriteString(fmt.Sprintf("                    <TD>%s</TD>\n", hour))

			var fileType string
			if contentInode.Type == [1]byte{1} {
				fileType = "File"
			} else {
				fileType = "Directory"
			}
			sb.WriteString(fmt.Sprintf("                    <TD>%s</TD>\n", fileType))
			sb.WriteString(fmt.Sprintf("                    <TD>%s</TD>\n", fileName))
			sb.WriteString(fmt.Sprintf("                </TR>\n"))
		}

	}

	sb.WriteString("            </TABLE>\n")
	sb.WriteString("        >\n")
	sb.WriteString("    ];\n")
	sb.WriteString("}\n")

	return sb.String(), nil
}

func (r *Rep) getPermString(perms [3]byte) string {
	var permString string
	for i := 0; i < 3; i++ {
		read, write, execute := env.CalculatePermissions(perms[i])
		if read {
			permString += "r"
		} else {
			permString += "-"
		}
		if write {
			permString += "w"
		} else {
			permString += "-"
		}
		if execute {
			permString += "x"
		} else {
			permString += "-"
		}
	}

	return permString
}

func (r *Rep) getOwnerNameByGID(GID int32, dirTree *Structs.DirTree) string {
	_, usersInode, err := dirTree.GetInodeByPath("/users.txt")
	if err != nil {
		return "error"
	}

	var content string
	content, err = dirTree.GetFileContentByInode(usersInode)
	if err != nil {
		return "error"
	}

	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1]
	for _, line := range lines {
		// get the coma separated values from the line and get the id
		words := strings.Split(line, ",")
		// if line is of type group and of length then is not a user
		if len(words) == 3 && words[1] == "G" {
			id, _ := strconv.Atoi(words[0])
			if int32(id) == GID {
				return words[2]
			}
		}
	}
	return "error"
}

func (r *Rep) getOwnerNameByUID(UID int32, dirTree *Structs.DirTree) string {
	_, usersInode, err := dirTree.GetInodeByPath("/users.txt")
	if err != nil {
		return "error"
	}

	var content string
	content, err = dirTree.GetFileContentByInode(usersInode)
	if err != nil {
		return "error"
	}

	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1]
	for _, line := range lines {
		// get the coma separated values from the line and get the id
		words := strings.Split(line, ",")
		// if line is of type group and of length then is not a user
		if len(words) == 5 && words[1] == "U" {
			id, _ := strconv.Atoi(words[0])
			if int32(id) == UID {
				return words[3]
			}
		}
	}
	return "error"
}

func (r *Rep) createInodeRep(dirTree *Structs.DirTree) (string, error) {
	var sb strings.Builder

	sb.WriteString("digraph FileSystem {\n")
	//sb.WriteString("    rankdir=LR;\n")
	//sb.WriteString("    splines=false;\n")
	sb.WriteString("    node [shape=plaintext, style=\"filled,rounded\", fillcolor=\"#f0f8ff\"];\n")

	// Loop through the BitMap of inodes
	for i, bit := range dirTree.InodeBitMap {
		index := int32(i)
		if bit != 1 {
			continue
		}

		var bitMapInode Structs.Inode
		if err := Utils.ReadObject(dirTree.File, &bitMapInode, int64(dirTree.SuperBlock.InodeStart+index*int32(binary.Size(Structs.Inode{})))); err != nil {
			return "", err
		}

		sb.WriteString(fmt.Sprintf("\n    inode%d [\n", index))
		sb.WriteString(fmt.Sprintf("        label=<\n"))
		sb.WriteString(fmt.Sprintf("        <TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\">\n"))
		sb.WriteString(fmt.Sprintf("            <TR><TD COLSPAN=\"2\" BGCOLOR=\"#cce5ff\"><B>Inode%d</B></TD></TR>\n", index))
		sb.WriteString(fmt.Sprintf("            <TR><TD>UID</TD><TD>%d</TD></TR>\n", bitMapInode.UID))
		sb.WriteString(fmt.Sprintf("            <TR><TD>GID</TD><TD>%d</TD></TR>\n", bitMapInode.GID))
		sb.WriteString(fmt.Sprintf("            <TR><TD>Size</TD><TD>%d</TD></TR>\n", bitMapInode.Size))
		time := strings.TrimRight(string(bitMapInode.ATime[:]), "\x00")
		sb.WriteString(fmt.Sprintf("            <TR><TD>ATime</TD><TD>%s</TD></TR>\n", time))
		time = strings.TrimRight(string(bitMapInode.CTime[:]), "\x00")
		sb.WriteString(fmt.Sprintf("            <TR><TD>CTime</TD><TD>%s</TD></TR>\n", time))
		time = strings.TrimRight(string(bitMapInode.MTime[:]), "\x00")
		sb.WriteString(fmt.Sprintf("            <TR><TD>MTime</TD><TD>%s</TD></TR>\n", time))
		sb.WriteString(fmt.Sprintf("            <TR>\n"))
		sb.WriteString(fmt.Sprintf("                <TD>IBlock</TD>\n"))
		sb.WriteString(fmt.Sprintf("                <TD>\n"))
		sb.WriteString(fmt.Sprintf("                    <TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\">\n"))
		for j, addr := range bitMapInode.IBlock {
			sb.WriteString(fmt.Sprintf("                        <TR><TD PORT=\"ib%d\">block_%d: %d</TD></TR>\n", j, j, addr))
			if addr == -1 {
				continue
			}
		}
		sb.WriteString(fmt.Sprintf("                    </TABLE>\n"))
		sb.WriteString(fmt.Sprintf("                </TD>\n"))
		sb.WriteString(fmt.Sprintf("            </TR>\n"))
		if bitMapInode.Type == [1]byte{0} {
			sb.WriteString(fmt.Sprintf("            <TR><TD>Type</TD><TD>%d</TD></TR>\n", 0))
		} else {
			sb.WriteString(fmt.Sprintf("            <TR><TD>Type</TD><TD>%d</TD></TR>\n", 1))
		}
		sb.WriteString(fmt.Sprintf("            <TR><TD>Perm</TD><TD>%s</TD></TR>\n", string(bitMapInode.Perm[:])))
		sb.WriteString(fmt.Sprintf("        </TABLE>\n"))

		sb.WriteString(fmt.Sprintf("        >\n"))
		sb.WriteString(fmt.Sprintf("    ];\n"))
	}

	sb.WriteString("}\n")
	return sb.String(), nil
}

func (r *Rep) createBMRep(bitmap []byte, reportName string) (string, error) {
	var sb strings.Builder

	sb.WriteString("digraph BitmapReport {\n")
	sb.WriteString("    rankdir=TB;\n")
	sb.WriteString("    node [shape=plaintext];\n\n")
	sb.WriteString(fmt.Sprintf("    %s [\n", reportName))
	sb.WriteString("        label=<\n")
	sb.WriteString("            <TABLE BORDER=\"1\" CELLBORDER=\"1\" CELLSPACING=\"0\" CELLPADDING=\"5\">\n")

	// Iterate through the bitmap and create rows of 20 columns
	for i := 0; i < len(bitmap); i++ {
		// If it's the beginning of a new row, open a TR tag.
		if i%100 == 0 {
			sb.WriteString("                <TR>\n")
		}

		// Add a table cell with the value.
		sb.WriteString(fmt.Sprintf("                    <TD>%d</TD>\n", bitmap[i]))

		// If it's the last column in the row, close the TR tag.
		if i%100 == 99 {
			sb.WriteString("                </TR>\n")
		}
	}

	// If the bitmap length is not a multiple of 20, close the last row.
	if len(bitmap)%100 != 0 {
		sb.WriteString("                </TR>\n")
	}

	sb.WriteString("            </TABLE>\n")
	sb.WriteString("        >\n")
	sb.WriteString("    ];\n")
	sb.WriteString("}\n")

	return sb.String(), nil
}
