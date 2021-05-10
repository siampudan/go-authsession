package database

import "github.com/go-pg/pg/v10"

func GetConnection() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "learn",
	})
	return db
}
