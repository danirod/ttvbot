package database

import (
	"database/sql"
)

type messages struct {
	db *sql.DB
}

func (m *messages) init() error {
	schema := `
	create table if not exists messages(
		id integer primary key autoincrement,
		timestamp default current_timestamp,
		username varchar(128) not null,
		channel varchar(128) not null,
		message text not null,
		first boolean default false
	);
	create index if not exists messages_username on messages(username);
	create index if not exists messages_channel on messages(channel);
	`
	_, err := m.db.Exec(schema)
	return err
}

func (m *messages) Message(username, channel, content string, first bool) error {
	sql := `insert into messages(username, channel, message, first) values(?, ?, ?, ?);`
	stmt, err := m.db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username, channel, content, first)
	return err
}
