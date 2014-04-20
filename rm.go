package main

import (
    "strconv"
    "log"
)
var cmdRm = &Command{
    UsageLine : "rm ids...",
}

func init() {
    cmdRm.Run = rmRun
    cmdRm.Flag.Usage = cmdRm.Usage
}

func rmRun(c *Command, args []string) {
    cmdRm.Flag.Parse(args)
    for _, sid := range cmdRm.Flag.Args() {
        id, _ := strconv.ParseInt(sid, 10, 64)
        err := RemoveSnippet(id)
        if err != nil {
            log.Fatal(err)
        }
    }
}
