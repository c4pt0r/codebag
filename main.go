package main

import (
    "flag"
    "os"
    "fmt"
    "strings"
    "log"
    "io/ioutil"
)

type Command struct {
    Run func(cmd *Command, args []string)
    UsageLine string
    Flag flag.FlagSet
}

func (c *Command) Name() string {
    name := c.UsageLine
    i := strings.Index(name, " ")
    if i >= 0 {
        name = name[:i]
    }
    return name
}


func (c *Command) Usage() {
    fmt.Fprintf(os.Stderr, "usage %s\n\n", c.UsageLine)
    os.Exit(2)
}

var commands = []*Command{
    cmdAdd,
    cmdGet,
    cmdList,
    cmdRm,
}

func init() {
    log.SetOutput(ioutil.Discard)
}

func main() {
    flag.Parse()
    args := flag.Args()
    if len(args) < 1 {
        os.Exit(2)
    }

    InitDb()
    for _, cmd := range commands {
        if cmd.Name() == args[0] && cmd.Run != nil {
            cmd.Run(cmd, args[1:])
        }
    }
}
