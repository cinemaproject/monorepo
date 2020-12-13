import AbstractBackend from "../../app/AbstractBackend.js";

export default class SuccessBackendMock extends AbstractBackend {
  async getFilm(filmId) {
    return {
      "id": "t000000001",
      "title": "Sample Film",
      "type": "movie",
      "start_year": "2019",
      "end_year": "",
      "runtime_minutes": "200"
    };
  }
}
