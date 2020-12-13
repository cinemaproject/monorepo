import AbstractView from "./AbstractView.js";

export default class About extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Dashboard");
    }

    async getHtml() {
      const template = require('../../templates/about.html');
      return template({ 'data': "foo" });
    }
}
