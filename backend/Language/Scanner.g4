lexer grammar Scanner;

options {
    language = Go;
    caseInsensitive = true;
}

fragment UNUSED   : [ \r\t]+                ;
fragment CONTENT  : (~('\n'|'"'|'\\')|'\\'.);
fragment STRING   : '"'(CONTENT)*'"'        ;
fragment INTEGER  : ('-')?[0-9]+            ;
fragment EXT      : [a-z0-9]+               ;
fragment ID1      : [a-z0-9_]+              ;
fragment ID2      : [a-z0-9_][a-z0-9_ ]*    ;
fragment COMMENTS : '#'(~[\r\n])*           ;
fragment PATH     : ((('.''.'|'.')?'/')*(ID1))+('.'(EXT))?|'"'((('.''.'|'.')?'/')*(ID2))+('.'(EXT))?'"'|'/';

// Reserved
RW_mkdisk       : 'mkdisk'                  ;
RW_rmdisk       : 'rmdisk'                  ;
RW_fdisk        : 'fdisk'                   ;
RW_mount        : 'mount'                   ;
RW_mounted      : 'mounted'                 ;
RW_mkfs         : 'mkfs'                    ;
RW_cat          : 'cat'                     ;
RW_login        : 'login'                   ;
RW_logout       : 'logout'                  ;
RW_mkgrp        : 'mkgrp'                   ;
RW_rmgrp        : 'rmgrp'                   ;
RW_mkusr        : 'mkusr'                   ;
RW_rmusr        : 'rmusr'                   ;
RW_chgrp        : 'chgrp'                   ;
RW_mkfile       : 'mkfile'                  ;
RW_mkdir        : 'mkdir'                   ;
RW_rep          : 'rep'                     ;

RW_size        : '-'([ ])*'size'          ;
RW_fit         : '-'([ ])*'fit'           ;
RW_unit        : '-'([ ])*'unit'          ;
RW_driveletter : '-'([ ])*'driveletter'   ;
RW_name        : '-'([ ])*'name'          ;
RW_type        : '-'([ ])*'type'          ;
RW_delete      : '-'([ ])*'delete'        ;
RW_add         : '-'([ ])*'add'           ;
RW_id          : '-'([ ])*'id'            ;
RW_fs          : '-'([ ])*'fs'            ;
RW_user        : '-'([ ])*'user'          ;
RW_pass        : '-'([ ])*'pass'          ;
RW_grp         : '-'([ ])*'grp'           ;
RW_path        : '-'([ ])*'path'          ;
RW_r           : '-'([ ])*'r'             ;
RW_cont        : '-'([ ])*'cont'          ;
RW_fileN       : '-'([ ])*'file'(INTEGER) ;
RW_destino     : '-'([ ])*'destino'       ;
RW_ugo         : '-'([ ])*'ugo'           ;
RW_ruta        : '-'([ ])*'ruta'          ;

// Tokens
TK_fit         : 'BF' | 'FF' | 'WF'       ;
TK_unit        : 'B' | 'K' | 'M'          ;
TK_type        : 'P' | 'E' | 'L'          ;
TK_fs          : '2fs' | '3fs'            ;
TK_number      : INTEGER                  ;
TK_id          : ID1 | '"'ID1'"'          ;
TK_ext         : EXT                      ;
TK_path        : PATH                     ;
TK_equ         : '='                      ;
TK_sym         : '?' | '*'                ;

COMMENTARY      : COMMENTS -> skip        ;
NEWLINE         : '\n' -> skip            ;
UNUSED_         : UNUSED -> skip          ;