import AbstractView from "./AbstractView.js";

export default class Person extends AbstractView {
    constructor(backendService, params, document) {
        super(backendService, params, document);
        this.setTitle("Film");
    }

    async getHtml() {
      const json = await this.backendService.getPerson(this.params.id);
      const template = require('../../templates/person.html');
      json.death_year = "-";
      if (json.person.death_year != 0) {
        json.death_year = json.person.death_year;
      }
      return template(json);
    }
}
