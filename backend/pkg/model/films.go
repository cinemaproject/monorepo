package model

import (
	"github.com/jmoiron/sqlx"
)

// Film describes a movie or series
type Film struct {
	ID int `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	PosterURL string `db:"poster_url" json:"poster_url"`
	Type string `db:"type" json:"type"`
	StartYear int `db:"start_year" json:"start_year"`
	EndYear int `db:"end_year" json:"end_year"`
	RuntimeMinutes int `db:"runtime_minutes" json:"runtime_minutes"`
	IMDBID string `db:"imdb_id" json:"imdb_id"`
}

func GetFilmByID(db *sqlx.DB, id int) (Film, error) {
	film := Film{}

	err := db.Get(&film, "SELECT * FROM films WHERE id=$1;", id)

	return film, err
}

func FindFilmsByTitle(db *sqlx.DB, title string) ([]Film, error) {
	var films []Film

	err := db.Select(&films, 
		"SELECT * FROM films WHERE lower(title) LIKE lower('%$1%');", title)

	return films, err
}
