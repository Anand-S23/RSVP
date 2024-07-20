package main

import (
	"time"

	"github.com/google/uuid"
)

type Person struct {
    ID            string    `db:"id" json:"id"`
    Name          string    `db:"name" json:"name"`
    Answered      bool      `db:"answered" json:"answered"`
    Attending     bool      `db:"attending" json:"attending"`
    DateArriving  time.Time `db:"date_arriving" json:"date_arriving"`
    DateDeparture time.Time `db:"date_departure" json:"date_departure"`
    Comment       string    `db:"comment" json:"comment"`
}

func NewPerson(name string) Person {
    return Person{
        ID: uuid.NewString(),
        Name: name,
        Answered: false,
        Attending: false,
        DateArriving: time.Now(),
        DateDeparture: time.Now(),
        Comment: "",
    }
}

func (p *Person) UpdatePerson(attending bool, arrival time.Time, departure time.Time, comment string) {
    p.Answered = true
    p.Attending = attending
    p.DateArriving = arrival
    p.DateDeparture = departure
    p.Comment = comment
}

