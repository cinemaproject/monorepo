package model

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

type Person struct {
	ID          int    `db:"id" json:"id"`
	PrimaryName string `db:"primary_name" json:"primary_name"`
	PhotoURL    string `db:"photo_url" json:"photo_url"`
	BirthYear   int    `db:"birth_year" json:"birth_year"`
	DeathYear   int    `db:"death_year" json:"death_year"`
	IMDBID      string `db:"imdb_id" json:"imdb_id"`
}

func GetPersonByID(db *sqlx.DB, id int) (Person, error) {
	person := Person{}

	err := db.Get(&person, "SELECT * FROM people WHERE id=$1;", id)

	return person, err
}

func FindPeopleByName(db *sqlx.DB, name string) ([]Person, error) {
	var people []Person

	err := db.Select(&people,
		"SELECT * FROM people WHERE lower(primary_name) LIKE $1;", "%"+strings.ToLower(name)+"%")

	return people, err
}
