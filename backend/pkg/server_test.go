package pkg

import (
	"database/sql"
	"encoding/json"
	"github.com/cinematic/monorepo/backend/pkg/model"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"testing"
)

var DefaultFilmsRow = []string{"id", "title", "poster_url", "type", "start_year", "end_year", "runtime_minutes", "imdb_id"}
var DefaultPeopleRow = []string{"id", "primary_name", "photo_url", "birth_year", "death_year", "imdb_id"}

const (
	ExpectPersonById = iota
	ExpectPersonByName = iota
	ExpectFilmsById = iota
	ExpectFilmsByTitle = iota
)

func _getTestServer(expected int) (*httptest.Server, *sql.DB, error) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	db := sqlx.NewDb(mockDB, "sqlmock")

	peopleRow := mock.NewRows(DefaultPeopleRow).AddRow(1, "John Doe", "", 1980, 0, "nm00001")
	filmsRow := mock.NewRows(DefaultFilmsRow).AddRow(1, "Test", "", "movie", 2015, 0, 100, "tt00001")
	if expected == ExpectPersonById {
		mock.ExpectQuery("^SELECT (.+) FROM people WHERE id=(.+)$").WithArgs(1).WillReturnRows(peopleRow)
	} else if expected == ExpectPersonByName {
		mock.ExpectQuery("^SELECT (.+) FROM people WHERE lower(.+)$").WithArgs("John").WillReturnRows(peopleRow)
	} else if expected == ExpectFilmsByTitle {
		mock.ExpectQuery("^SELECT (.+) FROM films WHERE (.+)$").WithArgs("Test").WillReturnRows(filmsRow)
	} else if expected == ExpectFilmsById {
		mock.ExpectQuery("^SELECT (.+) FROM films WHERE id=(.+)$").WithArgs(1).WillReturnRows(filmsRow)
	}

	router := InitializeRouter(db)
	ts := httptest.NewServer(router)

	return ts, mockDB, nil
}

func _generateRequest(t *testing.T, method, url string, body io.Reader) *http.Request {
	r, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatal(err)
	}

	return r
}

func TestGetPersonByIdRequest(t *testing.T) {
	// Arrange
	ts, db, err := _getTestServer(ExpectPersonById)
	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}
	defer ts.Close()
	defer db.Close()

	req := _generateRequest(t, "GET", ts.URL + "/api/people/1", nil)

	// Act
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to execute query. %s", err)
	}
	if resp == nil || resp.Body == nil {
		t.Fatal("Response is empty")
	}
	defer resp.Body.Close()

	// Assert
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to parse body: %s", err)
		}
		var person PersonWithFilms
		err = json.Unmarshal(bodyBytes, &person)
		if err != nil {
			bodyStr := string(bodyBytes)
			t.Fatalf("Failed to parse json: %s\n Body: %s", err, bodyStr)
		}
		if person.Person.ID != 1 {
			t.Errorf("Incorrect person id. Expected 1, got %d", person.Person.ID)
		}
	} else {
		t.Fatalf("Failed to execute query. Status is %s", resp.Status)
	}
}

func TestFindPeopleByNameRequest(t *testing.T) {
	ts, db, err := _getTestServer(ExpectPersonByName)
	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}
	defer ts.Close()
	defer db.Close()

	req := _generateRequest(t, "GET", ts.URL + "/api/people/search?name=John", nil)

	// Act
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to execute query. %s", err)
	}
	if resp == nil || resp.Body == nil {
		t.Fatal("Response is empty")
	}
	defer resp.Body.Close()

	// Assert
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to parse body: %s", err)
		}
		var people []model.Person
		err = json.Unmarshal(bodyBytes, &people)
		if err != nil {
			bodyStr := string(bodyBytes)
			t.Fatalf("Failed to parse json: %s\n Body: %s", err, bodyStr)
		}
		if len(people) != 1 {
			t.Errorf("Incorrect number of people. Expected 1, got %d", len(people))
		}
	} else {
		t.Fatalf("Failed to execute query. Status is %s", resp.Status)
	}
}

func TestFindFilmsByTitleRequest(t *testing.T) {
	ts, db, err := _getTestServer(ExpectFilmsByTitle)
	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}
	defer ts.Close()
	defer db.Close()

	req := _generateRequest(t, "GET", ts.URL + "/api/films/search?title=Test", nil)

	// Act
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to execute query. %s", err)
	}
	if resp == nil || resp.Body == nil {
		t.Fatal("Response is empty")
	}
	defer resp.Body.Close()

	// Assert
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to parse body: %s", err)
		}
		var films []model.Film
		err = json.Unmarshal(bodyBytes, &films)
		if err != nil {
			bodyStr := string(bodyBytes)
			t.Fatalf("Failed to parse json: %s\n Body: %s", err, bodyStr)
		}
		if len(films) != 1 {
			t.Errorf("Incorrect number of films. Expected 1, got %d", len(films))
		}
	} else {
		t.Fatalf("Failed to execute query. Status is %s", resp.Status)
	}
}

func TestGetFilmByIdRequest(t *testing.T) {
	// Arrange
	ts, db, err := _getTestServer(ExpectFilmsById)
	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}
	defer ts.Close()
	defer db.Close()

	req := _generateRequest(t, "GET", ts.URL + "/api/films/1", nil)

	// Act
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to execute query. %s", err)
	}
	if resp == nil || resp.Body == nil {
		t.Fatal("Response is empty")
	}
	defer resp.Body.Close()

	// Assert
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to parse body: %s", err)
		}
		var film FilmWithPeople
		err = json.Unmarshal(bodyBytes, &film)
		if err != nil {
			bodyStr := string(bodyBytes)
			t.Fatalf("Failed to parse json: %s\n Body: %s", err, bodyStr)
		}
		if film.Film.ID != 1 {
			t.Errorf("Incorrect person id. Expected 1, got %d", film.Film.ID)
		}
	} else {
		t.Fatalf("Failed to execute query. Status is %s", resp.Status)
	}
}
