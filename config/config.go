package config

import (
  "github.com/BurntSushi/toml"
  "os"
)

type Config struct {
  Db DBConfig
}

type DBConfig struct {
  Type string `toml:"type"`
  Port string `toml:"port"`
  Host string `toml:"host"`
  User string `toml:"user"`
  Pass string `toml:"pass"`
  Path string `toml:"path"`
}

func decode_toml(path string) Config {
  var config Config
  _, err := toml.DecodeFile(path, &config)
  if err != nil {
    panic(err)
  }
  return config
}

func Toml() Config {
  config := decode_toml("config.toml")
  return config
}

func Set(c *Config) {
  config := decode_toml("config.toml")
  os.Setenv("KINSOKU_TYPE", c.Db.Type)
  os.Setenv("KINSOKU_PORT", c.Db.Type)
  os.Setenv("KINSOKU_HOST", c.Db.Type)
  os.Setenv("KINSOKU_USER", c.Db.Type)
  os.Setenv("KINSOKU_PASS", c.Db.Type)
  os.Setenv("KINSOKU_PATH", c.Db.Type)
}

func Get() {}
