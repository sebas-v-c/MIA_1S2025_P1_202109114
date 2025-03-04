grammar Parser;

// TODO: import all the Classes
@header {
    import (
        commands "backend/Classes/Commands"
        interfaces "backend/Classes/Interfaces"
        "os"
    )
}


options {
    language = Go;
    tokenVocab = Lexer;
}

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
    ;

mkdisk returns[*commands.Mkdisk result]:
        m = RW_mkdisk p = mkdiskparams  {$result = commands.NewMkdisk($m.line, $m.pos, $p.result)}
    |   m = RW_mkdisk                   {$result = map[string]string{}}
    ;
