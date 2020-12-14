import AbstractBackend from "./AbstractBackend.js";

export default class SuccessBackendMock extends AbstractBackend {
  async getFilm(filmId) {
    return {
      "film": [{
        "id": "t000000001",
        "title": "Sample Film",
        "type": "movie",
        "start_year": 2019,
        "end_year": null,
        "runtime_minutes": 100
      }],
      "people": [{
        "id": "nm000000000001",
        "primary_name": "John Doe",
        "photo_url": "http://example.com",
        "birth_year": 1985,
        "death_year": null
      }]
    };
  }
}
