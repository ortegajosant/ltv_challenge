package persistence

import (
	"database/sql"
)

var db_context *sql.DB

var db_path = "./sql/jrdd.db"

func Db_init() *sql.DB {

	if db_context != nil {
		return db_context
	}

	var err error
	db_context, err := sql.Open("sqlite3", db_path)
	if err != nil {
		panic(err)
	}

	return db_context
}
