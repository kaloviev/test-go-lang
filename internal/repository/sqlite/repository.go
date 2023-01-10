package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	db *sql.DB
}

type Config struct {
	Path string
}

func MakeRepository(config Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", config.Path)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Sqlite{
		db: db,
	}, nil
}
