package database

import (
	"database/sql"
)

type joinpart struct {
	db *sql.DB
}

func (jp *joinpart) init() error {
	var schema = `
		create table if not exists joinpart(
			id integer primary key autoincrement,
			event varchar(16) not null,
			username varchar(128) not null,
			channel varchar(128) not null,
			timestamp default current_timestamp
		);
		create index if not exists joinpart_username on joinpart(username);
		create index if not exists joinpart_joins on joinpart(username) where event = 'join';
		create index if not exists joinpart_parts on joinpart(username) where event = 'part';
	`
	_, err := jp.db.Exec(schema)
	return err
}

func (jp *joinpart) Join(username, channel string) error {
	sql := `insert into joinpart(event, username, channel) values('join', ?, ?);`
	stmt, err := jp.db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username, channel)
	return err
}

func (jp *joinpart) Part(username, channel string) error {
	sql := `insert into joinpart(event, username, channel) values('part', ?, ?);`
	stmt, err := jp.db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username, channel)
	return err
}
