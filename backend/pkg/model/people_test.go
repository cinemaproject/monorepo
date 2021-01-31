package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"testing"
)

var DefaultPeopleRow = []string{"id", "primary_name", "photo_url", "birth_year", "death_year", "imdb_id"}

func TestGetPersonByID(t *testing.T) {
	// Arrange
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal()
	}
	defer mockDB.Close()

	db := sqlx.NewDb(mockDB, "sqlmock")

	row := mock.NewRows(DefaultPeopleRow).AddRow(1, "John Doe", "", 1980, 0, "nm00001")
	noRows := mock.NewRows(DefaultPeopleRow)

	mock.ExpectQuery("^SELECT (.+) FROM people WHERE id=(.+)$").WithArgs(1).WillReturnRows(row)
	mock.ExpectQuery("^SELECT (.+) FROM people WHERE id=(.+)$").WithArgs(2).WillReturnRows(noRows)

	// Act
	person, err := GetPersonByID(db, 1)

	// Assert
	if err != nil {
		t.Errorf("Failed to get person by id %s", err)
	}
	if person.PrimaryName != "John Doe" {
		t.Errorf("Person name didn't match")
	}

	// Act
	_, err = GetPersonByID(db, 2)

	// Assert
	if err == nil {
		t.Error("Fetched non-existent person")
	}
}

func TestFindPeopleByName(t *testing.T) {
	// Arrange
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal()
	}
	defer mockDB.Close()

	db := sqlx.NewDb(mockDB, "sqlmock")

	rows := mock.NewRows(DefaultPeopleRow).
		AddRow(1, "John Doe", "", 1980, 0, "nm00001").
		AddRow(2, "John Doe", "", 1980, 0, "nm00002").
		AddRow(3, "John Doe", "", 1980, 0, "nm00003")

	noRows := mock.NewRows(DefaultPeopleRow)

	mock.ExpectQuery("^SELECT (.+) FROM people WHERE (.+)$").WithArgs("%john%").WillReturnRows(rows)
	mock.ExpectQuery("^SELECT (.+) FROM people WHERE (.+)$").WithArgs("%jack%").WillReturnRows(noRows)

	// Act
	people, err := FindPeopleByName(db, "John")

	// Assert
	if err != nil {
		t.Errorf("Failed to find people: %s", err)
	}
	if len(people) != 3 {
		t.Errorf("Expected 3 people, got %d", len(people))
	}

	// Act
	people, err = FindPeopleByName(db, "Jack")

	// Assert
	if err != nil {
		t.Errorf("Failed to find people: %s", err)
	}
	if len(people) != 0 {
		t.Errorf("Expected 0 people, got %d", len(people))
	}
}
