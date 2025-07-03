package repository

import "database/sql"

const File string = "sqlite3.db"

type sqlite3repo struct {
	db *sql.DB
}

func newSqlite3Repo(db *sql.DB) *sqlite3repo {
	return &sqlite3repo{db}
}

// Creates necessary tables for required for NWDAF
func (r *sqlite3repo) Setup() error {

}
