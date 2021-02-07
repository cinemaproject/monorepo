package model

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

var DefaultFilmsRow = []string{"id", "title", "poster_url", "type", "start_year", "end_year", "runtime_minutes", "imdb_id", "descr", "album_id"}

func TestGetFilmByID(t *testing.T) {
	// Arrange
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal()
	}
	defer mockDB.Close()

	db := sqlx.NewDb(mockDB, "sqlmock")

	row := mock.NewRows(DefaultFilmsRow).AddRow(1, "Test", "", "movie", 2015, 0, 100, "tt00001", "", "")
	noRows := mock.NewRows(DefaultFilmsRow)

	mock.ExpectQuery("^SELECT (.+) FROM films WHERE id=(.+)$").WithArgs(1).WillReturnRows(row)
	mock.ExpectQuery("^SELECT (.+) FROM films WHERE id=(.+)$").WithArgs(2).WillReturnRows(noRows)

	// Act
	film, err := GetFilmByID(db, 1)

	// Assert
	if err != nil {
		t.Errorf("GetFilmByID returned an error %s", err)
	}
	if film.Title != "Test" {
		t.Error()
	}
	if film.ID != 1 {
		t.Error()
	}

	// Act
	_, err = GetFilmByID(db, 2)

	// Assert
	if err == nil {
		t.Error("GetFilmByID succeeded in returning non-existent row")
	}
}

func TestFindFilmsByTitle(t *testing.T) {
	// Arrange
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal()
	}
	defer mockDB.Close()

	db := sqlx.NewDb(mockDB, "sqlmock")

	rows := mock.NewRows(DefaultFilmsRow).
		AddRow(1, "Star Wars Episode 1", "", "movie", 1000, 0, 100, "tt00001", "", "").
		AddRow(2, "Star Wars Episode 2", "", "movie", 1000, 0, 100, "tt00002", "", "").
		AddRow(3, "Star Trek", "", "movie", 1000, 0, 100, "tt00003", "", "")

	noRows := mock.NewRows(DefaultFilmsRow)

	mock.ExpectQuery("^SELECT (.+) FROM films WHERE (.+)$").WithArgs("%star%").WillReturnRows(rows)
	mock.ExpectQuery("^SELECT (.+) FROM films WHERE (.+)$").WithArgs("%avengers%").WillReturnRows(noRows)

	// Act
	films, err := FindFilmsByTitle(db, "Star")

	// Assert
	if err != nil {
		t.Errorf("Failed to fetch films by title: %s", err)
	}
	if len(films) != 3 {
		t.Errorf("Expected 3 films to be returned, got %d", len(films))
	}

	// Act
	films, err = FindFilmsByTitle(db, "Avengers")

	// Assert
	if err != nil {
		t.Errorf("Failed to fetch films by title: %s", err)
	}
	if len(films) != 0 {
		t.Errorf("Expected 0 films to be returned, got %d", len(films))
	}
}
