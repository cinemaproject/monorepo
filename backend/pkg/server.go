package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/cinematic/monorepo/backend/pkg/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type PersonWithFilms struct {
	Person model.Person `json:"person"`
	Films  []model.Film `json:"films"`
}

type FilmWithPeople struct {
	Film   model.Film     `json:"film"`
	People []model.Person `json:"people"`
}

type RequestError struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func sendError(w http.ResponseWriter, errCode int, msg string) {
	w.WriteHeader(errCode)
	errMessage, err := json.Marshal(RequestError{errCode, msg})
	if err != nil {
		log.Fatal("Failed to convert to json")
	}
	w.Write(errMessage)
}

func handlePersonByIdRequest(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	vars := mux.Vars(r)
	personID, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, 400, "ID must be an integer")
		return
	}

	person, err := model.GetPersonByID(db, personID)
	if err != nil {
		log.Printf("Unhandled err: %s", err)
		sendError(w, 404, "Not found")
		return
	}

	relatedFilms, err := model.GetRelatedFilms(db, personID)
	if err != nil {
		log.Printf("Unhandled err: %s", err)
		sendError(w, 404, "Not found")
		return
	}

	personWithFilms := PersonWithFilms{person, relatedFilms}

	bytes, err := json.Marshal(personWithFilms)
	_, err = fmt.Fprint(w, string(bytes))
	if err != nil {
		log.Printf("Unhandled error: %s", err)
		sendError(w, 404, "Not found")
	}
}

func handlePeopleSearch(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	vars := r.URL.Query()
	name := vars.Get("name")

	if name == "" {
		sendError(w, 400, "Request must have non-empty name parameter")
		return
	}

	people, err := model.FindPeopleByName(db, name)

	if err != nil {
		sendError(w, 404, "Not found")
		log.Printf("Unhandled err: %s", err)
		return
	}

	bytes, err := json.Marshal(people)
	_, err = fmt.Fprint(w, string(bytes))
	if err != nil {
		sendError(w, 404, "Not found")
		log.Printf("Unhandled error: %s", err)
	}
}

func handleFilmByIDRequest(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	vars := mux.Vars(r)
	filmID, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, 400, "ID must be an integer")
		return
	}

	film, err := model.GetFilmByID(db, filmID)
	if err != nil {
		sendError(w, 404, "Not found")
		log.Printf("Unhandled err: %s", err)
		return
	}

	relatedPeople, err := model.GetRelatedPeople(db, filmID)
	if err != nil {
		sendError(w, 404, "Not found")
		log.Printf("Unhandled err: %s", err)
		return
	}

	filmWithPeople := FilmWithPeople{film, relatedPeople}

	bytes, err := json.Marshal(filmWithPeople)
	_, err = fmt.Fprint(w, string(bytes))
	if err != nil {
		sendError(w, 404, "Not found")
		log.Printf("Unhandled error: %s", err)
		return
	}
}

func handleFilmsSearch(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	vars := r.URL.Query()
	title := vars.Get("title")

	if title == "" {
		sendError(w, 400, "Request must have non-empty title parameter")
		return
	}

	films, err := model.FindFilmsByTitle(db, title)

	if err != nil {
		log.Printf("Unhandled err: %s", err)
		sendError(w, 404, "Not found")
		return
	}

	bytes, err := json.Marshal(films)
	_, err = fmt.Fprint(w, string(bytes))
	if err != nil {
		sendError(w, 404, "Not found")
		log.Printf("Unhandled error: %s", err)
		return
	}
}

// InitializeRouter creates a new router and adds handlers for REST APIs.
func InitializeRouter(db *sqlx.DB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/people/search", func(w http.ResponseWriter, r *http.Request) {
		handlePeopleSearch(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/api/people/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlePersonByIdRequest(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/api/films/search", func(w http.ResponseWriter, r *http.Request) {
		handleFilmsSearch(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/api/films/{id}", func(w http.ResponseWriter, r *http.Request) {
		handleFilmByIDRequest(w, r, db)
	}).Methods("GET")

	return r
}
