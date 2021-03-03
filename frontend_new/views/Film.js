import AbstractView from "./AbstractView.js";

export default class Film extends AbstractView {
    constructor(backendService, params, document) {
        super(backendService, params, document);
        this.setTitle("Film");
    }

    async getHtml() {
      const json = await this.backendService.getFilm(this.params.id);
      const template = require('../../templates/film.html');
      json.years = json.film.start_year;
      if (json.film.end_year != 0) {
        json.years += " &mdash " + json.film.end_year;
      }
      return template(json);
    }
}
