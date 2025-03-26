// Code generated from Parser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Parser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import (
	commands "backend/Classes/Commands"
	interfaces "backend/Classes/Interfaces"
	"strings"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ParserParser struct {
	*antlr.BaseParser
}

var ParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func parserParserInit() {
	staticData := &ParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'mkdisk'", "'rmdisk'", "'fdisk'", "'mount'", "'mounted'", "'mkfs'",
		"'cat'", "'login'", "'logout'", "'mkgrp'", "'rmgrp'", "'mkusr'", "'rmusr'",
		"'chgrp'", "'mkfile'", "'mkdir'", "'rep'", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"'full'", "", "", "", "", "'='", "", "", "'\\n'",
	}
	staticData.SymbolicNames = []string{
		"", "RW_mkdisk", "RW_rmdisk", "RW_fdisk", "RW_mount", "RW_mounted",
		"RW_mkfs", "RW_cat", "RW_login", "RW_logout", "RW_mkgrp", "RW_rmgrp",
		"RW_mkusr", "RW_rmusr", "RW_chgrp", "RW_mkfile", "RW_mkdir", "RW_rep",
		"RW_size", "RW_fit", "RW_unit", "RW_driveletter", "RW_name", "RW_type",
		"RW_delete", "RW_add", "RW_id", "RW_fs", "RW_user", "RW_pass", "RW_grp",
		"RW_path", "RW_r", "RW_cont", "RW_fileN", "RW_destino", "RW_ugo", "RW_ruta",
		"TK_fit", "TK_unit", "TK_type", "TK_fs", "TK_ftype", "TK_number", "TK_id",
		"TK_ext", "TK_path", "TK_equ", "TK_sym", "COMMENTARY", "NEWLINE", "UNUSED_",
	}
	staticData.RuleNames = []string{
		"init", "commands", "command", "mkdisk", "mkdiskparams", "mkdiskparam",
		"rmdisk", "fdisk", "fdiskparams", "fdiskparam", "mount", "mountparams",
		"mountparam", "mounted", "mkfs", "mkfsparams", "mkfsparam", "cat", "catparams",
		"catparam", "login", "loginparams", "loginparam", "logout", "mkgrp",
		"mkgrpparams", "mkgrpparam", "rmgrp", "rmgrpparams", "rmgrpparam", "mkusr",
		"mkusrparams", "mkusrparam",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 51, 463, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 73, 8, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 83, 8, 1, 10, 1, 12,
		1, 86, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 124, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3,
		132, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 142, 8,
		4, 10, 4, 12, 4, 145, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 163, 8, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 172, 8, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 3, 7, 180, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 5, 8, 190, 8, 8, 10, 8, 12, 8, 193, 9, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 219, 8,
		9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 227, 8, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 237, 8, 11, 10,
		11, 12, 11, 240, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 3, 12, 250, 8, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 3, 14, 261, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 5, 15, 271, 8, 15, 10, 15, 12, 15, 274, 9, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 284, 8, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 292, 8, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 302, 8, 18, 10, 18,
		12, 18, 305, 9, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 3, 20, 318, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 328, 8, 21, 10, 21, 12, 21, 331, 9,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 349, 8, 22, 1, 23, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 360, 8, 24,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 370, 8,
		25, 10, 25, 12, 25, 373, 9, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 3, 26, 383, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 3, 27, 391, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 5, 28, 401, 8, 28, 10, 28, 12, 28, 404, 9, 28, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 414, 8, 29, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 422, 8, 30, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 432, 8, 31, 10, 31, 12, 31, 435,
		9, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 461, 8, 32, 1, 32, 0, 10, 2, 8, 16,
		22, 30, 36, 42, 50, 56, 62, 33, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
		58, 60, 62, 64, 0, 0, 481, 0, 72, 1, 0, 0, 0, 2, 74, 1, 0, 0, 0, 4, 123,
		1, 0, 0, 0, 6, 131, 1, 0, 0, 0, 8, 133, 1, 0, 0, 0, 10, 162, 1, 0, 0, 0,
		12, 171, 1, 0, 0, 0, 14, 179, 1, 0, 0, 0, 16, 181, 1, 0, 0, 0, 18, 218,
		1, 0, 0, 0, 20, 226, 1, 0, 0, 0, 22, 228, 1, 0, 0, 0, 24, 249, 1, 0, 0,
		0, 26, 251, 1, 0, 0, 0, 28, 260, 1, 0, 0, 0, 30, 262, 1, 0, 0, 0, 32, 283,
		1, 0, 0, 0, 34, 291, 1, 0, 0, 0, 36, 293, 1, 0, 0, 0, 38, 306, 1, 0, 0,
		0, 40, 317, 1, 0, 0, 0, 42, 319, 1, 0, 0, 0, 44, 348, 1, 0, 0, 0, 46, 350,
		1, 0, 0, 0, 48, 359, 1, 0, 0, 0, 50, 361, 1, 0, 0, 0, 52, 382, 1, 0, 0,
		0, 54, 390, 1, 0, 0, 0, 56, 392, 1, 0, 0, 0, 58, 413, 1, 0, 0, 0, 60, 421,
		1, 0, 0, 0, 62, 423, 1, 0, 0, 0, 64, 460, 1, 0, 0, 0, 66, 67, 3, 2, 1,
		0, 67, 68, 5, 0, 0, 1, 68, 69, 6, 0, -1, 0, 69, 73, 1, 0, 0, 0, 70, 71,
		5, 0, 0, 1, 71, 73, 6, 0, -1, 0, 72, 66, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0,
		73, 1, 1, 0, 0, 0, 74, 75, 6, 1, -1, 0, 75, 76, 3, 4, 2, 0, 76, 77, 6,
		1, -1, 0, 77, 84, 1, 0, 0, 0, 78, 79, 10, 2, 0, 0, 79, 80, 3, 4, 2, 0,
		80, 81, 6, 1, -1, 0, 81, 83, 1, 0, 0, 0, 82, 78, 1, 0, 0, 0, 83, 86, 1,
		0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 3, 1, 0, 0, 0, 86,
		84, 1, 0, 0, 0, 87, 88, 3, 6, 3, 0, 88, 89, 6, 2, -1, 0, 89, 124, 1, 0,
		0, 0, 90, 91, 3, 12, 6, 0, 91, 92, 6, 2, -1, 0, 92, 124, 1, 0, 0, 0, 93,
		94, 3, 14, 7, 0, 94, 95, 6, 2, -1, 0, 95, 124, 1, 0, 0, 0, 96, 97, 3, 20,
		10, 0, 97, 98, 6, 2, -1, 0, 98, 124, 1, 0, 0, 0, 99, 100, 3, 26, 13, 0,
		100, 101, 6, 2, -1, 0, 101, 124, 1, 0, 0, 0, 102, 103, 3, 28, 14, 0, 103,
		104, 6, 2, -1, 0, 104, 124, 1, 0, 0, 0, 105, 106, 3, 34, 17, 0, 106, 107,
		6, 2, -1, 0, 107, 124, 1, 0, 0, 0, 108, 109, 3, 40, 20, 0, 109, 110, 6,
		2, -1, 0, 110, 124, 1, 0, 0, 0, 111, 112, 3, 46, 23, 0, 112, 113, 6, 2,
		-1, 0, 113, 124, 1, 0, 0, 0, 114, 115, 3, 48, 24, 0, 115, 116, 6, 2, -1,
		0, 116, 124, 1, 0, 0, 0, 117, 118, 3, 54, 27, 0, 118, 119, 6, 2, -1, 0,
		119, 124, 1, 0, 0, 0, 120, 121, 3, 60, 30, 0, 121, 122, 6, 2, -1, 0, 122,
		124, 1, 0, 0, 0, 123, 87, 1, 0, 0, 0, 123, 90, 1, 0, 0, 0, 123, 93, 1,
		0, 0, 0, 123, 96, 1, 0, 0, 0, 123, 99, 1, 0, 0, 0, 123, 102, 1, 0, 0, 0,
		123, 105, 1, 0, 0, 0, 123, 108, 1, 0, 0, 0, 123, 111, 1, 0, 0, 0, 123,
		114, 1, 0, 0, 0, 123, 117, 1, 0, 0, 0, 123, 120, 1, 0, 0, 0, 124, 5, 1,
		0, 0, 0, 125, 126, 5, 1, 0, 0, 126, 127, 3, 8, 4, 0, 127, 128, 6, 3, -1,
		0, 128, 132, 1, 0, 0, 0, 129, 130, 5, 1, 0, 0, 130, 132, 6, 3, -1, 0, 131,
		125, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 132, 7, 1, 0, 0, 0, 133, 134, 6,
		4, -1, 0, 134, 135, 3, 10, 5, 0, 135, 136, 6, 4, -1, 0, 136, 143, 1, 0,
		0, 0, 137, 138, 10, 2, 0, 0, 138, 139, 3, 10, 5, 0, 139, 140, 6, 4, -1,
		0, 140, 142, 1, 0, 0, 0, 141, 137, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143,
		141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 9, 1, 0, 0, 0, 145, 143, 1,
		0, 0, 0, 146, 147, 5, 18, 0, 0, 147, 148, 5, 47, 0, 0, 148, 149, 5, 43,
		0, 0, 149, 163, 6, 5, -1, 0, 150, 151, 5, 19, 0, 0, 151, 152, 5, 47, 0,
		0, 152, 153, 5, 38, 0, 0, 153, 163, 6, 5, -1, 0, 154, 155, 5, 20, 0, 0,
		155, 156, 5, 47, 0, 0, 156, 157, 5, 39, 0, 0, 157, 163, 6, 5, -1, 0, 158,
		159, 5, 31, 0, 0, 159, 160, 5, 47, 0, 0, 160, 161, 5, 46, 0, 0, 161, 163,
		6, 5, -1, 0, 162, 146, 1, 0, 0, 0, 162, 150, 1, 0, 0, 0, 162, 154, 1, 0,
		0, 0, 162, 158, 1, 0, 0, 0, 163, 11, 1, 0, 0, 0, 164, 165, 5, 2, 0, 0,
		165, 166, 5, 31, 0, 0, 166, 167, 5, 47, 0, 0, 167, 168, 5, 46, 0, 0, 168,
		172, 6, 6, -1, 0, 169, 170, 5, 2, 0, 0, 170, 172, 6, 6, -1, 0, 171, 164,
		1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 172, 13, 1, 0, 0, 0, 173, 174, 5, 3,
		0, 0, 174, 175, 3, 16, 8, 0, 175, 176, 6, 7, -1, 0, 176, 180, 1, 0, 0,
		0, 177, 178, 5, 3, 0, 0, 178, 180, 6, 7, -1, 0, 179, 173, 1, 0, 0, 0, 179,
		177, 1, 0, 0, 0, 180, 15, 1, 0, 0, 0, 181, 182, 6, 8, -1, 0, 182, 183,
		3, 18, 9, 0, 183, 184, 6, 8, -1, 0, 184, 191, 1, 0, 0, 0, 185, 186, 10,
		2, 0, 0, 186, 187, 3, 18, 9, 0, 187, 188, 6, 8, -1, 0, 188, 190, 1, 0,
		0, 0, 189, 185, 1, 0, 0, 0, 190, 193, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0,
		191, 192, 1, 0, 0, 0, 192, 17, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 194, 195,
		5, 18, 0, 0, 195, 196, 5, 47, 0, 0, 196, 197, 5, 43, 0, 0, 197, 219, 6,
		9, -1, 0, 198, 199, 5, 20, 0, 0, 199, 200, 5, 47, 0, 0, 200, 201, 5, 39,
		0, 0, 201, 219, 6, 9, -1, 0, 202, 203, 5, 31, 0, 0, 203, 204, 5, 47, 0,
		0, 204, 205, 5, 46, 0, 0, 205, 219, 6, 9, -1, 0, 206, 207, 5, 23, 0, 0,
		207, 208, 5, 47, 0, 0, 208, 209, 5, 40, 0, 0, 209, 219, 6, 9, -1, 0, 210,
		211, 5, 19, 0, 0, 211, 212, 5, 47, 0, 0, 212, 213, 5, 38, 0, 0, 213, 219,
		6, 9, -1, 0, 214, 215, 5, 22, 0, 0, 215, 216, 5, 47, 0, 0, 216, 217, 5,
		44, 0, 0, 217, 219, 6, 9, -1, 0, 218, 194, 1, 0, 0, 0, 218, 198, 1, 0,
		0, 0, 218, 202, 1, 0, 0, 0, 218, 206, 1, 0, 0, 0, 218, 210, 1, 0, 0, 0,
		218, 214, 1, 0, 0, 0, 219, 19, 1, 0, 0, 0, 220, 221, 5, 4, 0, 0, 221, 222,
		3, 22, 11, 0, 222, 223, 6, 10, -1, 0, 223, 227, 1, 0, 0, 0, 224, 225, 5,
		4, 0, 0, 225, 227, 6, 10, -1, 0, 226, 220, 1, 0, 0, 0, 226, 224, 1, 0,
		0, 0, 227, 21, 1, 0, 0, 0, 228, 229, 6, 11, -1, 0, 229, 230, 3, 24, 12,
		0, 230, 231, 6, 11, -1, 0, 231, 238, 1, 0, 0, 0, 232, 233, 10, 2, 0, 0,
		233, 234, 3, 24, 12, 0, 234, 235, 6, 11, -1, 0, 235, 237, 1, 0, 0, 0, 236,
		232, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 238, 239,
		1, 0, 0, 0, 239, 23, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 241, 242, 5, 31,
		0, 0, 242, 243, 5, 47, 0, 0, 243, 244, 5, 46, 0, 0, 244, 250, 6, 12, -1,
		0, 245, 246, 5, 22, 0, 0, 246, 247, 5, 47, 0, 0, 247, 248, 5, 44, 0, 0,
		248, 250, 6, 12, -1, 0, 249, 241, 1, 0, 0, 0, 249, 245, 1, 0, 0, 0, 250,
		25, 1, 0, 0, 0, 251, 252, 5, 5, 0, 0, 252, 253, 6, 13, -1, 0, 253, 27,
		1, 0, 0, 0, 254, 255, 5, 6, 0, 0, 255, 256, 3, 30, 15, 0, 256, 257, 6,
		14, -1, 0, 257, 261, 1, 0, 0, 0, 258, 259, 5, 6, 0, 0, 259, 261, 6, 14,
		-1, 0, 260, 254, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 261, 29, 1, 0, 0, 0,
		262, 263, 6, 15, -1, 0, 263, 264, 3, 32, 16, 0, 264, 265, 6, 15, -1, 0,
		265, 272, 1, 0, 0, 0, 266, 267, 10, 2, 0, 0, 267, 268, 3, 32, 16, 0, 268,
		269, 6, 15, -1, 0, 269, 271, 1, 0, 0, 0, 270, 266, 1, 0, 0, 0, 271, 274,
		1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 31, 1, 0,
		0, 0, 274, 272, 1, 0, 0, 0, 275, 276, 5, 23, 0, 0, 276, 277, 5, 47, 0,
		0, 277, 278, 5, 42, 0, 0, 278, 284, 6, 16, -1, 0, 279, 280, 5, 26, 0, 0,
		280, 281, 5, 47, 0, 0, 281, 282, 5, 44, 0, 0, 282, 284, 6, 16, -1, 0, 283,
		275, 1, 0, 0, 0, 283, 279, 1, 0, 0, 0, 284, 33, 1, 0, 0, 0, 285, 286, 5,
		7, 0, 0, 286, 287, 3, 36, 18, 0, 287, 288, 6, 17, -1, 0, 288, 292, 1, 0,
		0, 0, 289, 290, 5, 7, 0, 0, 290, 292, 6, 17, -1, 0, 291, 285, 1, 0, 0,
		0, 291, 289, 1, 0, 0, 0, 292, 35, 1, 0, 0, 0, 293, 294, 6, 18, -1, 0, 294,
		295, 3, 38, 19, 0, 295, 296, 6, 18, -1, 0, 296, 303, 1, 0, 0, 0, 297, 298,
		10, 2, 0, 0, 298, 299, 3, 38, 19, 0, 299, 300, 6, 18, -1, 0, 300, 302,
		1, 0, 0, 0, 301, 297, 1, 0, 0, 0, 302, 305, 1, 0, 0, 0, 303, 301, 1, 0,
		0, 0, 303, 304, 1, 0, 0, 0, 304, 37, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0,
		306, 307, 5, 34, 0, 0, 307, 308, 5, 47, 0, 0, 308, 309, 5, 46, 0, 0, 309,
		310, 6, 19, -1, 0, 310, 39, 1, 0, 0, 0, 311, 312, 5, 8, 0, 0, 312, 313,
		3, 42, 21, 0, 313, 314, 6, 20, -1, 0, 314, 318, 1, 0, 0, 0, 315, 316, 5,
		8, 0, 0, 316, 318, 6, 20, -1, 0, 317, 311, 1, 0, 0, 0, 317, 315, 1, 0,
		0, 0, 318, 41, 1, 0, 0, 0, 319, 320, 6, 21, -1, 0, 320, 321, 3, 44, 22,
		0, 321, 322, 6, 21, -1, 0, 322, 329, 1, 0, 0, 0, 323, 324, 10, 2, 0, 0,
		324, 325, 3, 44, 22, 0, 325, 326, 6, 21, -1, 0, 326, 328, 1, 0, 0, 0, 327,
		323, 1, 0, 0, 0, 328, 331, 1, 0, 0, 0, 329, 327, 1, 0, 0, 0, 329, 330,
		1, 0, 0, 0, 330, 43, 1, 0, 0, 0, 331, 329, 1, 0, 0, 0, 332, 333, 5, 28,
		0, 0, 333, 334, 5, 47, 0, 0, 334, 335, 5, 44, 0, 0, 335, 349, 6, 22, -1,
		0, 336, 337, 5, 29, 0, 0, 337, 338, 5, 47, 0, 0, 338, 339, 5, 44, 0, 0,
		339, 349, 6, 22, -1, 0, 340, 341, 5, 29, 0, 0, 341, 342, 5, 47, 0, 0, 342,
		343, 5, 43, 0, 0, 343, 349, 6, 22, -1, 0, 344, 345, 5, 26, 0, 0, 345, 346,
		5, 47, 0, 0, 346, 347, 5, 44, 0, 0, 347, 349, 6, 22, -1, 0, 348, 332, 1,
		0, 0, 0, 348, 336, 1, 0, 0, 0, 348, 340, 1, 0, 0, 0, 348, 344, 1, 0, 0,
		0, 349, 45, 1, 0, 0, 0, 350, 351, 5, 9, 0, 0, 351, 352, 6, 23, -1, 0, 352,
		47, 1, 0, 0, 0, 353, 354, 5, 10, 0, 0, 354, 355, 3, 50, 25, 0, 355, 356,
		6, 24, -1, 0, 356, 360, 1, 0, 0, 0, 357, 358, 5, 10, 0, 0, 358, 360, 6,
		24, -1, 0, 359, 353, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 360, 49, 1, 0, 0,
		0, 361, 362, 6, 25, -1, 0, 362, 363, 3, 52, 26, 0, 363, 364, 6, 25, -1,
		0, 364, 371, 1, 0, 0, 0, 365, 366, 10, 2, 0, 0, 366, 367, 3, 52, 26, 0,
		367, 368, 6, 25, -1, 0, 368, 370, 1, 0, 0, 0, 369, 365, 1, 0, 0, 0, 370,
		373, 1, 0, 0, 0, 371, 369, 1, 0, 0, 0, 371, 372, 1, 0, 0, 0, 372, 51, 1,
		0, 0, 0, 373, 371, 1, 0, 0, 0, 374, 375, 5, 22, 0, 0, 375, 376, 5, 47,
		0, 0, 376, 377, 5, 44, 0, 0, 377, 383, 6, 26, -1, 0, 378, 379, 5, 22, 0,
		0, 379, 380, 5, 47, 0, 0, 380, 381, 5, 43, 0, 0, 381, 383, 6, 26, -1, 0,
		382, 374, 1, 0, 0, 0, 382, 378, 1, 0, 0, 0, 383, 53, 1, 0, 0, 0, 384, 385,
		5, 11, 0, 0, 385, 386, 3, 56, 28, 0, 386, 387, 6, 27, -1, 0, 387, 391,
		1, 0, 0, 0, 388, 389, 5, 11, 0, 0, 389, 391, 6, 27, -1, 0, 390, 384, 1,
		0, 0, 0, 390, 388, 1, 0, 0, 0, 391, 55, 1, 0, 0, 0, 392, 393, 6, 28, -1,
		0, 393, 394, 3, 58, 29, 0, 394, 395, 6, 28, -1, 0, 395, 402, 1, 0, 0, 0,
		396, 397, 10, 2, 0, 0, 397, 398, 3, 58, 29, 0, 398, 399, 6, 28, -1, 0,
		399, 401, 1, 0, 0, 0, 400, 396, 1, 0, 0, 0, 401, 404, 1, 0, 0, 0, 402,
		400, 1, 0, 0, 0, 402, 403, 1, 0, 0, 0, 403, 57, 1, 0, 0, 0, 404, 402, 1,
		0, 0, 0, 405, 406, 5, 22, 0, 0, 406, 407, 5, 47, 0, 0, 407, 408, 5, 44,
		0, 0, 408, 414, 6, 29, -1, 0, 409, 410, 5, 22, 0, 0, 410, 411, 5, 47, 0,
		0, 411, 412, 5, 43, 0, 0, 412, 414, 6, 29, -1, 0, 413, 405, 1, 0, 0, 0,
		413, 409, 1, 0, 0, 0, 414, 59, 1, 0, 0, 0, 415, 416, 5, 12, 0, 0, 416,
		417, 3, 62, 31, 0, 417, 418, 6, 30, -1, 0, 418, 422, 1, 0, 0, 0, 419, 420,
		5, 12, 0, 0, 420, 422, 6, 30, -1, 0, 421, 415, 1, 0, 0, 0, 421, 419, 1,
		0, 0, 0, 422, 61, 1, 0, 0, 0, 423, 424, 6, 31, -1, 0, 424, 425, 3, 64,
		32, 0, 425, 426, 6, 31, -1, 0, 426, 433, 1, 0, 0, 0, 427, 428, 10, 2, 0,
		0, 428, 429, 3, 64, 32, 0, 429, 430, 6, 31, -1, 0, 430, 432, 1, 0, 0, 0,
		431, 427, 1, 0, 0, 0, 432, 435, 1, 0, 0, 0, 433, 431, 1, 0, 0, 0, 433,
		434, 1, 0, 0, 0, 434, 63, 1, 0, 0, 0, 435, 433, 1, 0, 0, 0, 436, 437, 5,
		28, 0, 0, 437, 438, 5, 47, 0, 0, 438, 439, 5, 44, 0, 0, 439, 461, 6, 32,
		-1, 0, 440, 441, 5, 28, 0, 0, 441, 442, 5, 47, 0, 0, 442, 443, 5, 43, 0,
		0, 443, 461, 6, 32, -1, 0, 444, 445, 5, 29, 0, 0, 445, 446, 5, 47, 0, 0,
		446, 447, 5, 44, 0, 0, 447, 461, 6, 32, -1, 0, 448, 449, 5, 29, 0, 0, 449,
		450, 5, 47, 0, 0, 450, 451, 5, 43, 0, 0, 451, 461, 6, 32, -1, 0, 452, 453,
		5, 30, 0, 0, 453, 454, 5, 47, 0, 0, 454, 455, 5, 44, 0, 0, 455, 461, 6,
		32, -1, 0, 456, 457, 5, 30, 0, 0, 457, 458, 5, 47, 0, 0, 458, 459, 5, 43,
		0, 0, 459, 461, 6, 32, -1, 0, 460, 436, 1, 0, 0, 0, 460, 440, 1, 0, 0,
		0, 460, 444, 1, 0, 0, 0, 460, 448, 1, 0, 0, 0, 460, 452, 1, 0, 0, 0, 460,
		456, 1, 0, 0, 0, 461, 65, 1, 0, 0, 0, 30, 72, 84, 123, 131, 143, 162, 171,
		179, 191, 218, 226, 238, 249, 260, 272, 283, 291, 303, 317, 329, 348, 359,
		371, 382, 390, 402, 413, 421, 433, 460,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ParserParserInit initializes any static state used to implement ParserParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewParserParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ParserParserInit() {
	staticData := &ParserParserStaticData
	staticData.once.Do(parserParserInit)
}

// NewParserParser produces a new parser instance for the optional input antlr.TokenStream.
func NewParserParser(input antlr.TokenStream) *ParserParser {
	ParserParserInit()
	this := new(ParserParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Parser.g4"

	return this
}

// ParserParser tokens.
const (
	ParserParserEOF            = antlr.TokenEOF
	ParserParserRW_mkdisk      = 1
	ParserParserRW_rmdisk      = 2
	ParserParserRW_fdisk       = 3
	ParserParserRW_mount       = 4
	ParserParserRW_mounted     = 5
	ParserParserRW_mkfs        = 6
	ParserParserRW_cat         = 7
	ParserParserRW_login       = 8
	ParserParserRW_logout      = 9
	ParserParserRW_mkgrp       = 10
	ParserParserRW_rmgrp       = 11
	ParserParserRW_mkusr       = 12
	ParserParserRW_rmusr       = 13
	ParserParserRW_chgrp       = 14
	ParserParserRW_mkfile      = 15
	ParserParserRW_mkdir       = 16
	ParserParserRW_rep         = 17
	ParserParserRW_size        = 18
	ParserParserRW_fit         = 19
	ParserParserRW_unit        = 20
	ParserParserRW_driveletter = 21
	ParserParserRW_name        = 22
	ParserParserRW_type        = 23
	ParserParserRW_delete      = 24
	ParserParserRW_add         = 25
	ParserParserRW_id          = 26
	ParserParserRW_fs          = 27
	ParserParserRW_user        = 28
	ParserParserRW_pass        = 29
	ParserParserRW_grp         = 30
	ParserParserRW_path        = 31
	ParserParserRW_r           = 32
	ParserParserRW_cont        = 33
	ParserParserRW_fileN       = 34
	ParserParserRW_destino     = 35
	ParserParserRW_ugo         = 36
	ParserParserRW_ruta        = 37
	ParserParserTK_fit         = 38
	ParserParserTK_unit        = 39
	ParserParserTK_type        = 40
	ParserParserTK_fs          = 41
	ParserParserTK_ftype       = 42
	ParserParserTK_number      = 43
	ParserParserTK_id          = 44
	ParserParserTK_ext         = 45
	ParserParserTK_path        = 46
	ParserParserTK_equ         = 47
	ParserParserTK_sym         = 48
	ParserParserCOMMENTARY     = 49
	ParserParserNEWLINE        = 50
	ParserParserUNUSED_        = 51
)

// ParserParser rules.
const (
	ParserParserRULE_init         = 0
	ParserParserRULE_commands     = 1
	ParserParserRULE_command      = 2
	ParserParserRULE_mkdisk       = 3
	ParserParserRULE_mkdiskparams = 4
	ParserParserRULE_mkdiskparam  = 5
	ParserParserRULE_rmdisk       = 6
	ParserParserRULE_fdisk        = 7
	ParserParserRULE_fdiskparams  = 8
	ParserParserRULE_fdiskparam   = 9
	ParserParserRULE_mount        = 10
	ParserParserRULE_mountparams  = 11
	ParserParserRULE_mountparam   = 12
	ParserParserRULE_mounted      = 13
	ParserParserRULE_mkfs         = 14
	ParserParserRULE_mkfsparams   = 15
	ParserParserRULE_mkfsparam    = 16
	ParserParserRULE_cat          = 17
	ParserParserRULE_catparams    = 18
	ParserParserRULE_catparam     = 19
	ParserParserRULE_login        = 20
	ParserParserRULE_loginparams  = 21
	ParserParserRULE_loginparam   = 22
	ParserParserRULE_logout       = 23
	ParserParserRULE_mkgrp        = 24
	ParserParserRULE_mkgrpparams  = 25
	ParserParserRULE_mkgrpparam   = 26
	ParserParserRULE_rmgrp        = 27
	ParserParserRULE_rmgrpparams  = 28
	ParserParserRULE_rmgrpparam   = 29
	ParserParserRULE_mkusr        = 30
	ParserParserRULE_mkusrparams  = 31
	ParserParserRULE_mkusrparam   = 32
)

// IInitContext is an interface to support dynamic dispatch.
type IInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC returns the c rule contexts.
	GetC() ICommandsContext

	// SetC sets the c rule contexts.
	SetC(ICommandsContext)

	// GetResult returns the result attribute.
	GetResult() []interfaces.Command

	// SetResult sets the result attribute.
	SetResult([]interfaces.Command)

	// Getter signatures
	EOF() antlr.TerminalNode
	Commands() ICommandsContext

	// IsInitContext differentiates from other interfaces.
	IsInitContext()
}

type InitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []interfaces.Command
	c      ICommandsContext
}

func NewEmptyInitContext() *InitContext {
	var p = new(InitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_init
	return p
}

func InitEmptyInitContext(p *InitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_init
}

func (*InitContext) IsInitContext() {}

func NewInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitContext {
	var p = new(InitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_init

	return p
}

func (s *InitContext) GetParser() antlr.Parser { return s.parser }

func (s *InitContext) GetC() ICommandsContext { return s.c }

func (s *InitContext) SetC(v ICommandsContext) { s.c = v }

func (s *InitContext) GetResult() []interfaces.Command { return s.result }

func (s *InitContext) SetResult(v []interfaces.Command) { s.result = v }

func (s *InitContext) EOF() antlr.TerminalNode {
	return s.GetToken(ParserParserEOF, 0)
}

func (s *InitContext) Commands() ICommandsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandsContext)
}

func (s *InitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterInit(s)
	}
}

func (s *InitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitInit(s)
	}
}

func (p *ParserParser) Init() (localctx IInitContext) {
	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ParserParserRULE_init)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_mkdisk, ParserParserRW_rmdisk, ParserParserRW_fdisk, ParserParserRW_mount, ParserParserRW_mounted, ParserParserRW_mkfs, ParserParserRW_cat, ParserParserRW_login, ParserParserRW_logout, ParserParserRW_mkgrp, ParserParserRW_rmgrp, ParserParserRW_mkusr:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)

			var _x = p.commands(0)

			localctx.(*InitContext).c = _x
		}
		{
			p.SetState(67)
			p.Match(ParserParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*InitContext).result = localctx.(*InitContext).GetC().GetResult()

	case ParserParserEOF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.Match(ParserParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*InitContext).result = []interfaces.Command{}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandsContext is an interface to support dynamic dispatch.
type ICommandsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() ICommandsContext

	// GetC returns the c rule contexts.
	GetC() ICommandContext

	// SetL sets the l rule contexts.
	SetL(ICommandsContext)

	// SetC sets the c rule contexts.
	SetC(ICommandContext)

	// GetResult returns the result attribute.
	GetResult() []interfaces.Command

	// SetResult sets the result attribute.
	SetResult([]interfaces.Command)

	// Getter signatures
	Command() ICommandContext
	Commands() ICommandsContext

	// IsCommandsContext differentiates from other interfaces.
	IsCommandsContext()
}

type CommandsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []interfaces.Command
	l      ICommandsContext
	c      ICommandContext
}

func NewEmptyCommandsContext() *CommandsContext {
	var p = new(CommandsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_commands
	return p
}

func InitEmptyCommandsContext(p *CommandsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_commands
}

func (*CommandsContext) IsCommandsContext() {}

func NewCommandsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandsContext {
	var p = new(CommandsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_commands

	return p
}

func (s *CommandsContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandsContext) GetL() ICommandsContext { return s.l }

func (s *CommandsContext) GetC() ICommandContext { return s.c }

func (s *CommandsContext) SetL(v ICommandsContext) { s.l = v }

func (s *CommandsContext) SetC(v ICommandContext) { s.c = v }

func (s *CommandsContext) GetResult() []interfaces.Command { return s.result }

func (s *CommandsContext) SetResult(v []interfaces.Command) { s.result = v }

func (s *CommandsContext) Command() ICommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *CommandsContext) Commands() ICommandsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandsContext)
}

func (s *CommandsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterCommands(s)
	}
}

func (s *CommandsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitCommands(s)
	}
}

func (p *ParserParser) Commands() (localctx ICommandsContext) {
	return p.commands(0)
}

func (p *ParserParser) commands(_p int) (localctx ICommandsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewCommandsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICommandsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, ParserParserRULE_commands, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)

		var _x = p.Command()

		localctx.(*CommandsContext).c = _x
	}
	localctx.(*CommandsContext).result = []interfaces.Command{localctx.(*CommandsContext).GetC().GetResult()}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCommandsContext(p, _parentctx, _parentState)
			localctx.(*CommandsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_commands)
			p.SetState(78)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(79)

				var _x = p.Command()

				localctx.(*CommandsContext).c = _x
			}
			localctx.(*CommandsContext).SetResult(localctx.(*CommandsContext).GetL().GetResult())
			localctx.(*CommandsContext).result = append(localctx.(*CommandsContext).result, localctx.(*CommandsContext).GetC().GetResult())

		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC1 returns the c1 rule contexts.
	GetC1() IMkdiskContext

	// GetC2 returns the c2 rule contexts.
	GetC2() IRmdiskContext

	// GetC3 returns the c3 rule contexts.
	GetC3() IFdiskContext

	// GetC4 returns the c4 rule contexts.
	GetC4() IMountContext

	// GetC5 returns the c5 rule contexts.
	GetC5() IMountedContext

	// GetC6 returns the c6 rule contexts.
	GetC6() IMkfsContext

	// GetC7 returns the c7 rule contexts.
	GetC7() ICatContext

	// GetC8 returns the c8 rule contexts.
	GetC8() ILoginContext

	// GetC9 returns the c9 rule contexts.
	GetC9() ILogoutContext

	// GetC10 returns the c10 rule contexts.
	GetC10() IMkgrpContext

	// GetC11 returns the c11 rule contexts.
	GetC11() IRmgrpContext

	// GetC12 returns the c12 rule contexts.
	GetC12() IMkusrContext

	// SetC1 sets the c1 rule contexts.
	SetC1(IMkdiskContext)

	// SetC2 sets the c2 rule contexts.
	SetC2(IRmdiskContext)

	// SetC3 sets the c3 rule contexts.
	SetC3(IFdiskContext)

	// SetC4 sets the c4 rule contexts.
	SetC4(IMountContext)

	// SetC5 sets the c5 rule contexts.
	SetC5(IMountedContext)

	// SetC6 sets the c6 rule contexts.
	SetC6(IMkfsContext)

	// SetC7 sets the c7 rule contexts.
	SetC7(ICatContext)

	// SetC8 sets the c8 rule contexts.
	SetC8(ILoginContext)

	// SetC9 sets the c9 rule contexts.
	SetC9(ILogoutContext)

	// SetC10 sets the c10 rule contexts.
	SetC10(IMkgrpContext)

	// SetC11 sets the c11 rule contexts.
	SetC11(IRmgrpContext)

	// SetC12 sets the c12 rule contexts.
	SetC12(IMkusrContext)

	// GetResult returns the result attribute.
	GetResult() interfaces.Command

	// SetResult sets the result attribute.
	SetResult(interfaces.Command)

	// Getter signatures
	Mkdisk() IMkdiskContext
	Rmdisk() IRmdiskContext
	Fdisk() IFdiskContext
	Mount() IMountContext
	Mounted() IMountedContext
	Mkfs() IMkfsContext
	Cat() ICatContext
	Login() ILoginContext
	Logout() ILogoutContext
	Mkgrp() IMkgrpContext
	Rmgrp() IRmgrpContext
	Mkusr() IMkusrContext

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result interfaces.Command
	c1     IMkdiskContext
	c2     IRmdiskContext
	c3     IFdiskContext
	c4     IMountContext
	c5     IMountedContext
	c6     IMkfsContext
	c7     ICatContext
	c8     ILoginContext
	c9     ILogoutContext
	c10    IMkgrpContext
	c11    IRmgrpContext
	c12    IMkusrContext
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) GetC1() IMkdiskContext { return s.c1 }

func (s *CommandContext) GetC2() IRmdiskContext { return s.c2 }

func (s *CommandContext) GetC3() IFdiskContext { return s.c3 }

func (s *CommandContext) GetC4() IMountContext { return s.c4 }

func (s *CommandContext) GetC5() IMountedContext { return s.c5 }

func (s *CommandContext) GetC6() IMkfsContext { return s.c6 }

func (s *CommandContext) GetC7() ICatContext { return s.c7 }

func (s *CommandContext) GetC8() ILoginContext { return s.c8 }

func (s *CommandContext) GetC9() ILogoutContext { return s.c9 }

func (s *CommandContext) GetC10() IMkgrpContext { return s.c10 }

func (s *CommandContext) GetC11() IRmgrpContext { return s.c11 }

func (s *CommandContext) GetC12() IMkusrContext { return s.c12 }

func (s *CommandContext) SetC1(v IMkdiskContext) { s.c1 = v }

func (s *CommandContext) SetC2(v IRmdiskContext) { s.c2 = v }

func (s *CommandContext) SetC3(v IFdiskContext) { s.c3 = v }

func (s *CommandContext) SetC4(v IMountContext) { s.c4 = v }

func (s *CommandContext) SetC5(v IMountedContext) { s.c5 = v }

func (s *CommandContext) SetC6(v IMkfsContext) { s.c6 = v }

func (s *CommandContext) SetC7(v ICatContext) { s.c7 = v }

func (s *CommandContext) SetC8(v ILoginContext) { s.c8 = v }

func (s *CommandContext) SetC9(v ILogoutContext) { s.c9 = v }

func (s *CommandContext) SetC10(v IMkgrpContext) { s.c10 = v }

func (s *CommandContext) SetC11(v IRmgrpContext) { s.c11 = v }

func (s *CommandContext) SetC12(v IMkusrContext) { s.c12 = v }

func (s *CommandContext) GetResult() interfaces.Command { return s.result }

func (s *CommandContext) SetResult(v interfaces.Command) { s.result = v }

func (s *CommandContext) Mkdisk() IMkdiskContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdiskContext)
}

func (s *CommandContext) Rmdisk() IRmdiskContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmdiskContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmdiskContext)
}

func (s *CommandContext) Fdisk() IFdiskContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFdiskContext)
}

func (s *CommandContext) Mount() IMountContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountContext)
}

func (s *CommandContext) Mounted() IMountedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountedContext)
}

func (s *CommandContext) Mkfs() IMkfsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfsContext)
}

func (s *CommandContext) Cat() ICatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICatContext)
}

func (s *CommandContext) Login() ILoginContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoginContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoginContext)
}

func (s *CommandContext) Logout() ILogoutContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogoutContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogoutContext)
}

func (s *CommandContext) Mkgrp() IMkgrpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkgrpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkgrpContext)
}

func (s *CommandContext) Rmgrp() IRmgrpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmgrpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmgrpContext)
}

func (s *CommandContext) Mkusr() IMkusrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkusrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkusrContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *ParserParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ParserParserRULE_command)
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_mkdisk:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)

			var _x = p.Mkdisk()

			localctx.(*CommandContext).c1 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC1().GetResult()

	case ParserParserRW_rmdisk:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)

			var _x = p.Rmdisk()

			localctx.(*CommandContext).c2 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC2().GetResult()

	case ParserParserRW_fdisk:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)

			var _x = p.Fdisk()

			localctx.(*CommandContext).c3 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC3().GetResult()

	case ParserParserRW_mount:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(96)

			var _x = p.Mount()

			localctx.(*CommandContext).c4 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC4().GetResult()

	case ParserParserRW_mounted:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(99)

			var _x = p.Mounted()

			localctx.(*CommandContext).c5 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC5().GetResult()

	case ParserParserRW_mkfs:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(102)

			var _x = p.Mkfs()

			localctx.(*CommandContext).c6 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC6().GetResult()

	case ParserParserRW_cat:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(105)

			var _x = p.Cat()

			localctx.(*CommandContext).c7 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC7().GetResult()

	case ParserParserRW_login:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(108)

			var _x = p.Login()

			localctx.(*CommandContext).c8 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC8().GetResult()

	case ParserParserRW_logout:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(111)

			var _x = p.Logout()

			localctx.(*CommandContext).c9 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC9().GetResult()

	case ParserParserRW_mkgrp:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(114)

			var _x = p.Mkgrp()

			localctx.(*CommandContext).c10 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC10().GetResult()

	case ParserParserRW_rmgrp:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(117)

			var _x = p.Rmgrp()

			localctx.(*CommandContext).c11 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC11().GetResult()

	case ParserParserRW_mkusr:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(120)

			var _x = p.Mkusr()

			localctx.(*CommandContext).c12 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC12().GetResult()

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkdiskContext is an interface to support dynamic dispatch.
type IMkdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetM returns the m token.
	GetM() antlr.Token

	// SetM sets the m token.
	SetM(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IMkdiskparamsContext

	// SetP sets the p rule contexts.
	SetP(IMkdiskparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Mkdisk

	// SetResult sets the result attribute.
	SetResult(*commands.Mkdisk)

	// Getter signatures
	RW_mkdisk() antlr.TerminalNode
	Mkdiskparams() IMkdiskparamsContext

	// IsMkdiskContext differentiates from other interfaces.
	IsMkdiskContext()
}

type MkdiskContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Mkdisk
	m      antlr.Token
	p      IMkdiskparamsContext
}

func NewEmptyMkdiskContext() *MkdiskContext {
	var p = new(MkdiskContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdisk
	return p
}

func InitEmptyMkdiskContext(p *MkdiskContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdisk
}

func (*MkdiskContext) IsMkdiskContext() {}

func NewMkdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkdiskContext {
	var p = new(MkdiskContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkdisk

	return p
}

func (s *MkdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *MkdiskContext) GetM() antlr.Token { return s.m }

func (s *MkdiskContext) SetM(v antlr.Token) { s.m = v }

func (s *MkdiskContext) GetP() IMkdiskparamsContext { return s.p }

func (s *MkdiskContext) SetP(v IMkdiskparamsContext) { s.p = v }

func (s *MkdiskContext) GetResult() *commands.Mkdisk { return s.result }

func (s *MkdiskContext) SetResult(v *commands.Mkdisk) { s.result = v }

func (s *MkdiskContext) RW_mkdisk() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mkdisk, 0)
}

func (s *MkdiskContext) Mkdiskparams() IMkdiskparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdiskparamsContext)
}

func (s *MkdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkdisk(s)
	}
}

func (s *MkdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkdisk(s)
	}
}

func (p *ParserParser) Mkdisk() (localctx IMkdiskContext) {
	localctx = NewMkdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ParserParserRULE_mkdisk)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)

			var _m = p.Match(ParserParserRW_mkdisk)

			localctx.(*MkdiskContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)

			var _x = p.mkdiskparams(0)

			localctx.(*MkdiskContext).p = _x
		}
		localctx.(*MkdiskContext).result = commands.NewMkdisk((func() int {
			if localctx.(*MkdiskContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkdiskContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkdiskContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkdiskContext).GetM().GetColumn()
			}
		}()), localctx.(*MkdiskContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)

			var _m = p.Match(ParserParserRW_mkdisk)

			localctx.(*MkdiskContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkdiskContext).result = commands.NewMkdisk((func() int {
			if localctx.(*MkdiskContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkdiskContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkdiskContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkdiskContext).GetM().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkdiskparamsContext is an interface to support dynamic dispatch.
type IMkdiskparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IMkdiskparamsContext

	// GetP returns the p rule contexts.
	GetP() IMkdiskparamContext

	// SetL sets the l rule contexts.
	SetL(IMkdiskparamsContext)

	// SetP sets the p rule contexts.
	SetP(IMkdiskparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Mkdiskparam() IMkdiskparamContext
	Mkdiskparams() IMkdiskparamsContext

	// IsMkdiskparamsContext differentiates from other interfaces.
	IsMkdiskparamsContext()
}

type MkdiskparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IMkdiskparamsContext
	p      IMkdiskparamContext
}

func NewEmptyMkdiskparamsContext() *MkdiskparamsContext {
	var p = new(MkdiskparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdiskparams
	return p
}

func InitEmptyMkdiskparamsContext(p *MkdiskparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdiskparams
}

func (*MkdiskparamsContext) IsMkdiskparamsContext() {}

func NewMkdiskparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkdiskparamsContext {
	var p = new(MkdiskparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkdiskparams

	return p
}

func (s *MkdiskparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *MkdiskparamsContext) GetL() IMkdiskparamsContext { return s.l }

func (s *MkdiskparamsContext) GetP() IMkdiskparamContext { return s.p }

func (s *MkdiskparamsContext) SetL(v IMkdiskparamsContext) { s.l = v }

func (s *MkdiskparamsContext) SetP(v IMkdiskparamContext) { s.p = v }

func (s *MkdiskparamsContext) GetResult() map[string]string { return s.result }

func (s *MkdiskparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *MkdiskparamsContext) Mkdiskparam() IMkdiskparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdiskparamContext)
}

func (s *MkdiskparamsContext) Mkdiskparams() IMkdiskparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdiskparamsContext)
}

func (s *MkdiskparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdiskparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkdiskparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkdiskparams(s)
	}
}

func (s *MkdiskparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkdiskparams(s)
	}
}

func (p *ParserParser) Mkdiskparams() (localctx IMkdiskparamsContext) {
	return p.mkdiskparams(0)
}

func (p *ParserParser) mkdiskparams(_p int) (localctx IMkdiskparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMkdiskparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMkdiskparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, ParserParserRULE_mkdiskparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)

		var _x = p.Mkdiskparam()

		localctx.(*MkdiskparamsContext).p = _x
	}
	localctx.(*MkdiskparamsContext).result = map[string]string{localctx.(*MkdiskparamsContext).GetP().GetResult()[0]: localctx.(*MkdiskparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMkdiskparamsContext(p, _parentctx, _parentState)
			localctx.(*MkdiskparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_mkdiskparams)
			p.SetState(137)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(138)

				var _x = p.Mkdiskparam()

				localctx.(*MkdiskparamsContext).p = _x
			}
			localctx.(*MkdiskparamsContext).SetResult(localctx.(*MkdiskparamsContext).GetL().GetResult())
			localctx.(*MkdiskparamsContext).result[localctx.(*MkdiskparamsContext).GetP().GetResult()[0]] = localctx.(*MkdiskparamsContext).GetP().GetResult()[1]

		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkdiskparamContext is an interface to support dynamic dispatch.
type IMkdiskparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV1 returns the v1 token.
	GetV1() antlr.Token

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// GetV3 returns the v3 token.
	GetV3() antlr.Token

	// GetV4 returns the v4 token.
	GetV4() antlr.Token

	// SetV1 sets the v1 token.
	SetV1(antlr.Token)

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// SetV3 sets the v3 token.
	SetV3(antlr.Token)

	// SetV4 sets the v4 token.
	SetV4(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_size() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_number() antlr.TerminalNode
	RW_fit() antlr.TerminalNode
	TK_fit() antlr.TerminalNode
	RW_unit() antlr.TerminalNode
	TK_unit() antlr.TerminalNode
	RW_path() antlr.TerminalNode
	TK_path() antlr.TerminalNode

	// IsMkdiskparamContext differentiates from other interfaces.
	IsMkdiskparamContext()
}

type MkdiskparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     antlr.Token
	v2     antlr.Token
	v3     antlr.Token
	v4     antlr.Token
}

func NewEmptyMkdiskparamContext() *MkdiskparamContext {
	var p = new(MkdiskparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdiskparam
	return p
}

func InitEmptyMkdiskparamContext(p *MkdiskparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdiskparam
}

func (*MkdiskparamContext) IsMkdiskparamContext() {}

func NewMkdiskparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkdiskparamContext {
	var p = new(MkdiskparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkdiskparam

	return p
}

func (s *MkdiskparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MkdiskparamContext) GetV1() antlr.Token { return s.v1 }

func (s *MkdiskparamContext) GetV2() antlr.Token { return s.v2 }

func (s *MkdiskparamContext) GetV3() antlr.Token { return s.v3 }

func (s *MkdiskparamContext) GetV4() antlr.Token { return s.v4 }

func (s *MkdiskparamContext) SetV1(v antlr.Token) { s.v1 = v }

func (s *MkdiskparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *MkdiskparamContext) SetV3(v antlr.Token) { s.v3 = v }

func (s *MkdiskparamContext) SetV4(v antlr.Token) { s.v4 = v }

func (s *MkdiskparamContext) GetResult() []string { return s.result }

func (s *MkdiskparamContext) SetResult(v []string) { s.result = v }

func (s *MkdiskparamContext) RW_size() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_size, 0)
}

func (s *MkdiskparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *MkdiskparamContext) TK_number() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_number, 0)
}

func (s *MkdiskparamContext) RW_fit() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_fit, 0)
}

func (s *MkdiskparamContext) TK_fit() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_fit, 0)
}

func (s *MkdiskparamContext) RW_unit() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_unit, 0)
}

func (s *MkdiskparamContext) TK_unit() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_unit, 0)
}

func (s *MkdiskparamContext) RW_path() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_path, 0)
}

func (s *MkdiskparamContext) TK_path() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_path, 0)
}

func (s *MkdiskparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdiskparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkdiskparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkdiskparam(s)
	}
}

func (s *MkdiskparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkdiskparam(s)
	}
}

func (p *ParserParser) Mkdiskparam() (localctx IMkdiskparamContext) {
	localctx = NewMkdiskparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ParserParserRULE_mkdiskparam)
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_size:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(146)
			p.Match(ParserParserRW_size)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*MkdiskparamContext).v1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkdiskparamContext).result = []string{"size", (func() string {
			if localctx.(*MkdiskparamContext).GetV1() == nil {
				return ""
			} else {
				return localctx.(*MkdiskparamContext).GetV1().GetText()
			}
		}())}

	case ParserParserRW_fit:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.Match(ParserParserRW_fit)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)

			var _m = p.Match(ParserParserTK_fit)

			localctx.(*MkdiskparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkdiskparamContext).result = []string{"fit", (func() string {
			if localctx.(*MkdiskparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*MkdiskparamContext).GetV2().GetText()
			}
		}())}

	case ParserParserRW_unit:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(154)
			p.Match(ParserParserRW_unit)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)

			var _m = p.Match(ParserParserTK_unit)

			localctx.(*MkdiskparamContext).v3 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkdiskparamContext).result = []string{"unit", (func() string {
			if localctx.(*MkdiskparamContext).GetV3() == nil {
				return ""
			} else {
				return localctx.(*MkdiskparamContext).GetV3().GetText()
			}
		}())}

	case ParserParserRW_path:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(158)
			p.Match(ParserParserRW_path)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*MkdiskparamContext).v4 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkdiskparamContext).result = []string{"path", strings.Trim((func() string {
			if localctx.(*MkdiskparamContext).GetV4() == nil {
				return ""
			} else {
				return localctx.(*MkdiskparamContext).GetV4().GetText()
			}
		}()), "\"")}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRmdiskContext is an interface to support dynamic dispatch.
type IRmdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetR returns the r token.
	GetR() antlr.Token

	// GetP returns the p token.
	GetP() antlr.Token

	// SetR sets the r token.
	SetR(antlr.Token)

	// SetP sets the p token.
	SetP(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *commands.Rmdisk

	// SetResult sets the result attribute.
	SetResult(*commands.Rmdisk)

	// Getter signatures
	RW_path() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	RW_rmdisk() antlr.TerminalNode
	TK_path() antlr.TerminalNode

	// IsRmdiskContext differentiates from other interfaces.
	IsRmdiskContext()
}

type RmdiskContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Rmdisk
	r      antlr.Token
	p      antlr.Token
}

func NewEmptyRmdiskContext() *RmdiskContext {
	var p = new(RmdiskContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rmdisk
	return p
}

func InitEmptyRmdiskContext(p *RmdiskContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rmdisk
}

func (*RmdiskContext) IsRmdiskContext() {}

func NewRmdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RmdiskContext {
	var p = new(RmdiskContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_rmdisk

	return p
}

func (s *RmdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *RmdiskContext) GetR() antlr.Token { return s.r }

func (s *RmdiskContext) GetP() antlr.Token { return s.p }

func (s *RmdiskContext) SetR(v antlr.Token) { s.r = v }

func (s *RmdiskContext) SetP(v antlr.Token) { s.p = v }

func (s *RmdiskContext) GetResult() *commands.Rmdisk { return s.result }

func (s *RmdiskContext) SetResult(v *commands.Rmdisk) { s.result = v }

func (s *RmdiskContext) RW_path() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_path, 0)
}

func (s *RmdiskContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *RmdiskContext) RW_rmdisk() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_rmdisk, 0)
}

func (s *RmdiskContext) TK_path() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_path, 0)
}

func (s *RmdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RmdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RmdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterRmdisk(s)
	}
}

func (s *RmdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitRmdisk(s)
	}
}

func (p *ParserParser) Rmdisk() (localctx IRmdiskContext) {
	localctx = NewRmdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ParserParserRULE_rmdisk)
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)

			var _m = p.Match(ParserParserRW_rmdisk)

			localctx.(*RmdiskContext).r = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Match(ParserParserRW_path)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*RmdiskContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RmdiskContext).result = commands.NewRmdisk((func() int {
			if localctx.(*RmdiskContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RmdiskContext).GetR().GetLine()
			}
		}()), (func() int {
			if localctx.(*RmdiskContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RmdiskContext).GetR().GetColumn()
			}
		}()), map[string]string{"path": strings.Trim((func() string {
			if localctx.(*RmdiskContext).GetP() == nil {
				return ""
			} else {
				return localctx.(*RmdiskContext).GetP().GetText()
			}
		}()), "\"")})

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(169)

			var _m = p.Match(ParserParserRW_rmdisk)

			localctx.(*RmdiskContext).r = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RmdiskContext).result = commands.NewRmdisk((func() int {
			if localctx.(*RmdiskContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RmdiskContext).GetR().GetLine()
			}
		}()), (func() int {
			if localctx.(*RmdiskContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RmdiskContext).GetR().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFdiskContext is an interface to support dynamic dispatch.
type IFdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetF returns the f token.
	GetF() antlr.Token

	// SetF sets the f token.
	SetF(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IFdiskparamsContext

	// SetP sets the p rule contexts.
	SetP(IFdiskparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Fdisk

	// SetResult sets the result attribute.
	SetResult(*commands.Fdisk)

	// Getter signatures
	RW_fdisk() antlr.TerminalNode
	Fdiskparams() IFdiskparamsContext

	// IsFdiskContext differentiates from other interfaces.
	IsFdiskContext()
}

type FdiskContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Fdisk
	f      antlr.Token
	p      IFdiskparamsContext
}

func NewEmptyFdiskContext() *FdiskContext {
	var p = new(FdiskContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdisk
	return p
}

func InitEmptyFdiskContext(p *FdiskContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdisk
}

func (*FdiskContext) IsFdiskContext() {}

func NewFdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FdiskContext {
	var p = new(FdiskContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_fdisk

	return p
}

func (s *FdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *FdiskContext) GetF() antlr.Token { return s.f }

func (s *FdiskContext) SetF(v antlr.Token) { s.f = v }

func (s *FdiskContext) GetP() IFdiskparamsContext { return s.p }

func (s *FdiskContext) SetP(v IFdiskparamsContext) { s.p = v }

func (s *FdiskContext) GetResult() *commands.Fdisk { return s.result }

func (s *FdiskContext) SetResult(v *commands.Fdisk) { s.result = v }

func (s *FdiskContext) RW_fdisk() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_fdisk, 0)
}

func (s *FdiskContext) Fdiskparams() IFdiskparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFdiskparamsContext)
}

func (s *FdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterFdisk(s)
	}
}

func (s *FdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitFdisk(s)
	}
}

func (p *ParserParser) Fdisk() (localctx IFdiskContext) {
	localctx = NewFdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ParserParserRULE_fdisk)
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)

			var _m = p.Match(ParserParserRW_fdisk)

			localctx.(*FdiskContext).f = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)

			var _x = p.fdiskparams(0)

			localctx.(*FdiskContext).p = _x
		}
		localctx.(*FdiskContext).result = commands.NewFdisk((func() int {
			if localctx.(*FdiskContext).GetF() == nil {
				return 0
			} else {
				return localctx.(*FdiskContext).GetF().GetLine()
			}
		}()), (func() int {
			if localctx.(*FdiskContext).GetF() == nil {
				return 0
			} else {
				return localctx.(*FdiskContext).GetF().GetColumn()
			}
		}()), localctx.(*FdiskContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)

			var _m = p.Match(ParserParserRW_fdisk)

			localctx.(*FdiskContext).f = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskContext).result = commands.NewFdisk((func() int {
			if localctx.(*FdiskContext).GetF() == nil {
				return 0
			} else {
				return localctx.(*FdiskContext).GetF().GetLine()
			}
		}()), (func() int {
			if localctx.(*FdiskContext).GetF() == nil {
				return 0
			} else {
				return localctx.(*FdiskContext).GetF().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFdiskparamsContext is an interface to support dynamic dispatch.
type IFdiskparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IFdiskparamsContext

	// GetP returns the p rule contexts.
	GetP() IFdiskparamContext

	// SetL sets the l rule contexts.
	SetL(IFdiskparamsContext)

	// SetP sets the p rule contexts.
	SetP(IFdiskparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Fdiskparam() IFdiskparamContext
	Fdiskparams() IFdiskparamsContext

	// IsFdiskparamsContext differentiates from other interfaces.
	IsFdiskparamsContext()
}

type FdiskparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IFdiskparamsContext
	p      IFdiskparamContext
}

func NewEmptyFdiskparamsContext() *FdiskparamsContext {
	var p = new(FdiskparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdiskparams
	return p
}

func InitEmptyFdiskparamsContext(p *FdiskparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdiskparams
}

func (*FdiskparamsContext) IsFdiskparamsContext() {}

func NewFdiskparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FdiskparamsContext {
	var p = new(FdiskparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_fdiskparams

	return p
}

func (s *FdiskparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *FdiskparamsContext) GetL() IFdiskparamsContext { return s.l }

func (s *FdiskparamsContext) GetP() IFdiskparamContext { return s.p }

func (s *FdiskparamsContext) SetL(v IFdiskparamsContext) { s.l = v }

func (s *FdiskparamsContext) SetP(v IFdiskparamContext) { s.p = v }

func (s *FdiskparamsContext) GetResult() map[string]string { return s.result }

func (s *FdiskparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *FdiskparamsContext) Fdiskparam() IFdiskparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFdiskparamContext)
}

func (s *FdiskparamsContext) Fdiskparams() IFdiskparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFdiskparamsContext)
}

func (s *FdiskparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FdiskparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FdiskparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterFdiskparams(s)
	}
}

func (s *FdiskparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitFdiskparams(s)
	}
}

func (p *ParserParser) Fdiskparams() (localctx IFdiskparamsContext) {
	return p.fdiskparams(0)
}

func (p *ParserParser) fdiskparams(_p int) (localctx IFdiskparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewFdiskparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFdiskparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, ParserParserRULE_fdiskparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)

		var _x = p.Fdiskparam()

		localctx.(*FdiskparamsContext).p = _x
	}
	localctx.(*FdiskparamsContext).result = map[string]string{localctx.(*FdiskparamsContext).GetP().GetResult()[0]: localctx.(*FdiskparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFdiskparamsContext(p, _parentctx, _parentState)
			localctx.(*FdiskparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_fdiskparams)
			p.SetState(185)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(186)

				var _x = p.Fdiskparam()

				localctx.(*FdiskparamsContext).p = _x
			}
			localctx.(*FdiskparamsContext).SetResult(localctx.(*FdiskparamsContext).GetL().GetResult())
			localctx.(*FdiskparamsContext).result[localctx.(*FdiskparamsContext).GetP().GetResult()[0]] = localctx.(*FdiskparamsContext).GetP().GetResult()[1]

		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFdiskparamContext is an interface to support dynamic dispatch.
type IFdiskparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV1 returns the v1 token.
	GetV1() antlr.Token

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// GetV3 returns the v3 token.
	GetV3() antlr.Token

	// GetV4 returns the v4 token.
	GetV4() antlr.Token

	// GetV5 returns the v5 token.
	GetV5() antlr.Token

	// GetV6 returns the v6 token.
	GetV6() antlr.Token

	// SetV1 sets the v1 token.
	SetV1(antlr.Token)

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// SetV3 sets the v3 token.
	SetV3(antlr.Token)

	// SetV4 sets the v4 token.
	SetV4(antlr.Token)

	// SetV5 sets the v5 token.
	SetV5(antlr.Token)

	// SetV6 sets the v6 token.
	SetV6(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_size() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_number() antlr.TerminalNode
	RW_unit() antlr.TerminalNode
	TK_unit() antlr.TerminalNode
	RW_path() antlr.TerminalNode
	TK_path() antlr.TerminalNode
	RW_type() antlr.TerminalNode
	TK_type() antlr.TerminalNode
	RW_fit() antlr.TerminalNode
	TK_fit() antlr.TerminalNode
	RW_name() antlr.TerminalNode
	TK_id() antlr.TerminalNode

	// IsFdiskparamContext differentiates from other interfaces.
	IsFdiskparamContext()
}

type FdiskparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     antlr.Token
	v2     antlr.Token
	v3     antlr.Token
	v4     antlr.Token
	v5     antlr.Token
	v6     antlr.Token
}

func NewEmptyFdiskparamContext() *FdiskparamContext {
	var p = new(FdiskparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdiskparam
	return p
}

func InitEmptyFdiskparamContext(p *FdiskparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdiskparam
}

func (*FdiskparamContext) IsFdiskparamContext() {}

func NewFdiskparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FdiskparamContext {
	var p = new(FdiskparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_fdiskparam

	return p
}

func (s *FdiskparamContext) GetParser() antlr.Parser { return s.parser }

func (s *FdiskparamContext) GetV1() antlr.Token { return s.v1 }

func (s *FdiskparamContext) GetV2() antlr.Token { return s.v2 }

func (s *FdiskparamContext) GetV3() antlr.Token { return s.v3 }

func (s *FdiskparamContext) GetV4() antlr.Token { return s.v4 }

func (s *FdiskparamContext) GetV5() antlr.Token { return s.v5 }

func (s *FdiskparamContext) GetV6() antlr.Token { return s.v6 }

func (s *FdiskparamContext) SetV1(v antlr.Token) { s.v1 = v }

func (s *FdiskparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *FdiskparamContext) SetV3(v antlr.Token) { s.v3 = v }

func (s *FdiskparamContext) SetV4(v antlr.Token) { s.v4 = v }

func (s *FdiskparamContext) SetV5(v antlr.Token) { s.v5 = v }

func (s *FdiskparamContext) SetV6(v antlr.Token) { s.v6 = v }

func (s *FdiskparamContext) GetResult() []string { return s.result }

func (s *FdiskparamContext) SetResult(v []string) { s.result = v }

func (s *FdiskparamContext) RW_size() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_size, 0)
}

func (s *FdiskparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *FdiskparamContext) TK_number() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_number, 0)
}

func (s *FdiskparamContext) RW_unit() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_unit, 0)
}

func (s *FdiskparamContext) TK_unit() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_unit, 0)
}

func (s *FdiskparamContext) RW_path() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_path, 0)
}

func (s *FdiskparamContext) TK_path() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_path, 0)
}

func (s *FdiskparamContext) RW_type() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_type, 0)
}

func (s *FdiskparamContext) TK_type() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_type, 0)
}

func (s *FdiskparamContext) RW_fit() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_fit, 0)
}

func (s *FdiskparamContext) TK_fit() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_fit, 0)
}

func (s *FdiskparamContext) RW_name() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_name, 0)
}

func (s *FdiskparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *FdiskparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FdiskparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FdiskparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterFdiskparam(s)
	}
}

func (s *FdiskparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitFdiskparam(s)
	}
}

func (p *ParserParser) Fdiskparam() (localctx IFdiskparamContext) {
	localctx = NewFdiskparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ParserParserRULE_fdiskparam)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_size:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(194)
			p.Match(ParserParserRW_size)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*FdiskparamContext).v1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"size", (func() string {
			if localctx.(*FdiskparamContext).GetV1() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV1().GetText()
			}
		}())}

	case ParserParserRW_unit:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
			p.Match(ParserParserRW_unit)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)

			var _m = p.Match(ParserParserTK_unit)

			localctx.(*FdiskparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"unit", (func() string {
			if localctx.(*FdiskparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV2().GetText()
			}
		}())}

	case ParserParserRW_path:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(202)
			p.Match(ParserParserRW_path)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*FdiskparamContext).v3 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"path", strings.Trim((func() string {
			if localctx.(*FdiskparamContext).GetV3() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV3().GetText()
			}
		}()), "\"")}

	case ParserParserRW_type:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(206)
			p.Match(ParserParserRW_type)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)

			var _m = p.Match(ParserParserTK_type)

			localctx.(*FdiskparamContext).v4 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"type", (func() string {
			if localctx.(*FdiskparamContext).GetV4() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV4().GetText()
			}
		}())}

	case ParserParserRW_fit:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(210)
			p.Match(ParserParserRW_fit)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)

			var _m = p.Match(ParserParserTK_fit)

			localctx.(*FdiskparamContext).v5 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"fit", (func() string {
			if localctx.(*FdiskparamContext).GetV5() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV5().GetText()
			}
		}())}

	case ParserParserRW_name:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(214)
			p.Match(ParserParserRW_name)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*FdiskparamContext).v6 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"name", strings.Trim((func() string {
			if localctx.(*FdiskparamContext).GetV6() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV6().GetText()
			}
		}()), "\"")}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMountContext is an interface to support dynamic dispatch.
type IMountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetM returns the m token.
	GetM() antlr.Token

	// SetM sets the m token.
	SetM(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IMountparamsContext

	// SetP sets the p rule contexts.
	SetP(IMountparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Mount

	// SetResult sets the result attribute.
	SetResult(*commands.Mount)

	// Getter signatures
	RW_mount() antlr.TerminalNode
	Mountparams() IMountparamsContext

	// IsMountContext differentiates from other interfaces.
	IsMountContext()
}

type MountContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Mount
	m      antlr.Token
	p      IMountparamsContext
}

func NewEmptyMountContext() *MountContext {
	var p = new(MountContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mount
	return p
}

func InitEmptyMountContext(p *MountContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mount
}

func (*MountContext) IsMountContext() {}

func NewMountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MountContext {
	var p = new(MountContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mount

	return p
}

func (s *MountContext) GetParser() antlr.Parser { return s.parser }

func (s *MountContext) GetM() antlr.Token { return s.m }

func (s *MountContext) SetM(v antlr.Token) { s.m = v }

func (s *MountContext) GetP() IMountparamsContext { return s.p }

func (s *MountContext) SetP(v IMountparamsContext) { s.p = v }

func (s *MountContext) GetResult() *commands.Mount { return s.result }

func (s *MountContext) SetResult(v *commands.Mount) { s.result = v }

func (s *MountContext) RW_mount() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mount, 0)
}

func (s *MountContext) Mountparams() IMountparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountparamsContext)
}

func (s *MountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMount(s)
	}
}

func (s *MountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMount(s)
	}
}

func (p *ParserParser) Mount() (localctx IMountContext) {
	localctx = NewMountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ParserParserRULE_mount)
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)

			var _m = p.Match(ParserParserRW_mount)

			localctx.(*MountContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)

			var _x = p.mountparams(0)

			localctx.(*MountContext).p = _x
		}
		localctx.(*MountContext).result = commands.NewMount((func() int {
			if localctx.(*MountContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MountContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MountContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MountContext).GetM().GetColumn()
			}
		}()), localctx.(*MountContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)

			var _m = p.Match(ParserParserRW_mount)

			localctx.(*MountContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MountContext).result = commands.NewMount((func() int {
			if localctx.(*MountContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MountContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MountContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MountContext).GetM().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMountparamsContext is an interface to support dynamic dispatch.
type IMountparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IMountparamsContext

	// GetP returns the p rule contexts.
	GetP() IMountparamContext

	// SetL sets the l rule contexts.
	SetL(IMountparamsContext)

	// SetP sets the p rule contexts.
	SetP(IMountparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Mountparam() IMountparamContext
	Mountparams() IMountparamsContext

	// IsMountparamsContext differentiates from other interfaces.
	IsMountparamsContext()
}

type MountparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IMountparamsContext
	p      IMountparamContext
}

func NewEmptyMountparamsContext() *MountparamsContext {
	var p = new(MountparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mountparams
	return p
}

func InitEmptyMountparamsContext(p *MountparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mountparams
}

func (*MountparamsContext) IsMountparamsContext() {}

func NewMountparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MountparamsContext {
	var p = new(MountparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mountparams

	return p
}

func (s *MountparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *MountparamsContext) GetL() IMountparamsContext { return s.l }

func (s *MountparamsContext) GetP() IMountparamContext { return s.p }

func (s *MountparamsContext) SetL(v IMountparamsContext) { s.l = v }

func (s *MountparamsContext) SetP(v IMountparamContext) { s.p = v }

func (s *MountparamsContext) GetResult() map[string]string { return s.result }

func (s *MountparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *MountparamsContext) Mountparam() IMountparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountparamContext)
}

func (s *MountparamsContext) Mountparams() IMountparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountparamsContext)
}

func (s *MountparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MountparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MountparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMountparams(s)
	}
}

func (s *MountparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMountparams(s)
	}
}

func (p *ParserParser) Mountparams() (localctx IMountparamsContext) {
	return p.mountparams(0)
}

func (p *ParserParser) mountparams(_p int) (localctx IMountparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMountparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMountparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, ParserParserRULE_mountparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)

		var _x = p.Mountparam()

		localctx.(*MountparamsContext).p = _x
	}
	localctx.(*MountparamsContext).result = map[string]string{localctx.(*MountparamsContext).GetP().GetResult()[0]: localctx.(*MountparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMountparamsContext(p, _parentctx, _parentState)
			localctx.(*MountparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_mountparams)
			p.SetState(232)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(233)

				var _x = p.Mountparam()

				localctx.(*MountparamsContext).p = _x
			}
			localctx.(*MountparamsContext).SetResult(localctx.(*MountparamsContext).GetL().GetResult())
			localctx.(*MountparamsContext).result[localctx.(*MountparamsContext).GetP().GetResult()[0]] = localctx.(*MountparamsContext).GetP().GetResult()[1]

		}
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMountparamContext is an interface to support dynamic dispatch.
type IMountparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV1 returns the v1 token.
	GetV1() antlr.Token

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// SetV1 sets the v1 token.
	SetV1(antlr.Token)

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_path() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_path() antlr.TerminalNode
	RW_name() antlr.TerminalNode
	TK_id() antlr.TerminalNode

	// IsMountparamContext differentiates from other interfaces.
	IsMountparamContext()
}

type MountparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     antlr.Token
	v2     antlr.Token
}

func NewEmptyMountparamContext() *MountparamContext {
	var p = new(MountparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mountparam
	return p
}

func InitEmptyMountparamContext(p *MountparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mountparam
}

func (*MountparamContext) IsMountparamContext() {}

func NewMountparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MountparamContext {
	var p = new(MountparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mountparam

	return p
}

func (s *MountparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MountparamContext) GetV1() antlr.Token { return s.v1 }

func (s *MountparamContext) GetV2() antlr.Token { return s.v2 }

func (s *MountparamContext) SetV1(v antlr.Token) { s.v1 = v }

func (s *MountparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *MountparamContext) GetResult() []string { return s.result }

func (s *MountparamContext) SetResult(v []string) { s.result = v }

func (s *MountparamContext) RW_path() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_path, 0)
}

func (s *MountparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *MountparamContext) TK_path() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_path, 0)
}

func (s *MountparamContext) RW_name() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_name, 0)
}

func (s *MountparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *MountparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MountparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MountparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMountparam(s)
	}
}

func (s *MountparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMountparam(s)
	}
}

func (p *ParserParser) Mountparam() (localctx IMountparamContext) {
	localctx = NewMountparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ParserParserRULE_mountparam)
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_path:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			p.Match(ParserParserRW_path)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*MountparamContext).v1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MountparamContext).result = []string{"path", strings.Trim((func() string {
			if localctx.(*MountparamContext).GetV1() == nil {
				return ""
			} else {
				return localctx.(*MountparamContext).GetV1().GetText()
			}
		}()), "\"")}

	case ParserParserRW_name:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(245)
			p.Match(ParserParserRW_name)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*MountparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MountparamContext).result = []string{"name", strings.Trim((func() string {
			if localctx.(*MountparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*MountparamContext).GetV2().GetText()
			}
		}()), "\"")}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMountedContext is an interface to support dynamic dispatch.
type IMountedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetM returns the m token.
	GetM() antlr.Token

	// SetM sets the m token.
	SetM(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *commands.Mounted

	// SetResult sets the result attribute.
	SetResult(*commands.Mounted)

	// Getter signatures
	RW_mounted() antlr.TerminalNode

	// IsMountedContext differentiates from other interfaces.
	IsMountedContext()
}

type MountedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Mounted
	m      antlr.Token
}

func NewEmptyMountedContext() *MountedContext {
	var p = new(MountedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mounted
	return p
}

func InitEmptyMountedContext(p *MountedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mounted
}

func (*MountedContext) IsMountedContext() {}

func NewMountedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MountedContext {
	var p = new(MountedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mounted

	return p
}

func (s *MountedContext) GetParser() antlr.Parser { return s.parser }

func (s *MountedContext) GetM() antlr.Token { return s.m }

func (s *MountedContext) SetM(v antlr.Token) { s.m = v }

func (s *MountedContext) GetResult() *commands.Mounted { return s.result }

func (s *MountedContext) SetResult(v *commands.Mounted) { s.result = v }

func (s *MountedContext) RW_mounted() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mounted, 0)
}

func (s *MountedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MountedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MountedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMounted(s)
	}
}

func (s *MountedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMounted(s)
	}
}

func (p *ParserParser) Mounted() (localctx IMountedContext) {
	localctx = NewMountedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ParserParserRULE_mounted)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)

		var _m = p.Match(ParserParserRW_mounted)

		localctx.(*MountedContext).m = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*MountedContext).result = commands.NewMounted((func() int {
		if localctx.(*MountedContext).GetM() == nil {
			return 0
		} else {
			return localctx.(*MountedContext).GetM().GetLine()
		}
	}()), (func() int {
		if localctx.(*MountedContext).GetM() == nil {
			return 0
		} else {
			return localctx.(*MountedContext).GetM().GetColumn()
		}
	}()))

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkfsContext is an interface to support dynamic dispatch.
type IMkfsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetM returns the m token.
	GetM() antlr.Token

	// SetM sets the m token.
	SetM(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IMkfsparamsContext

	// SetP sets the p rule contexts.
	SetP(IMkfsparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Mkfs

	// SetResult sets the result attribute.
	SetResult(*commands.Mkfs)

	// Getter signatures
	RW_mkfs() antlr.TerminalNode
	Mkfsparams() IMkfsparamsContext

	// IsMkfsContext differentiates from other interfaces.
	IsMkfsContext()
}

type MkfsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Mkfs
	m      antlr.Token
	p      IMkfsparamsContext
}

func NewEmptyMkfsContext() *MkfsContext {
	var p = new(MkfsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfs
	return p
}

func InitEmptyMkfsContext(p *MkfsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfs
}

func (*MkfsContext) IsMkfsContext() {}

func NewMkfsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkfsContext {
	var p = new(MkfsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkfs

	return p
}

func (s *MkfsContext) GetParser() antlr.Parser { return s.parser }

func (s *MkfsContext) GetM() antlr.Token { return s.m }

func (s *MkfsContext) SetM(v antlr.Token) { s.m = v }

func (s *MkfsContext) GetP() IMkfsparamsContext { return s.p }

func (s *MkfsContext) SetP(v IMkfsparamsContext) { s.p = v }

func (s *MkfsContext) GetResult() *commands.Mkfs { return s.result }

func (s *MkfsContext) SetResult(v *commands.Mkfs) { s.result = v }

func (s *MkfsContext) RW_mkfs() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mkfs, 0)
}

func (s *MkfsContext) Mkfsparams() IMkfsparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfsparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfsparamsContext)
}

func (s *MkfsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkfsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkfs(s)
	}
}

func (s *MkfsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkfs(s)
	}
}

func (p *ParserParser) Mkfs() (localctx IMkfsContext) {
	localctx = NewMkfsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ParserParserRULE_mkfs)
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(254)

			var _m = p.Match(ParserParserRW_mkfs)

			localctx.(*MkfsContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)

			var _x = p.mkfsparams(0)

			localctx.(*MkfsContext).p = _x
		}
		localctx.(*MkfsContext).result = commands.NewMkfs((func() int {
			if localctx.(*MkfsContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkfsContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkfsContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkfsContext).GetM().GetColumn()
			}
		}()), localctx.(*MkfsContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(258)

			var _m = p.Match(ParserParserRW_mkfs)

			localctx.(*MkfsContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkfsContext).result = commands.NewMkfs((func() int {
			if localctx.(*MkfsContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkfsContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkfsContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkfsContext).GetM().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkfsparamsContext is an interface to support dynamic dispatch.
type IMkfsparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IMkfsparamsContext

	// GetP returns the p rule contexts.
	GetP() IMkfsparamContext

	// SetL sets the l rule contexts.
	SetL(IMkfsparamsContext)

	// SetP sets the p rule contexts.
	SetP(IMkfsparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Mkfsparam() IMkfsparamContext
	Mkfsparams() IMkfsparamsContext

	// IsMkfsparamsContext differentiates from other interfaces.
	IsMkfsparamsContext()
}

type MkfsparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IMkfsparamsContext
	p      IMkfsparamContext
}

func NewEmptyMkfsparamsContext() *MkfsparamsContext {
	var p = new(MkfsparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfsparams
	return p
}

func InitEmptyMkfsparamsContext(p *MkfsparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfsparams
}

func (*MkfsparamsContext) IsMkfsparamsContext() {}

func NewMkfsparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkfsparamsContext {
	var p = new(MkfsparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkfsparams

	return p
}

func (s *MkfsparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *MkfsparamsContext) GetL() IMkfsparamsContext { return s.l }

func (s *MkfsparamsContext) GetP() IMkfsparamContext { return s.p }

func (s *MkfsparamsContext) SetL(v IMkfsparamsContext) { s.l = v }

func (s *MkfsparamsContext) SetP(v IMkfsparamContext) { s.p = v }

func (s *MkfsparamsContext) GetResult() map[string]string { return s.result }

func (s *MkfsparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *MkfsparamsContext) Mkfsparam() IMkfsparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfsparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfsparamContext)
}

func (s *MkfsparamsContext) Mkfsparams() IMkfsparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfsparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfsparamsContext)
}

func (s *MkfsparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfsparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkfsparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkfsparams(s)
	}
}

func (s *MkfsparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkfsparams(s)
	}
}

func (p *ParserParser) Mkfsparams() (localctx IMkfsparamsContext) {
	return p.mkfsparams(0)
}

func (p *ParserParser) mkfsparams(_p int) (localctx IMkfsparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMkfsparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMkfsparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, ParserParserRULE_mkfsparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)

		var _x = p.Mkfsparam()

		localctx.(*MkfsparamsContext).p = _x
	}
	localctx.(*MkfsparamsContext).result = map[string]string{localctx.(*MkfsparamsContext).GetP().GetResult()[0]: localctx.(*MkfsparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMkfsparamsContext(p, _parentctx, _parentState)
			localctx.(*MkfsparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_mkfsparams)
			p.SetState(266)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(267)

				var _x = p.Mkfsparam()

				localctx.(*MkfsparamsContext).p = _x
			}
			localctx.(*MkfsparamsContext).SetResult(localctx.(*MkfsparamsContext).GetL().GetResult())
			localctx.(*MkfsparamsContext).result[localctx.(*MkfsparamsContext).GetP().GetResult()[0]] = localctx.(*MkfsparamsContext).GetP().GetResult()[1]

		}
		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkfsparamContext is an interface to support dynamic dispatch.
type IMkfsparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV1 returns the v1 token.
	GetV1() antlr.Token

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// SetV1 sets the v1 token.
	SetV1(antlr.Token)

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_type() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_ftype() antlr.TerminalNode
	RW_id() antlr.TerminalNode
	TK_id() antlr.TerminalNode

	// IsMkfsparamContext differentiates from other interfaces.
	IsMkfsparamContext()
}

type MkfsparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     antlr.Token
	v2     antlr.Token
}

func NewEmptyMkfsparamContext() *MkfsparamContext {
	var p = new(MkfsparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfsparam
	return p
}

func InitEmptyMkfsparamContext(p *MkfsparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfsparam
}

func (*MkfsparamContext) IsMkfsparamContext() {}

func NewMkfsparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkfsparamContext {
	var p = new(MkfsparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkfsparam

	return p
}

func (s *MkfsparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MkfsparamContext) GetV1() antlr.Token { return s.v1 }

func (s *MkfsparamContext) GetV2() antlr.Token { return s.v2 }

func (s *MkfsparamContext) SetV1(v antlr.Token) { s.v1 = v }

func (s *MkfsparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *MkfsparamContext) GetResult() []string { return s.result }

func (s *MkfsparamContext) SetResult(v []string) { s.result = v }

func (s *MkfsparamContext) RW_type() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_type, 0)
}

func (s *MkfsparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *MkfsparamContext) TK_ftype() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_ftype, 0)
}

func (s *MkfsparamContext) RW_id() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_id, 0)
}

func (s *MkfsparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *MkfsparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfsparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkfsparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkfsparam(s)
	}
}

func (s *MkfsparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkfsparam(s)
	}
}

func (p *ParserParser) Mkfsparam() (localctx IMkfsparamContext) {
	localctx = NewMkfsparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ParserParserRULE_mkfsparam)
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_type:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.Match(ParserParserRW_type)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)

			var _m = p.Match(ParserParserTK_ftype)

			localctx.(*MkfsparamContext).v1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkfsparamContext).result = []string{"type", (func() string {
			if localctx.(*MkfsparamContext).GetV1() == nil {
				return ""
			} else {
				return localctx.(*MkfsparamContext).GetV1().GetText()
			}
		}())}

	case ParserParserRW_id:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(279)
			p.Match(ParserParserRW_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*MkfsparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkfsparamContext).result = []string{"id", strings.Trim((func() string {
			if localctx.(*MkfsparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*MkfsparamContext).GetV2().GetText()
			}
		}()), "\"")}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICatContext is an interface to support dynamic dispatch.
type ICatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC returns the c token.
	GetC() antlr.Token

	// SetC sets the c token.
	SetC(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() ICatparamsContext

	// SetP sets the p rule contexts.
	SetP(ICatparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Cat

	// SetResult sets the result attribute.
	SetResult(*commands.Cat)

	// Getter signatures
	RW_cat() antlr.TerminalNode
	Catparams() ICatparamsContext

	// IsCatContext differentiates from other interfaces.
	IsCatContext()
}

type CatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Cat
	c      antlr.Token
	p      ICatparamsContext
}

func NewEmptyCatContext() *CatContext {
	var p = new(CatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_cat
	return p
}

func InitEmptyCatContext(p *CatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_cat
}

func (*CatContext) IsCatContext() {}

func NewCatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatContext {
	var p = new(CatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_cat

	return p
}

func (s *CatContext) GetParser() antlr.Parser { return s.parser }

func (s *CatContext) GetC() antlr.Token { return s.c }

func (s *CatContext) SetC(v antlr.Token) { s.c = v }

func (s *CatContext) GetP() ICatparamsContext { return s.p }

func (s *CatContext) SetP(v ICatparamsContext) { s.p = v }

func (s *CatContext) GetResult() *commands.Cat { return s.result }

func (s *CatContext) SetResult(v *commands.Cat) { s.result = v }

func (s *CatContext) RW_cat() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_cat, 0)
}

func (s *CatContext) Catparams() ICatparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICatparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICatparamsContext)
}

func (s *CatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterCat(s)
	}
}

func (s *CatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitCat(s)
	}
}

func (p *ParserParser) Cat() (localctx ICatContext) {
	localctx = NewCatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ParserParserRULE_cat)
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)

			var _m = p.Match(ParserParserRW_cat)

			localctx.(*CatContext).c = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)

			var _x = p.catparams(0)

			localctx.(*CatContext).p = _x
		}
		localctx.(*CatContext).result = commands.NewCat((func() int {
			if localctx.(*CatContext).GetC() == nil {
				return 0
			} else {
				return localctx.(*CatContext).GetC().GetLine()
			}
		}()), (func() int {
			if localctx.(*CatContext).GetC() == nil {
				return 0
			} else {
				return localctx.(*CatContext).GetC().GetColumn()
			}
		}()), localctx.(*CatContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(289)

			var _m = p.Match(ParserParserRW_cat)

			localctx.(*CatContext).c = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*CatContext).result = commands.NewCat((func() int {
			if localctx.(*CatContext).GetC() == nil {
				return 0
			} else {
				return localctx.(*CatContext).GetC().GetLine()
			}
		}()), (func() int {
			if localctx.(*CatContext).GetC() == nil {
				return 0
			} else {
				return localctx.(*CatContext).GetC().GetColumn()
			}
		}()), []string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICatparamsContext is an interface to support dynamic dispatch.
type ICatparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() ICatparamsContext

	// GetP returns the p rule contexts.
	GetP() ICatparamContext

	// SetL sets the l rule contexts.
	SetL(ICatparamsContext)

	// SetP sets the p rule contexts.
	SetP(ICatparamContext)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	Catparam() ICatparamContext
	Catparams() ICatparamsContext

	// IsCatparamsContext differentiates from other interfaces.
	IsCatparamsContext()
}

type CatparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	l      ICatparamsContext
	p      ICatparamContext
}

func NewEmptyCatparamsContext() *CatparamsContext {
	var p = new(CatparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_catparams
	return p
}

func InitEmptyCatparamsContext(p *CatparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_catparams
}

func (*CatparamsContext) IsCatparamsContext() {}

func NewCatparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatparamsContext {
	var p = new(CatparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_catparams

	return p
}

func (s *CatparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *CatparamsContext) GetL() ICatparamsContext { return s.l }

func (s *CatparamsContext) GetP() ICatparamContext { return s.p }

func (s *CatparamsContext) SetL(v ICatparamsContext) { s.l = v }

func (s *CatparamsContext) SetP(v ICatparamContext) { s.p = v }

func (s *CatparamsContext) GetResult() []string { return s.result }

func (s *CatparamsContext) SetResult(v []string) { s.result = v }

func (s *CatparamsContext) Catparam() ICatparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICatparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICatparamContext)
}

func (s *CatparamsContext) Catparams() ICatparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICatparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICatparamsContext)
}

func (s *CatparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CatparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterCatparams(s)
	}
}

func (s *CatparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitCatparams(s)
	}
}

func (p *ParserParser) Catparams() (localctx ICatparamsContext) {
	return p.catparams(0)
}

func (p *ParserParser) catparams(_p int) (localctx ICatparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewCatparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICatparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, ParserParserRULE_catparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)

		var _x = p.Catparam()

		localctx.(*CatparamsContext).p = _x
	}
	localctx.(*CatparamsContext).result = []string{localctx.(*CatparamsContext).GetP().GetResult()}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCatparamsContext(p, _parentctx, _parentState)
			localctx.(*CatparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_catparams)
			p.SetState(297)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(298)

				var _x = p.Catparam()

				localctx.(*CatparamsContext).p = _x
			}
			localctx.(*CatparamsContext).SetResult(localctx.(*CatparamsContext).GetL().GetResult())
			localctx.(*CatparamsContext).result = append(localctx.(*CatparamsContext).result, localctx.(*CatparamsContext).GetP().GetResult())

		}
		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICatparamContext is an interface to support dynamic dispatch.
type ICatparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP1 returns the p1 token.
	GetP1() antlr.Token

	// SetP1 sets the p1 token.
	SetP1(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() string

	// SetResult sets the result attribute.
	SetResult(string)

	// Getter signatures
	RW_fileN() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_path() antlr.TerminalNode

	// IsCatparamContext differentiates from other interfaces.
	IsCatparamContext()
}

type CatparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result string
	p1     antlr.Token
}

func NewEmptyCatparamContext() *CatparamContext {
	var p = new(CatparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_catparam
	return p
}

func InitEmptyCatparamContext(p *CatparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_catparam
}

func (*CatparamContext) IsCatparamContext() {}

func NewCatparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatparamContext {
	var p = new(CatparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_catparam

	return p
}

func (s *CatparamContext) GetParser() antlr.Parser { return s.parser }

func (s *CatparamContext) GetP1() antlr.Token { return s.p1 }

func (s *CatparamContext) SetP1(v antlr.Token) { s.p1 = v }

func (s *CatparamContext) GetResult() string { return s.result }

func (s *CatparamContext) SetResult(v string) { s.result = v }

func (s *CatparamContext) RW_fileN() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_fileN, 0)
}

func (s *CatparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *CatparamContext) TK_path() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_path, 0)
}

func (s *CatparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CatparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterCatparam(s)
	}
}

func (s *CatparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitCatparam(s)
	}
}

func (p *ParserParser) Catparam() (localctx ICatparamContext) {
	localctx = NewCatparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ParserParserRULE_catparam)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.Match(ParserParserRW_fileN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.Match(ParserParserTK_equ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(308)

		var _m = p.Match(ParserParserTK_path)

		localctx.(*CatparamContext).p1 = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*CatparamContext).result = strings.Trim((func() string {
		if localctx.(*CatparamContext).GetP1() == nil {
			return ""
		} else {
			return localctx.(*CatparamContext).GetP1().GetText()
		}
	}()), "\"")

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoginContext is an interface to support dynamic dispatch.
type ILoginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l token.
	GetL() antlr.Token

	// SetL sets the l token.
	SetL(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() ILoginparamsContext

	// SetP sets the p rule contexts.
	SetP(ILoginparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Login

	// SetResult sets the result attribute.
	SetResult(*commands.Login)

	// Getter signatures
	RW_login() antlr.TerminalNode
	Loginparams() ILoginparamsContext

	// IsLoginContext differentiates from other interfaces.
	IsLoginContext()
}

type LoginContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Login
	l      antlr.Token
	p      ILoginparamsContext
}

func NewEmptyLoginContext() *LoginContext {
	var p = new(LoginContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_login
	return p
}

func InitEmptyLoginContext(p *LoginContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_login
}

func (*LoginContext) IsLoginContext() {}

func NewLoginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoginContext {
	var p = new(LoginContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_login

	return p
}

func (s *LoginContext) GetParser() antlr.Parser { return s.parser }

func (s *LoginContext) GetL() antlr.Token { return s.l }

func (s *LoginContext) SetL(v antlr.Token) { s.l = v }

func (s *LoginContext) GetP() ILoginparamsContext { return s.p }

func (s *LoginContext) SetP(v ILoginparamsContext) { s.p = v }

func (s *LoginContext) GetResult() *commands.Login { return s.result }

func (s *LoginContext) SetResult(v *commands.Login) { s.result = v }

func (s *LoginContext) RW_login() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_login, 0)
}

func (s *LoginContext) Loginparams() ILoginparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoginparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoginparamsContext)
}

func (s *LoginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterLogin(s)
	}
}

func (s *LoginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitLogin(s)
	}
}

func (p *ParserParser) Login() (localctx ILoginContext) {
	localctx = NewLoginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ParserParserRULE_login)
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(311)

			var _m = p.Match(ParserParserRW_login)

			localctx.(*LoginContext).l = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)

			var _x = p.loginparams(0)

			localctx.(*LoginContext).p = _x
		}
		localctx.(*LoginContext).result = commands.NewLogin((func() int {
			if localctx.(*LoginContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*LoginContext).GetL().GetLine()
			}
		}()), (func() int {
			if localctx.(*LoginContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*LoginContext).GetL().GetColumn()
			}
		}()), localctx.(*LoginContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(315)

			var _m = p.Match(ParserParserRW_login)

			localctx.(*LoginContext).l = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*LoginContext).result = commands.NewLogin((func() int {
			if localctx.(*LoginContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*LoginContext).GetL().GetLine()
			}
		}()), (func() int {
			if localctx.(*LoginContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*LoginContext).GetL().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoginparamsContext is an interface to support dynamic dispatch.
type ILoginparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() ILoginparamsContext

	// GetP returns the p rule contexts.
	GetP() ILoginparamContext

	// SetL sets the l rule contexts.
	SetL(ILoginparamsContext)

	// SetP sets the p rule contexts.
	SetP(ILoginparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Loginparam() ILoginparamContext
	Loginparams() ILoginparamsContext

	// IsLoginparamsContext differentiates from other interfaces.
	IsLoginparamsContext()
}

type LoginparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      ILoginparamsContext
	p      ILoginparamContext
}

func NewEmptyLoginparamsContext() *LoginparamsContext {
	var p = new(LoginparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_loginparams
	return p
}

func InitEmptyLoginparamsContext(p *LoginparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_loginparams
}

func (*LoginparamsContext) IsLoginparamsContext() {}

func NewLoginparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoginparamsContext {
	var p = new(LoginparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_loginparams

	return p
}

func (s *LoginparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *LoginparamsContext) GetL() ILoginparamsContext { return s.l }

func (s *LoginparamsContext) GetP() ILoginparamContext { return s.p }

func (s *LoginparamsContext) SetL(v ILoginparamsContext) { s.l = v }

func (s *LoginparamsContext) SetP(v ILoginparamContext) { s.p = v }

func (s *LoginparamsContext) GetResult() map[string]string { return s.result }

func (s *LoginparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *LoginparamsContext) Loginparam() ILoginparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoginparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoginparamContext)
}

func (s *LoginparamsContext) Loginparams() ILoginparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoginparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoginparamsContext)
}

func (s *LoginparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoginparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoginparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterLoginparams(s)
	}
}

func (s *LoginparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitLoginparams(s)
	}
}

func (p *ParserParser) Loginparams() (localctx ILoginparamsContext) {
	return p.loginparams(0)
}

func (p *ParserParser) loginparams(_p int) (localctx ILoginparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewLoginparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILoginparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, ParserParserRULE_loginparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)

		var _x = p.Loginparam()

		localctx.(*LoginparamsContext).p = _x
	}
	localctx.(*LoginparamsContext).result = map[string]string{localctx.(*LoginparamsContext).GetP().GetResult()[0]: localctx.(*LoginparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLoginparamsContext(p, _parentctx, _parentState)
			localctx.(*LoginparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_loginparams)
			p.SetState(323)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(324)

				var _x = p.Loginparam()

				localctx.(*LoginparamsContext).p = _x
			}
			localctx.(*LoginparamsContext).SetResult(localctx.(*LoginparamsContext).GetL().GetResult())
			localctx.(*LoginparamsContext).result[localctx.(*LoginparamsContext).GetP().GetResult()[0]] = localctx.(*LoginparamsContext).GetP().GetResult()[1]

		}
		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoginparamContext is an interface to support dynamic dispatch.
type ILoginparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP1 returns the p1 token.
	GetP1() antlr.Token

	// GetP2 returns the p2 token.
	GetP2() antlr.Token

	// GetP3 returns the p3 token.
	GetP3() antlr.Token

	// SetP1 sets the p1 token.
	SetP1(antlr.Token)

	// SetP2 sets the p2 token.
	SetP2(antlr.Token)

	// SetP3 sets the p3 token.
	SetP3(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_user() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	RW_pass() antlr.TerminalNode
	TK_number() antlr.TerminalNode
	RW_id() antlr.TerminalNode

	// IsLoginparamContext differentiates from other interfaces.
	IsLoginparamContext()
}

type LoginparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	p1     antlr.Token
	p2     antlr.Token
	p3     antlr.Token
}

func NewEmptyLoginparamContext() *LoginparamContext {
	var p = new(LoginparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_loginparam
	return p
}

func InitEmptyLoginparamContext(p *LoginparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_loginparam
}

func (*LoginparamContext) IsLoginparamContext() {}

func NewLoginparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoginparamContext {
	var p = new(LoginparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_loginparam

	return p
}

func (s *LoginparamContext) GetParser() antlr.Parser { return s.parser }

func (s *LoginparamContext) GetP1() antlr.Token { return s.p1 }

func (s *LoginparamContext) GetP2() antlr.Token { return s.p2 }

func (s *LoginparamContext) GetP3() antlr.Token { return s.p3 }

func (s *LoginparamContext) SetP1(v antlr.Token) { s.p1 = v }

func (s *LoginparamContext) SetP2(v antlr.Token) { s.p2 = v }

func (s *LoginparamContext) SetP3(v antlr.Token) { s.p3 = v }

func (s *LoginparamContext) GetResult() []string { return s.result }

func (s *LoginparamContext) SetResult(v []string) { s.result = v }

func (s *LoginparamContext) RW_user() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_user, 0)
}

func (s *LoginparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *LoginparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *LoginparamContext) RW_pass() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_pass, 0)
}

func (s *LoginparamContext) TK_number() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_number, 0)
}

func (s *LoginparamContext) RW_id() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_id, 0)
}

func (s *LoginparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoginparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoginparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterLoginparam(s)
	}
}

func (s *LoginparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitLoginparam(s)
	}
}

func (p *ParserParser) Loginparam() (localctx ILoginparamContext) {
	localctx = NewLoginparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ParserParserRULE_loginparam)
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(332)
			p.Match(ParserParserRW_user)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*LoginparamContext).p1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*LoginparamContext).result = []string{"user", strings.Trim((func() string {
			if localctx.(*LoginparamContext).GetP1() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).GetP1().GetText()
			}
		}()), "\"")}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(336)
			p.Match(ParserParserRW_pass)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*LoginparamContext).p2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*LoginparamContext).result = []string{"pass", strings.Trim((func() string {
			if localctx.(*LoginparamContext).GetP2() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).GetP2().GetText()
			}
		}()), "\"")}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(340)
			p.Match(ParserParserRW_pass)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*LoginparamContext).p2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*LoginparamContext).result = []string{"pass", (func() string {
			if localctx.(*LoginparamContext).GetP2() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).GetP2().GetText()
			}
		}())}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(344)
			p.Match(ParserParserRW_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*LoginparamContext).p3 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*LoginparamContext).result = []string{"id", strings.ToUpper(strings.Trim((func() string {
			if localctx.(*LoginparamContext).GetP3() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).GetP3().GetText()
			}
		}()), "\""))}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogoutContext is an interface to support dynamic dispatch.
type ILogoutContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l token.
	GetL() antlr.Token

	// SetL sets the l token.
	SetL(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *commands.Logout

	// SetResult sets the result attribute.
	SetResult(*commands.Logout)

	// Getter signatures
	RW_logout() antlr.TerminalNode

	// IsLogoutContext differentiates from other interfaces.
	IsLogoutContext()
}

type LogoutContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Logout
	l      antlr.Token
}

func NewEmptyLogoutContext() *LogoutContext {
	var p = new(LogoutContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_logout
	return p
}

func InitEmptyLogoutContext(p *LogoutContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_logout
}

func (*LogoutContext) IsLogoutContext() {}

func NewLogoutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogoutContext {
	var p = new(LogoutContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_logout

	return p
}

func (s *LogoutContext) GetParser() antlr.Parser { return s.parser }

func (s *LogoutContext) GetL() antlr.Token { return s.l }

func (s *LogoutContext) SetL(v antlr.Token) { s.l = v }

func (s *LogoutContext) GetResult() *commands.Logout { return s.result }

func (s *LogoutContext) SetResult(v *commands.Logout) { s.result = v }

func (s *LogoutContext) RW_logout() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_logout, 0)
}

func (s *LogoutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogoutContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogoutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterLogout(s)
	}
}

func (s *LogoutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitLogout(s)
	}
}

func (p *ParserParser) Logout() (localctx ILogoutContext) {
	localctx = NewLogoutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ParserParserRULE_logout)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)

		var _m = p.Match(ParserParserRW_logout)

		localctx.(*LogoutContext).l = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*LogoutContext).result = commands.NewLogout((func() int {
		if localctx.(*LogoutContext).GetL() == nil {
			return 0
		} else {
			return localctx.(*LogoutContext).GetL().GetLine()
		}
	}()), (func() int {
		if localctx.(*LogoutContext).GetL() == nil {
			return 0
		} else {
			return localctx.(*LogoutContext).GetL().GetColumn()
		}
	}()))

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkgrpContext is an interface to support dynamic dispatch.
type IMkgrpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l token.
	GetL() antlr.Token

	// SetL sets the l token.
	SetL(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IMkgrpparamsContext

	// SetP sets the p rule contexts.
	SetP(IMkgrpparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Mkgrp

	// SetResult sets the result attribute.
	SetResult(*commands.Mkgrp)

	// Getter signatures
	RW_mkgrp() antlr.TerminalNode
	Mkgrpparams() IMkgrpparamsContext

	// IsMkgrpContext differentiates from other interfaces.
	IsMkgrpContext()
}

type MkgrpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Mkgrp
	l      antlr.Token
	p      IMkgrpparamsContext
}

func NewEmptyMkgrpContext() *MkgrpContext {
	var p = new(MkgrpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkgrp
	return p
}

func InitEmptyMkgrpContext(p *MkgrpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkgrp
}

func (*MkgrpContext) IsMkgrpContext() {}

func NewMkgrpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkgrpContext {
	var p = new(MkgrpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkgrp

	return p
}

func (s *MkgrpContext) GetParser() antlr.Parser { return s.parser }

func (s *MkgrpContext) GetL() antlr.Token { return s.l }

func (s *MkgrpContext) SetL(v antlr.Token) { s.l = v }

func (s *MkgrpContext) GetP() IMkgrpparamsContext { return s.p }

func (s *MkgrpContext) SetP(v IMkgrpparamsContext) { s.p = v }

func (s *MkgrpContext) GetResult() *commands.Mkgrp { return s.result }

func (s *MkgrpContext) SetResult(v *commands.Mkgrp) { s.result = v }

func (s *MkgrpContext) RW_mkgrp() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mkgrp, 0)
}

func (s *MkgrpContext) Mkgrpparams() IMkgrpparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkgrpparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkgrpparamsContext)
}

func (s *MkgrpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkgrpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkgrpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkgrp(s)
	}
}

func (s *MkgrpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkgrp(s)
	}
}

func (p *ParserParser) Mkgrp() (localctx IMkgrpContext) {
	localctx = NewMkgrpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ParserParserRULE_mkgrp)
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(353)

			var _m = p.Match(ParserParserRW_mkgrp)

			localctx.(*MkgrpContext).l = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)

			var _x = p.mkgrpparams(0)

			localctx.(*MkgrpContext).p = _x
		}
		localctx.(*MkgrpContext).result = commands.NewMkgrp((func() int {
			if localctx.(*MkgrpContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*MkgrpContext).GetL().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkgrpContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*MkgrpContext).GetL().GetColumn()
			}
		}()), localctx.(*MkgrpContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(357)

			var _m = p.Match(ParserParserRW_mkgrp)

			localctx.(*MkgrpContext).l = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkgrpContext).result = commands.NewMkgrp((func() int {
			if localctx.(*MkgrpContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*MkgrpContext).GetL().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkgrpContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*MkgrpContext).GetL().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkgrpparamsContext is an interface to support dynamic dispatch.
type IMkgrpparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IMkgrpparamsContext

	// GetP returns the p rule contexts.
	GetP() IMkgrpparamContext

	// SetL sets the l rule contexts.
	SetL(IMkgrpparamsContext)

	// SetP sets the p rule contexts.
	SetP(IMkgrpparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Mkgrpparam() IMkgrpparamContext
	Mkgrpparams() IMkgrpparamsContext

	// IsMkgrpparamsContext differentiates from other interfaces.
	IsMkgrpparamsContext()
}

type MkgrpparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IMkgrpparamsContext
	p      IMkgrpparamContext
}

func NewEmptyMkgrpparamsContext() *MkgrpparamsContext {
	var p = new(MkgrpparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkgrpparams
	return p
}

func InitEmptyMkgrpparamsContext(p *MkgrpparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkgrpparams
}

func (*MkgrpparamsContext) IsMkgrpparamsContext() {}

func NewMkgrpparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkgrpparamsContext {
	var p = new(MkgrpparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkgrpparams

	return p
}

func (s *MkgrpparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *MkgrpparamsContext) GetL() IMkgrpparamsContext { return s.l }

func (s *MkgrpparamsContext) GetP() IMkgrpparamContext { return s.p }

func (s *MkgrpparamsContext) SetL(v IMkgrpparamsContext) { s.l = v }

func (s *MkgrpparamsContext) SetP(v IMkgrpparamContext) { s.p = v }

func (s *MkgrpparamsContext) GetResult() map[string]string { return s.result }

func (s *MkgrpparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *MkgrpparamsContext) Mkgrpparam() IMkgrpparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkgrpparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkgrpparamContext)
}

func (s *MkgrpparamsContext) Mkgrpparams() IMkgrpparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkgrpparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkgrpparamsContext)
}

func (s *MkgrpparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkgrpparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkgrpparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkgrpparams(s)
	}
}

func (s *MkgrpparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkgrpparams(s)
	}
}

func (p *ParserParser) Mkgrpparams() (localctx IMkgrpparamsContext) {
	return p.mkgrpparams(0)
}

func (p *ParserParser) mkgrpparams(_p int) (localctx IMkgrpparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMkgrpparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMkgrpparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, ParserParserRULE_mkgrpparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)

		var _x = p.Mkgrpparam()

		localctx.(*MkgrpparamsContext).p = _x
	}
	localctx.(*MkgrpparamsContext).result = map[string]string{localctx.(*MkgrpparamsContext).GetP().GetResult()[0]: localctx.(*MkgrpparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMkgrpparamsContext(p, _parentctx, _parentState)
			localctx.(*MkgrpparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_mkgrpparams)
			p.SetState(365)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(366)

				var _x = p.Mkgrpparam()

				localctx.(*MkgrpparamsContext).p = _x
			}
			localctx.(*MkgrpparamsContext).SetResult(localctx.(*MkgrpparamsContext).GetL().GetResult())
			localctx.(*MkgrpparamsContext).result[localctx.(*MkgrpparamsContext).GetP().GetResult()[0]] = localctx.(*MkgrpparamsContext).GetP().GetResult()[1]

		}
		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkgrpparamContext is an interface to support dynamic dispatch.
type IMkgrpparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP1 returns the p1 token.
	GetP1() antlr.Token

	// GetP2 returns the p2 token.
	GetP2() antlr.Token

	// SetP1 sets the p1 token.
	SetP1(antlr.Token)

	// SetP2 sets the p2 token.
	SetP2(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_name() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	TK_number() antlr.TerminalNode

	// IsMkgrpparamContext differentiates from other interfaces.
	IsMkgrpparamContext()
}

type MkgrpparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	p1     antlr.Token
	p2     antlr.Token
}

func NewEmptyMkgrpparamContext() *MkgrpparamContext {
	var p = new(MkgrpparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkgrpparam
	return p
}

func InitEmptyMkgrpparamContext(p *MkgrpparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkgrpparam
}

func (*MkgrpparamContext) IsMkgrpparamContext() {}

func NewMkgrpparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkgrpparamContext {
	var p = new(MkgrpparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkgrpparam

	return p
}

func (s *MkgrpparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MkgrpparamContext) GetP1() antlr.Token { return s.p1 }

func (s *MkgrpparamContext) GetP2() antlr.Token { return s.p2 }

func (s *MkgrpparamContext) SetP1(v antlr.Token) { s.p1 = v }

func (s *MkgrpparamContext) SetP2(v antlr.Token) { s.p2 = v }

func (s *MkgrpparamContext) GetResult() []string { return s.result }

func (s *MkgrpparamContext) SetResult(v []string) { s.result = v }

func (s *MkgrpparamContext) RW_name() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_name, 0)
}

func (s *MkgrpparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *MkgrpparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *MkgrpparamContext) TK_number() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_number, 0)
}

func (s *MkgrpparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkgrpparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkgrpparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkgrpparam(s)
	}
}

func (s *MkgrpparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkgrpparam(s)
	}
}

func (p *ParserParser) Mkgrpparam() (localctx IMkgrpparamContext) {
	localctx = NewMkgrpparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ParserParserRULE_mkgrpparam)
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(374)
			p.Match(ParserParserRW_name)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*MkgrpparamContext).p1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkgrpparamContext).result = []string{"name", strings.Trim((func() string {
			if localctx.(*MkgrpparamContext).GetP1() == nil {
				return ""
			} else {
				return localctx.(*MkgrpparamContext).GetP1().GetText()
			}
		}()), "\"")}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(378)
			p.Match(ParserParserRW_name)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(379)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*MkgrpparamContext).p2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkgrpparamContext).result = []string{"name", (func() string {
			if localctx.(*MkgrpparamContext).GetP2() == nil {
				return ""
			} else {
				return localctx.(*MkgrpparamContext).GetP2().GetText()
			}
		}())}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRmgrpContext is an interface to support dynamic dispatch.
type IRmgrpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l token.
	GetL() antlr.Token

	// SetL sets the l token.
	SetL(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IRmgrpparamsContext

	// SetP sets the p rule contexts.
	SetP(IRmgrpparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Rmgrp

	// SetResult sets the result attribute.
	SetResult(*commands.Rmgrp)

	// Getter signatures
	RW_rmgrp() antlr.TerminalNode
	Rmgrpparams() IRmgrpparamsContext

	// IsRmgrpContext differentiates from other interfaces.
	IsRmgrpContext()
}

type RmgrpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Rmgrp
	l      antlr.Token
	p      IRmgrpparamsContext
}

func NewEmptyRmgrpContext() *RmgrpContext {
	var p = new(RmgrpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rmgrp
	return p
}

func InitEmptyRmgrpContext(p *RmgrpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rmgrp
}

func (*RmgrpContext) IsRmgrpContext() {}

func NewRmgrpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RmgrpContext {
	var p = new(RmgrpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_rmgrp

	return p
}

func (s *RmgrpContext) GetParser() antlr.Parser { return s.parser }

func (s *RmgrpContext) GetL() antlr.Token { return s.l }

func (s *RmgrpContext) SetL(v antlr.Token) { s.l = v }

func (s *RmgrpContext) GetP() IRmgrpparamsContext { return s.p }

func (s *RmgrpContext) SetP(v IRmgrpparamsContext) { s.p = v }

func (s *RmgrpContext) GetResult() *commands.Rmgrp { return s.result }

func (s *RmgrpContext) SetResult(v *commands.Rmgrp) { s.result = v }

func (s *RmgrpContext) RW_rmgrp() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_rmgrp, 0)
}

func (s *RmgrpContext) Rmgrpparams() IRmgrpparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmgrpparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmgrpparamsContext)
}

func (s *RmgrpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RmgrpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RmgrpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterRmgrp(s)
	}
}

func (s *RmgrpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitRmgrp(s)
	}
}

func (p *ParserParser) Rmgrp() (localctx IRmgrpContext) {
	localctx = NewRmgrpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ParserParserRULE_rmgrp)
	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(384)

			var _m = p.Match(ParserParserRW_rmgrp)

			localctx.(*RmgrpContext).l = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)

			var _x = p.rmgrpparams(0)

			localctx.(*RmgrpContext).p = _x
		}
		localctx.(*RmgrpContext).result = commands.NewRmgrp((func() int {
			if localctx.(*RmgrpContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*RmgrpContext).GetL().GetLine()
			}
		}()), (func() int {
			if localctx.(*RmgrpContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*RmgrpContext).GetL().GetColumn()
			}
		}()), localctx.(*RmgrpContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(388)

			var _m = p.Match(ParserParserRW_rmgrp)

			localctx.(*RmgrpContext).l = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RmgrpContext).result = commands.NewRmgrp((func() int {
			if localctx.(*RmgrpContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*RmgrpContext).GetL().GetLine()
			}
		}()), (func() int {
			if localctx.(*RmgrpContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*RmgrpContext).GetL().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRmgrpparamsContext is an interface to support dynamic dispatch.
type IRmgrpparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IRmgrpparamsContext

	// GetP returns the p rule contexts.
	GetP() IRmgrpparamContext

	// SetL sets the l rule contexts.
	SetL(IRmgrpparamsContext)

	// SetP sets the p rule contexts.
	SetP(IRmgrpparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Rmgrpparam() IRmgrpparamContext
	Rmgrpparams() IRmgrpparamsContext

	// IsRmgrpparamsContext differentiates from other interfaces.
	IsRmgrpparamsContext()
}

type RmgrpparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IRmgrpparamsContext
	p      IRmgrpparamContext
}

func NewEmptyRmgrpparamsContext() *RmgrpparamsContext {
	var p = new(RmgrpparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rmgrpparams
	return p
}

func InitEmptyRmgrpparamsContext(p *RmgrpparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rmgrpparams
}

func (*RmgrpparamsContext) IsRmgrpparamsContext() {}

func NewRmgrpparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RmgrpparamsContext {
	var p = new(RmgrpparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_rmgrpparams

	return p
}

func (s *RmgrpparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *RmgrpparamsContext) GetL() IRmgrpparamsContext { return s.l }

func (s *RmgrpparamsContext) GetP() IRmgrpparamContext { return s.p }

func (s *RmgrpparamsContext) SetL(v IRmgrpparamsContext) { s.l = v }

func (s *RmgrpparamsContext) SetP(v IRmgrpparamContext) { s.p = v }

func (s *RmgrpparamsContext) GetResult() map[string]string { return s.result }

func (s *RmgrpparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *RmgrpparamsContext) Rmgrpparam() IRmgrpparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmgrpparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmgrpparamContext)
}

func (s *RmgrpparamsContext) Rmgrpparams() IRmgrpparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmgrpparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmgrpparamsContext)
}

func (s *RmgrpparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RmgrpparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RmgrpparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterRmgrpparams(s)
	}
}

func (s *RmgrpparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitRmgrpparams(s)
	}
}

func (p *ParserParser) Rmgrpparams() (localctx IRmgrpparamsContext) {
	return p.rmgrpparams(0)
}

func (p *ParserParser) rmgrpparams(_p int) (localctx IRmgrpparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewRmgrpparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRmgrpparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, ParserParserRULE_rmgrpparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(393)

		var _x = p.Rmgrpparam()

		localctx.(*RmgrpparamsContext).p = _x
	}
	localctx.(*RmgrpparamsContext).result = map[string]string{localctx.(*RmgrpparamsContext).GetP().GetResult()[0]: localctx.(*RmgrpparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRmgrpparamsContext(p, _parentctx, _parentState)
			localctx.(*RmgrpparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_rmgrpparams)
			p.SetState(396)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(397)

				var _x = p.Rmgrpparam()

				localctx.(*RmgrpparamsContext).p = _x
			}
			localctx.(*RmgrpparamsContext).SetResult(localctx.(*RmgrpparamsContext).GetL().GetResult())
			localctx.(*RmgrpparamsContext).result[localctx.(*RmgrpparamsContext).GetP().GetResult()[0]] = localctx.(*RmgrpparamsContext).GetP().GetResult()[1]

		}
		p.SetState(404)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRmgrpparamContext is an interface to support dynamic dispatch.
type IRmgrpparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP1 returns the p1 token.
	GetP1() antlr.Token

	// GetP2 returns the p2 token.
	GetP2() antlr.Token

	// SetP1 sets the p1 token.
	SetP1(antlr.Token)

	// SetP2 sets the p2 token.
	SetP2(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_name() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	TK_number() antlr.TerminalNode

	// IsRmgrpparamContext differentiates from other interfaces.
	IsRmgrpparamContext()
}

type RmgrpparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	p1     antlr.Token
	p2     antlr.Token
}

func NewEmptyRmgrpparamContext() *RmgrpparamContext {
	var p = new(RmgrpparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rmgrpparam
	return p
}

func InitEmptyRmgrpparamContext(p *RmgrpparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rmgrpparam
}

func (*RmgrpparamContext) IsRmgrpparamContext() {}

func NewRmgrpparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RmgrpparamContext {
	var p = new(RmgrpparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_rmgrpparam

	return p
}

func (s *RmgrpparamContext) GetParser() antlr.Parser { return s.parser }

func (s *RmgrpparamContext) GetP1() antlr.Token { return s.p1 }

func (s *RmgrpparamContext) GetP2() antlr.Token { return s.p2 }

func (s *RmgrpparamContext) SetP1(v antlr.Token) { s.p1 = v }

func (s *RmgrpparamContext) SetP2(v antlr.Token) { s.p2 = v }

func (s *RmgrpparamContext) GetResult() []string { return s.result }

func (s *RmgrpparamContext) SetResult(v []string) { s.result = v }

func (s *RmgrpparamContext) RW_name() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_name, 0)
}

func (s *RmgrpparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *RmgrpparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *RmgrpparamContext) TK_number() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_number, 0)
}

func (s *RmgrpparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RmgrpparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RmgrpparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterRmgrpparam(s)
	}
}

func (s *RmgrpparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitRmgrpparam(s)
	}
}

func (p *ParserParser) Rmgrpparam() (localctx IRmgrpparamContext) {
	localctx = NewRmgrpparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ParserParserRULE_rmgrpparam)
	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(405)
			p.Match(ParserParserRW_name)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*RmgrpparamContext).p1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RmgrpparamContext).result = []string{"name", strings.Trim((func() string {
			if localctx.(*RmgrpparamContext).GetP1() == nil {
				return ""
			} else {
				return localctx.(*RmgrpparamContext).GetP1().GetText()
			}
		}()), "\"")}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(409)
			p.Match(ParserParserRW_name)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(411)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*RmgrpparamContext).p2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RmgrpparamContext).result = []string{"name", (func() string {
			if localctx.(*RmgrpparamContext).GetP2() == nil {
				return ""
			} else {
				return localctx.(*RmgrpparamContext).GetP2().GetText()
			}
		}())}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkusrContext is an interface to support dynamic dispatch.
type IMkusrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l token.
	GetL() antlr.Token

	// SetL sets the l token.
	SetL(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IMkusrparamsContext

	// SetP sets the p rule contexts.
	SetP(IMkusrparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Mkusr

	// SetResult sets the result attribute.
	SetResult(*commands.Mkusr)

	// Getter signatures
	RW_mkusr() antlr.TerminalNode
	Mkusrparams() IMkusrparamsContext

	// IsMkusrContext differentiates from other interfaces.
	IsMkusrContext()
}

type MkusrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Mkusr
	l      antlr.Token
	p      IMkusrparamsContext
}

func NewEmptyMkusrContext() *MkusrContext {
	var p = new(MkusrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkusr
	return p
}

func InitEmptyMkusrContext(p *MkusrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkusr
}

func (*MkusrContext) IsMkusrContext() {}

func NewMkusrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkusrContext {
	var p = new(MkusrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkusr

	return p
}

func (s *MkusrContext) GetParser() antlr.Parser { return s.parser }

func (s *MkusrContext) GetL() antlr.Token { return s.l }

func (s *MkusrContext) SetL(v antlr.Token) { s.l = v }

func (s *MkusrContext) GetP() IMkusrparamsContext { return s.p }

func (s *MkusrContext) SetP(v IMkusrparamsContext) { s.p = v }

func (s *MkusrContext) GetResult() *commands.Mkusr { return s.result }

func (s *MkusrContext) SetResult(v *commands.Mkusr) { s.result = v }

func (s *MkusrContext) RW_mkusr() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mkusr, 0)
}

func (s *MkusrContext) Mkusrparams() IMkusrparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkusrparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkusrparamsContext)
}

func (s *MkusrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkusrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkusrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkusr(s)
	}
}

func (s *MkusrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkusr(s)
	}
}

func (p *ParserParser) Mkusr() (localctx IMkusrContext) {
	localctx = NewMkusrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ParserParserRULE_mkusr)
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(415)

			var _m = p.Match(ParserParserRW_mkusr)

			localctx.(*MkusrContext).l = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)

			var _x = p.mkusrparams(0)

			localctx.(*MkusrContext).p = _x
		}
		localctx.(*MkusrContext).result = commands.NewMkusr((func() int {
			if localctx.(*MkusrContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*MkusrContext).GetL().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkusrContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*MkusrContext).GetL().GetColumn()
			}
		}()), localctx.(*MkusrContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(419)

			var _m = p.Match(ParserParserRW_mkusr)

			localctx.(*MkusrContext).l = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkusrContext).result = commands.NewMkusr((func() int {
			if localctx.(*MkusrContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*MkusrContext).GetL().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkusrContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*MkusrContext).GetL().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkusrparamsContext is an interface to support dynamic dispatch.
type IMkusrparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IMkusrparamsContext

	// GetP returns the p rule contexts.
	GetP() IMkusrparamContext

	// SetL sets the l rule contexts.
	SetL(IMkusrparamsContext)

	// SetP sets the p rule contexts.
	SetP(IMkusrparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Mkusrparam() IMkusrparamContext
	Mkusrparams() IMkusrparamsContext

	// IsMkusrparamsContext differentiates from other interfaces.
	IsMkusrparamsContext()
}

type MkusrparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IMkusrparamsContext
	p      IMkusrparamContext
}

func NewEmptyMkusrparamsContext() *MkusrparamsContext {
	var p = new(MkusrparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkusrparams
	return p
}

func InitEmptyMkusrparamsContext(p *MkusrparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkusrparams
}

func (*MkusrparamsContext) IsMkusrparamsContext() {}

func NewMkusrparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkusrparamsContext {
	var p = new(MkusrparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkusrparams

	return p
}

func (s *MkusrparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *MkusrparamsContext) GetL() IMkusrparamsContext { return s.l }

func (s *MkusrparamsContext) GetP() IMkusrparamContext { return s.p }

func (s *MkusrparamsContext) SetL(v IMkusrparamsContext) { s.l = v }

func (s *MkusrparamsContext) SetP(v IMkusrparamContext) { s.p = v }

func (s *MkusrparamsContext) GetResult() map[string]string { return s.result }

func (s *MkusrparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *MkusrparamsContext) Mkusrparam() IMkusrparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkusrparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkusrparamContext)
}

func (s *MkusrparamsContext) Mkusrparams() IMkusrparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkusrparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkusrparamsContext)
}

func (s *MkusrparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkusrparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkusrparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkusrparams(s)
	}
}

func (s *MkusrparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkusrparams(s)
	}
}

func (p *ParserParser) Mkusrparams() (localctx IMkusrparamsContext) {
	return p.mkusrparams(0)
}

func (p *ParserParser) mkusrparams(_p int) (localctx IMkusrparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMkusrparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMkusrparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, ParserParserRULE_mkusrparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(424)

		var _x = p.Mkusrparam()

		localctx.(*MkusrparamsContext).p = _x
	}
	localctx.(*MkusrparamsContext).result = map[string]string{localctx.(*MkusrparamsContext).GetP().GetResult()[0]: localctx.(*MkusrparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMkusrparamsContext(p, _parentctx, _parentState)
			localctx.(*MkusrparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_mkusrparams)
			p.SetState(427)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(428)

				var _x = p.Mkusrparam()

				localctx.(*MkusrparamsContext).p = _x
			}
			localctx.(*MkusrparamsContext).SetResult(localctx.(*MkusrparamsContext).GetL().GetResult())
			localctx.(*MkusrparamsContext).result[localctx.(*MkusrparamsContext).GetP().GetResult()[0]] = localctx.(*MkusrparamsContext).GetP().GetResult()[1]

		}
		p.SetState(435)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkusrparamContext is an interface to support dynamic dispatch.
type IMkusrparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP1 returns the p1 token.
	GetP1() antlr.Token

	// GetP2 returns the p2 token.
	GetP2() antlr.Token

	// GetP3 returns the p3 token.
	GetP3() antlr.Token

	// GetP4 returns the p4 token.
	GetP4() antlr.Token

	// GetP5 returns the p5 token.
	GetP5() antlr.Token

	// GetP6 returns the p6 token.
	GetP6() antlr.Token

	// SetP1 sets the p1 token.
	SetP1(antlr.Token)

	// SetP2 sets the p2 token.
	SetP2(antlr.Token)

	// SetP3 sets the p3 token.
	SetP3(antlr.Token)

	// SetP4 sets the p4 token.
	SetP4(antlr.Token)

	// SetP5 sets the p5 token.
	SetP5(antlr.Token)

	// SetP6 sets the p6 token.
	SetP6(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_user() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	TK_number() antlr.TerminalNode
	RW_pass() antlr.TerminalNode
	RW_grp() antlr.TerminalNode

	// IsMkusrparamContext differentiates from other interfaces.
	IsMkusrparamContext()
}

type MkusrparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	p1     antlr.Token
	p2     antlr.Token
	p3     antlr.Token
	p4     antlr.Token
	p5     antlr.Token
	p6     antlr.Token
}

func NewEmptyMkusrparamContext() *MkusrparamContext {
	var p = new(MkusrparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkusrparam
	return p
}

func InitEmptyMkusrparamContext(p *MkusrparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkusrparam
}

func (*MkusrparamContext) IsMkusrparamContext() {}

func NewMkusrparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkusrparamContext {
	var p = new(MkusrparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkusrparam

	return p
}

func (s *MkusrparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MkusrparamContext) GetP1() antlr.Token { return s.p1 }

func (s *MkusrparamContext) GetP2() antlr.Token { return s.p2 }

func (s *MkusrparamContext) GetP3() antlr.Token { return s.p3 }

func (s *MkusrparamContext) GetP4() antlr.Token { return s.p4 }

func (s *MkusrparamContext) GetP5() antlr.Token { return s.p5 }

func (s *MkusrparamContext) GetP6() antlr.Token { return s.p6 }

func (s *MkusrparamContext) SetP1(v antlr.Token) { s.p1 = v }

func (s *MkusrparamContext) SetP2(v antlr.Token) { s.p2 = v }

func (s *MkusrparamContext) SetP3(v antlr.Token) { s.p3 = v }

func (s *MkusrparamContext) SetP4(v antlr.Token) { s.p4 = v }

func (s *MkusrparamContext) SetP5(v antlr.Token) { s.p5 = v }

func (s *MkusrparamContext) SetP6(v antlr.Token) { s.p6 = v }

func (s *MkusrparamContext) GetResult() []string { return s.result }

func (s *MkusrparamContext) SetResult(v []string) { s.result = v }

func (s *MkusrparamContext) RW_user() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_user, 0)
}

func (s *MkusrparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *MkusrparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *MkusrparamContext) TK_number() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_number, 0)
}

func (s *MkusrparamContext) RW_pass() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_pass, 0)
}

func (s *MkusrparamContext) RW_grp() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_grp, 0)
}

func (s *MkusrparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkusrparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkusrparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkusrparam(s)
	}
}

func (s *MkusrparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkusrparam(s)
	}
}

func (p *ParserParser) Mkusrparam() (localctx IMkusrparamContext) {
	localctx = NewMkusrparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ParserParserRULE_mkusrparam)
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(436)
			p.Match(ParserParserRW_user)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(437)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(438)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*MkusrparamContext).p1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkusrparamContext).result = []string{"user", strings.Trim((func() string {
			if localctx.(*MkusrparamContext).GetP1() == nil {
				return ""
			} else {
				return localctx.(*MkusrparamContext).GetP1().GetText()
			}
		}()), "\"")}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(440)
			p.Match(ParserParserRW_user)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(441)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(442)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*MkusrparamContext).p2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkusrparamContext).result = []string{"user", (func() string {
			if localctx.(*MkusrparamContext).GetP2() == nil {
				return ""
			} else {
				return localctx.(*MkusrparamContext).GetP2().GetText()
			}
		}())}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(444)
			p.Match(ParserParserRW_pass)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(446)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*MkusrparamContext).p3 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkusrparamContext).result = []string{"pass", strings.Trim((func() string {
			if localctx.(*MkusrparamContext).GetP3() == nil {
				return ""
			} else {
				return localctx.(*MkusrparamContext).GetP3().GetText()
			}
		}()), "\"")}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(448)
			p.Match(ParserParserRW_pass)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(449)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(450)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*MkusrparamContext).p4 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkusrparamContext).result = []string{"pass", (func() string {
			if localctx.(*MkusrparamContext).GetP4() == nil {
				return ""
			} else {
				return localctx.(*MkusrparamContext).GetP4().GetText()
			}
		}())}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(452)
			p.Match(ParserParserRW_grp)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(453)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(454)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*MkusrparamContext).p5 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkusrparamContext).result = []string{"grp", strings.Trim((func() string {
			if localctx.(*MkusrparamContext).GetP5() == nil {
				return ""
			} else {
				return localctx.(*MkusrparamContext).GetP5().GetText()
			}
		}()), "\"")}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(456)
			p.Match(ParserParserRW_grp)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(458)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*MkusrparamContext).p6 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkusrparamContext).result = []string{"grp", (func() string {
			if localctx.(*MkusrparamContext).GetP6() == nil {
				return ""
			} else {
				return localctx.(*MkusrparamContext).GetP6().GetText()
			}
		}())}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *ParserParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *CommandsContext = nil
		if localctx != nil {
			t = localctx.(*CommandsContext)
		}
		return p.Commands_Sempred(t, predIndex)

	case 4:
		var t *MkdiskparamsContext = nil
		if localctx != nil {
			t = localctx.(*MkdiskparamsContext)
		}
		return p.Mkdiskparams_Sempred(t, predIndex)

	case 8:
		var t *FdiskparamsContext = nil
		if localctx != nil {
			t = localctx.(*FdiskparamsContext)
		}
		return p.Fdiskparams_Sempred(t, predIndex)

	case 11:
		var t *MountparamsContext = nil
		if localctx != nil {
			t = localctx.(*MountparamsContext)
		}
		return p.Mountparams_Sempred(t, predIndex)

	case 15:
		var t *MkfsparamsContext = nil
		if localctx != nil {
			t = localctx.(*MkfsparamsContext)
		}
		return p.Mkfsparams_Sempred(t, predIndex)

	case 18:
		var t *CatparamsContext = nil
		if localctx != nil {
			t = localctx.(*CatparamsContext)
		}
		return p.Catparams_Sempred(t, predIndex)

	case 21:
		var t *LoginparamsContext = nil
		if localctx != nil {
			t = localctx.(*LoginparamsContext)
		}
		return p.Loginparams_Sempred(t, predIndex)

	case 25:
		var t *MkgrpparamsContext = nil
		if localctx != nil {
			t = localctx.(*MkgrpparamsContext)
		}
		return p.Mkgrpparams_Sempred(t, predIndex)

	case 28:
		var t *RmgrpparamsContext = nil
		if localctx != nil {
			t = localctx.(*RmgrpparamsContext)
		}
		return p.Rmgrpparams_Sempred(t, predIndex)

	case 31:
		var t *MkusrparamsContext = nil
		if localctx != nil {
			t = localctx.(*MkusrparamsContext)
		}
		return p.Mkusrparams_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ParserParser) Commands_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Mkdiskparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Fdiskparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Mountparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Mkfsparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Catparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Loginparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Mkgrpparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Rmgrpparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Mkusrparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
