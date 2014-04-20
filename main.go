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
    fmt.Fprintf(os.Stderr, "%s\n", c.UsageLine)
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
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "[get|add|ls|rm]\n")
        for _, cmd := range commands {
            cmd.Usage()
        }
        os.Exit(2)
    }
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
