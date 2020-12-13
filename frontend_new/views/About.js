import AbstractView from "./AbstractView.js";

export default class About extends AbstractView {
    constructor(backendService, params, document) {
        super(backendService, params, document);
        this.setTitle("About us");
    }

    async getHtml() {
      const template = require('../../templates/about.html');
      return template({ 'data': "foo" });
    }
}
