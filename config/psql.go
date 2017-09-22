package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DB_USER = "ipriver"
	DB_PASS = "root"
	DB_NAME = "twitch_bot"
)

var Db *sql.DB

func init() {
	Db, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASS, DB_NAME))
	if err != nil {
		panic(err)
	}
}
