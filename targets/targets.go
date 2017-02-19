package Targets

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

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
	pos text,
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
	defer db.Close(db)
	sql := fmt.Sprintf("SELECT id FROM targets WHERE name = '%s';", t.Name)
	rows, err := db.Query(sql)
	if err != nil {
		return errors.New(fmt.Sprintf("SQL is bad request.detai: %s", err.Error()))
	}
	defer rows.Close()
	for rows.Next() {
		return errors.New(fmt.Sprintf("%s is already exists.", t.Name))
	}
	return nil
}

func _create_targets(t Targets) error {
	db, err = sql.Open("sqlite3", "./targets.db")
	if err != nil {
		return errors.New("database file is not found.")
	}
	defer db.Close(db)
	tx, err := db.Begin()
	if err != nil {
		return errors.New("cannot open database.")
	}
	stmt, err := tx.Prepare("INSERT INTO targets(name, tag) values(?, ?)")
	if err != nil {
		return errors.New(fmt.Sprintf("catch error: %s", err.Error()))
	}
	defer stmt.Close()
	res, err = stmt.Exec(t.Name, strings.Join(t.Tag, ","))
	id, _ := res.LastInsertId()
	if err != nil {
		return errors.New(fmt.Sprintf("cannot run sql : %s", err.Error()))
	}
	tx.Commit()
	if len(t.Tergets) > 0 {
		tar_res := _insert_target(id, t.Targets)
	}
	return nil
}

func _insert_target(index int64, ts []Target) bool {
	db, err = sql.Open("sqlite3", "./targets.db")
	if err != nil {
		return false
	}
	defer db.Close(db)
	tx, err := db.Begin()
	if err != nil {
		return false
	}
	stmt, err := tx.Prepare("INSERT INTO targets(name, tag) values(?, ?)")
	if err != nil {
		return false
	}
	defer stmt.Close()
	for _, v := range ts {
		_, err = stmt.Exec(index, v.Surf, v.Pos, v.Proc)
		if err != nil {
			return false
		}
	}
	tx.Commit()
	return true
}

// Update is update targets and target/
func Update(targets Targets) (bool, error) {
	if err = _updateTargets(targets); err != nil {
		return false, errors.New("cannot update targets")
	}
	if err = _updateTarget(targets.Targets); err != nil {
		return false, errors.New("cannot update target")
	}
	return true, nil
}

func _updateTargets(ts Targets) error {
	db, err = sql.Open("sqlite3", "./targets.db")
	if err != nil {
		return errors.New("database file is not found.")
	}
	defer db.Close(db)
	tx, err := db.Begin()
	if err != nil {
		return errors.New("cannot open database.")
	}
	stmt, err := tx.Prepare("INSERT INTO targets(name, tag) values(?, ?)")
	if err != nil {
		return errors.New(fmt.Sprintf("catch error: %s", err.Error()))
	}
	defer stmt.Close()
	res, err = stmt.Exec(t.Name, strings.Join(t.Tag, ","))
	id, _ := res.LastInsertId()
	if err != nil {
		return errors.New(fmt.Sprintf("cannot run sql : %s", err.Error()))
	}
	tx.Commit()
	return nil
}

func _updateTarget(t []Target) error {
	db, err = sql.Open("sqlite3", "./targets.db")
	if err != nil {
		return errors.New("database file is not found.")
	}
	defer db.Close(db)
	tx, err := db.Begin()
	if err != nil {
		return errors.New("cannot open database.")
	}
	stmt, err := tx.Prepare("INSERT INTO targets(name, tag) values(?, ?)")
	if err != nil {
		return errors.New(fmt.Sprintf("catch error: %s", err.Error()))
	}
	defer stmt.Close()
	for _, v := range ts {
		_, err = stmt.Exec(index, v.Surf, v.Pos, v.Proc)
		if err != nil {
			return errors.New(fmt.Sprintf("cannot run sql : %s", err.Error()))
		}
	}
	tx.Commit()
	return nil
}

// Read is show targets data
func Read(name string, tag []string) (Targets, error) {
	db, err = sql.Open("sqlite3", "./targets.db")
	if err != nil {
		return errors.New("database file is not found.")
	}
	defer db.Close(db)
	sql := `SELECT
		targets.id, targets.name, target.surf, target.pos, target.process
	FROM
		target JOIN target on targets.id = target.target_id
	WHERE
		targets.name = "%s" AND targets.tag = "%s";`
	sql := fmt.Sprintf(sql, t.Name, strings.Join(t.Tag, ","))
	rows, err := db.Query(sql)
	if err != nil {
		return errors.New(fmt.Sprintf("SQL is bad request.detai: %s", err.Error()))
	}
	defer rows.Close()

	var ts = Targets{}
	var id int64
	var name string
	var surf string
	var pos string
	var proc int64
	for rows.Next() {
		err = rows.Scan(&id, &name, &surf, &pos, &proc)
		if err != nil {
			return ts, errors.New("cannot parse rows.")
		}
		ts.Name = name
		ts.Targets = append(ts.Targets, Target{surf, pos, proc})
	}
	return ts, nil

}

// Delete is delete targets or target
func Delete(del string, targets Targets) (bool, error) {}
