package main

type Config struct {
    Editor string `json:editor`
}

var GlobalCfg *Config
func init() {
    LoadConfig()
}

func LoadConfig() error {
    GlobalCfg = &Config {
        Editor : "vim",
    }
    return nil
}

func WriteConfig() error {
    return nil
}
