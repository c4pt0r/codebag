package main

import (
    "fmt"
)
var cmdList = &Command{
    UsageLine : "list",
}

var limit, offset int64

func init() {
    cmdList.Flag.Int64Var(&limit, "n", 0, "count")
    cmdList.Flag.Int64Var(&offset, "offset", 0, "offset")
    cmdList.Run = listRun
}

func listRun(c *Command, args []string) {
    cmdList.Flag.Parse(args)
    if snippets, err := ListSnippets(offset, limit); err == nil {
        for _, s := range snippets {
            fmt.Printf("%d\t%s\t%s\t%s\n", s.id, s.desc, s.date.Format(SimpleTimeFmt), s.ftype)
        }
    }
}
