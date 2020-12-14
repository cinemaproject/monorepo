import AbstractBackend from './AbstractBackend.js'

export default class Backend extends AbstractBackend {
  constructor(hostname) {
    super();
    this.baseURL = "/api/";
  }

  async __doRequest(method, endpoint, parameters) {
    let xhr = new XMLHttpRequest();
    xhr.open(method, this.baseURL + endpoint);
    xhr.send(parameters);
    return new Promise((resolve, reject) => {
      xhr.onload = function () {
        resolve(JSON.parse(xhr.response));
      }
      xhr.onerror = function () {
        reject();
      }
    });
  }

  async getFilm(filmId) {
    return await this.__doRequest("GET", "/films/" + filmId, null);
  }

  async getPerson(personId) {
    return await this.__doRequest("GET", "/people/" + personId, null);
  }

  async searchFilms(titlePattern) {
    return await this.__doRequest("GET", "/films/search?title=" + titlePattern, null);
  }

  async searchPeople(namePattern) {
    return await this.__doRequest("GET", "/people/search?name=" + namePattern, null);
  }
}
