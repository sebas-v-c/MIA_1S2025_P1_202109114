package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Structs"
	"backend/Classes/Utils"
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
	//var dirTree = Structs.NewDirTree(superBlock, file)
	var repContent string

	switch r.Params["name"] {
	case "mbr":
		repContent, err = r.createMBRRep(file)
		if err != nil {
			r.AppendError(err.Error())
			return
		}
	case "disk":
		//r.createDiskRep()
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
		//r.createTreeRep()
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
		sb.WriteString(fmt.Sprintf("                <tr><td><b>part_start</b></td><td>%d</td></tr>\n", part.Start))
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

				ebrName := strings.TrimRight(string(ebr.Name[:]), "\x00")
				sb.WriteString("                <tr><td colspan=\"2\" bgcolor=\"#FF9999\"><b>Particion Logica</b></td></tr>\n")
				sb.WriteString(fmt.Sprintf("                <tr><td><b>part_mount</b></td><td>%c</td></tr>\n", ebr.Mount))
				sb.WriteString(fmt.Sprintf("                <tr><td><b>part_next</b></td><td>%d</td></tr>\n", ebr.Next))
				sb.WriteString(fmt.Sprintf("                <tr><td><b>part_fit</b></td><td>%c</td></tr>\n", ebr.Fit))
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
