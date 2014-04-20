package main

import (
    "fmt"
    "strconv"
)
var cmdGet = &Command{
    UsageLine : "get",
}

var id int64

func init() {
    cmdGet.Run = getRun
}

func getRun(c *Command, args []string) {
    cmdGet.Flag.Parse(args)
    for _, sid := range cmdGet.Flag.Args() {
        id, _ := strconv.ParseInt(sid, 10, 64)
        s, err := FetchSnippet(id)
        if err == nil {
            fmt.Println("########## " + s.desc + " " + s.date.Format(SimpleTimeFmt) + " ##########")
            fmt.Printf(string(s.content))
            fmt.Println()
        }
    }
}
