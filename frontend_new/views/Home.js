import AbstractView from "./AbstractView.js";

export default class Home extends AbstractView {
    constructor(backendService, params, document) {
        super(backendService, params, document);
        this.setTitle("Home");
    }

    async getHtml() {
      const template = require('../../templates/home.html');
      return template();
    }
}
