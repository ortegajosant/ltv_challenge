package persistence

import (
	"database/sql"
)

var db_context *sql.DB

func db_init() *sql.DB {

	if db_context != nil {
		return db_context
	}

	var err error

	db_context, err := sql.Open("sqlite", "./sql/jrdd.db")
	if err != nil {
		panic(err)
	}

	return db_context
}
