package main

import (
	"database/sql"
	"fmt"
)

type DB struct {
	username string
	password string
	hostname string
	name     string
	charset  string
	db       *sql.DB
}

func NewDB() (*DB, error) {
	db_user := "root"
	db_pass := ""
	db_host := "localhost"
	db_name := "serviced"
	db_charset := "utf8"

	var db *DB = &DB{
		username: db_user,
		password: db_pass,
		hostname: db_host,
		name:     db_name,
		charset:  db_charset}

	return db, nil
}

func (db *DB) Connect() error {
	sqldb, err := sql.Open("mysql", fmt.Sprintf("%v:%v@%v/%v?charset=%v",
		db.username, db.password, db.hostname, db.charset))

	if err != nil {
		return err
	}

	db.db = sqldb
	return nil

	// TODO: run db migrations
	// db_sql_init := "TODO"
	// db.Query(db_sql_init)
}
