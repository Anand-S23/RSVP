package main

import "github.com/jmoiron/sqlx"

type Store struct {
    DB *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
    return &Store{ DB: db }
}

func (s *Store) CreatePerson() {
}

func (s *Store) ReadPerson()  {
}

func (s *Store) ReadPeople() {
}

func (s *Store) UpdatePerson() {
}

