grammar Parser;

@header {
    import (
        commands "backend/Classes/Commands"
        interfaces "backend/Classes/Interfaces"
    )
}


options {
    language = Go;
    tokenVocab = Scanner;
}

// =============== START ===============
init returns[[]interfaces.Command result] :
        c = commands EOF    {$result = $c.result}
    |   EOF                 {$result = []interfaces.Command{}}
    ;

commands returns[[]interfaces.Command result] :
        l = commands c = command    {$result = $l.result;; $result = append($result, $c.result)}
    |   c = command                 {$result = []interfaces.Command{$c.result}}
    ;

command returns[interfaces.Command result]:
        c1 = mkdisk     {$result = $c1.result}
    |   c2 = rmdisk     {$result = $c2.result}
    |   c3 = fdisk      {$result = $c3.result}
    ;

// =============== MKDISK ===============
mkdisk returns[*commands.Mkdisk result]:
        m = RW_mkdisk p = mkdiskparams  {$result = commands.NewMkdisk($m.line, $m.pos, $p.result)} // mkdisk {parametros}
    |   m = RW_mkdisk                   {$result = commands.NewMkdisk($m.line, $m.pos, map[string]string{})} // mkdisk
    ;

mkdiskparams returns[map[string]string result]:
        l = mkdiskparams p = mkdiskparam    {$result = $l.result;; $result[$p.result[0]] = $p.result[1]}
    |   p = mkdiskparam                     {$result = map[string]string{$p.result[0]: $p.result[1] }}
    ;

mkdiskparam returns[[]string result]:
        RW_size TK_equ v1 = TK_number   {$result = []string{"size", $v1.text}}
    |   RW_fit TK_equ v2 = TK_fit       {$result = []string{"fit", $v2.text}}
    |   RW_unit TK_equ v3 = TK_unit     {$result = []string{"unit", $v3.text}}
    |   RW_path TK_equ v4 = TK_path     {$result = []string{"path", $v4.text}}
    ;

// =============== RMDISK ===============
rmdisk returns[*commands.Rmdisk result]:
        r = RW_rmdisk RW_path TK_equ p = TK_path    {$result = commands.NewRmdisk($r.line, $r.pos, map[string]string{"path": $p.text})}
    |   r = RW_rmdisk                               {$result = commands.NewRmdisk($r.line, $r.pos, map[string]string{})}
    ;

// =============== FDISK ===============
fdisk returns[*commands.Fdisk result]:
        f = RW_fdisk p = fdiskparams    {$result = commands.NewFdisk($f.line, $f.pos, $p.result)}
    |   f = RW_fdisk                    {$result = commands.NewFdisk($f.line, $f.pos, map[string]string{})}
    ;

fdiskparams returns[map[string]string result]:
        l = fdiskparams p = fdiskparam  {$result = $l.result;; $result[$p.result[0]] = $p.result[1]}
    |   p = fdiskparam                  {$result = map[string]string{$p.result[0]: $p.result[1]}}
        ;

fdiskparam returns[[]string result] :
        RW_size TK_equ v1 = TK_number   {$result = []string{"size", $v1.text}}
    |   RW_unit TK_equ v2 = TK_unit     {$result = []string{"unit", $v2.text}}
    |   RW_path TK_equ v3 = TK_path     {$result = []string{"path", $v3.text}}
    |   RW_type TK_equ v4 = TK_type     {$result = []string{"type", $v4.text}}
    |   RW_fit  TK_equ v5 = TK_fit      {$result = []string{"fit", $v5.text}}
    |   RW_name TK_equ v6 = TK_id       {$result = []string{"name", $v6.text}}
    ;