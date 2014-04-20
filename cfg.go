package main
import (
    "os"
    "encoding/json"
    "io/ioutil"
)

type Config struct {
    Editor string `json:"editor"`
    DbFile string `json:"db_file"`
}

var GlobalCfg *Config

func init() {
    LoadConfig()
}

func LoadConfig() error {
    filename := GetHomeDir() + "/.codebag.conf"
    if _, err := os.Stat(filename); err == nil {
        b, _ := ioutil.ReadFile(filename)
        json.Unmarshal(b, &GlobalCfg)
    } else {
        GlobalCfg = &Config {
            Editor : "vim",
            DbFile : "/tmp/codebag.db",
        }
        b, _ := json.MarshalIndent(GlobalCfg, "", "    ")
        ioutil.WriteFile(filename, b, 0644)
    }
    return nil
}

func WriteConfig() error {
    return nil
}
