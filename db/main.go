package main

import (
  "os"
  "../config"
  "database/sql"
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
  case "sqlite3":
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
  db, err := conncect_sqlite(c)
  if err != nil {
    panic(err)
  }
  _, err := init_sqlite(db)
  if err != nil {
    panic(err)
  }
  return db
}

func connect_sqlite(c *config.Config) {
  return sql.Open(c.Db.Type, c.Db.Path)
}

func init_sqlite(db *sql.DB) {
  create := `
    create table if not exists list(id int, name text, created_at datetime, updated_at datetime);
    create table if not exists item(id int, list_id int, item text, created_at datetime, updated_at datetime);
  `
  _, err = db.Exec(create)
  return "", err
}

func insert_list_sqlite(db *sql.DB, list List) {
  insert :=
}
