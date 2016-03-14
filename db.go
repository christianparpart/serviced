package main

import (
	"database/sql"
	"fmt"
)

type DB struct {
	hostname string
	port     int
	username string
	password string
	name     string
	charset  string
	db       *sql.DB
}

func OpenDB(hostname string, port int,
	username string, password string, name string) (*DB, error) {

	db_host := "localhost"
	db_name := "serviced"
	db_user := "root"
	db_pass := ""
	db_charset := "utf8"

	var db *DB = &DB{
		username: db_user,
		password: db_pass,
		hostname: db_host,
		name:     db_name,
		charset:  db_charset}

	err := db.Connect()

	return db, err
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
