package app

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDB(dbUrl string, production bool) *sqlx.DB {
    db, err := sqlx.Open("postgres", dbUrl)
    if err != nil {
        log.Fatal("could not open db :: ", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal("could not ping db :: ", err)
    }

    return db
}

