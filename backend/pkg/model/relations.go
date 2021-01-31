package model

import "github.com/jmoiron/sqlx"

func GetRelatedPeople(db *sqlx.DB, filmID int) ([]Person, error) {
	var people []Person

	err := db.Select(&people, "SELECT * FROM people WHERE id IN (SELECT person_id FROM relations WHERE film_id = $1);",
		filmID)

	return people, err
}

func GetRelatedFilms(db *sqlx.DB, personID int) ([]Film, error) {
	var films []Film

	err := db.Select(&films, "SELECT * FROM films WHERE id IN (SELECT film_id FROM relations WHERE person_id = $1);",
		personID)

	return films, err
}
