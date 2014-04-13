package main
import (
    "log"
    "os"
    "os/exec"
    "io/ioutil"
)

var cmdAdd = &Command{
    UsageLine : "add files... [-message] [-edit] [-type]",
}

var message string
var editId string
var ftype string

func init() {
    cmdAdd.Run = addRun
    cmdAdd.Flag.StringVar(&message, "message", "", "message")
    cmdAdd.Flag.StringVar(&editId, "id", "", "snippet id")
    cmdAdd.Flag.StringVar(&ftype, "type", "", "type (language)")
}

func addRun(cmd *Command, args []string) {
    cmd.Flag.Parse(args)
    log.Println("message:", message)
    log.Println("snippet_id:", editId)
    log.Println("args", cmd.Flag.Args())
    addSnippet(nil, "", "")
}

func readFromEditor() (string, error) {
    tmpFileName := os.TempDir() + RandString(16)
    defer os.Remove(tmpFileName)
    log.Println(tmpFileName)
    cmd := exec.Command(GlobalCfg.Editor, tmpFileName)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Start(); err != nil {
        return "", err
    }
    if err := cmd.Wait(); err != nil {
        return "", err
    }

    b, err := ioutil.ReadFile(tmpFileName)
    return string(b), err
}

func addSnippet(files []string, desc string, ftype string) {
    if files == nil || len(files) == 0 {
        content, err := readFromEditor()
        if err == nil {
            log.Println(content)
        }
    }
}

