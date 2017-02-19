package targets

import (
	"os"
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

// Target is Mask Target
type Target struct {
	Surf string
	Pos  string
	Proc string
}

// Targets is Slice any Target
type Targets struct {
	Name    string   // ターゲット名
	Tag     []string // タグ名リスト
	Targets []Target // ターゲットリスト
}

/*
create table targets(
	id integer not null primary key,
	name text not null,
	tag text
);
create table target(
	id integer not null primary key,
	targets_id integer not null,
	surf text,
	proccess integer
);
*/

func Create(targets Targets) (bool, error) {
	ch_err := _check_data(targets)
	if ch_err != nil {
		return false, ch_err
	}
}

func _check_data(t Targets) error {
	db, err = sql.Open("sqlite3", "./targets.db")
	if err != nil {
		return errors.New("database file is not found.")		
	}
	defer db.Close(err)
	sql := fmt.Sprintf("SELECT id FROM targets WHERE name = '%s';", t.Name)
	rows, err := db.Query(sql)
	if err != nil {
		return errors.New(fmt.Sprintf("SQL is bad request.detai: %s", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		return errors.New(fmt.Sprintf("%s is already exists.", t.Name))
	}
	return nil
}

func Update(targets Targets) (bool, error)             {}
func Read(name string, tag []string) (Targets, error)  {}
func Delete(del string, targets Targets) (bool, error) {}
