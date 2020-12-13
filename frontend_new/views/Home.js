import AbstractView from "./AbstractView.js";

export default class Home extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Dashboard");
    }

    async getHtml() {
      const template = require('../../templates/home.html');
      return template();
    }
}

