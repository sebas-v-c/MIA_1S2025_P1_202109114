// Code generated from /home/zibas/Documents/USAC/SEMESTRE-9/MIA/MIA_1S2025_P1_202109114/backend/Language/Parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
    "", "", "", "", "'='", "", "", "'\\n'",
  }
  staticData.SymbolicNames = []string{
    "", "RW_mkdisk", "RW_rmdisk", "RW_fdisk", "RW_mount", "RW_mounted", 
    "RW_mkfs", "RW_cat", "RW_login", "RW_logout", "RW_mkgrp", "RW_rmgrp", 
    "RW_mkusr", "RW_rmusr", "RW_chgrp", "RW_mkfile", "RW_mkdir", "RW_rep", 
    "RW_size", "RW_fit", "RW_unit", "RW_driveletter", "RW_name", "RW_type", 
    "RW_delete", "RW_add", "RW_id", "RW_fs", "RW_user", "RW_pass", "RW_grp", 
    "RW_path", "RW_r", "RW_cont", "RW_fileN", "RW_destino", "RW_ugo", "RW_ruta", 
    "TK_fit", "TK_unit", "TK_type", "TK_fs", "TK_number", "TK_id", "TK_ext", 
    "TK_path", "TK_equ", "TK_sym", "COMMENTARY", "NEWLINE", "UNUSED_",
  }
  staticData.RuleNames = []string{
    "init", "commands", "command", "mkdisk", "mkdiskparams", "mkdiskparam", 
    "rmdisk", "fdisk", "fdiskparams", "fdiskparam", "mount", "mountparams", 
    "mountparam",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 50, 188, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 
	0, 33, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 43, 
	8, 1, 10, 1, 12, 1, 46, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 
	1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 60, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 
	3, 1, 3, 3, 3, 68, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 
	5, 4, 78, 8, 4, 10, 4, 12, 4, 81, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 
	5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 99, 
	8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 108, 8, 6, 1, 7, 
	1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 116, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 
	1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 126, 8, 8, 10, 8, 12, 8, 129, 9, 8, 1, 9, 
	1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 
	1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 
	155, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 163, 8, 10, 
	1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 173, 8, 
	11, 10, 11, 12, 11, 176, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 
	1, 12, 1, 12, 3, 12, 186, 8, 12, 1, 12, 0, 4, 2, 8, 16, 22, 13, 0, 2, 4, 
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 0, 195, 0, 32, 1, 0, 0, 0, 2, 
	34, 1, 0, 0, 0, 4, 59, 1, 0, 0, 0, 6, 67, 1, 0, 0, 0, 8, 69, 1, 0, 0, 0, 
	10, 98, 1, 0, 0, 0, 12, 107, 1, 0, 0, 0, 14, 115, 1, 0, 0, 0, 16, 117, 
	1, 0, 0, 0, 18, 154, 1, 0, 0, 0, 20, 162, 1, 0, 0, 0, 22, 164, 1, 0, 0, 
	0, 24, 185, 1, 0, 0, 0, 26, 27, 3, 2, 1, 0, 27, 28, 5, 0, 0, 1, 28, 29, 
	6, 0, -1, 0, 29, 33, 1, 0, 0, 0, 30, 31, 5, 0, 0, 1, 31, 33, 6, 0, -1, 
	0, 32, 26, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 33, 1, 1, 0, 0, 0, 34, 35, 6, 
	1, -1, 0, 35, 36, 3, 4, 2, 0, 36, 37, 6, 1, -1, 0, 37, 44, 1, 0, 0, 0, 
	38, 39, 10, 2, 0, 0, 39, 40, 3, 4, 2, 0, 40, 41, 6, 1, -1, 0, 41, 43, 1, 
	0, 0, 0, 42, 38, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 
	45, 1, 0, 0, 0, 45, 3, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 48, 3, 6, 3, 
	0, 48, 49, 6, 2, -1, 0, 49, 60, 1, 0, 0, 0, 50, 51, 3, 12, 6, 0, 51, 52, 
	6, 2, -1, 0, 52, 60, 1, 0, 0, 0, 53, 54, 3, 14, 7, 0, 54, 55, 6, 2, -1, 
	0, 55, 60, 1, 0, 0, 0, 56, 57, 3, 20, 10, 0, 57, 58, 6, 2, -1, 0, 58, 60, 
	1, 0, 0, 0, 59, 47, 1, 0, 0, 0, 59, 50, 1, 0, 0, 0, 59, 53, 1, 0, 0, 0, 
	59, 56, 1, 0, 0, 0, 60, 5, 1, 0, 0, 0, 61, 62, 5, 1, 0, 0, 62, 63, 3, 8, 
	4, 0, 63, 64, 6, 3, -1, 0, 64, 68, 1, 0, 0, 0, 65, 66, 5, 1, 0, 0, 66, 
	68, 6, 3, -1, 0, 67, 61, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 68, 7, 1, 0, 0, 
	0, 69, 70, 6, 4, -1, 0, 70, 71, 3, 10, 5, 0, 71, 72, 6, 4, -1, 0, 72, 79, 
	1, 0, 0, 0, 73, 74, 10, 2, 0, 0, 74, 75, 3, 10, 5, 0, 75, 76, 6, 4, -1, 
	0, 76, 78, 1, 0, 0, 0, 77, 73, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 
	1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 9, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 
	82, 83, 5, 18, 0, 0, 83, 84, 5, 46, 0, 0, 84, 85, 5, 42, 0, 0, 85, 99, 
	6, 5, -1, 0, 86, 87, 5, 19, 0, 0, 87, 88, 5, 46, 0, 0, 88, 89, 5, 38, 0, 
	0, 89, 99, 6, 5, -1, 0, 90, 91, 5, 20, 0, 0, 91, 92, 5, 46, 0, 0, 92, 93, 
	5, 39, 0, 0, 93, 99, 6, 5, -1, 0, 94, 95, 5, 31, 0, 0, 95, 96, 5, 46, 0, 
	0, 96, 97, 5, 45, 0, 0, 97, 99, 6, 5, -1, 0, 98, 82, 1, 0, 0, 0, 98, 86, 
	1, 0, 0, 0, 98, 90, 1, 0, 0, 0, 98, 94, 1, 0, 0, 0, 99, 11, 1, 0, 0, 0, 
	100, 101, 5, 2, 0, 0, 101, 102, 5, 31, 0, 0, 102, 103, 5, 46, 0, 0, 103, 
	104, 5, 45, 0, 0, 104, 108, 6, 6, -1, 0, 105, 106, 5, 2, 0, 0, 106, 108, 
	6, 6, -1, 0, 107, 100, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 13, 1, 0, 
	0, 0, 109, 110, 5, 3, 0, 0, 110, 111, 3, 16, 8, 0, 111, 112, 6, 7, -1, 
	0, 112, 116, 1, 0, 0, 0, 113, 114, 5, 3, 0, 0, 114, 116, 6, 7, -1, 0, 115, 
	109, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 15, 1, 0, 0, 0, 117, 118, 6, 
	8, -1, 0, 118, 119, 3, 18, 9, 0, 119, 120, 6, 8, -1, 0, 120, 127, 1, 0, 
	0, 0, 121, 122, 10, 2, 0, 0, 122, 123, 3, 18, 9, 0, 123, 124, 6, 8, -1, 
	0, 124, 126, 1, 0, 0, 0, 125, 121, 1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127, 
	125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 17, 1, 0, 0, 0, 129, 127, 1, 
	0, 0, 0, 130, 131, 5, 18, 0, 0, 131, 132, 5, 46, 0, 0, 132, 133, 5, 42, 
	0, 0, 133, 155, 6, 9, -1, 0, 134, 135, 5, 20, 0, 0, 135, 136, 5, 46, 0, 
	0, 136, 137, 5, 39, 0, 0, 137, 155, 6, 9, -1, 0, 138, 139, 5, 31, 0, 0, 
	139, 140, 5, 46, 0, 0, 140, 141, 5, 45, 0, 0, 141, 155, 6, 9, -1, 0, 142, 
	143, 5, 23, 0, 0, 143, 144, 5, 46, 0, 0, 144, 145, 5, 40, 0, 0, 145, 155, 
	6, 9, -1, 0, 146, 147, 5, 19, 0, 0, 147, 148, 5, 46, 0, 0, 148, 149, 5, 
	38, 0, 0, 149, 155, 6, 9, -1, 0, 150, 151, 5, 22, 0, 0, 151, 152, 5, 46, 
	0, 0, 152, 153, 5, 43, 0, 0, 153, 155, 6, 9, -1, 0, 154, 130, 1, 0, 0, 
	0, 154, 134, 1, 0, 0, 0, 154, 138, 1, 0, 0, 0, 154, 142, 1, 0, 0, 0, 154, 
	146, 1, 0, 0, 0, 154, 150, 1, 0, 0, 0, 155, 19, 1, 0, 0, 0, 156, 157, 5, 
	4, 0, 0, 157, 158, 3, 22, 11, 0, 158, 159, 6, 10, -1, 0, 159, 163, 1, 0, 
	0, 0, 160, 161, 5, 4, 0, 0, 161, 163, 6, 10, -1, 0, 162, 156, 1, 0, 0, 
	0, 162, 160, 1, 0, 0, 0, 163, 21, 1, 0, 0, 0, 164, 165, 6, 11, -1, 0, 165, 
	166, 3, 24, 12, 0, 166, 167, 6, 11, -1, 0, 167, 174, 1, 0, 0, 0, 168, 169, 
	10, 2, 0, 0, 169, 170, 3, 24, 12, 0, 170, 171, 6, 11, -1, 0, 171, 173, 
	1, 0, 0, 0, 172, 168, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174, 172, 1, 0, 
	0, 0, 174, 175, 1, 0, 0, 0, 175, 23, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 
	177, 178, 5, 31, 0, 0, 178, 179, 5, 46, 0, 0, 179, 180, 5, 45, 0, 0, 180, 
	186, 6, 12, -1, 0, 181, 182, 5, 22, 0, 0, 182, 183, 5, 46, 0, 0, 183, 184, 
	5, 43, 0, 0, 184, 186, 6, 12, -1, 0, 185, 177, 1, 0, 0, 0, 185, 181, 1, 
	0, 0, 0, 186, 25, 1, 0, 0, 0, 13, 32, 44, 59, 67, 79, 98, 107, 115, 127, 
	154, 162, 174, 185,
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
	ParserParserEOF = antlr.TokenEOF
	ParserParserRW_mkdisk = 1
	ParserParserRW_rmdisk = 2
	ParserParserRW_fdisk = 3
	ParserParserRW_mount = 4
	ParserParserRW_mounted = 5
	ParserParserRW_mkfs = 6
	ParserParserRW_cat = 7
	ParserParserRW_login = 8
	ParserParserRW_logout = 9
	ParserParserRW_mkgrp = 10
	ParserParserRW_rmgrp = 11
	ParserParserRW_mkusr = 12
	ParserParserRW_rmusr = 13
	ParserParserRW_chgrp = 14
	ParserParserRW_mkfile = 15
	ParserParserRW_mkdir = 16
	ParserParserRW_rep = 17
	ParserParserRW_size = 18
	ParserParserRW_fit = 19
	ParserParserRW_unit = 20
	ParserParserRW_driveletter = 21
	ParserParserRW_name = 22
	ParserParserRW_type = 23
	ParserParserRW_delete = 24
	ParserParserRW_add = 25
	ParserParserRW_id = 26
	ParserParserRW_fs = 27
	ParserParserRW_user = 28
	ParserParserRW_pass = 29
	ParserParserRW_grp = 30
	ParserParserRW_path = 31
	ParserParserRW_r = 32
	ParserParserRW_cont = 33
	ParserParserRW_fileN = 34
	ParserParserRW_destino = 35
	ParserParserRW_ugo = 36
	ParserParserRW_ruta = 37
	ParserParserTK_fit = 38
	ParserParserTK_unit = 39
	ParserParserTK_type = 40
	ParserParserTK_fs = 41
	ParserParserTK_number = 42
	ParserParserTK_id = 43
	ParserParserTK_ext = 44
	ParserParserTK_path = 45
	ParserParserTK_equ = 46
	ParserParserTK_sym = 47
	ParserParserCOMMENTARY = 48
	ParserParserNEWLINE = 49
	ParserParserUNUSED_ = 50
)

// ParserParser rules.
const (
	ParserParserRULE_init = 0
	ParserParserRULE_commands = 1
	ParserParserRULE_command = 2
	ParserParserRULE_mkdisk = 3
	ParserParserRULE_mkdiskparams = 4
	ParserParserRULE_mkdiskparam = 5
	ParserParserRULE_rmdisk = 6
	ParserParserRULE_fdisk = 7
	ParserParserRULE_fdiskparams = 8
	ParserParserRULE_fdiskparam = 9
	ParserParserRULE_mount = 10
	ParserParserRULE_mountparams = 11
	ParserParserRULE_mountparam = 12
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
	c ICommandsContext 
}

func NewEmptyInitContext() *InitContext {
	var p = new(InitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_init
	return p
}

func InitEmptyInitContext(p *InitContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandsContext); ok {
			t = ctx.(antlr.RuleContext);
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




func (p *ParserParser) Init() (localctx IInitContext) {
	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ParserParserRULE_init)
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_mkdisk, ParserParserRW_rmdisk, ParserParserRW_fdisk, ParserParserRW_mount:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(26)

			var _x = p.commands(0)

			localctx.(*InitContext).c = _x
		}
		{
			p.SetState(27)
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
			p.SetState(30)
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
	l ICommandsContext 
	c ICommandContext 
}

func NewEmptyCommandsContext() *CommandsContext {
	var p = new(CommandsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_commands
	return p
}

func InitEmptyCommandsContext(p *CommandsContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *CommandsContext) Commands() ICommandsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandsContext); ok {
			t = ctx.(antlr.RuleContext);
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
		p.SetState(35)

		var _x = p.Command()


		localctx.(*CommandsContext).c = _x
	}
	localctx.(*CommandsContext).result = []interfaces.Command{localctx.(*CommandsContext).GetC().GetResult()}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(44)
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
			p.SetState(38)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(39)

				var _x = p.Command()


				localctx.(*CommandsContext).c = _x
			}
			localctx.(*CommandsContext).SetResult( localctx.(*CommandsContext).GetL().GetResult()); localctx.(*CommandsContext).result = append(localctx.(*CommandsContext).result, localctx.(*CommandsContext).GetC().GetResult())


		}
		p.SetState(46)
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


	// SetC1 sets the c1 rule contexts.
	SetC1(IMkdiskContext)

	// SetC2 sets the c2 rule contexts.
	SetC2(IRmdiskContext)

	// SetC3 sets the c3 rule contexts.
	SetC3(IFdiskContext)

	// SetC4 sets the c4 rule contexts.
	SetC4(IMountContext)


	// GetResult returns the result attribute.
	GetResult() interfaces.Command


	// SetResult sets the result attribute.
	SetResult(interfaces.Command)


	// Getter signatures
	Mkdisk() IMkdiskContext
	Rmdisk() IRmdiskContext
	Fdisk() IFdiskContext
	Mount() IMountContext

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result interfaces.Command
	c1 IMkdiskContext 
	c2 IRmdiskContext 
	c3 IFdiskContext 
	c4 IMountContext 
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext)  {
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


func (s *CommandContext) SetC1(v IMkdiskContext) { s.c1 = v }

func (s *CommandContext) SetC2(v IRmdiskContext) { s.c2 = v }

func (s *CommandContext) SetC3(v IFdiskContext) { s.c3 = v }

func (s *CommandContext) SetC4(v IMountContext) { s.c4 = v }


func (s *CommandContext) GetResult() interfaces.Command { return s.result }


func (s *CommandContext) SetResult(v interfaces.Command) { s.result = v }


func (s *CommandContext) Mkdisk() IMkdiskContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdiskContext)
}

func (s *CommandContext) Rmdisk() IRmdiskContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmdiskContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmdiskContext)
}

func (s *CommandContext) Fdisk() IFdiskContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFdiskContext)
}

func (s *CommandContext) Mount() IMountContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *ParserParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ParserParserRULE_command)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_mkdisk:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)

			var _x = p.Mkdisk()


			localctx.(*CommandContext).c1 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC1().GetResult()


	case ParserParserRW_rmdisk:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)

			var _x = p.Rmdisk()


			localctx.(*CommandContext).c2 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC2().GetResult()


	case ParserParserRW_fdisk:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(53)

			var _x = p.Fdisk()


			localctx.(*CommandContext).c3 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC3().GetResult()


	case ParserParserRW_mount:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)

			var _x = p.Mount()


			localctx.(*CommandContext).c4 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC4().GetResult()



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
	m antlr.Token
	p IMkdiskparamsContext 
}

func NewEmptyMkdiskContext() *MkdiskContext {
	var p = new(MkdiskContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdisk
	return p
}

func InitEmptyMkdiskContext(p *MkdiskContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskparamsContext); ok {
			t = ctx.(antlr.RuleContext);
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




func (p *ParserParser) Mkdisk() (localctx IMkdiskContext) {
	localctx = NewMkdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ParserParserRULE_mkdisk)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)

			var _m = p.Match(ParserParserRW_mkdisk)

			localctx.(*MkdiskContext).m = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(62)

			var _x = p.mkdiskparams(0)

			localctx.(*MkdiskContext).p = _x
		}
		localctx.(*MkdiskContext).result = commands.NewMkdisk((func() int { if localctx.(*MkdiskContext).GetM() == nil { return 0 } else { return localctx.(*MkdiskContext).GetM().GetLine() }}()), (func() int { if localctx.(*MkdiskContext).GetM() == nil { return 0 } else { return localctx.(*MkdiskContext).GetM().GetColumn() }}()), localctx.(*MkdiskContext).GetP().GetResult())


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)

			var _m = p.Match(ParserParserRW_mkdisk)

			localctx.(*MkdiskContext).m = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*MkdiskContext).result = commands.NewMkdisk((func() int { if localctx.(*MkdiskContext).GetM() == nil { return 0 } else { return localctx.(*MkdiskContext).GetM().GetLine() }}()), (func() int { if localctx.(*MkdiskContext).GetM() == nil { return 0 } else { return localctx.(*MkdiskContext).GetM().GetColumn() }}()), map[string]string{})

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
	l IMkdiskparamsContext 
	p IMkdiskparamContext 
}

func NewEmptyMkdiskparamsContext() *MkdiskparamsContext {
	var p = new(MkdiskparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdiskparams
	return p
}

func InitEmptyMkdiskparamsContext(p *MkdiskparamsContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskparamContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdiskparamContext)
}

func (s *MkdiskparamsContext) Mkdiskparams() IMkdiskparamsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskparamsContext); ok {
			t = ctx.(antlr.RuleContext);
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
		p.SetState(70)

		var _x = p.Mkdiskparam()


		localctx.(*MkdiskparamsContext).p = _x
	}
	localctx.(*MkdiskparamsContext).result = map[string]string{localctx.(*MkdiskparamsContext).GetP().GetResult()[0]: localctx.(*MkdiskparamsContext).GetP().GetResult()[1] }

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(79)
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
			p.SetState(73)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(74)

				var _x = p.Mkdiskparam()


				localctx.(*MkdiskparamsContext).p = _x
			}
			localctx.(*MkdiskparamsContext).SetResult( localctx.(*MkdiskparamsContext).GetL().GetResult()); localctx.(*MkdiskparamsContext).result[localctx.(*MkdiskparamsContext).GetP().GetResult()[0]] = localctx.(*MkdiskparamsContext).GetP().GetResult()[1]


		}
		p.SetState(81)
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
	v1 antlr.Token
	v2 antlr.Token
	v3 antlr.Token
	v4 antlr.Token
}

func NewEmptyMkdiskparamContext() *MkdiskparamContext {
	var p = new(MkdiskparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdiskparam
	return p
}

func InitEmptyMkdiskparamContext(p *MkdiskparamContext)  {
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




func (p *ParserParser) Mkdiskparam() (localctx IMkdiskparamContext) {
	localctx = NewMkdiskparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ParserParserRULE_mkdiskparam)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_size:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Match(ParserParserRW_size)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(84)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*MkdiskparamContext).v1 = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*MkdiskparamContext).result = []string{"size", (func() string { if localctx.(*MkdiskparamContext).GetV1() == nil { return "" } else { return localctx.(*MkdiskparamContext).GetV1().GetText() }}())}


	case ParserParserRW_fit:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
			p.Match(ParserParserRW_fit)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(88)

			var _m = p.Match(ParserParserTK_fit)

			localctx.(*MkdiskparamContext).v2 = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*MkdiskparamContext).result = []string{"fit", (func() string { if localctx.(*MkdiskparamContext).GetV2() == nil { return "" } else { return localctx.(*MkdiskparamContext).GetV2().GetText() }}())}


	case ParserParserRW_unit:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.Match(ParserParserRW_unit)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(92)

			var _m = p.Match(ParserParserTK_unit)

			localctx.(*MkdiskparamContext).v3 = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*MkdiskparamContext).result = []string{"unit", (func() string { if localctx.(*MkdiskparamContext).GetV3() == nil { return "" } else { return localctx.(*MkdiskparamContext).GetV3().GetText() }}())}


	case ParserParserRW_path:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(94)
			p.Match(ParserParserRW_path)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(96)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*MkdiskparamContext).v4 = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*MkdiskparamContext).result = []string{"path", strings.Trim((func() string { if localctx.(*MkdiskparamContext).GetV4() == nil { return "" } else { return localctx.(*MkdiskparamContext).GetV4().GetText() }}()), "\"")}



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
	r antlr.Token
	p antlr.Token
}

func NewEmptyRmdiskContext() *RmdiskContext {
	var p = new(RmdiskContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rmdisk
	return p
}

func InitEmptyRmdiskContext(p *RmdiskContext)  {
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




func (p *ParserParser) Rmdisk() (localctx IRmdiskContext) {
	localctx = NewRmdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ParserParserRULE_rmdisk)
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)

			var _m = p.Match(ParserParserRW_rmdisk)

			localctx.(*RmdiskContext).r = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(101)
			p.Match(ParserParserRW_path)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(103)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*RmdiskContext).p = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*RmdiskContext).result = commands.NewRmdisk((func() int { if localctx.(*RmdiskContext).GetR() == nil { return 0 } else { return localctx.(*RmdiskContext).GetR().GetLine() }}()), (func() int { if localctx.(*RmdiskContext).GetR() == nil { return 0 } else { return localctx.(*RmdiskContext).GetR().GetColumn() }}()), map[string]string{"path": strings.Trim((func() string { if localctx.(*RmdiskContext).GetP() == nil { return "" } else { return localctx.(*RmdiskContext).GetP().GetText() }}()), "\"")})


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)

			var _m = p.Match(ParserParserRW_rmdisk)

			localctx.(*RmdiskContext).r = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*RmdiskContext).result = commands.NewRmdisk((func() int { if localctx.(*RmdiskContext).GetR() == nil { return 0 } else { return localctx.(*RmdiskContext).GetR().GetLine() }}()), (func() int { if localctx.(*RmdiskContext).GetR() == nil { return 0 } else { return localctx.(*RmdiskContext).GetR().GetColumn() }}()), map[string]string{})

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
	f antlr.Token
	p IFdiskparamsContext 
}

func NewEmptyFdiskContext() *FdiskContext {
	var p = new(FdiskContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdisk
	return p
}

func InitEmptyFdiskContext(p *FdiskContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskparamsContext); ok {
			t = ctx.(antlr.RuleContext);
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




func (p *ParserParser) Fdisk() (localctx IFdiskContext) {
	localctx = NewFdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ParserParserRULE_fdisk)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)

			var _m = p.Match(ParserParserRW_fdisk)

			localctx.(*FdiskContext).f = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(110)

			var _x = p.fdiskparams(0)

			localctx.(*FdiskContext).p = _x
		}
		localctx.(*FdiskContext).result = commands.NewFdisk((func() int { if localctx.(*FdiskContext).GetF() == nil { return 0 } else { return localctx.(*FdiskContext).GetF().GetLine() }}()), (func() int { if localctx.(*FdiskContext).GetF() == nil { return 0 } else { return localctx.(*FdiskContext).GetF().GetColumn() }}()), localctx.(*FdiskContext).GetP().GetResult())


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(113)

			var _m = p.Match(ParserParserRW_fdisk)

			localctx.(*FdiskContext).f = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*FdiskContext).result = commands.NewFdisk((func() int { if localctx.(*FdiskContext).GetF() == nil { return 0 } else { return localctx.(*FdiskContext).GetF().GetLine() }}()), (func() int { if localctx.(*FdiskContext).GetF() == nil { return 0 } else { return localctx.(*FdiskContext).GetF().GetColumn() }}()), map[string]string{})

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
	l IFdiskparamsContext 
	p IFdiskparamContext 
}

func NewEmptyFdiskparamsContext() *FdiskparamsContext {
	var p = new(FdiskparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdiskparams
	return p
}

func InitEmptyFdiskparamsContext(p *FdiskparamsContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskparamContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFdiskparamContext)
}

func (s *FdiskparamsContext) Fdiskparams() IFdiskparamsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskparamsContext); ok {
			t = ctx.(antlr.RuleContext);
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
		p.SetState(118)

		var _x = p.Fdiskparam()


		localctx.(*FdiskparamsContext).p = _x
	}
	localctx.(*FdiskparamsContext).result = map[string]string{localctx.(*FdiskparamsContext).GetP().GetResult()[0]: localctx.(*FdiskparamsContext).GetP().GetResult()[1] }

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(127)
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
			p.SetState(121)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(122)

				var _x = p.Fdiskparam()


				localctx.(*FdiskparamsContext).p = _x
			}
			localctx.(*FdiskparamsContext).SetResult( localctx.(*FdiskparamsContext).GetL().GetResult()); localctx.(*FdiskparamsContext).result[localctx.(*FdiskparamsContext).GetP().GetResult()[0]] = localctx.(*FdiskparamsContext).GetP().GetResult()[1]


		}
		p.SetState(129)
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
	v1 antlr.Token
	v2 antlr.Token
	v3 antlr.Token
	v4 antlr.Token
	v5 antlr.Token
	v6 antlr.Token
}

func NewEmptyFdiskparamContext() *FdiskparamContext {
	var p = new(FdiskparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdiskparam
	return p
}

func InitEmptyFdiskparamContext(p *FdiskparamContext)  {
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




func (p *ParserParser) Fdiskparam() (localctx IFdiskparamContext) {
	localctx = NewFdiskparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ParserParserRULE_fdiskparam)
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_size:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.Match(ParserParserRW_size)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(132)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*FdiskparamContext).v1 = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"size", (func() string { if localctx.(*FdiskparamContext).GetV1() == nil { return "" } else { return localctx.(*FdiskparamContext).GetV1().GetText() }}())}


	case ParserParserRW_unit:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(134)
			p.Match(ParserParserRW_unit)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(136)

			var _m = p.Match(ParserParserTK_unit)

			localctx.(*FdiskparamContext).v2 = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"unit", (func() string { if localctx.(*FdiskparamContext).GetV2() == nil { return "" } else { return localctx.(*FdiskparamContext).GetV2().GetText() }}())}


	case ParserParserRW_path:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(138)
			p.Match(ParserParserRW_path)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(140)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*FdiskparamContext).v3 = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"path", strings.Trim((func() string { if localctx.(*FdiskparamContext).GetV3() == nil { return "" } else { return localctx.(*FdiskparamContext).GetV3().GetText() }}()), "\"")}


	case ParserParserRW_type:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(142)
			p.Match(ParserParserRW_type)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(144)

			var _m = p.Match(ParserParserTK_type)

			localctx.(*FdiskparamContext).v4 = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"type", (func() string { if localctx.(*FdiskparamContext).GetV4() == nil { return "" } else { return localctx.(*FdiskparamContext).GetV4().GetText() }}())}


	case ParserParserRW_fit:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(146)
			p.Match(ParserParserRW_fit)
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

			var _m = p.Match(ParserParserTK_fit)

			localctx.(*FdiskparamContext).v5 = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"fit", (func() string { if localctx.(*FdiskparamContext).GetV5() == nil { return "" } else { return localctx.(*FdiskparamContext).GetV5().GetText() }}())}


	case ParserParserRW_name:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(150)
			p.Match(ParserParserRW_name)
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

			var _m = p.Match(ParserParserTK_id)

			localctx.(*FdiskparamContext).v6 = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"name", strings.Trim((func() string { if localctx.(*FdiskparamContext).GetV6() == nil { return "" } else { return localctx.(*FdiskparamContext).GetV6().GetText() }}()), "\"")}



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
	m antlr.Token
	p IMountparamsContext 
}

func NewEmptyMountContext() *MountContext {
	var p = new(MountContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mount
	return p
}

func InitEmptyMountContext(p *MountContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountparamsContext); ok {
			t = ctx.(antlr.RuleContext);
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




func (p *ParserParser) Mount() (localctx IMountContext) {
	localctx = NewMountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ParserParserRULE_mount)
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)

			var _m = p.Match(ParserParserRW_mount)

			localctx.(*MountContext).m = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(157)

			var _x = p.mountparams(0)

			localctx.(*MountContext).p = _x
		}
		localctx.(*MountContext).result = commands.NewMount((func() int { if localctx.(*MountContext).GetM() == nil { return 0 } else { return localctx.(*MountContext).GetM().GetLine() }}()), (func() int { if localctx.(*MountContext).GetM() == nil { return 0 } else { return localctx.(*MountContext).GetM().GetColumn() }}()), localctx.(*MountContext).GetP().GetResult())


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(160)

			var _m = p.Match(ParserParserRW_mount)

			localctx.(*MountContext).m = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*MountContext).result = commands.NewMount((func() int { if localctx.(*MountContext).GetM() == nil { return 0 } else { return localctx.(*MountContext).GetM().GetLine() }}()), (func() int { if localctx.(*MountContext).GetM() == nil { return 0 } else { return localctx.(*MountContext).GetM().GetColumn() }}()), map[string]string{})

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
	l IMountparamsContext 
	p IMountparamContext 
}

func NewEmptyMountparamsContext() *MountparamsContext {
	var p = new(MountparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mountparams
	return p
}

func InitEmptyMountparamsContext(p *MountparamsContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountparamContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountparamContext)
}

func (s *MountparamsContext) Mountparams() IMountparamsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountparamsContext); ok {
			t = ctx.(antlr.RuleContext);
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
		p.SetState(165)

		var _x = p.Mountparam()


		localctx.(*MountparamsContext).p = _x
	}
	localctx.(*MountparamsContext).result = map[string]string{localctx.(*MountparamsContext).GetP().GetResult()[0]: localctx.(*MountparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(174)
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
			p.SetState(168)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(169)

				var _x = p.Mountparam()


				localctx.(*MountparamsContext).p = _x
			}
			localctx.(*MountparamsContext).SetResult( localctx.(*MountparamsContext).GetL().GetResult()); localctx.(*MountparamsContext).result[localctx.(*MountparamsContext).GetP().GetResult()[0]] = localctx.(*MountparamsContext).GetP().GetResult()[1]


		}
		p.SetState(176)
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
	v1 antlr.Token
	v2 antlr.Token
}

func NewEmptyMountparamContext() *MountparamContext {
	var p = new(MountparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mountparam
	return p
}

func InitEmptyMountparamContext(p *MountparamContext)  {
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




func (p *ParserParser) Mountparam() (localctx IMountparamContext) {
	localctx = NewMountparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ParserParserRULE_mountparam)
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_path:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.Match(ParserParserRW_path)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(179)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*MountparamContext).v1 = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*MountparamContext).result = []string{"path", strings.Trim((func() string { if localctx.(*MountparamContext).GetV1() == nil { return "" } else { return localctx.(*MountparamContext).GetV1().GetText() }}()), "\"")}


	case ParserParserRW_name:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.Match(ParserParserRW_name)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(183)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*MountparamContext).v2 = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		localctx.(*MountparamContext).result = []string{"name", strings.Trim((func() string { if localctx.(*MountparamContext).GetV2() == nil { return "" } else { return localctx.(*MountparamContext).GetV2().GetText() }}()), "\"")}



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


func (p *ParserParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
			var t *CommandsContext = nil
			if localctx != nil { t = localctx.(*CommandsContext) }
			return p.Commands_Sempred(t, predIndex)

	case 4:
			var t *MkdiskparamsContext = nil
			if localctx != nil { t = localctx.(*MkdiskparamsContext) }
			return p.Mkdiskparams_Sempred(t, predIndex)

	case 8:
			var t *FdiskparamsContext = nil
			if localctx != nil { t = localctx.(*FdiskparamsContext) }
			return p.Fdiskparams_Sempred(t, predIndex)

	case 11:
			var t *MountparamsContext = nil
			if localctx != nil { t = localctx.(*MountparamsContext) }
			return p.Mountparams_Sempred(t, predIndex)


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

