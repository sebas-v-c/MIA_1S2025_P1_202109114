lexer grammar Lexer;

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

// Tokens
TK_option       : '-' EXT+                      ;
TK_value        : STRING | PATH | INTEGER | ID1 ;
TK_equ          : '='                           ;
COMMENTARY      : COMMENTS                      ;
NEWLINE         : '\n' -> skip                  ;
UNUSED_         : UNUSED -> skip                ;