package main

type Config struct {
    Editor string `json:editor`
    DbFile string `json:db_file`
}

var GlobalCfg *Config

func init() {
    LoadConfig()
}

func LoadConfig() error {
    GlobalCfg = &Config {
        Editor : "vim",
        DbFile : "/tmp/codebag.db",
    }
    return nil
}

func WriteConfig() error {
    return nil
}
