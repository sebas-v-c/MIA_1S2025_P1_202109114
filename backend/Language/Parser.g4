// code inspired from @brandonT2002
grammar Parser;

@header {
    import (
        commands "backend/Classes/Commands"
        interfaces "backend/Classes/Interfaces"
        "strings"
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
    |   c4 = mount      {$result = $c4.result}
    |   c5 = mounted    {$result = $c5.result}
    |   c6 = mkfs       {$result = $c6.result}
    |   c7 = cat        {$result = $c7.result}
    |   c8 = login      {$result = $c8.result}
    |   c9 = logout     {$result = $c9.result}
    |   c10 = mkgrp     {$result = $c10.result}
    |   c11 = rmgrp     {$result = $c11.result}
    |   c12 = mkusr     {$result = $c12.result}
    |   c13 = rmusr     {$result = $c13.result}
    |   c14 = chgrp     {$result = $c14.result}
    |   c15 = mkdir     {$result = $c15.result}
    |   c16 = mkfile    {$result = $c16.result}
    |   c17 = rep       {$result = $c17.result}
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
    |   RW_path TK_equ v4 = TK_path     {$result = []string{"path", strings.Trim($v4.text, "\"")}}
    ;

// =============== RMDISK ===============
rmdisk returns[*commands.Rmdisk result]:
        r = RW_rmdisk RW_path TK_equ p = TK_path    {$result = commands.NewRmdisk($r.line, $r.pos, map[string]string{"path": strings.Trim($p.text, "\"")})}
    |   r = RW_rmdisk                               {$result = commands.NewRmdisk($r.line, $r.pos, map[string]string{})}
    ;

// =============== FDISK ===============
fdisk returns[*commands.Fdisk result]:
        f = RW_fdisk p = fdiskparams    {$result = commands.NewFdisk($f.line, $f.pos, $p.result)}
    |   f = RW_fdisk                    {$result = commands.NewFdisk($f.line, $f.pos, map[string]string{})}
    ;

fdiskparams returns[map[string]string result]:
        l = fdiskparams p = fdiskparam  {$result = $l.result;; $result[$p.result[0]] = $p.result[1]}
    |   p = fdiskparam                  {$result = map[string]string{$p.result[0]: $p.result[1] }}
    ;

fdiskparam returns[[]string result] :
        RW_size TK_equ v1 = TK_number   {$result = []string{"size", $v1.text}}
    |   RW_unit TK_equ v2 = TK_unit     {$result = []string{"unit", $v2.text}}
    |   RW_path TK_equ v3 = TK_path     {$result = []string{"path", strings.Trim($v3.text, "\"")}}
    |   RW_type TK_equ v4 = TK_type     {$result = []string{"type", $v4.text}}
    |   RW_fit  TK_equ v5 = TK_fit      {$result = []string{"fit", $v5.text}}
    |   RW_name TK_equ v6 = TK_id      {$result = []string{"name", strings.Trim($v6.text, "\"")}}
    ;

// =============== MOUNT ===============
// From now on I'm going make all parameteres uppercase for easier analyzis in the code
mount returns[*commands.Mount result]:
        m = RW_mount p = mountparams    {$result = commands.NewMount($m.line, $m.pos, $p.result)}
    |   m = RW_mount                    {$result = commands.NewMount($m.line, $m.pos, map[string]string{})}
    ;

mountparams returns[map[string]string result]:
        l = mountparams p = mountparam  {$result = $l.result;; $result[$p.result[0]] = $p.result[1]}
    |   p = mountparam                  {$result = map[string]string{$p.result[0]: $p.result[1]}}
    ;

mountparam returns[[]string result]:
        RW_path TK_equ v1 = TK_path     {$result = []string{"path", strings.Trim($v1.text, "\"")}}
    |   RW_name TK_equ v2 = TK_id      {$result = []string{"name", strings.Trim($v2.text, "\"")}}
    ;

// =============== MOUNTED ===============
mounted returns[*commands.Mounted result]:
        m = RW_mounted  {$result = commands.NewMounted($m.line, $m.pos)}
        ;
// =============== MKFS ===============
mkfs returns[*commands.Mkfs result]:
        m = RW_mkfs p = mkfsparams      {$result = commands.NewMkfs($m.line, $m.pos, $p.result)}
    |   m = RW_mkfs                     {$result = commands.NewMkfs($m.line, $m.pos, map[string]string{})}
    ;

mkfsparams returns[map[string]string result]:
        l = mkfsparams p = mkfsparam    {$result = $l.result;; $result[$p.result[0]] = $p.result[1]}
    |   p = mkfsparam                   {$result = map[string]string{$p.result[0]: $p.result[1]}}
    ;

mkfsparam returns[[]string result]:
        RW_type TK_equ v1 = TK_ftype    {$result = []string{"type", $v1.text}}
    |   RW_id TK_equ v2 = TK_id        {$result = []string{"id", strings.Trim($v2.text, "\"")}}
    ;

// =============== CAT ===============
cat returns[*commands.Cat result]:
        c = RW_cat p = catparams    {$result = commands.NewCat($c.line, $c.pos, $p.result)}
    |   c = RW_cat                  {$result = commands.NewCat($c.line, $c.pos, []string{})}
    ;

catparams returns[[]string result]:
        l = catparams p = catparam  {$result = $l.result;; $result = append($result, $p.result)}
    |   p = catparam                {$result = []string{$p.result}}
    ;

catparam returns[string result]:
        RW_fileN TK_equ p1 = TK_path    {$result = strings.Trim($p1.text, "\"")}
        ;


// =============== LOGIN ===============
login returns[*commands.Login result]:
        l = RW_login p = loginparams    {$result = commands.NewLogin($l.line, $l.pos, $p.result)}
    |   l = RW_login                    {$result = commands.NewLogin($l.line, $l.pos, map[string]string{})}
    ;

loginparams returns[map[string]string result]:
        l = loginparams p = loginparam      {$result = $l.result;; $result[$p.result[0]] = $p.result[1]}
    |   p = loginparam                      {$result = map[string]string{$p.result[0]: $p.result[1]}}
    ;

loginparam returns[[]string result]:
        RW_user TK_equ p1 = TK_id       {$result = []string{"user", strings.Trim($p1.text, "\"")}}
    |   RW_pass TK_equ p2 = TK_id       {$result = []string{"pass", strings.Trim($p2.text, "\"")}}
    |   RW_pass TK_equ p2 = TK_number   {$result = []string{"pass", $p2.text}}
    |   RW_id TK_equ p3 = TK_id         {$result = []string{"id", strings.ToUpper(strings.Trim($p3.text, "\""))}}
    ;

// =============== LOGOUT ===============
logout returns[*commands.Logout result]:
        l = RW_logout   {$result = commands.NewLogout($l.line, $l.pos)}
    ;

// =============== MKGRP ===============
mkgrp returns[*commands.Mkgrp result]:
        l = RW_mkgrp p = mkgrpparams    {$result = commands.NewMkgrp($l.line, $l.pos, $p.result)}
    |   l = RW_mkgrp                    {$result = commands.NewMkgrp($l.line, $l.pos, map[string]string{})}
    ;

mkgrpparams returns[map[string]string result]:
        l = mkgrpparams p = mkgrpparam      {$result = $l.result;; $result[$p.result[0]] = $p.result[1]}
    |   p = mkgrpparam                      {$result = map[string]string{$p.result[0]: $p.result[1]}}
    ;

mkgrpparam returns[[]string result]:
        RW_name TK_equ p1 = TK_id       {$result = []string{"name", strings.Trim($p1.text, "\"")}}
    |   RW_name TK_equ p2 = TK_number   {$result = []string{"name", $p2.text}}
    ;

// =============== RMGRP ===============
rmgrp returns[*commands.Rmgrp result]:
        l = RW_rmgrp p = rmgrpparams    {$result = commands.NewRmgrp($l.line, $l.pos, $p.result)}
    |   l = RW_rmgrp                    {$result = commands.NewRmgrp($l.line, $l.pos, map[string]string{})}
    ;

rmgrpparams returns[map[string]string result]:
        l = rmgrpparams p = rmgrpparam      {$result = $l.result;; $result[$p.result[0]] = $p.result[1]}
    |   p = rmgrpparam                      {$result = map[string]string{$p.result[0]: $p.result[1]}}
    ;

rmgrpparam returns[[]string result]:
        RW_name TK_equ p1 = TK_id       {$result = []string{"name", strings.Trim($p1.text, "\"")}}
    |   RW_name TK_equ p2 = TK_number   {$result = []string{"name", $p2.text}}
    ;

// =============== MKUSR ===============
mkusr returns[*commands.Mkusr result]:
        l = RW_mkusr p = mkusrparams    {$result = commands.NewMkusr($l.line, $l.pos, $p.result)}
    |   l = RW_mkusr                    {$result = commands.NewMkusr($l.line, $l.pos, map[string]string{})}
    ;

mkusrparams returns[map[string]string result]:
        l = mkusrparams p = mkusrparam      {$result = $l.result;; $result[$p.result[0]] = $p.result[1]}
    |   p = mkusrparam                      {$result = map[string]string{$p.result[0]: $p.result[1]}}
    ;

mkusrparam returns[[]string result]:
        RW_user TK_equ p1 = TK_id       {$result = []string{"user", strings.Trim($p1.text, "\"")}}
    |   RW_user TK_equ p2 = TK_number   {$result = []string{"user", $p2.text}}
    |   RW_pass TK_equ p3 = TK_id       {$result = []string{"pass", strings.Trim($p3.text, "\"")}}
    |   RW_pass TK_equ p4 = TK_number   {$result = []string{"pass", $p4.text}}
    |   RW_grp TK_equ p5 = TK_id        {$result = []string{"grp", strings.Trim($p5.text, "\"")}}
    |   RW_grp TK_equ p6 = TK_number    {$result = []string{"grp", $p6.text}}
    ;

// =============== RMUSR ===============
rmusr returns[*commands.Rmusr result]:
        l = RW_rmusr p = rmusrparams    {$result = commands.NewRmusr($l.line, $l.pos, $p.result)}
    |   l = RW_rmusr                    {$result = commands.NewRmusr($l.line, $l.pos, map[string]string{})}
    ;

rmusrparams returns[map[string]string result]:
        l = rmusrparams p = rmusrparam      {$result = $l.result;; $result[$p.result[0]] = $p.result[1]}
    |   p = rmusrparam                      {$result = map[string]string{$p.result[0]: $p.result[1]}}
    ;

rmusrparam returns[[]string result]:
        RW_user TK_equ p1 = TK_id       {$result = []string{"user", strings.Trim($p1.text, "\"")}}
    |   RW_user TK_equ p2 = TK_number   {$result = []string{"user", $p2.text}}
    ;

// =============== CHGRP ===============
chgrp returns[*commands.Chgrp result]:
        l = RW_chgrp p = chgrpparams    {$result = commands.NewChgrp($l.line, $l.pos, $p.result)}
    |   l = RW_chgrp                    {$result = commands.NewChgrp($l.line, $l.pos, map[string]string{})}
    ;

chgrpparams returns[map[string]string result]:
        l = chgrpparams p = chgrpparam      {$result = $l.result;; $result[$p.result[0]] = $p.result[1]}
    |   p = chgrpparam                      {$result = map[string]string{$p.result[0]: $p.result[1]}}
    ;

chgrpparam returns[[]string result]:
        RW_user TK_equ p1 = TK_id       {$result = []string{"user", strings.Trim($p1.text, "\"")}}
    |   RW_user TK_equ p2 = TK_number   {$result = []string{"user", $p2.text}}
    |   RW_grp TK_equ p3 = TK_id        {$result = []string{"grp", strings.Trim($p3.text, "\"")}}
    |   RW_grp TK_equ p4 = TK_number    {$result = []string{"grp", $p4.text}}
    ;

// =============== MKDIR ===============
mkdir returns[*commands.Mkdir result]: l = RW_mkdir p = mkdirparams    {$result = commands.NewMkdir($l.line, $l.pos, $p.result)} |   l = RW_mkdir                    {$result = commands.NewMkdir($l.line, $l.pos, map[string]string{})} ;

mkdirparams returns[map[string]string result]:
        l = mkdirparams p = mkdirparam      {$result = $l.result;; $result[$p.result[0]] = $p.result[1]}
    |   p = mkdirparam                      {$result = map[string]string{$p.result[0]: $p.result[1]}}
    ;

mkdirparam returns[[]string result]:
        RW_path TK_equ p1 = TK_path     {$result = []string{"path", strings.Trim($p1.text, "\"")}}
    |   RW_p                            {$result = []string{"p", ""}}
    ;

// =============== MKFILE ===============
mkfile returns[*commands.Mkfile result]:
        l = RW_mkfile p = mkfileparams  {$result = commands.NewMkfile($l.line, $l.pos, $p.result)}
    |   l = RW_mkfile                   {$result = commands.NewMkfile($l.line, $l.pos, map[string]string{})}
    ;

mkfileparams returns[map[string]string result]:
        l = mkfileparams p = mkfileparam    {$result = $l.result;; $result[$p.result[0]] = $p.result[1]}
    |   p = mkfileparam                     {$result = map[string]string{$p.result[0]: $p.result[1]}}
    ;

mkfileparam returns[[]string result]:
        RW_path TK_equ p1 = TK_path     {$result = []string{"path", strings.Trim($p1.text, "\"")}}
    |   RW_r                            {$result = []string{"r", ""}}
    |   RW_size TK_equ p2 = TK_number   {$result = []string{"size", $p2.text}}
    |   RW_cont TK_equ p3 = TK_path     {$result = []string{"cont", strings.Trim($p3.text, "\"")}}
    ;

// =============== REP ===============
rep returns[*commands.Rep result]:
        l = RW_rep p = repparams  {$result = commands.NewRep($l.line, $l.pos, $p.result)}
    |   l = RW_rep                   {$result = commands.NewRep($l.line, $l.pos, map[string]string{})}
    ;

repparams returns[map[string]string result]:
        l = repparams p = repparam    {$result = $l.result;; $result[$p.result[0]] = $p.result[1]}
    |   p = repparam                     {$result = map[string]string{$p.result[0]: $p.result[1]}}
    ;

repparam returns[[]string result]:
        RW_name TK_equ p1 = TK_id       {$result = []string{"name", strings.Trim($p1.text, "\"")}}
    |   RW_path TK_equ p2 = TK_path     {$result = []string{"path", strings.Trim($p2.text, "\"")}}
    |   RW_id TK_equ p3 = TK_id         {$result = []string{"id", strings.Trim($p3.text, "\"")}}
    |   RW_pfl TK_equ p4 = TK_path     {$result = []string{"pfl", strings.Trim($p4.text, "\"")}}
    ;
