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
		//r.createInodeRep()
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
		//r.createBMInodeRep()
		if err != nil {
			r.AppendError(err.Error())
			return
		}
	case "bm_block":
		//r.createBMBlockRep()
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
		//r.createSbRep()
		if err != nil {
			r.AppendError(err.Error())
			return
		}
	case "file":
		//r.createFileRep()
		if err != nil {
			r.AppendError(err.Error())
			return
		}
	case "ls":
		//r.createLsRep()
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
