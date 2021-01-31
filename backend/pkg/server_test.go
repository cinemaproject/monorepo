package pkg

import (
	"database/sql"
	"encoding/json"

	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cinematic/monorepo/backend/pkg/model"
	"github.com/jmoiron/sqlx"
)

var DefaultFilmsRow = []string{"id", "title", "poster_url", "type", "start_year", "end_year", "runtime_minutes", "imdb_id"}
var DefaultPeopleRow = []string{"id", "primary_name", "photo_url", "birth_year", "death_year", "imdb_id"}

const (
	ExpectPersonById   = iota
	ExpectPersonByName = iota
	ExpectFilmsById    = iota
	ExpectFilmsByTitle = iota
)

func _getTestServer() (*httptest.Server, *sql.DB, error) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	db := sqlx.NewDb(mockDB, "sqlmock")

	//////////////////////////////////////////////////////////////////////////////
	// Set up mock queries
	//////////////////////////////////////////////////////////////////////////////

	// Request person with correct ID
	mock.ExpectQuery("^SELECT (.+) FROM people WHERE id=(.+)$").WithArgs(1).WillReturnRows(mock.NewRows(
		DefaultPeopleRow).AddRow(1, "John Doe", "", 1980, 0, "nm00001"))
	mock.ExpectQuery("^SELECT (.+) FROM films WHERE id IN (.+)$").WithArgs(1).WillReturnRows(mock.NewRows(DefaultFilmsRow).AddRow(1, "Test", "", "movie", 2015, 0, 100, "tt00001"))

	// Request person with missing ID
	mock.ExpectQuery("^SELECT (.+) FROM people WHERE id=(.+)$").WithArgs(2).WillReturnRows(mock.NewRows(DefaultPeopleRow))

	// Request film with correct ID
	mock.ExpectQuery("^SELECT (.+) FROM films WHERE id=(.+)$").WithArgs(1).WillReturnRows(mock.NewRows(DefaultFilmsRow).AddRow(1, "Test", "", "movie", 2015, 0, 100, "tt00001"))
	mock.ExpectQuery("^SELECT (.+) FROM people WHERE id IN (.+)$").WithArgs(1).WillReturnRows(mock.NewRows(DefaultPeopleRow).AddRow(1, "John Doe", "", 1980, 0, "nm00001"))

	// Find person by name
	mock.ExpectQuery("^SELECT (.+) FROM people WHERE lower(.+)$").WithArgs("John").WillReturnRows(mock.NewRows(DefaultPeopleRow).AddRow(1, "John Doe", "", 1980, 0, "nm00001"))

	// Find person by missing name
	mock.ExpectQuery("^SELECT (.+) FROM people WHERE lower(.+)$").WithArgs("Jack").WillReturnRows(mock.NewRows(DefaultPeopleRow))

	// Find films by title
	mock.ExpectQuery("^SELECT (.+) FROM films WHERE lower(.+)$").WithArgs("Star").WillReturnRows(
		mock.NewRows(DefaultFilmsRow).
			AddRow(1, "Star Wars Episode 1", "", "movie", 1980, 0, 100, "tt00001").
			AddRow(2, "Star Wars Episode 2", "", "movie", 1980, 0, 100, "tt00001").
			AddRow(3, "Star Trek", "", "movie", 1980, 0, 100, "tt00001"))

	// Request film with missing ID
	mock.ExpectQuery("^SELECT (.+) FROM films WHERE lower(.+)$").WithArgs("Avengers").WillReturnRows(mock.NewRows(DefaultFilmsRow))

	router := InitializeRouter(db)
	ts := httptest.NewServer(router)

	return ts, mockDB, nil
}

func generateRequest(t *testing.T, method, url string, body io.Reader) *http.Request {
	r, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatal(err)
	}

	return r
}

type TestCheckType = func(*testing.T, []byte)

type TestCase struct {
	Name   string
	Req    *http.Request
	Verify TestCheckType
}

func verifyPersonWithID(t *testing.T, bodyBytes []byte) {
	var person PersonWithFilms
	err := json.Unmarshal(bodyBytes, &person)
	if err != nil {
		bodyStr := string(bodyBytes)
		t.Fatalf("Failed to parse json: %s\n Body: %s", err, bodyStr)
	}
	if person.Person.ID != 1 {
		t.Errorf("Incorrect person id. Expected 1, got %d", person.Person.ID)
	}
}

func verifyPeopleArray(length int) TestCheckType {
	return func(t *testing.T, bodyBytes []byte) {
		var people []model.Person
		err := json.Unmarshal(bodyBytes, &people)
		if err != nil {
			bodyStr := string(bodyBytes)
			t.Fatalf("Failed to parse json: %s\n Body: %s", err, bodyStr)
		}
		if len(people) != length {
			t.Errorf("Incorrect number of people. Expected %d, got %d", length, len(people))
		}
	}
}

func verifyFilmsArray(length int) TestCheckType {
	return func(t *testing.T, bodyBytes []byte) {
		var films []model.Film
		err := json.Unmarshal(bodyBytes, &films)
		if err != nil {
			bodyStr := string(bodyBytes)
			t.Fatalf("Failed to parse json: %s\n Body: %s", err, bodyStr)
		}
		if len(films) != length {
			t.Errorf("Incorrect number of people. Expected %d, got %d", length, len(films))
		}
	}
}

func verifySmthMissing(t *testing.T, bodyBytes []byte) {
	var reqErr RequestError
	err := json.Unmarshal(bodyBytes, &reqErr)
	if err != nil {
		bodyStr := string(bodyBytes)
		t.Fatalf("Failed to parse json: %s\nBody: %s", err, bodyStr)
	}
	if reqErr.ErrorCode != 404 {
		t.Errorf("Incorrect error code. Expected 404, got %d", reqErr.ErrorCode)
	}
}

func verifySmthIncorrect(t *testing.T, bodyBytes []byte) {
	var reqErr RequestError
	err := json.Unmarshal(bodyBytes, &reqErr)
	if err != nil {
		bodyStr := string(bodyBytes)
		t.Fatalf("Failed to parse json: %s\nBody: %s", err, bodyStr)
	}
	if reqErr.ErrorCode != 400 {
		t.Errorf("Incorrect error code. Expected 400, got %d", reqErr.ErrorCode)
	}
}

func verifyFilmWithID(t *testing.T, bodyBytes []byte) {
	var film FilmWithPeople
	err := json.Unmarshal(bodyBytes, &film)
	if err != nil {
		bodyStr := string(bodyBytes)
		t.Fatalf("Failed to parse json: %s\n Body: %s", err, bodyStr)
	}
	if film.Film.ID != 1 {
		t.Errorf("Incorrect film id. Expected 1, got %d", film.Film.ID)
	}
}

func TestServerRequests(t *testing.T) {
	// Arrange
	ts, db, err := _getTestServer()
	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}
	defer ts.Close()
	defer db.Close()

	// Describe test cases
	tests := []TestCase{
		{"Request person with ID", generateRequest(t, "GET", ts.URL+"/api/people/1", nil), verifyPersonWithID},
		{"Request person with missing ID", generateRequest(t, "GET", ts.URL+"/api/people/2", nil), verifySmthMissing},
		{"Request person with incorrect ID", generateRequest(t, "GET", ts.URL+"/api/people/foo", nil), verifySmthIncorrect},
		{"Request film with ID", generateRequest(t, "GET", ts.URL+"/api/films/1", nil), verifyFilmWithID},
		{"Request film with missing ID", generateRequest(t, "GET", ts.URL+"/api/films/2", nil), verifySmthMissing},
		{"Request film with incorrect ID", generateRequest(t, "GET", ts.URL+"/api/films/foo", nil), verifySmthIncorrect},
		{"Find person by name", generateRequest(t, "GET", ts.URL+"/api/people/search?name=John", nil), verifyPeopleArray(1)},
		{"Find person by missing name", generateRequest(t, "GET", ts.URL+"/api/people/search?name=Jack", nil), verifyPeopleArray(0)},
		{"Find person incorrect request", generateRequest(t, "GET", ts.URL+"/api/people/search", nil), verifySmthIncorrect},
		{"Find films by title", generateRequest(t, "GET", ts.URL+"/api/films/search?title=Star", nil), verifyFilmsArray(3)},
		{"Find films by missing title", generateRequest(t, "GET", ts.URL+"/api/films/search?title=Avengers", nil), verifyFilmsArray(0)},
		{"Find films incorrect request", generateRequest(t, "GET", ts.URL+"/api/films/search", nil), verifySmthIncorrect},
	}

	for _, tstCase := range tests {
		t.Run(tstCase.Name, func(t *testing.T) {
			// Act
			resp, _ := http.DefaultClient.Do(tstCase.Req)
			if resp == nil || resp.Body == nil {
				t.Fatal("Response is empty")
			}
			defer resp.Body.Close()

			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to parse body: %s", err)
			}

			// Assert
			tstCase.Verify(t, bodyBytes)
		})
	}
}
