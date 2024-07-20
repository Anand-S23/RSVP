package app

import "github.com/jmoiron/sqlx"

type Store struct {
    DB *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
    return &Store{ DB: db }
}

func (s *Store) CreatePerson(person Person) error {
    query := `
		INSERT INTO people (id, name, answered, attending, date_arriving, date_departure, comment)
		VALUES (:id, :name, :answered, :attending, :date_arriving, :date_departure, :comment)
	`
	_, err := s.DB.NamedExec(query, person)
	return err
}

func (s *Store) ReadPerson(id string) (*Person, error) {
    query := `SELECT id, name, answered, attending, date_arriving, date_departure, comment FROM people WHERE id = $1`

	var person Person
	err := s.DB.Get(&person, query, id)
	if err != nil {
        return nil, err
	}

	return &person, nil
}

func (s *Store) ReadPeople() ([]Person, error) {
    query := `SELECT id, name, answered, attending, date_arriving, date_departure, comment FROM people`

	var people []Person
	err := s.DB.Select(&people, query)
	if err != nil {
		return nil, err
	}

	return people, nil
}

func (s *Store) UpdatePerson(person Person) error {
    query := `
		UPDATE people
		SET answered = :answered,
			attending = :attending,
			date_arriving = :date_arriving,
			date_departure = :date_departure,
			comment = :comment
		WHERE id = :id
	`

	_, err := s.DB.NamedExec(query, person)
    return err
}

