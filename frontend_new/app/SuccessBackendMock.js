import AbstractBackend from "./AbstractBackend.js";

export default class SuccessBackendMock extends AbstractBackend {
  async getFilm(filmId) {
    return {
      "film": {
        "id": 1,
        "title": "Sample Film",
        "poster_url": "sample_url",
        "type": "movie",
        "start_year": 2019,
        "end_year": 0,
        "runtime_minutes": 100,
        "imdb_id": "tt000000001"
      },
      "people": [{
        "primary_name": "John Doe",
        "photo_url": "http://example.com",
        "birth_year": 1985,
        "death_year": 0,
        "imdb_id": "nm000000000001"
      }]
    };
  }
}
