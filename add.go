package main
import (
    "os"
    "time"
    "fmt"
    "os/exec"
    "io/ioutil"
    "encoding/json"
    "path/filepath"
    "errors"
)

var cmdAdd = &Command{
    UsageLine : "add [-m] [-t] [-r] files...",
}

var message string
var tags string
var recurisive bool

func init() {
    cmdAdd.Run = addRun
    cmdAdd.Flag.StringVar(&message, "m", "", "desc message")
    cmdAdd.Flag.StringVar(&tags, "t", "", "tags")
    cmdAdd.Flag.BoolVar(&recurisive, "r", false, "recurisive")
}

func addRun(cmd *Command, args []string) {
    cmdAdd.Flag.Parse(args)
    if len(message) > 0 {
        addSnippet(cmd.Flag.Args(), message, tags)
    } else {
        fmt.Println("message is missing")
        os.Exit(-1)
    }
}

func readFromEditor() (string, error) {
    tmpFileName := os.TempDir() + RandString(16)
    defer os.Remove(tmpFileName)
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

func getFileInfo(filename string) (map[string]interface{}, error) {
    v := make(map[string]interface{})
    b , err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    v["content"] = string(b)
    if err != nil {
        return nil, err
    }
    v["filename"] = filename
    return v, nil
}

func wrapperFilesToJson(baseDir string, r bool, args ... string) ([]byte, error) {
    // get filename from arguments
    var files []map[string]interface{}
    for _, fname := range args {
        fpath, err := filepath.Abs(fname)
        if err == nil {
            if stat, err := os.Stat(fpath); err == nil {
                // check if file extists
                rel, err := filepath.Rel(baseDir, fpath)
                if err != nil {
                    return nil, err
                }
                if !stat.IsDir() {
                    info, err := getFileInfo(rel)
                    if err != nil {
                        return nil, err
                    }
                    files = append(files, info)
                } else if r == true {
                    // recurisive mode
                    err := filepath.Walk(rel, func(path string, info os.FileInfo, err error) error {
                        i, err := getFileInfo(path)
                        if err != nil {
                            return err
                        }
                        files = append(files, i)
                        return nil
                    })
                    if err != nil {
                        return nil ,err
                    }
                }
            }
        } else {
            return nil, err
        }
    }
    if len(files) == 0 {
        return nil, errors.New("invalid file(s)")
    }
    b, _ := json.Marshal(files)
    return b, nil
}

func addSnippet(files []string, desc string, tags string) *Snippet {
    if files == nil || len(files) == 0 {
        content, err := readFromEditor()
        v := make(map[string]interface{})
        v["content"] = content
        v["filename"] = nil
        var lst []map[string]interface{}
        lst = append(lst, v)
        b, _ := json.Marshal(lst)
        if err == nil {
            s := NewSnippet(tags, desc, time.Now(), b)
            if id, err := AddSnippet(s); err == nil {
                fmt.Println(id)
            } else {
                Fatal(err)
            }
            return s
        }
    } else {
        baseDir, err := os.Getwd()
        if err != nil {
            Fatal(err)
        }
        b, err := wrapperFilesToJson(baseDir, recurisive, files...)
        if err != nil {
            Fatal(err)
        }
        s := NewSnippet(tags, desc, time.Now(), b)
        if id, err := AddSnippet(s); err == nil {
            fmt.Println(id)
        } else {
            Fatal(err)
        }
        return s
    }
    return nil
}

