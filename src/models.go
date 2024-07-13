package main

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
    Email    string `db:"email"`
    Password string `db:"password_hash"`
}

type Person struct {
    ID            string    `db:"id" json:"id"`
    Name          string    `db:"name" json:"name"`
    Answered      bool      `db:"answered" json:"answered"`
    Attending     bool      `db:"attending" json:"attending"`
    DateArriving  time.Time `db:"date_arriving" json:"date_arriving"`
    DateDeparture time.Time `db:"date_departure" json:"date_departure"`
    Comment       string    `db:"comment" json:"comment"`
}

func NewAdmin(email string, password string) Admin {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
        log.Fatal("could not hash password :: ", err)
	}

    return Admin{
        Email: email,
        Password: string(hashedPassword),
    }
}


