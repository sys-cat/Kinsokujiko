package main

import (
  "../config"
  _ "github.com/mattn/go-sqlite3"
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
  switch config.Db.Type {
    case "SQLite":
      return SQLite(config)
    case "MySQL":
      return MySQL(config)
    case "MariaDB":
      return MariaDB(config)
    case "ES":
      return ES(config)
    }
}

func SQLite(c *config.Config) {
  db, err := sql.Open(c.Db.Type, c.Db.Path)
}
