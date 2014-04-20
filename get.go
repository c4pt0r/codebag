package main

import (
    "fmt"
    "strconv"
    "encoding/json"
)
var cmdGet = &Command{
    UsageLine : "get ids...",
}

var id int64

func init() {
    cmdGet.Run = getRun
    cmdGet.Flag.Usage = cmdGet.Usage
}

func getRun(c *Command, args []string) {
    cmdGet.Flag.Parse(args)
    for _, sid := range cmdGet.Flag.Args() {
        id, _ := strconv.ParseInt(sid, 10, 64)
        s, err := FetchSnippet(id)
        if err == nil {
            fmt.Println("#" + strconv.FormatInt(s.id, 10), s.desc, s.date.Format(SimpleTimeFmt), "\n")
            var v []map[string]interface{}
            json.Unmarshal(s.content, &v)
            for _, file := range v {
                fmt.Printf(file["content"].(string))
                fmt.Println()
            }
            fmt.Println()
        }
    }
}
