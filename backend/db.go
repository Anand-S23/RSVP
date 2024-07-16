package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDB(dbUrl string, production bool, email string, password string) *sqlx.DB {
    db, err := sqlx.Open("postgres", dbUrl)
    if err != nil {
        log.Fatal("could not open db :: ", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal("could not ping db :: ", err)
    }

    err = CreateDBAdmin(db, email, password)
    if err != nil {
        log.Fatal("could not create db admin :: ", err)
    }

    return db
}

func CreateDBAdmin(db *sqlx.DB, email string, password string) error {
    admin := NewAdmin(email, password)
    insertUserCommand := `insert into admins (email, password_hash) values (:email, :password_hash)`
    _, err := db.NamedExec(insertUserCommand, admin)
    return err
}

func PopulatePeople(db *sqlx.DB) error {
    return nil
}
