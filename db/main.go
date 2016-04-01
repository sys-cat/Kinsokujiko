package main

import (
  "../config"
)

const DBTYPE = []string{
  "SQLite",
  "MySQL",
  "MariaDB",
  "ES",
}

type List struct {
  Id    int64
  Name  string
  CreatedAt string
  UpdatedAt string
}

type Item struct{
  Id int64
  ListId int64
  Item string
  CreatedAt string
  UpdatedAt string
}

func main {
  config := config.Toml()
  for key, value := range DBTYPE {
    if config.Db.Type == value {
      // 変数の値を関数名にする方法
    }
  }
}
