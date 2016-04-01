package config

import "github.com/BurntSushi/toml"

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
