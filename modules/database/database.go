package database

import (
	"database/sql"

	"danirod.es/pkg/ttvbot"

	_ "github.com/mattn/go-sqlite3"
)

type databaseModule interface {
	init() error
}

type Database struct {
	db         *sql.DB
	JoinPartDB *joinpart
	MessageDB  *messages
}

func openSqlConnection(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./roster.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (db *Database) init() error {
	modules := []databaseModule{db.JoinPartDB, db.MessageDB}
	for _, mod := range modules {
		err := mod.init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *Database) close() error {
	return db.db.Close()
}

func newDatabase(config *ttvbot.Config) *Database {
	db, err := openSqlConnection(config.DatabaseUrl)
	if err != nil {
		panic(err)
	}
	return &Database{
		db:         db,
		JoinPartDB: &joinpart{db: db},
		MessageDB:  &messages{db: db},
	}
}
