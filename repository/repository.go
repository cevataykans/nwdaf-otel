package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Repository interface {
	Setup() error
}

type repository struct {
	repo Repository
}

func NewSQLiteRepo() (Repository, error) {
	db, err := sql.Open("sqlite3", File)
	if err != nil {
		return nil, fmt.Errorf("cannot open sqlite3 conn: %w", err)
	}
	return repository{
		repo: newSqlite3Repo(db),
	}, nil
}
