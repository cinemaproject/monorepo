import AbstractView from "./AbstractView.js";

export default class Film extends AbstractView {
    constructor(backendService, params, document) {
        super(backendService, params, document);
        this.setTitle("Film");
    }

    async getHtml() {
      const json = await this.backendService.getFilm(this.params.id);
      const template = require('../../templates/film.html');
      return template(json);
    }
}
