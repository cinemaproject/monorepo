package model

import "github.com/jmoiron/sqlx"

func GetRelatedPeople(db *sqlx.DB, filmID int) ([]Person, error) {
	var people []Person

	err := db.Select(&people, "select * from people where id in (select person_id from relations where film_id = $1);", filmID)

	return people, err
}

func GetRelatedFilms(db *sqlx.DB, personID int) ([]Film, error) {
	var films []Film

	err := db.Select(&films, "select * from films where id in (select film_id from relations where person_id = $1);", personID)

	return films, err
}
